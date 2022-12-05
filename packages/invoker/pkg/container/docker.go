package container

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/hanchchch/gimi/packages/invoker/pkg/utils"
)

type DockerContainer struct {
	ID      string
	manager *Manager
}

type DockerContainerConfig struct {
	AttachStdin  bool
	AttachStdout bool
	AttachStderr bool
	StopTimeout  *int
	Env          []string
	Cmd          []string
	Image        string
}

func (c *DockerContainer) GetID() string {
	return c.ID
}

func (c *DockerContainer) Start() error {
	return c.manager.docker.ContainerStart(context.Background(), c.ID, types.ContainerStartOptions{})
}

func (c *DockerContainer) Wait(timeout time.Duration) (container.ContainerWaitOKBody, error) {
	statusCh, errCh := c.manager.docker.ContainerWait(context.Background(), c.ID, container.WaitConditionNotRunning)

	if timeout > 0 {
		timer := time.NewTimer(timeout)
		select {
		case err := <-errCh:
			if !timer.Stop() {
				<-timer.C
			}
			return container.ContainerWaitOKBody{}, err
		case body := <-statusCh:
			return body, nil
		case <-timer.C:
			c.Stop()
			return container.ContainerWaitOKBody{}, fmt.Errorf("timeout waiting for container %s", c.ID)
		}
	} else {
		select {
		case err := <-errCh:
			return container.ContainerWaitOKBody{}, err
		case body := <-statusCh:
			return body, nil
		}
	}
}

func (c *DockerContainer) Stop() error {
	return c.manager.docker.ContainerStop(context.Background(), c.ID, nil)
}

func (c *DockerContainer) Logs() ([]byte, []byte, error) {
	r, err := c.manager.docker.ContainerLogs(context.Background(), c.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		return nil, nil, err
	}

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	stdcopy.StdCopy(stdout, stderr, r)
	return stdout.Bytes(), stderr.Bytes(), nil
}

func (c *DockerContainer) Run(timeout time.Duration) error {
	if err := c.Start(); err != nil {
		return err
	}

	_, err := c.Wait(timeout)
	return err
}

func (c *DockerContainer) Remove() error {
	return c.manager.docker.ContainerRemove(context.Background(), c.GetID(), types.ContainerRemoveOptions{})
}

func (m *Manager) CreateDockerContainer(config *DockerContainerConfig) (Container, error) {
	resp, err := m.docker.ContainerCreate(context.Background(), &container.Config{
		AttachStdin:  config.AttachStdin,
		AttachStdout: config.AttachStdout,
		AttachStderr: config.AttachStderr,
		Env:          config.Env,
		Cmd:          config.Cmd,
		StopTimeout:  config.StopTimeout,
		Image:        config.Image,
	}, nil, nil, nil, m.namespace+"_"+utils.RandString(12))
	if err != nil {
		return nil, err
	}

	c := &DockerContainer{
		ID:      resp.ID,
		manager: m,
	}

	m.AppendContainer(c)

	return c, nil
}

func (m *Manager) PullImage(refStr string) (io.ReadCloser, error) {
	return m.docker.ImagePull(context.Background(), refStr, types.ImagePullOptions{})
}

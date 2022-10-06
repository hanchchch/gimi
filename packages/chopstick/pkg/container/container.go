package container

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

type Container struct {
	ID      string
	manager *Manager
}

func (c *Container) Start() error {
	return c.manager.docker.ContainerStart(context.Background(), c.ID, types.ContainerStartOptions{})
}

func (c *Container) Wait(timeout time.Duration) (container.ContainerWaitOKBody, error) {
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

func (c *Container) Stop() error {
	return c.manager.docker.ContainerStop(context.Background(), c.ID, nil)
}

func (c *Container) Logs() (io.ReadCloser, error) {
	return c.manager.docker.ContainerLogs(context.Background(), c.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
}

func (c *Container) Run(timeout time.Duration) error {
	if err := c.Start(); err != nil {
		return err
	}

	_, err := c.Wait(timeout)
	return err
}

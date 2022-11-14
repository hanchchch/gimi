package container

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/hanchchch/gimi/packages/invoker/pkg/utils"
	pb "github.com/hanchchch/gimi/packages/proto/go/messages"
)

const (
	TryImageName     = "gimi-inspection"
	ManagerNamespace = "invoker"
)

type Manager struct {
	docker     *client.Client
	containers map[string]*Container
	namespace  string
}

type ContainerConfig struct {
	AttachStdin  bool
	AttachStdout bool
	AttachStderr bool
	StopTimeout  *int
	Env          []string
	Cmd          []string
	Image        string
}

type InspectionContainerConfig struct {
	AttachStdin  bool
	AttachStdout bool
	AttachStderr bool
	StopTimeout  *int
	Env          []string
	Args         *pb.InspectionArgs
}

func NewManager() *Manager {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	return &Manager{
		docker:     cli,
		containers: make(map[string]*Container),
		namespace:  ManagerNamespace,
	}
}

func (m *Manager) PullImage(refStr string) (io.ReadCloser, error) {
	return m.docker.ImagePull(context.Background(), refStr, types.ImagePullOptions{})
}

func (m *Manager) CreateContainer(config *ContainerConfig) (*Container, error) {
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

	c := &Container{
		ID:      resp.ID,
		manager: m,
	}

	m.containers[c.ID] = c

	return c, nil
}

func (m *Manager) CreateInspectionContainer(config *InspectionContainerConfig) (*Container, error) {
	return m.CreateContainer(&ContainerConfig{
		AttachStdin:  config.AttachStdin,
		AttachStdout: config.AttachStdout,
		AttachStderr: config.AttachStderr,
		Env:          config.Env,
		Image:        TryImageName,
		Cmd:          []string{"inspection", "-url", config.Args.Url},
		StopTimeout:  config.StopTimeout,
	})
}

func (m *Manager) RemoveContainer(c *Container) error {
	delete(m.containers, c.ID)
	return m.docker.ContainerRemove(context.Background(), c.ID, types.ContainerRemoveOptions{})
}

func (m *Manager) ListContainers() ([]*Container, error) {
	containers, err := m.docker.ContainerList(context.Background(), types.ContainerListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{Key: "name", Value: m.namespace}),
		All:     true,
	})

	if err != nil {
		return nil, err
	}

	var result []*Container
	for _, container := range containers {
		c := &Container{
			ID:      container.ID,
			manager: m,
		}
		result = append(result, c)
	}

	return result, nil
}

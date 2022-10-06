package container

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/hanchchch/gimi/packages/chopstick/pkg/utils"
)

type Manager struct {
	docker     *client.Client
	containers map[string]*Container
	namespace  string
}

func NewManager() *Manager {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	return &Manager{
		docker:     cli,
		containers: make(map[string]*Container),
		namespace:  "chopstick",
	}
}

func (m *Manager) PullImage(refStr string) (io.ReadCloser, error) {
	return m.docker.ImagePull(context.Background(), refStr, types.ImagePullOptions{})
}

func (m *Manager) CreateContainer(config *container.Config) *Container {
	resp, err := m.docker.ContainerCreate(context.Background(), config, nil, nil, nil, m.namespace+"_"+utils.RandString(12))
	if err != nil {
		panic(err)
	}

	c := &Container{
		ID:      resp.ID,
		manager: m,
	}

	m.containers[c.ID] = c

	return c
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

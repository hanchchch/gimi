package container

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/docker/docker/client"
	pb "github.com/hanchchch/gimi/packages/proto/go/messages"
)

const (
	TryImageName     = "gimi-inspection"
	ManagerNamespace = "invoker"
)

type Manager struct {
	docker     *client.Client
	aws        *session.Session
	containers map[string]Container
	namespace  string
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
		containers: make(map[string]Container),
		namespace:  ManagerNamespace,
	}
}

func (m *Manager) AppendContainer(c Container) (Container, error) {
	m.containers[c.GetID()] = c
	return c, nil
}

func (m *Manager) RemoveContainer(c Container) error {
	err := c.Remove()
	if err != nil {
		return err
	}
	delete(m.containers, c.GetID())
	return nil
}

func (m *Manager) ListContainers() ([]Container, error) {
	var result []Container
	for _, container := range m.containers {
		result = append(result, container)
	}

	return result, nil
}

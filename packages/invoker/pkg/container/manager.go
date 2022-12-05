package container

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/docker/docker/client"
)

const (
	TryImageName     = "gimi-inspection"
	ManagerNamespace = "invoker"
)

type AWSCredentials struct {
	AccessKeyId     string
	SecretAccessKey string
}

type Manager struct {
	docker         *client.Client
	aws            *session.Session
	containers     map[string]Container
	namespace      string
	awsCredentials AWSCredentials
}

type ManagerConfig struct {
	AWS AWSCredentials
}

func NewManager(config *ManagerConfig) (*Manager, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	m := &Manager{
		docker:     cli,
		containers: make(map[string]Container),
		namespace:  ManagerNamespace,
	}
	if config.AWS.AccessKeyId != "" {
		static_credentials := credentials.NewStaticCredentials(
			config.AWS.AccessKeyId,
			config.AWS.SecretAccessKey,
			"",
		)

		sess, err := session.NewSession(&aws.Config{
			Credentials: static_credentials,
			Region:      aws.String("ap-northeast-2"),
		})

		if err != nil {
			return nil, fmt.Errorf("unable to start aws session: %v", err)
		}

		m.aws = sess
		m.awsCredentials = config.AWS
	}

	return m, nil
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

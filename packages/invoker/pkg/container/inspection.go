package container

import (
	"fmt"

	pb "github.com/hanchchch/gimi/packages/proto/go/messages"
)

type InspectionContainerConfig struct {
	Provider string
	Name     string
	Args     *pb.InspectionArgs
}

func (m *Manager) CreateInspectionContainer(config *InspectionContainerConfig) (Container, error) {
	if config.Provider == "docker" {
		return m.CreateDockerContainer(&DockerContainerConfig{
			Image: config.Name,
			Cmd:   []string{"inspection", "-url", config.Args.Url},
			Env: []string{
				"AWS_ACCESS_KEY_ID=" + m.awsCredentials.AccessKeyId,
				"AWS_SECRET_ACCESS_KEY=" + m.awsCredentials.SecretAccessKey,
			},
		})
	} else if config.Provider == "lambda" {
		return m.CreateLambdaContainer(&LambdaContainerConfig{
			FunctionName: config.Name,
			Payload:      map[string]string{"url": config.Args.Url},
			Env: []string{
				"AWS_ACCESS_KEY_ID=" + m.awsCredentials.AccessKeyId,
				"AWS_SECRET_ACCESS_KEY=" + m.awsCredentials.SecretAccessKey,
			},
		})
	} else {
		return nil, fmt.Errorf("unknown provider: %s", config.Provider)
	}
}

package container

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/lambda"
	pb "github.com/hanchchch/gimi/packages/proto/go/messages"
)

type LambdaContainer struct {
	ID           string
	manager      *Manager
	result       *lambda.InvokeOutput
	functionName string
	Payload      []byte
}

type LambdaContainerConfig struct {
	Env          []string
	Cmd          []string
	FunctionName string
	Payload      map[string]string
}

func (c *LambdaContainer) GetID() string {
	return c.ID
}

func (c *LambdaContainer) Logs() ([]byte, []byte, error) {
	if c.result.FunctionError != nil {
		return nil, c.result.Payload, nil
	} else {
		return c.result.Payload, nil, nil
	}
}

func (c *LambdaContainer) GetResult() (*pb.InspectionResult, error) {
	if c.result.FunctionError != nil {
		return nil, fmt.Errorf("lambda error: %s, %v", *c.result.FunctionError, string(c.result.Payload))
	}

	var r pb.InspectionResult
	err := json.Unmarshal(c.result.Payload, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *LambdaContainer) Run(timeout time.Duration) error {
	svc := lambda.New(c.manager.aws)
	input := &lambda.InvokeInput{
		FunctionName: &c.functionName,
		Payload:      c.Payload,
	}

	result, err := svc.Invoke(input)
	if err != nil {
		return err
	}

	c.result = result
	return nil
}

func (c *LambdaContainer) Remove() error {
	return nil
}

func (m *Manager) CreateLambdaContainer(config *LambdaContainerConfig) (Container, error) {
	if m.aws == nil {
		return nil, fmt.Errorf("AWS not configured")
	}

	pbytes, err := json.Marshal(config.Payload)
	if err != nil {
		return nil, err
	}

	c := &LambdaContainer{
		ID:           "ID",
		manager:      m,
		functionName: config.FunctionName,
		Payload:      pbytes,
	}

	m.AppendContainer(c)

	return c, nil
}

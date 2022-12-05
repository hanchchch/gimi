package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/hanchchch/gimi/packages/invoker/pkg/container"
	"github.com/hanchchch/gimi/packages/invoker/pkg/listener"
	"github.com/hanchchch/gimi/packages/invoker/pkg/utils"
	"github.com/joho/godotenv"

	pb "github.com/hanchchch/gimi/packages/proto/go/messages"
)

const (
	inspectionName = "gimi-inspection"
	bound          = "-----"
	timeout        = 60 * time.Second
)

func main() {
	err := godotenv.Load()
	if err == nil {
		fmt.Println("loading .env file")
	}

	listeners, err := listener.NewListeners(listener.ListenerOptions{
		HTTP: listener.HTTPListenerOptions{
			Addr: utils.Getenv("HTTP_ADDR", ""),
		},
		Redis: listener.RedisListenerOptions{
			Url: utils.Getenv("REDIS_URL", ""),
		},
	})

	if err != nil {
		panic(err)
	}

	awsAccessKey := utils.Getenv("AWS_ACCESS_KEY_ID", "")
	awsSecretKey := utils.Getenv("AWS_SECRET_ACCESS_KEY", "")

	m, err := container.NewManager(&container.ManagerConfig{
		AWS: container.AWSCredentials{
			AccessKeyId:     awsAccessKey,
			SecretAccessKey: awsSecretKey,
		},
	})

	if err != nil {
		panic(err)
	}

	handler := func(args *pb.HandlerArgs) (*pb.HandlerResult, error) {
		log.Printf("received request: %s", args.Args.Url)
		c, err := m.CreateInspectionContainer(&container.InspectionContainerConfig{
			Provider: "lambda",
			Name:     inspectionName,
			Args:     args.Args,
		})

		if err != nil {
			return nil, fmt.Errorf("failed to create container: %w", err)
		}

		if err := c.Run(timeout); err != nil {
			return nil, fmt.Errorf("failed to run container: %w", err)
		}

		r, err := c.GetResult()

		if err != nil {
			return nil, fmt.Errorf("failed to get logs from container: %w", err)
		}

		if err := m.RemoveContainer(c); err != nil {
			return nil, fmt.Errorf("failed to remove container: %w", err)
		}

		fmt.Printf("result: %v\n", r.String())

		return &pb.HandlerResult{
			RequestId: args.RequestId,
			Result:    r,
		}, nil
	}

	for _, l := range listeners {
		if err := l.OnInvoke(handler); err != nil {
			panic(err)
		}
	}

	var wg sync.WaitGroup
	for _, l := range listeners {
		wg.Add(1)
		go func(l listener.Listener) {
			if err := l.Listen(); err != nil {
				panic(err)
			}
			wg.Done()
		}(l)
	}
	wg.Wait()
}

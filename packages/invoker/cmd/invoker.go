package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"github.com/hanchchch/gimi/packages/invoker/pkg/container"
	"github.com/hanchchch/gimi/packages/invoker/pkg/listener"
	"github.com/hanchchch/gimi/packages/invoker/pkg/utils"
	"github.com/joho/godotenv"
	"google.golang.org/protobuf/proto"

	pb "github.com/hanchchch/gimi/packages/proto/go/messages"
)

const (
	bound = "-----"
)

func main() {
	m := container.NewManager()
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

	handler := func(args *pb.HandlerArgs) (*pb.HandlerResult, error) {
		c, err := m.CreateInspectionContainer(&container.InspectionContainerConfig{
			Args: args.Args,
			Env: []string{
				"AWS_ACCESS_KEY_ID=" + awsAccessKey,
				"AWS_SECRET_ACCESS_KEY=" + awsSecretKey,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create container: %w", err)
		}

		if err := c.Run(60 * time.Second); err != nil {
			return nil, fmt.Errorf("failed to run container: %w", err)
		}

		stdout, stderr, err := c.Logs()

		if err != nil {
			return nil, fmt.Errorf("failed to get logs from container: %w", err)
		}

		if err := m.RemoveContainer(c); err != nil {
			return nil, fmt.Errorf("failed to remove container: %w", err)
		}

		r := &pb.InspectionResult{}
		splited := bytes.Split(stdout, []byte(bound))
		if len(splited) < 2 {
			return nil, fmt.Errorf("failed to parse stdout:\n%s\n\nstderr:\n%s", stdout, stderr)
		}
		if err := proto.Unmarshal(splited[1], r); err != nil {
			return nil, fmt.Errorf("failed to unmarshal inspection result: %w", err)
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

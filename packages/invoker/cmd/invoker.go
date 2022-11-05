package main

import (
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

	handler := func(args *pb.HandlerArgs) (*pb.HandlerResult, error) {
		c, err := m.CreateTryContainer(&container.TryContainerConfig{
			Args: args.Args,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create container: %w", err)
		}

		if err := c.Run(10 * time.Second); err != nil {
			return nil, fmt.Errorf("failed to run container: %w", err)
		}

		stdout, _, err := c.Logs()

		if err != nil {
			return nil, fmt.Errorf("failed to get logs from container: %w", err)
		}

		if err := m.RemoveContainer(c); err != nil {
			return nil, fmt.Errorf("failed to remove container: %w", err)
		}

		r := &pb.InspectionResult{}
		proto.Unmarshal(stdout, r)

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

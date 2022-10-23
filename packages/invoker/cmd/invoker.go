package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/hanchchch/gimi/packages/invoker/pkg/container"
	"github.com/hanchchch/gimi/packages/invoker/pkg/listener"
	"github.com/hanchchch/gimi/packages/invoker/pkg/utils"
	"github.com/joho/godotenv"
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

	handler := func(args listener.HandlerArgs) (*listener.HandlerResult, error) {
		c, err := m.CreateTryContainer(&container.TryContainerConfig{
			Args: args.InspectionArgs,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create container: %w", err)
		}

		if err := c.Run(10 * time.Second); err != nil {
			return nil, fmt.Errorf("failed to run container: %w", err)
		}

		stdout, stderr, err := c.Logs()
		if err != nil {
			return nil, fmt.Errorf("failed to get logs from container: %w", err)
		}

		if err := m.RemoveContainer(c); err != nil {
			return nil, fmt.Errorf("failed to remove container: %w", err)
		}

		return &listener.HandlerResult{
			RequestId: args.RequestId,
			InspectionResult: container.InspectionResult{
				Url:    args.InspectionArgs.Url,
				Stdout: string(stdout),
				Stderr: string(stderr),
			},
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

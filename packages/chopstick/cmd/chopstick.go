package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/hanchchch/gimi/packages/chopstick/pkg/container"
	"github.com/hanchchch/gimi/packages/chopstick/pkg/listener"
)

func main() {
	m := container.NewManager()

	listeners, err := listener.NewListeners(listener.ListenerOptions{
		HTTP: listener.HTTPListenerOptions{
			Addr: ":8080",
		},
		Redis: listener.RedisListenerOptions{
			Url: "redis://localhost:6379",
		},
	})

	if err != nil {
		panic(err)
	}

	handler := func(args listener.HandlerArgs) (map[string]string, error) {
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

		return map[string]string{"request_id": args.RequestId, "stdout": string(stdout), "stderr": string(stderr)}, nil
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

package main

import (
	"fmt"
	"time"

	"github.com/hanchchch/gimi/packages/chopstick/pkg/container"
	"github.com/hanchchch/gimi/packages/chopstick/pkg/listener"
)

func main() {
	m := container.NewManager()

	l := listener.NewListener(listener.ListenerOptions{
		HTTP: listener.HTTPListenerOptions{
			Addr: ":8080",
		},
	})

	l.OnInvoke(func(args container.TryArgs) (interface{}, error) {
		c, err := m.CreateTryContainer(&container.TryContainerConfig{
			Args: args,
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

		return map[string]string{"stdout": string(stdout), "stderr": string(stderr)}, nil
	})

	l.Listen()
}

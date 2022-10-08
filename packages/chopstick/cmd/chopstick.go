package main

import (
	"time"

	"github.com/hanchchch/gimi/packages/chopstick/pkg/container"
	"github.com/hanchchch/gimi/packages/chopstick/pkg/listener"
)

func main() {
	l := listener.NewListener(listener.ListenerOptions{
		HTTP: listener.HTTPListenerOptions{
			Addr: ":8080",
		},
	})

	l.OnInvoke(func(args container.TryArgs) (interface{}, error) {
		m := container.NewManager()
		c := m.CreateTryContainer(&container.TryContainerConfig{
			Args: args,
		})

		if err := c.Run(10 * time.Second); err != nil {
			return nil, err
		}

		stdout, stderr, err := c.Logs()
		if err != nil {
			return nil, err
		}

		if err := m.RemoveContainer(c); err != nil {
			return nil, err
		}

		return map[string]string{"stdout": string(stdout), "stderr": string(stderr)}, nil
	})

	l.Listen()
}

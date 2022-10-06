package main

import (
	"io"
	"os"
	"time"

	dc "github.com/docker/docker/api/types/container"
	"github.com/hanchchch/gimi/packages/chopstick/pkg/container"
)

func main() {
	m := container.NewManager()
	c := m.CreateContainer(&dc.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
	})

	if err := c.Run(10 * time.Second); err != nil {
		panic(err)
	}

	out, err := c.Logs()
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)

	if err := m.RemoveContainer(c); err != nil {
		panic(err)
	}
}

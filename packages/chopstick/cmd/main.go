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

	c.Start()
	_, err := c.Wait(10 * time.Second)
	if err != nil {
		panic(err)
	}

	out, err := c.Log()
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)

	err = m.RemoveContainer(c)
	if err != nil {
		panic(err)
	}
}

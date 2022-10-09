package container

import (
	"strings"
	"testing"
	"time"
)

func contains(s []string, substr string) bool {
	for _, v := range s {
		if v == substr {
			return true
		}
	}
	return false
}

func TestContainerCreateRemove(t *testing.T) {
	m := NewManager()
	c1, _ := m.CreateContainer(&ContainerConfig{
		Image: "alpine",
	})
	c2, _ := m.CreateContainer(&ContainerConfig{
		Image: "alpine",
	})

	if c1.ID == "" || c2.ID == "" {
		t.Fatal("container id is empty")
	}

	if c1.manager == nil || c2.manager == nil {
		t.Fatal("container manager is nil")
	}

	if strings.Compare(m.containers[c1.ID].ID, c1.ID) != 0 || strings.Compare(m.containers[c2.ID].ID, c2.ID) != 0 {
		t.Fatalf("map containers of manager is not correct\nexpected %v, got %v\nexpected %v, got %v", c1.ID, m.containers[c1.ID].ID, c2.ID, m.containers[c2.ID].ID)
	}

	<-time.After(1 * time.Second)

	containers, err := m.ListContainers()
	if err != nil {
		t.Fatal(err)
	}

	if len(containers) != 2 {
		t.Fatal("expected 2 containers, got", len(containers))
	}

	var cIds []string
	for _, c := range containers {
		cIds = append(cIds, c.ID)
	}

	if !contains(cIds, c1.ID) {
		t.Fatal("expected container", c1.ID, "not found")
	}

	if !contains(cIds, c2.ID) {
		t.Fatal("expected container", c2.ID, "not found")
	}

	if err := m.RemoveContainer(c1); err != nil {
		t.Fatal(err)
	}

	if err := m.RemoveContainer(c2); err != nil {
		panic(err)
	}

	containersAfterRemove, err := m.ListContainers()
	if err != nil {
		t.Fatal(err)
	}

	if len(containersAfterRemove) != 0 {
		t.Fatal("expected 0 containers, got", len(containersAfterRemove))
	}
}

func TestContainerRun(t *testing.T) {
	m := NewManager()
	c, _ := m.CreateContainer(&ContainerConfig{
		Image: "alpine",
	})

	if err := c.Run(10 * time.Second); err != nil {
		t.Fatal(err)
	}

	if err := m.RemoveContainer(c); err != nil {
		panic(err)
	}
}

func TestContainerLog(t *testing.T) {
	m := NewManager()
	c, _ := m.CreateContainer(&ContainerConfig{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
	})

	if err := c.Run(10 * time.Second); err != nil {
		t.Fatal(err)
	}

	_, _, err := c.Logs()
	if err != nil {
		t.Fatal(err)
	}

	if err := m.RemoveContainer(c); err != nil {
		panic(err)
	}
}

package container

import "time"

type Container interface {
	GetID() string
	Run(timeout time.Duration) error
	Logs() ([]byte, []byte, error)
	Remove() error
}

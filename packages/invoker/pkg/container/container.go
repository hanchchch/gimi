package container

import (
	"time"

	pb "github.com/hanchchch/gimi/packages/proto/go/messages"
)

type Container interface {
	GetID() string
	Run(timeout time.Duration) error
	Logs() ([]byte, []byte, error)
	GetResult() (*pb.InspectionResult, error)
	Remove() error
}

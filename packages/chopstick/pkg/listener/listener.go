package listener

import (
	"github.com/hanchchch/gimi/packages/chopstick/pkg/container"
)

type HandlerArgs struct {
	RequestId      string                   `json:"request_id"`
	InspectionArgs container.InspectionArgs `json:"inspection_args"`
}

type HandlerFunc func(HandlerArgs) (map[string]string, error)

type Listener interface {
	Listen() error
	OnInvoke(callback HandlerFunc) error
}

type ListenerOptions struct {
	HTTP  HTTPListenerOptions
	Redis RedisListenerOptions
}

func NewListeners(options ListenerOptions) ([]Listener, error) {
	listeners := make([]Listener, 0, 2)
	if options.HTTP.Addr != "" {
		l := NewHTTPListener(options.HTTP)
		listeners = append(listeners, l)
	}
	if options.Redis.Url != "" {
		l, err := NewRedisListener(options.Redis)
		if err != nil {
			return nil, err
		}
		listeners = append(listeners, l)
	}
	return listeners, nil
}

package listener

type HandlerFunc func() (interface{}, error)

type Listener interface {
	Listen() error
	OnInvoke(callback HandlerFunc) error
}

type ListenerOptions struct {
	HTTP HTTPListenerOptions
}

func NewListener(options ListenerOptions) Listener {
	if options.HTTP.Addr != "" {
		return NewHTTPListener(options.HTTP)
	}
	return nil
}

package listener

type RedisListener struct {
}

func NewRedisListener() *RedisListener {
	return &RedisListener{}
}

func (l *RedisListener) Listen() error {
	return nil
}

func (l *RedisListener) OnInvoke(callback func() error) error {
	return nil
}

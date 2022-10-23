package listener

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const keyPrefix = "gimi:inspection"

type RedisListener struct {
	id        string
	url       string
	callback  HandlerFunc
	redis     *redis.Client
	keyPrefix string
}

type RedisListenerOptions struct {
	Url string
}

func NewRedisListener(options RedisListenerOptions) (*RedisListener, error) {
	opt, err := redis.ParseURL(options.Url)
	if err != nil {
		return nil, err
	}

	return &RedisListener{
		id:        uuid.New().String(),
		url:       options.Url,
		redis:     redis.NewClient(opt),
		keyPrefix: keyPrefix,
	}, nil
}

func (l *RedisListener) buildKey(keys ...string) string {
	return strings.Join(append([]string{l.keyPrefix}, keys...), ":")
}

func (l *RedisListener) ListenerKey() string {
	return l.buildKey("listener", runtime.GOOS, l.id)
}

func (l *RedisListener) QueueKey() string {
	return l.buildKey("request", runtime.GOOS)
}

func (l *RedisListener) ResultKey(requestId string) string {
	return l.buildKey("results", requestId)
}

func (l *RedisListener) KeepListenerKey() error {
	for {
		l.redis.SetEX(context.Background(), l.ListenerKey(), "OK", 3*time.Second)
		time.Sleep(1 * time.Second)
	}
}

func (l *RedisListener) Listen() error {
	log.Println("Listening to", l.url)
	go l.KeepListenerKey()
	for {
		ctx := context.Background()
		cmd := l.redis.BRPop(ctx, 0, l.QueueKey())
		if cmd.Err() != nil {
			log.Printf("error while brpop: %v", cmd.Err())
			continue
		}
		body := cmd.Val()[1]

		var args HandlerArgs
		if err := json.NewDecoder(bytes.NewBufferString(body)).Decode(&args); err != nil {
			log.Printf("error while decoding request: %v", err)
			continue
		}

		data, err := l.callback(args)
		if err != nil {
			log.Printf("error while processing request: %v", err)
			continue
		}

		resp, _ := json.Marshal(data)
		_, err = l.redis.Set(ctx, l.ResultKey(args.RequestId), resp, 3*time.Hour).Result()
		if err != nil {
			log.Printf("error while lpush: %v", err)
			continue
		}
	}
}

func (l *RedisListener) OnInvoke(callback HandlerFunc) error {
	l.callback = callback
	return nil
}

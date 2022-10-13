package listener

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisListener struct {
	url      string
	callback HandlerFunc
	redis    *redis.Client
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
		url:   options.Url,
		redis: redis.NewClient(opt),
	}, nil
}

func (l *RedisListener) Listen() error {
	log.Println("Listening to", l.url)
	for {
		ctx := context.Background()
		cmd := l.redis.BRPop(ctx, 0, "gimi:inspection:queue")
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
		_, err = l.redis.LPush(ctx, "gimi:inspection:results", resp).Result()
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

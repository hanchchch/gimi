package listener

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	pb "github.com/hanchchch/gimi/packages/proto/go/messages"
	"google.golang.org/protobuf/proto"
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

func (l *RedisListener) ResultPubKey(requestId string) string {
	return l.buildKey("results", "pub", requestId)
}

func (l *RedisListener) KeepListenerKey() error {
	for {
		l.redis.SetEX(context.Background(), l.ListenerKey(), "OK", 3*time.Second)
		time.Sleep(1 * time.Second)
	}
}

func (l *RedisListener) Process(ctx context.Context, args *pb.HandlerArgs) error {
	data, err := l.callback(args)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		estr := err.Error()
		data = &pb.HandlerResult{
			RequestId: args.RequestId,
			Error:     &estr,
		}
	}

	resp, err := proto.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}

	_, err = l.redis.Set(ctx, l.ResultKey(args.RequestId), resp, 3*time.Hour).Result()
	if err != nil {
		return fmt.Errorf("failed to lpush: %v", err)
	}

	return nil
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

		args := &pb.HandlerArgs{}
		if err := proto.Unmarshal([]byte(body), args); err != nil {
			log.Printf("failed to parse request: %v", err)
			continue
		}

		go func() {
			if err := l.Process(ctx, args); err != nil {
				log.Printf("error while processing request: %v", err)
			}
		}()
	}
}

func (l *RedisListener) OnInvoke(callback HandlerFunc) error {
	l.callback = callback
	return nil
}

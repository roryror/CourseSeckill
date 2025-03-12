package interfaces

import (
	"context"
	"time"
)

type Database interface {
	Create(ctx context.Context, model interface{}, value ...interface{}) error
	Find(ctx context.Context, model interface{}, query string, args ...interface{}) error
	Update(ctx context.Context, model interface{}, updates map[string]interface{}, query string, args ...interface{}) error
	Delete(ctx context.Context, model interface{}, query string, args ...interface{}) error

	Transaction(ctx context.Context, fn func(ctx context.Context, client Database, args ...interface{}) error, rb func(ctx context.Context, client Database, err error, args ...interface{}) error, args ...interface{}) error

	NotFoundError() error
	Expr(ctx context.Context, expr string) interface{}
	Close() error
}

type Cache interface {
	// Hash functions
	HSet(ctx context.Context, hashName string, key string, value interface{}) error
	HGet(ctx context.Context, hashName string, key string) (interface{}, error)
	HIncrBy(ctx context.Context, hashName string, key string, incr int64) (int64, error)
	
	// Set function
	SAdd(ctx context.Context, setName string, values ...interface{}) error
	SIsMember(ctx context.Context, setName string, value interface{}) (bool, error)
	
	Expire(ctx context.Context, object string, expiration time.Duration) error
	
	// Function to handle complex logic
	RunScript(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error)
	
	Del(ctx context.Context, object string) error

	Close() error
}

type MsgQueue interface {
	Write(ctx context.Context, msg []interface{}) error
	Read(ctx context.Context, rid int) (interface{}, error)
	GroupSize() int
	Close() error
}

type Channel interface {
	Ch() chan interface{}
	BatchSize() int
	Duration() int

	Close() error
}


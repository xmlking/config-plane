package config

import "context"


type Client interface {
	Subscription(pattern string) Subscription
	Get(key string) (val string, err error)
	GetAll(pattern string) (changelog []string, err error)
	Put(key string, val string) (err error)
}

type Subscription interface {
	Receive(ctx context.Context, f func(ctx context.Context, cSet []string)) (err error)
}
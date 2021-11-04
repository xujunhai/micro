package redis

import "github.com/go-redis/redis/v8"

type trace struct {
	withTrace bool
}

type Options struct {
	*redis.Options
	trace
}

type ClusterOptions struct {
	*redis.ClusterOptions
	trace
}

type FailoverOptions struct {
	*redis.FailoverOptions
	trace
}

var traceHook = &TraceHook{}

// NewClient returns a client to the Redis Server specified by Options.
func NewClient(opt *Options) *redis.Client {
	c := redis.NewClient(opt.Options)
	if opt.withTrace {
		c.AddHook(traceHook)
	}
	return c
}

// NewClient returns a client to the Redis Server specified by Options.
func NewClusterClient(opt *ClusterOptions) *redis.ClusterClient {
	c := redis.NewClusterClient(opt.ClusterOptions)
	if opt.withTrace {
		c.AddHook(traceHook)
	}
	return c
}

func NewSentinelClient(opt *Options) *redis.SentinelClient {
	c := redis.NewSentinelClient(opt.Options)
	if opt.withTrace {
		c.AddHook(traceHook)
	}
	return c
}

func NewFailoverClient(opt *FailoverOptions) *redis.Client {
	c := redis.NewFailoverClient(opt.FailoverOptions)
	if opt.withTrace {
		c.AddHook(traceHook)
	}
	return c
}

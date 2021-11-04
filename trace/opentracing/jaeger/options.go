package jaeger

type Options struct {
	env     bool
	service string
	agent   string // host:port or send spans to collector at this URL
}

type Option func(*Options)

func Service(service string) Option {
	return func(options *Options) {
		options.service = service
	}
}

func Agent(agent string) Option {
	return func(options *Options) {
		options.agent = agent
	}
}

func Env(e bool) Option {
	return func(options *Options) {
		options.env = e
	}
}

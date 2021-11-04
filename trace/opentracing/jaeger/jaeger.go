package jaeger

import (
	"context"
	"fmt"
	"github.com/uber/jaeger-client-go/config"
	"gitlab.ziroom.com/rent-web/micro/trace"
	"io"
	"net/url"
	"time"
)

type jaegerTrace struct {
}

func NewJaegerTrace() *jaegerTrace {
	return &jaegerTrace{}
}

// init service root tracer, using in main goroutine
func (j *jaegerTrace) Init(opts ...Option) io.Closer {
	var options Options
	for _, v := range opts {
		v(&options)
	}

	var cfg *config.Configuration
	if options.env {
		var err error
		if cfg, err = config.FromEnv(); err != nil {
			panic(fmt.Sprintf("InitJaeger Error: %s", err.Error()))
		}
	} else {
		cfg = &config.Configuration{
			ServiceName: options.service,
			Sampler: &config.SamplerConfig{
				Type:  "const",
				Param: 1,
			},
			Reporter: &config.ReporterConfig{
				LogSpans: true,
			},
		}
		endpoint, err := url.ParseRequestURI(options.agent)
		if err == nil {
			cfg.Reporter.CollectorEndpoint = endpoint.String()
		} else {
			cfg.Reporter.LocalAgentHostPort = options.agent
		}
	}
	loggerOpt := config.Logger(JaegerLogger)
	closer, err := cfg.InitGlobalTracer(options.service, loggerOpt)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	return closer
}

// Start a trace
func (j *jaegerTrace) Start(ctx context.Context, name string) (context.Context, *trace.Span) {
	return ctx, &trace.Span{
		Trace:    "",
		Name:     name,
		Id:       "",
		Parent:   "",
		Started:  time.Time{},
		Duration: 0,
		Metadata: nil,
		Type:     0,
	}
}

// Finish the trace
func (j *jaegerTrace) Finish(*trace.Span) error {
	return nil
}

// Read the traces
func (j *jaegerTrace) Read(...*trace.ReadOption) ([]*trace.Span, error) {
	return nil, nil
}

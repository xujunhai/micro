package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"strings"
)

type TraceHook struct {
}

func (t TraceHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	spanName := strings.ToUpper(cmd.Name())
	span, _ := opentracing.StartSpanFromContext(ctx, spanName)
	ext.DBType.Set(span, "redis")
	ext.DBStatement.Set(span, cmd.String())
	ctx = opentracing.ContextWithSpan(ctx, span)
	return ctx, nil
}

func (t TraceHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span := opentracing.SpanFromContext(ctx)
	span.Finish()
	return nil
}

func (t TraceHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	pipelineSpan, ctx := opentracing.StartSpanFromContext(ctx, "(pipeline)")

	ext.DBType.Set(pipelineSpan, "redis")

	for i := len(cmds); i > 0; i-- {
		cmdName := strings.ToUpper(cmds[i-1].Name())
		if cmdName == "" {
			cmdName = "(empty command)"
		}

		span, _ := opentracing.StartSpanFromContext(ctx, cmdName)
		ext.DBType.Set(span, "redis")
		ext.DBStatement.Set(span, fmt.Sprintf("%v", cmds[i-1].Args()))
		defer span.Finish() //create heap{funval} so caution!!
	}
	//defer pipelineSpan.Finish()
	ctx = opentracing.ContextWithSpan(ctx, pipelineSpan)
	return ctx, nil
}

func (t TraceHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	pipelineSpan := opentracing.SpanFromContext(ctx)
	pipelineSpan.Finish()
	return nil
}

func process(ctx context.Context) func(oldProcess func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
	return func(oldProcess func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
		return func(cmd redis.Cmder) error {
			spanName := strings.ToUpper(cmd.Name())
			span, _ := opentracing.StartSpanFromContext(ctx, spanName)
			ext.DBType.Set(span, "redis")
			ext.DBStatement.Set(span, fmt.Sprintf("%v", cmd.Args()))
			defer span.Finish()

			return oldProcess(cmd)
		}
	}
}

func processPipeline(ctx context.Context) func(oldProcess func(cmds []redis.Cmder) error) func(cmds []redis.Cmder) error {
	return func(oldProcess func(cmds []redis.Cmder) error) func(cmds []redis.Cmder) error {
		return func(cmds []redis.Cmder) error {
			pipelineSpan, ctx := opentracing.StartSpanFromContext(ctx, "(pipeline)")

			ext.DBType.Set(pipelineSpan, "redis")

			for i := len(cmds); i > 0; i-- {
				cmdName := strings.ToUpper(cmds[i-1].Name())
				if cmdName == "" {
					cmdName = "(empty command)"
				}

				span, _ := opentracing.StartSpanFromContext(ctx, cmdName)
				ext.DBType.Set(span, "redis")
				ext.DBStatement.Set(span, fmt.Sprintf("%v", cmds[i-1].Args()))
				defer span.Finish() //create heap{funval} so caution!!
			}

			defer pipelineSpan.Finish()
			return oldProcess(cmds)
		}
	}
}

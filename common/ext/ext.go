package ext

//use to package init
import (
	_ "xmicro/config/source/nacos"
	_ "xmicro/registry/nacos"
	_ "xmicro/trace/opentracing/jaeger"
)

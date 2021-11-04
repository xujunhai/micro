package ext

//use to package init
import (
	_ "gitlab.ziroom.com/rent-web/micro/config/source/nacos"
	_ "gitlab.ziroom.com/rent-web/micro/registry/nacos"
	_ "gitlab.ziroom.com/rent-web/micro/trace/opentracing/jaeger"
)

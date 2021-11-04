package configuration

import (
	"fmt"
	"xmicro/client"
	"xmicro/common/component"
	"xmicro/server"

	"os"
	"strings"
	"sync"
)

import (
	"xmicro/common/constant"
	"xmicro/logger"
	"xmicro/logger/core"
	"xmicro/service"
	"xmicro/service/rpc"
	"xmicro/trace/opentracing/jaeger"
	"xmicro/transport/grpc"
)

var (
	clientConfig *ClientConfig
	serverConfig *ServerConfig
	// baseConfig = providerConfig.BaseConfig or consumerConfig
	baseConfig *BaseConfig
	sslEnabled = false

	// configAccessMutex is used to make sure that xxxxConfig will only be created once if needed.
	// it should be used combine with double-check to avoid the race condition
	configAccessMutex sync.Mutex
)

//Load from config file build the service instance
func InitServer(confFiles ...string) service.Service {
	var confFile string
	if len(confFiles) == 0 {
		confFile = os.Getenv(constant.ConfServerFilePath)
	} else {
		confFile = confFiles[0]
	}

	if confFile == "" {
		panic("application configure(server) file name is nil")
	}

	if errPro := ServerInit(confFile); errPro != nil {
		logger.Errorf("[serverInit] %#v", errPro)
		serverConfig = nil
		panic("application configure(server) serverInit err" + errPro.Error())
	} else {
		// Even though baseConfig has been initialized, we override it
		// because we think read from config file is correct config
		baseConfig = &serverConfig.BaseConfig
	}

	//logger init
	if serverConfig.LoggerConfig != nil {
		//normalize logFile=logDir+serviceName
		logFile := serverConfig.LoggerConfig.LogDir + string(os.PathSeparator) + serverConfig.ApplicationConfig.Name
		loggerNew, _ := logger.NewZapLogger(
			core.WithLogFile(logFile),
			core.WithLevel(core.GetLevel(serverConfig.LoggerConfig.Level)),
		)
		logger.SetLogger(loggerNew)
	}

	options := []service.Option{
		service.Name(serverConfig.ApplicationConfig.Name),
		service.Version(serverConfig.ApplicationConfig.Version),
		//tcp bind address to, private > public > loopback
		service.Address(fmt.Sprintf("0.0.0.0:%d", serverConfig.Port)),
		service.Transport(grpc.NewTransport()),
	}

	//trace init
	if serverConfig.TraceConfig != nil {
		closer := jaeger.NewJaegerTrace().Init(
			jaeger.Service(serverConfig.ApplicationConfig.Name),
			jaeger.Env(serverConfig.TraceConfig.FromEnv),
			jaeger.Agent(serverConfig.TraceConfig.Agent),
		)
		//closure io.closer
		b := service.BeforeStop(func() error {
			closer.Close()
			return nil
		})
		options = append(options, b)
	}

	//create service instance
	registryBuilder := &registryBuilder{}
	registry, err := registryBuilder.buildRegistry(serverConfig.BaseConfig)
	if err != nil {
		panic("application configure(server) registryBuilder err" + err.Error())
	}
	options = append(options, service.Registry(registry))

	//start config center
	if err := serverConfig.startConfigCenter(serverConfig.BaseConfig); err != nil {
		panic("application configure(server) startConfigCenter err" + err.Error())
	}

	//wrapper handler load,all request used for
	if serverConfig.BaseConfig.Wrapper != "" {
		var handlerWrappers []server.HandlerWrapper
		wrappers := strings.Split(serverConfig.BaseConfig.Wrapper, ",")
		for _, v := range wrappers {
			if h := component.GetServerWrapper(v); h != nil {
				handlerWrappers = append(handlerWrappers, h)
			}
		}
		if len(handlerWrappers) > 0 {
			options = append(options, service.WrapHandler(handlerWrappers...))
		}
	}

	return rpc.NewApp(
		options...,
	)
}

func InitClient(confFiles ...string) service.Service {
	var confFile string
	if len(confFiles) == 0 {
		confFile = os.Getenv(constant.ConfClientFilePath)
	} else {
		confFile = confFiles[0]
	}

	if confFile == "" {
		panic("application configure(client) file name is nil")
	}

	if errPro := ClientInit(confFile); errPro != nil {
		logger.Errorf("[serverInit] %#v", errPro)
		clientConfig = nil
		panic("application configure(server) serverInit err" + errPro.Error())
	} else {
		// Even though baseConfig has been initialized, we override it
		// because we think read from config file is correct config
		baseConfig = &clientConfig.BaseConfig
	}

	//logger init
	if clientConfig.LoggerConfig != nil {
		//normalize logFile=logDir+serviceName
		logFile := clientConfig.LoggerConfig.LogDir + string(os.PathSeparator) + clientConfig.ApplicationConfig.Name
		loggerNew, _ := logger.NewZapLogger(
			core.WithLogFile(logFile),
			core.WithLevel(core.GetLevel(clientConfig.LoggerConfig.Level)),
		)
		logger.SetLogger(loggerNew)
	}

	options := []service.Option{
		service.Name(clientConfig.ApplicationConfig.Name),
		service.Version(clientConfig.ApplicationConfig.Version),
		service.Transport(grpc.NewTransport()),
	}

	//trace init
	if clientConfig.TraceConfig != nil {
		closer := jaeger.NewJaegerTrace().Init(
			jaeger.Service(clientConfig.ApplicationConfig.Name),
			jaeger.Env(clientConfig.TraceConfig.FromEnv),
			jaeger.Agent(clientConfig.TraceConfig.Agent),
		)
		//closure io.closer
		b := service.BeforeStop(func() error {
			closer.Close()
			return nil
		})
		options = append(options, b)
	}

	//create service instance
	registryBuilder := &registryBuilder{}
	registry, err := registryBuilder.buildRegistry(clientConfig.BaseConfig)
	if err != nil {
		panic("application configure(server) registryBuilder err" + err.Error())
	}
	options = append(options, service.Registry(registry))

	//start config center
	if err := clientConfig.startConfigCenter(clientConfig.BaseConfig); err != nil {
		panic("application configure(server) startConfigCenter err" + err.Error())
	}

	//wrapper handler load,all request used for
	if clientConfig.BaseConfig.Wrapper != "" {
		var clientWrappers []client.Wrapper
		wrappers := strings.Split(clientConfig.BaseConfig.Wrapper, ",")
		for _, v := range wrappers {
			if h := component.GetClientWrapper(v); h != nil {
				clientWrappers = append(clientWrappers, h)
			}
		}
		if len(clientWrappers) > 0 {
			options = append(options, service.WrapClient(clientWrappers...))
		}
	}

	return rpc.NewApp(
		options...,
	)
}

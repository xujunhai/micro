package component

import (
	"xmicro/client"
	"xmicro/server"
)

//wrapper
var serverWrapper = make(map[string]server.HandlerWrapper)

func SetServerWrapper(com string, wrapper server.HandlerWrapper) {
	serverWrapper[com] = wrapper
}

//get
func GetServerWrapper(com string) server.HandlerWrapper {
	if handler, ok := serverWrapper[com]; ok {
		return handler
	}
	return nil
}

//wrapper
var clientWrapper = make(map[string]client.Wrapper)

func SetClientWrapper(com string, wrapper client.Wrapper) {
	clientWrapper[com] = wrapper
}

//get
func GetClientWrapper(com string) client.Wrapper {
	if handler, ok := clientWrapper[com]; ok {
		return handler
	}
	return nil
}

package component

import (
	"gitlab.ziroom.com/rent-web/micro/service"
)

var serviceInstance = make(map[string]service.Factory)

func SetServiceFactory(com string, factory service.Factory) {
	serviceInstance[com] = factory
}

func GetServiceFactory(com string) service.Factory {
	if serviceInstance[com] == nil {
		panic("serviceInstance for " + com + " is not existing, make sure you have import the package.")
	}
	return serviceInstance[com]
}

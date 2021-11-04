package component

import (
	"xmicro/config"
	"xmicro/registry"
)

// component manager
var configCenterFactories = make(map[string]config.DynamicConfigurationFactory)

func SetConfigCenterFactory(com string, factory config.DynamicConfigurationFactory) {
	configCenterFactories[com] = factory
}

func GetConfigCenterFactory(com string) config.DynamicConfigurationFactory {
	if configCenterFactories[com] == nil {
		panic("config center for " + com + " is not existing, make sure you have import the package.")
	}
	return configCenterFactories[com]
}

var registryFactories = make(map[string]registry.RegistryFactory)

func SetRegistryFactory(com string, factory registry.RegistryFactory) {
	registryFactories[com] = factory
}

func GetRegistryFactory(com string) registry.RegistryFactory {
	if registryFactories[com] == nil {
		panic("RegistryFactory for " + com + " is not existing, make sure you have import the package.")
	}
	return registryFactories[com]
}

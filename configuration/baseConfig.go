package configuration

import "bytes"

// BaseConfig is the common configuration for provider and consumer
type BaseConfig struct {
	// application config
	ApplicationConfig  *ApplicationConfig  `yaml:"application" json:"application,omitempty" property:"application"`
	ConfigCenterConfig *ConfigCenterConfig `yaml:"configCenter" json:"configCenter,omitempty"`
	RegistryConfig     *RegistryConfig     `yaml:"registry" json:"registry,omitempty"`
	LoggerConfig       *LoggerConfig       `yaml:"logger" json:"logger,omitempty"`
	TraceConfig        *TraceConfig        `yaml:"trace" json:"trace,omitempty"`

	//wrapper
	Wrapper string `yaml:"wrapper" json:"wrapper,omitempty" property:"wrapper"`
	//configuration yaml file buffer
	fileStream *bytes.Buffer
	//config center
	configCenter
}

package configuration

import (
	"bytes"
	"github.com/creasty/defaults"
	perrors "github.com/pkg/errors"
	"xmicro/common/constant"
	"xmicro/util/yaml"
)

// ServerConfig is the default configuration of service provider
type ServerConfig struct {
	BaseConfig `yaml:",inline"`
	Port       int         `default:"8090" yaml:"port" json:"port,omitempty" property:"port"`
	Consumer   interface{} `yaml:"consumer"`
}

// UnmarshalYAML unmarshals the ServerConfig by @unmarshal function
func (c *ServerConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := defaults.Set(c); err != nil {
		return err
	}
	type plain ServerConfig
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}
	return nil
}

// nolint
func (*ServerConfig) Prefix() string {
	return constant.ServerConfigPrefix
}

// SetServerConfig sets provider config by @p
func SetServerConfig(p ServerConfig) {
	serverConfig = &p
}

// ServerInit loads config file to init provider config
func ServerInit(confProFile string) error {
	if len(confProFile) == 0 {
		return perrors.Errorf("application configure(provider) file name is nil")
	}
	serverConfig = &ServerConfig{}
	fileStream, err := yaml.UnmarshalYMLConfig(confProFile, serverConfig)
	if err != nil {
		return perrors.Errorf("unmarshalYmlConfig error %v", perrors.WithStack(err))
	}

	serverConfig.fileStream = bytes.NewBuffer(fileStream)
	return nil
}

func configCenterRefreshServer() error {
	// fresh it
	return nil
}

package configuration

import (
	"bytes"
	"github.com/creasty/defaults"
	perrors "github.com/pkg/errors"
	"time"
)

import (
	"gitlab.ziroom.com/rent-web/micro/common/constant"
	"gitlab.ziroom.com/rent-web/micro/logger"
	"gitlab.ziroom.com/rent-web/micro/util/yaml"
)

const (
	MaxWheelTimeSpan = 900e9 // 900s, 15 minute
)

// ClientConfig is Consumer default configuration
type ClientConfig struct {
	BaseConfig `yaml:",inline"`
	// client
	ConnectTimeoutStr string `default:"100ms"  yaml:"connectTimeout" json:"connectTimeout,omitempty" property:"connectTimeout"`
	ConnectTimeout    time.Duration

	RequestTimeoutStr string `yaml:"requestTimeout" default:"5s" json:"requestTimeout,omitempty" property:"requestTimeout"`
	RequestTimeout    time.Duration

	//async communicate
	Producer interface{} `yaml:"producer"`
}

// UnmarshalYAML unmarshals the ClientConfig by @unmarshal function
func (c *ClientConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := defaults.Set(c); err != nil {
		return err
	}
	type plain ClientConfig
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}
	return nil
}

// nolint
func (*ClientConfig) Prefix() string {
	return constant.ClientConfigPrefix
}

// SetClientConfig sets clientConfig by @c
func SetClientConfig(c ClientConfig) {
	clientConfig = &c
}

// ClientInit loads config file to init consumer config
func ClientInit(confConFile string) error {
	if confConFile == "" {
		return perrors.Errorf("application configure(consumer) file name is nil")
	}
	clientConfig = &ClientConfig{}
	fileStream, err := yaml.UnmarshalYMLConfig(confConFile, clientConfig)
	if err != nil {
		return perrors.Errorf("unmarshalYmlConfig error %v", perrors.WithStack(err))
	}
	clientConfig.fileStream = bytes.NewBuffer(fileStream)

	if clientConfig.RequestTimeoutStr != "" {
		if clientConfig.RequestTimeout, err = time.ParseDuration(clientConfig.RequestTimeoutStr); err != nil {
			return perrors.WithMessagef(err, "time.ParseDuration(Request_Timeout{%#v})", clientConfig.RequestTimeoutStr)
		}
		if clientConfig.RequestTimeout >= time.Duration(MaxWheelTimeSpan) {
			return perrors.WithMessagef(err, "request_timeout %s should be less than %s",
				clientConfig.RequestTimeoutStr, time.Duration(MaxWheelTimeSpan))
		}
	}
	if clientConfig.ConnectTimeoutStr != "" {
		if clientConfig.ConnectTimeout, err = time.ParseDuration(clientConfig.ConnectTimeoutStr); err != nil {
			return perrors.WithMessagef(err, "time.ParseDuration(Connect_Timeout{%#v})", clientConfig.ConnectTimeoutStr)
		}
	}

	logger.Debugf("consumer config{%#v}\n", clientConfig)
	return nil
}

func configCenterRefreshConsumer() error {
	//fresh it
	return nil
}

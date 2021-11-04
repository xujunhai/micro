package configuration

import (
	"github.com/creasty/defaults"
	"xmicro/common/constant"
)

// Only support one service application, so need only one port to export
// ApplicationConfig is a configuration for current application, whether the application is a provider or a consumer
type ApplicationConfig struct {
	Organization string `yaml:"organization" json:"organization,omitempty" property:"organization"`
	Name         string `yaml:"name" json:"name,omitempty" property:"name"`
	Module       string `yaml:"module" json:"module,omitempty" property:"module"`
	Version      string `yaml:"version" json:"version,omitempty" property:"version"`
	Owner        string `yaml:"owner" json:"owner,omitempty" property:"owner"`
	Environment  string `yaml:"env" json:"env,omitempty" property:"env"`
	// the metadata type. remote or local
	MetadataType string `default:"local" yaml:"metadataType" json:"metadataType,omitempty" property:"metadataType"`
}

// nolint
func (*ApplicationConfig) Prefix() string {
	return constant.Micro + ".application."
}

// UnmarshalYAML unmarshals the ApplicationConfig by @unmarshal function
func (c *ApplicationConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := defaults.Set(c); err != nil {
		return err
	}
	type plain ApplicationConfig
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}
	return nil
}

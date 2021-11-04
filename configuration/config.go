/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package configuration

import (
	"context"
	"github.com/creasty/defaults"
	perrors "github.com/pkg/errors"
	"net/url"
	"time"
)

import (
	"gitlab.ziroom.com/rent-web/micro/common"
	"gitlab.ziroom.com/rent-web/micro/common/component"
	"gitlab.ziroom.com/rent-web/micro/common/constant"
	"gitlab.ziroom.com/rent-web/micro/logger"
)

// ConfigCenterConfig is configuration for config center
//
// ConfigCenter also introduced concepts of namespace and group to better manage Key-Value pairs by group,
// those configs are already built-in in many professional third-party configuration centers.
// In most cases, namespace is used to isolate different tenants, while group is used to divide the key set from one tenant into groups.
//
// ConfigCenter protocol has currently supported nacos,apollo
type ConfigCenterConfig struct {
	context       context.Context
	Protocol      string `required:"true"  yaml:"protocol"  json:"protocol,omitempty"`
	Address       string `yaml:"address" json:"address,omitempty"`
	Cluster       string `yaml:"cluster" json:"cluster,omitempty"`
	Group         string `default:"micro" yaml:"group" json:"group,omitempty"`
	Username      string `yaml:"username" json:"username,omitempty"`
	Password      string `yaml:"password" json:"password,omitempty"`
	LogDir        string `yaml:"logDir" json:"logDir,omitempty"`
	ConfigFile    string `default:"micro.properties" yaml:"configFile"  json:"configFile,omitempty"`
	Namespace     string `default:"micro" yaml:"namespace"  json:"namespace,omitempty"`
	AppConfigFile string `default:"micro.properties" yaml:"appConfigFile"  json:"appConfigFile,omitempty"`
	AppId         string `default:"micro" yaml:"appId"  json:"appId,omitempty"`
	TimeoutStr    string `yaml:"timeout"  json:"timeout,omitempty"`
	timeout       time.Duration
}

// UnmarshalYAML unmarshals the ConfigCenterConfig by @unmarshal function
func (c *ConfigCenterConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := defaults.Set(c); err != nil {
		return err
	}
	type plain ConfigCenterConfig
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}
	return nil
}

// GetUrlMap gets url map from ConfigCenterConfig
func (c *ConfigCenterConfig) GetUrlMap() url.Values {
	urlMap := url.Values{}
	urlMap.Set(constant.ConfigNamespaceKey, c.Namespace)
	urlMap.Set(constant.ConfigGroupKey, c.Group)
	urlMap.Set(constant.ConfigClusterKey, c.Cluster)
	urlMap.Set(constant.ConfigAppIdKey, c.AppId)
	urlMap.Set(constant.ConfigLogDirKey, c.LogDir)
	return urlMap
}

type configCenter struct {
}

// toURL will compatible with baseConfig.ConfigCenterConfig.Address and baseConfig.ConfigCenterConfig.RemoteRef before 1.6.0
// After 1.6.0 will not compatible, only baseConfig.ConfigCenterConfig.RemoteRef
func (b *configCenter) toURL(baseConfig BaseConfig) (*common.URL, error) {
	if len(baseConfig.ConfigCenterConfig.Address) > 0 {
		return common.NewURL(baseConfig.ConfigCenterConfig.Address,
			common.WithProtocol(baseConfig.ConfigCenterConfig.Protocol), common.WithParams(baseConfig.ConfigCenterConfig.GetUrlMap()))
	}
	return nil, perrors.New("baseConfig.ConfigCenterConfig.Address Empty")
}

// startConfigCenter will start the config center.
// it will prepare the environment
func (b *configCenter) startConfigCenter(baseConfig BaseConfig) error {
	newUrl, err := b.toURL(baseConfig)
	if err != nil {
		return err
	}

	//init config center factory from service locator
	factory := component.GetConfigCenterFactory(newUrl.Protocol)
	dynamicConfig, err := factory.GetDynamicConfiguration(newUrl)
	if err != nil {
		logger.Errorf("Get dynamic configuration error , error message is %v", err)
		return perrors.WithStack(err)
	}

	GetEnvInstance().SetDynamicConfiguration(dynamicConfig)

	//TODO get from dynamicConfig
	//content, err := dynamicConfig.GetProperties(baseConfig.ConfigCenterConfig.ConfigFile, config.WithGroup(baseConfig.ConfigCenterConfig.Group))
	//if err != nil {
	//	logger.Errorf("Get config content in dynamic configuration error , error message is %v", err)
	//	return perrors.WithStack(err)
	//}
	//
	//var appGroup string
	//var appContent string
	//
	//TODO whether the application is a provider or a consumer
	//appGroup = baseConfig.ApplicationConfig.Name
	//if len(appGroup) != 0 {
	//	configFile := baseConfig.ConfigCenterConfig.AppConfigFile
	//	if len(configFile) == 0 {
	//		configFile = baseConfig.ConfigCenterConfig.ConfigFile
	//	}
	//	appContent, err = dynamicConfig.GetProperties(configFile, config.WithGroup(appGroup))
	//	if err != nil {
	//		return perrors.WithStack(err)
	//	}
	//}
	//// global config file
	//mapContent, err := dynamicConfig.Parser().Parse(content)
	//if err != nil {
	//	return perrors.WithStack(err)
	//}
	//GetEnvInstance().UpdateExternalConfigMap(mapContent)
	//
	//// appGroup config file
	//if len(appContent) != 0 {
	//	appMapConent, err := dynamicConfig.Parser().Parse(appContent)
	//	if err != nil {
	//		return perrors.WithStack(err)
	//	}
	//	GetEnvInstance().UpdateAppExternalConfigMap(appMapConent)
	//}
	return nil
}

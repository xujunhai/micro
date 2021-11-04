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
	perrors "github.com/pkg/errors"

	"net/url"
	"strconv"
	"strings"
)

import (
	"github.com/creasty/defaults"
)

import (
	"xmicro/common"
	"xmicro/common/component"
	"xmicro/common/constant"
	"xmicro/logger"
	mregistry "xmicro/registry"
)

// RegistryConfig is the configuration of the registry center
type RegistryConfig struct {
	//option nacos/zookeeper/etcd/kubernetes
	Protocol   string `required:"true" yaml:"protocol"  json:"protocol,omitempty" property:"protocol"`
	TimeoutStr string `yaml:"timeout" default:"5s" json:"timeout,omitempty" property:"timeout"` // unit: second
	Group      string `yaml:"group" json:"group,omitempty" property:"group"`
	TTL        string `yaml:"ttl" default:"10m" json:"ttl,omitempty" property:"ttl"` // unit: minute
	// for registry
	Address    string `yaml:"address" json:"address,omitempty" property:"address"`
	Username   string `yaml:"username" json:"username,omitempty" property:"username"`
	Password   string `yaml:"password" json:"password,omitempty"  property:"password"`
	Simplified bool   `yaml:"simplified" json:"simplified,omitempty"  property:"simplified"`
	// Always use this registry first if set to true, useful when subscribe to multiple registries
	Preferred bool `yaml:"preferred" json:"preferred,omitempty" property:"preferred"`
	// The region where the registry belongs, usually used to isolate traffics
	Zone string `yaml:"zone" json:"zone,omitempty" property:"zone"`
	// Affects traffic distribution among registries,
	Params map[string]string `yaml:"params" json:"params,omitempty" property:"params"`
}

// UnmarshalYAML unmarshals the RegistryConfig by @unmarshal function
func (r *RegistryConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := defaults.Set(r); err != nil {
		return err
	}
	type plain RegistryConfig
	if err := unmarshal((*plain)(r)); err != nil {
		return err
	}
	return nil
}

// nolint
func (r *RegistryConfig) Prefix() string {
	return constant.RegistryConfigPrefix + "|" + constant.SingleRegistryConfigPrefix
}

func (r *RegistryConfig) getUrlMap(roleType common.RoleType) url.Values {
	urlMap := url.Values{}
	urlMap.Set(constant.GroupKey, r.Group)
	urlMap.Set(constant.RoleKey, strconv.Itoa(int(roleType)))
	urlMap.Set(constant.RegistryKey, r.Protocol)
	urlMap.Set(constant.RegistryTimeoutKey, r.TimeoutStr)
	// multi registry invoker weight label for load balance
	urlMap.Set(constant.RegistryKey+"."+constant.RegistryLabelKey, strconv.FormatBool(true))
	urlMap.Set(constant.RegistryKey+"."+constant.PreferredKey, strconv.FormatBool(r.Preferred))
	urlMap.Set(constant.RegistryKey+"."+constant.ZoneKey, r.Zone)
	urlMap.Set(constant.RegistryTtlKey, r.TTL)
	for k, v := range r.Params {
		urlMap.Set(k, v)
	}
	return urlMap
}

type registryBuilder struct {
}

func (r *registryBuilder) toURL(baseConfig BaseConfig) (*common.URL, error) {
	address := baseConfig.RegistryConfig.Address
	//address = translateRegistryConf(address, baseConfig.RegistryConfig)
	url, err := common.NewURL(constant.RegistryProtocol+"://"+address,
		common.WithParamsValue("simplified", strconv.FormatBool(baseConfig.RegistryConfig.Simplified)),
		common.WithUsername(baseConfig.RegistryConfig.Username),
		common.WithPassword(baseConfig.RegistryConfig.Password),
		common.WithLocation(baseConfig.RegistryConfig.Address),
		common.WithParams(baseConfig.RegistryConfig.getUrlMap(common.CONSUMER)),
	)

	if err != nil {
		logger.Errorf("The registry id: %s url is invalid, error: %#v", baseConfig.RegistryConfig.Address, err)
		panic(err)
	}
	return url, err
}

func (r *registryBuilder) buildRegistry(baseConfig BaseConfig) (mregistry.Registry, error) {
	newUrl, err := r.toURL(baseConfig)
	if err != nil {
		return nil, err
	}

	//init config center factory from service locator
	factory := component.GetRegistryFactory(newUrl.GetParam(constant.RegistryKey, constant.NacosKey))
	registry, err := factory.GetRegistry(newUrl)
	if err != nil {
		logger.Errorf("Get registry from configuration error , error message is %v", err)
		return nil, perrors.WithStack(err)
	}
	return registry, nil
}

func translateRegistryConf(address string, registryConf *RegistryConfig) string {
	if strings.Contains(address, "://") {
		translatedUrl, err := url.Parse(address)
		if err != nil {
			logger.Errorf("The registry url is invalid, error: %#v", err)
			panic(err)
		}
		address = translatedUrl.Host
		registryConf.Protocol = translatedUrl.Scheme
		registryConf.Address = strings.Replace(registryConf.Address, translatedUrl.Scheme+"://", "", -1)
	}
	return address
}

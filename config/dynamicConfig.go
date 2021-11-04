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

package config

import (
	"time"
)

import (
	gxset "github.com/dubbogo/gost/container/set"
)

import (
	"xmicro/common"
	"xmicro/config/parser"
)

// ////////////////////////////////////////
// DynamicConfiguration
// ////////////////////////////////////////
const (
	// DEFAULT_GROUP: default group
	DefaultGroup = "micro"
	// DEFAULT_CONFIG_TIMEOUT: default config timeout
	DefaultConfigTimeout = "10s"
)

// DynamicConfiguration for modify listener and get properties file
type DynamicConfiguration interface {
	//config parser
	Parser() parser.ConfigurationParser
	SetParser(parser.ConfigurationParser)
	//change listener
	AddListener(string, ConfigurationListener, ...Option)
	RemoveListener(string, ConfigurationListener, ...Option)
	// GetProperties get properties file
	GetProperties(string, ...Option) (string, error)

	// GetRule get Router rule properties file
	GetRule(string, ...Option) (string, error)

	// GetInternalProperty get value by key in Default properties file(dubbo.properties)
	GetInternalProperty(string, ...Option) (string, error)

	// PublishConfig will publish the config with the (key, group, value) pair
	PublishConfig(string, string, string) error

	// RemoveConfig will remove the config white the (key, group) pair
	RemoveConfig(string, string) error

	// GetConfigKeysByGroup will return all keys with the group
	GetConfigKeysByGroup(group string) (*gxset.HashSet, error)
}

type DynamicConfigurationFactory interface {
	GetDynamicConfiguration(*common.URL) (DynamicConfiguration, error)
}

// Options ...
type Options struct {
	Group   string
	Timeout time.Duration
}

// Option ...
type Option func(*Options)

// WithGroup assigns group to opt.Group
func WithGroup(group string) Option {
	return func(opt *Options) {
		opt.Group = group
	}
}

// WithTimeout assigns time to opt.Timeout
func WithTimeout(time time.Duration) Option {
	return func(opt *Options) {
		opt.Timeout = time
	}
}

// GetRuleKey The format is '{interfaceName}:[version]:[group]'
func GetRuleKey(url *common.URL) string {
	return url.ColonSeparatedKey()
}

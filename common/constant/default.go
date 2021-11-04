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

package constant

const (
	Micro             = "micro"
	PROVIDER_PROTOCOL = "provider"
	//compatible with 2.6.x
	OverrideProtocol = "override"
	EMPTY_PROTOCOL   = "empty"
	ROUTER_PROTOCOL  = "router"
)

const (
	DefaultWeight = 100     //
	DefaultWarmup = 10 * 60 // in java here is 10*60*1000 because of System.currentTimeMillis() is measured in milliseconds & in go time.Unix() is second
)

const (
	DefaultLoadbalance      = "roundrobin"
	DefaultRetries          = "2"
	DefaultRetriesInt       = 2
	DefaultProtocol         = "grpc"
	DefaultRegTimeout       = "10s"
	DefaultRegTtl           = "15m"
	DefaultFailbackTimes    = "3"
	DefaultFailbackTimesInt = 3
	DefaultFailbackTasks    = 100
	DefaultRestClient       = "resty"
	DefaultRestServer       = "go-restful"
	DefaultPort             = 20000
	DefaultSerialization    = Hessian2Serialization
)

const (
	DefaultKey                = "default"
	PREFIX_DEFAULT_KEY        = DefaultKey + "."
	DEFAULT_SERVICE_FILTERS   = "echo,token,accesslog,tps,generic_service,execute,pshutdown"
	DEFAULT_REFERENCE_FILTERS = "cshutdown"
	GENERIC_REFERENCE_FILTERS = "generic"
	GENERIC                   = "$invoke"
	ECHO                      = "$echo"
)

const (
	AnyValue          = "*"
	AnyhostValue      = "0.0.0.0"
	LocalHostValue    = "127.0.0.1"
	RemoveValuePrefix = "-"
)

const (
	ConfiguratorsCategory           = "configurators"
	RouterCategory                  = "category"
	DefaultCategory                 = ProviderCategory
	DynamicConfiguratorsCategory    = "dynamicconfigurators"
	AppDynamicConfiguratorsCategory = "appdynamicconfigurators"
	ProviderCategory                = "providers"
	ConsumerCategory                = "consumers"
)

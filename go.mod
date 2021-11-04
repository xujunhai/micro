module xmicro

go 1.14

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20191206172530-e9b2fee46413
	golang.org/x/net => github.com/golang/net v0.0.0-20191207000613-e7e4b65ae663
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys => github.com/golang/sys v0.0.0-20191206220618-eeba5f6aabab
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20191206204035-259af5ff87bd
)

require (
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/Shopify/sarama v1.19.0
	github.com/Shopify/toxiproxy v2.1.4+incompatible // indirect
	github.com/Workiva/go-datastructures v1.0.52
	github.com/apache/dubbo-getty v1.4.1
	github.com/creasty/defaults v1.5.1
	github.com/dubbogo/gost v1.10.1
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/frankban/quicktest v1.11.2 // indirect
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-redis/redis/v8 v8.4.4
	github.com/golang/mock v1.4.4 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/golang/snappy v0.0.2 // indirect
	github.com/google/uuid v1.1.2
	github.com/jinzhu/copier v0.1.0
	github.com/json-iterator/go v1.1.9
	github.com/magiconair/properties v1.8.4
	github.com/mitchellh/mapstructure v1.2.3 // indirect
	github.com/nacos-group/nacos-sdk-go v1.0.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0 // indirect
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	github.com/zouyx/agollo/v3 v3.4.5
	go.uber.org/zap v1.15.0
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/sys v0.0.0-20200523222454-059865788121 // indirect
	golang.org/x/tools v0.0.0-20200426102838-f3a5411a4c3b // indirect
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.3.0
	gorm.io/gorm v1.20.8
)

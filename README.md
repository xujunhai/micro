#说明
 -基于micro2.9.0进行扩展
#示例
  ##配置文件：
  port: 8090
  #服务元数据
  application:
    organization: "your group"
    name: "micro_service_xxx"
    version: "v0.2.1"
    owner: "your name"
    env: "prod"
  #配置中心
  configCenter:
    protocol: "nacos"
    address: "10.16.37.xxx:8848"
    group: "micro"
    username:
    password:
    logDir: "/tmp/logs"
    timeout: "2s"
  #注册中心
  registry:
    protocol: "nacos"
    timeout: "4s"
    group: "micro"
    ttl: "10m"
    address: "10.16.37.xxx:8848"
    username:
    password:
  #默认zap
  logger:
    level: "debug"
    logDir: "/tmp/logs"
  #默认jaeger
trace:
  fromEnv: false
  agent: "10.16.37.xxx:6831"
#handler
wrapper: "tracing"

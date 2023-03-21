### 简要介绍
一个单体的后端服务框架，包含一以下技术代码，下载即用。


### 技术选型及开发环境
| 序号 | 技术                   | 版本              | 说明            | 官网                                         |
|----|----------------------|-----------------|---------------|--------------------------------------------|
| 1  | Go                   | 1.18.3          | 开发语言          | https://go.dev/                            |
| 2  | Lua                  | 5.3             | Redis Lua 脚本  | http://www.lua.org/                        |
| 3  | Mysql                | 8.0.29          | 关系型数据库        | https://www.mysql.com/                     |
| 4  | Redis                | 6.2.7           | KV 数据库        | https://redis.io/                          |
| 5  | Docker               | 20.10.17        | 应用容器引擎        | https://www.docker.com                     |
| 6  | Nginx                | 1.23.1          | Web/反向代理      | https://www.nginx.com/                     |
| 7  | JWT                  | v4.4.2          | JWT登录支持       | https://github.com/golang-jwt/jwt          |
| 8  | Validator            | v10.11.0        | 验证框架          | https://github.com/go-playground/validator |
| 9  | Jaeger               | all-in-one:1.37 | 链路追踪          | https://www.jaegertracing.io               |
| 10 | ETCD                 | 3.5             | 服务发现          | https://etcd.io/                           |
| 11 | ShardingSphere-Proxy | 5.2.1           | 数据库中间件        | https://shardingsphere.apache.org/         |
| 12 | Gin                  | v1.8.1          | Web 框架        | https://github.com/gin-gonic/gin           |
| 13 | gRpc                 | v1.49.0         | 远程过程调用系统      | https://pkg.go.dev/google.golang.org/grpc  |
| 14 | GORM                 | v1.23.8         | ORM           | https://gorm.io/                           |


### 使用技术：
语言：Golang
数据库：MySQL 
web框架：Gin  
日志框架：Zap  
ORM框架：Gorm
分库分表：sharding-gorm
配置框架：Viper
权限管理框架：Casbin
用户认证：Jwt
接口文档：Swagger
分布式缓存:reis
消息队列:kafka
资源监控(普罗米修斯)：Prometheus
发送邮件：gomail.v2
雪花算法：tool/snowflake       
布隆过滤器：tool/bloom_filter     
一致性哈希算法：tool/consistent_hash
大文件上传：app/controller/file
验证码：captcha
校验器：validata
定时器：task
文件系统：minio
分布式追踪系统：jeager
k8s二次开发：k8s.io/client-go


### 项目目录简介
项目结构遵循：https://github.com/golang-standards/project-layout

conf:                服务配置
pkg:                 服务共有代码
docs:                swag文档
static:              保存的静态文件
internal:            服务私有代码
logs:                日志文件
script:              脚本文件


### 项目代码规范
google代码规范：https://google.github.io/styleguide/go/guide
uber代码规范：https://github.com/xxjwxc/uber_go_guide_cn


### 启动项目


### 接口文档用swag框架
启动后端服务后:http://127.0.0.1:8888/swagger/index.html#/





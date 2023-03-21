### 简要介绍
一个单体的后端服务框架，包含一以下技术代码，下载即用。

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
启动背包后端服务后
http://127.0.0.1:8888/swagger/index.html#/





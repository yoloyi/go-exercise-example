## Go-Micro

> 学习 Go-Micro 首先要理解 Go-Micro 与 Micro 之间的关系。

### 什么是 `Go-Micro` ?

Micro 是一个着眼于分布式系统开发的微服务生态系统。

* [Micro 是框架](https://github.com/micro/go-micro)
* [Micro 也是工具集](https://github.com/micro/micro)
* [Micro 有社区](http://slack.micro.mu/)
* [Micro 还是一个生态系统](https://micro.mu/explore/)

### 什么是 `Go-Micro` ?

基于 `Go` 语言的可插拔 `RPC` 微服务开发框架；包含服务发现、`RPC` 客户/服务端、广播/订阅机制等等，他是组成 Micro 的一个库。
大致总结就是 Micro 是 Go-Micro 的一个运行环境（Micro runtime)。

## 如何学习 `Go-Micro` ？

下面从官方实例 `Hello world` 来编写第一个 `go-micro` 程序。


### 前期准备

1、安装 [proto](https://github.com/protocolbuffers/protobuf/releases)

2、安装 [proto-gen-go](https://github.com/golang/protobuf/protoc-gen-go)

3、安装 [protoc-gen-micro](https://github.com/micro/protoc-gen-micro)

使用命令 `.proto` 文件生成 `.pb.go` 文件与 `.pb.micro.go` 文件。
```bash
 $ protoc -I go-mirco/helloworld/proto  --micro_out=plugins=grpc:./go-mirco/helloworld/proto  --go_out=plugins=grpc:./go-mirco/helloworld/proto helloworld.proto
```

### Hello world 例子

1、[Service](cmd/srv/main.go)

2、[Client](cmd/client/main.go) 














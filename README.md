## Wallet Manager Svr

依据 [gRPC 系统架构解决方案](https://github.com/ZuoFuhong/grpc-system-design) 创建的标准 gRPC 服务，使用 [monica](https://github.com/ZuoFuhong/monica) 服务注册中心，
全链路日志上报到 [grpc-datacollector](https://github.com/ZuoFuhong/grpc-datacollector).

### Running

```shell
go run main.go
```

在服务启动前，请先确保下列依赖服务已经在本地安装且可以正常访问：

- [Monica 服务注册中心](https://github.com/ZuoFuhong/monica)

### 部署部署

1.二进制包部署

```shell
go build -o go_wallet_manage_svr
chmod +x restart.sh
./restart.sh
```

2.Docker 部署

```shell
# 编译二进制
GOOS=linux go build -o main .

# 编译镜像
docker build -t ccr.ccs.tencentyun.com/tcb-xxx-xupz/prod-xxx-online:[tag] .

# 推送镜像
docker push ccr.ccs.tencentyun.com/tcb-xxx-xupz/prod-xxxx-online:[tag]
```

### License

This project is licensed under the [Apache 2.0 license](https://github.com/ZuoFuhong/go_wallet_manage_svr/blob/master/LICENSE).
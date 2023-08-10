# Kratos Project Template

## Install Kratos

```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

## Create a service

```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```

## Generate other auxiliary files by Makefile

```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```

## Automated Initialization (wire)

```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker

```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

## K8S-kratos
-   前提条件: 安装K8S istio consul
1. 测试k8s deployment异常退出后重启,在`info` `main.go +86`设置一分钟后程序退出
2. 测试k8s istio网关配置, 访istio网关端口查看是否能正常访问两个service
3. 测试k8s 内部consul发现注册,注册discovery到`user` `/v1/info/create/1` 路由请求到 `/v1/user/get`, 成功返回`{"data":{"code":200,"msg":"get_user"}}`
4. 测试k8s 内部访问本地redis, 访问`/v1/info/update/1`向redis写入`{update:1}`后再访问`/v1/info/get/1`获得写入redis的数据`{"data":{"code":200,"msg":"{update:1}"}}`

```bash
# 查看注册信息 http://localhost:8500/v1/catalog/service/info
# upload docker images
kubectl apply -f k8s-apply.yaml

```
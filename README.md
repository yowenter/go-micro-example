# go micro example




## 配置环境变量

```bash
source env.sh
```

## 安装依赖

```bash
go get github.com/micro/protoc-gen-micro/v2  # 在 bin 目录下会有 protoc-gen-micro
go get -u github.com/golang/protobuf/protoc-gen-go  # 在 bin 目录下会有 protoc-gen-go

go get github.com/micro/go-micro  # 下载 go micro 服务框架，下载完在 pkg 目录下会有一吨东西
```

## 生成 proto pb go 文件

```bash
protoc --plugin=protoc-gen-go=$GOPATH/bin/protoc-gen-go --plugin=protoc-gen-micro=$GOPATH/bin/protoc-gen-micro --proto_path=$GOPATH/src/wuu.com/example/hello:. --micro_out=$GOPATH/src/wuu.com/example/hello --go_out=$GOPATH/src/wuu.com/example/hello $GOPATH/src/wuu.com/example/hello/hello.proto
```

```bash
cd src/wuu.com/example
go mod init 
go run server/server.go    # 启动 server
go run client/client.go    # 在新的 terminal 启动 客户端
```

## 清除 go mod 缓存
```bash
go clean -modcache
```


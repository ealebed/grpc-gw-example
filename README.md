# Greeter Service

An example Go-Micro based gRPC service

## What's here?

- **server** - a gRPC greeter service
- **gateway** - a grpc-gateway


## Generate proto for service

```
protoc -I/usr/local/include -I. \
	-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I server/proto/hello/hello.proto \
	--go_out=plugins=grpc:. \
	server/proto/hello/hello.proto
```

## Generate proto for gateway

```
protoc -I/usr/local/include -I. \
	-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I gateway/proto/hello/hello.proto \
	--go_out=plugins=grpc:. \
	gateway/proto/hello/hello.proto
```

```
protoc -I/usr/local/include -I. \
	-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I gateway/proto/hello/hello.proto \
	--grpc-gateway_out=logtostderr=true:. \
	gateway/proto/hello/hello.proto
```

## Test Gateway

Run Service
```
$ go run server/main.go
```

Run gateway

```
$ go run gateway/main.go
```

Curl gateway

```
$ curl -d '{"name": "John"}' http://localhost/greeter/hello
```

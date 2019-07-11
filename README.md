# grpc-sample

## how to running
```$xslt
1. go get github.com/hakuna86/grpc-sample
2. running server : go run $PROJECT_ROOT/cmd/server/server.go
3. running client : go run $PROJECT_ROOT/cmd/client/client.go

all done.
```

## modify proto source
```$xslt
$PROJECT_ROOT/proto_build.sh
```

## Setting for grpc

```$xslt
1. Install gRPC
  $ go get -u google.golang.org/grpc
  
2. Install Protocol Buffers v3
  $ go get -u github.com/golang/protobuf/protoc-gen-go
  
3. Installing protoc
   https://github.com/protocolbuffers/protobuf/releases
   > download and unzip, and setting
```


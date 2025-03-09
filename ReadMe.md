This is a simple REST API server built in Go lang. The data packets make use of Protocol Buffer messages to transfer data between server and client. Protocol Buffers are designed to be light weight. 

The following commands can be used to generate stubs and reverse proxy for conversion of a gRPC into a REST protocol and vice versa.

Generating the stub with:
```
protoc -I=helloworld/ \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 helloworld/service.proto \
 --go_out=plugins=grpc:HelloWorld
```

Generating reverse-proxy with:
```
protoc -I=helloworld \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --grpc-gateway_out=logtostderr=true:HelloWorld/ \
 helloworld/service.proto
```

This uses necessary apis for creating the stub in Go.
Server entrypoint should then be written to utilize the stub

Then build your code using:
```
go build
```

Execute the binary n you are good to go!


Ref:
1) https://github.com/grpc-ecosystem/grpc-gateway

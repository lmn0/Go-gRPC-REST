Running protocol buffer for Go should be done with the following command:

Generating the stub with:
```
protoc -I=HelloWorld/ \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 HelloWorld/service.proto \
 --go_out=plugins=grpc:HelloWorld
```

Generating reverse-proxy with:
```
protoc -I=HelloWorld \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --grpc-gateway_out=logtostderr=true:HelloWorld/ \
 HelloWorld/service.proto
```

This uses necessary apis for creating the stub in Go.
Server entrypoint should then be written to utilize the stub

Ref:
1) https://github.com/grpc-ecosystem/grpc-gateway
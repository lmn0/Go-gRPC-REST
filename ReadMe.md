Running protocol buffer for Go should be done with the following command:

```
protoc -I=HelloWorld/ \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 HelloWorld/service.proto \
 --go_out=plugins=grpc:HelloWorld
```

This uses necessary apis for creating the stub in Go.
Server entrypoint should then be written to utilize the stub

Ref:
1) https://github.com/grpc-ecosystem/grpc-gateway
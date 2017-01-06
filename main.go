package main
import (
 "flag"
 "net/http"
 "sync"

 "github.com/golang/glog"
 "golang.org/x/net/context"
 "github.com/grpc-ecosystem/grpc-gateway/runtime"
 "google.golang.org/grpc"
 "google.golang.org/grpc/codes"
 "google.golang.org/grpc/metadata"

 gw "helloworld"
)

var (
 echoEndpoint = flag.String("echo_endpoint", "localhost:9090", "/v1/echo")
)
type _EchoMessage struct{
	v *gw.EchoMessage
	m sync.Mutex
}

type EchoMessage interface{
  }

func (s *_EchoMessage) Msg(ctx context.Context, msg *gw.EchoMessage) (*gw.EchoMessage, error) {
	s.m.Lock();
	defer s.m.Unlock()

	err:= grpc.SendHeader(ctx,metadata.New(map[string]string{
		"uuid":msg.Value,
		}))
	if err != nil {
		return nil, err
	}

	if a := s.v; a !=nil{
		return a, nil
	}
	grpc.SetTrailer(ctx, metadata.New(map[string]string{
		"foo" : "foo2",
		"bar" : "bar2",
		}))
	return nil, grpc.Errorf(codes.NotFound, "not found")
}

func run() error {
 ctx := context.Background()
 ctx, cancel := context.WithCancel(ctx)
 defer cancel()

 mux := runtime.NewServeMux()

 opts := []grpc.DialOption{grpc.WithInsecure()}
 err := gw.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
 if err != nil {
   return err
 }

 http.ListenAndServe(":8085", mux)
 return nil
}

func main(){
 flag.Parse()
 defer glog.Flush()

 if err := run(); err != nil {
   glog.Fatal(err)
 }



}
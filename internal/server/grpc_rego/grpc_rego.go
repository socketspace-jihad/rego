package grpc_rego

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/socketspace-jihad/rego/internal/core"
	"github.com/socketspace-jihad/rego/internal/server"
	"github.com/socketspace-jihad/rego/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
)

type GRPCRego struct {
	proto.UnimplementedKeyValueStorageServer
}

func (g *GRPCRego) Set(ctx context.Context, in *proto.KeyValue) (*proto.Status, error) {
	fmt.Println("SET OPERATION TRIGGERED", in.Key, in.Value)
	core.Set(in.Key, in.Value)
	return &proto.Status{}, nil
}

func (g *GRPCRego) Get(ctx context.Context, in *proto.Key) (*proto.Value, error) {
	fmt.Println("GET OPERATION TRIGGERED")
	val, _ := core.Get(in.Key)
	fmt.Println(val)
	anyVal, _ := val.(*anypb.Any)
	return &proto.Value{
		Value: anyVal,
	}, nil
}

func (g *GRPCRego) Serve() {
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()

	proto.RegisterKeyValueStorageServer(srv, &GRPCRego{})

	log.Println("grpc server is listening on :50050..")

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}

}

func NewGRPCRego() server.Server {
	return &GRPCRego{}
}

func init() {
	server.RegisterServer("grpc_rego", NewGRPCRego())
}

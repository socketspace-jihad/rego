package server

import (
	"context"
	"errors"

	"github.com/socketspace-jihad/rego/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	goproto "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type GRPCRego struct {
	Host     string
	conn     *grpc.ClientConn
	grpcConn proto.KeyValueStorageClient
}

type GRPCRegoConfig func(*GRPCRego)

func WithHostname(host string) GRPCRegoConfig {
	return func(g *GRPCRego) {
		g.Host = host
	}
}

func NewGRPCConnection(configs ...GRPCRegoConfig) *GRPCRego {
	grpcConn := &GRPCRego{
		Host: "localhost:50050",
	}
	for _, config := range configs {
		config(grpcConn)
	}
	return grpcConn
}

func (g *GRPCRego) Connect() error {
	conn, err := grpc.NewClient(g.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	g.conn = conn
	g.grpcConn = proto.NewKeyValueStorageClient(g.conn)
	return nil
}

func (g *GRPCRego) Disconnect() error {
	return g.conn.Close()
}

func (g *GRPCRego) Get(key string) (any, error) {
	val, err := g.grpcConn.Get(context.Background(), &proto.Key{Key: key})
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (g *GRPCRego) Set(key string, value any) error {
	msg, ok := value.(goproto.Message)
	if !ok {
		return errors.New("failed to convert to anypb.Any")
	}
	anyMsg, err := anypb.New(msg)
	if err != nil {
		return err
	}
	_, err = g.grpcConn.Set(context.Background(), &proto.KeyValue{Key: key, Value: anyMsg})
	if err != nil {
		return err
	}
	return nil
}

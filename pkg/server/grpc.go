package server

import (
	"context"
	"fmt"

	"github.com/socketspace-jihad/rego/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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

func (g *GRPCRego) GetString(key string) (string, error) {
	val, err := g.grpcConn.Get(context.Background(), &proto.Key{Key: key})
	if err != nil {
		return "", err
	}
	fmt.Println("any value", val)
	value := &wrapperspb.StringValue{}
	if err := val.Value.UnmarshalTo(value); err != nil {
		return "", err
	}
	return value.Value, nil
}

func (g *GRPCRego) SetString(key string, value string) error {
	stringValue := &wrapperspb.StringValue{
		Value: value,
	}
	anyMsg, err := anypb.New(stringValue)
	if err != nil {
		return err
	}
	_, err = g.grpcConn.Set(context.Background(), &proto.KeyValue{Key: key, Value: anyMsg})
	if err != nil {
		return err
	}
	return nil
}

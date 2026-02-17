package server

import (
	"context"
	"errors"

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
	if val.Value == nil {
		return "", errors.New("key doesn't exists")
	}
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

func (g *GRPCRego) GetBool(key string) (bool, error) {
	val, err := g.grpcConn.Get(context.Background(), &proto.Key{Key: key})
	if err != nil {
		return false, err
	}
	if val.Value == nil {
		return false, errors.New("key doesn't exists")
	}
	value := &wrapperspb.BoolValue{}
	if err := val.Value.UnmarshalTo(value); err != nil {
		return false, err
	}
	return value.Value, nil
}

func (g *GRPCRego) SetBool(key string, value bool) error {
	stringValue := &wrapperspb.BoolValue{
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

func (g *GRPCRego) GetInt(key string) (int64, error) {
	val, err := g.grpcConn.Get(context.Background(), &proto.Key{Key: key})
	if err != nil {
		return 0, err
	}
	if val.Value == nil {
		return 0, errors.New("key doesn't exists")
	}
	value := &wrapperspb.Int64Value{}
	if err := val.Value.UnmarshalTo(value); err != nil {
		return 0, err
	}
	return value.Value, nil
}

func (g *GRPCRego) SetInt(key string, value int64) error {
	stringValue := &wrapperspb.Int64Value{
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

func (g *GRPCRego) GetFloat(key string) (float32, error) {
	val, err := g.grpcConn.Get(context.Background(), &proto.Key{Key: key})
	if err != nil {
		return 0, err
	}
	if val.Value == nil {
		return 0, errors.New("key doesn't exists")
	}
	value := &wrapperspb.FloatValue{}
	if err := val.Value.UnmarshalTo(value); err != nil {
		return 0, err
	}
	return value.Value, nil
}

func (g *GRPCRego) SetFloat(key string, value float32) error {
	stringValue := &wrapperspb.FloatValue{
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

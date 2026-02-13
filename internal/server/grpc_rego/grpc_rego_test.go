package grpc_rego

import (
	"context"
	"errors"
	"testing"

	"github.com/socketspace-jihad/rego/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func Test_GRPCRego(t *testing.T) {
	srv := &GRPCRego{}

	t.Run("GRPC Set operation", func(t *testing.T) {

		req := &proto.KeyValue{
			Key:   "test",
			Value: &anypb.Any{Value: []byte("this is value")},
		}

		_, err := srv.Set(context.Background(), req)
		if err != nil {
			t.Error(err)
			return
		}
	})

	t.Run("GRPC Set is not working", func(t *testing.T) {
		req := &proto.Key{
			Key: "test",
		}

		resp, err := srv.Get(context.Background(), req)
		if err != nil {
			t.Error(err)
			return
		}

		if string(resp.Value.Value) == "" {
			t.Error(errors.New("Set operation doesnt work"))
		}
	})

	t.Run("GRPC Get operation", func(t *testing.T) {

		req := &proto.Key{
			Key: "test",
		}

		resp, err := srv.Get(context.Background(), req)
		if err != nil {
			t.Error(err)
			return
		}

		if string(resp.Value.Value) != "this is value" {
			t.Error(errors.New("Get operation is not returning correct value"))
		}
	})

}

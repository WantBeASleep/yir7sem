package grpc

import (
	"context"
	"fmt"

	pb "auth/internal/generated/grpc/service"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type ValidatorInterceptor struct {
	validator *protovalidate.Validator
}

func InitValidator() (*ValidatorInterceptor, error) { //TODO: messages ...interface{} - нужно ли передавать в метод
	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			&pb.LoginIn{},
			&pb.RefreshIn{},
			&pb.RegisterIn{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("init validator: %v", err)
	}
	return &ValidatorInterceptor{
			validator: validator,
		},
		nil
}

func (vi *ValidatorInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if err := vi.validator.Validate(req.(proto.Message)); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %v", err))
		}
		return handler(ctx, req)
	}
}

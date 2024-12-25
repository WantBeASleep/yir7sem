package grpc

import (
	"context"
	"fmt"

	pb "uzi/internal/generated/grpc/service"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type ValidatorInterceptor struct {
	validator *protovalidate.Validator
}

func InitValidator() (*ValidatorInterceptor, error) {
	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			&pb.CreateUziIn{},
			&pb.UpdateUziIn{},
			&pb.GetUziIn{},
			&pb.UpdateEchographicIn{},
			&pb.GetEchographicIn{},

			&pb.GetUziImagesIn{},
			&pb.GetImageSegmentsWithNodesIn{},

			&pb.CreateSegmentIn{},
			&pb.DeleteSegmentIn{},
			&pb.UpdateSegmentIn{},

			&pb.CreateNodeIn{},
			&pb.UpdateNodeIn{},
			&pb.DeleteNodeIn{},
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

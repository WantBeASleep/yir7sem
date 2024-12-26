package grpc

import (
	"context"
	"fmt"
	"time"

	pb "med/internal/generated/grpc/service"

	"github.com/bufbuild/protovalidate-go"

	obslib "github.com/senorUVE/observer-yir/observerlib"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type ValidatorInterceptor struct {
	validator *protovalidate.Validator
	observer  *obslib.Observer
}

func InitValidator(observer *obslib.Observer) (*ValidatorInterceptor, error) {
	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			&pb.CreatePatientIn{},
			&pb.GetPatientIn{},
			&pb.UpdatePatientIn{},
			&pb.GetDoctorPatientsIn{},

			&pb.GetDoctorIn{},
			&pb.UpdateDoctorIn{},
			&pb.RegisterDoctorIn{},

			&pb.CreateCardIn{},
			&pb.UpdateCardIn{},
			&pb.GetCardIn{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("init validator: %v", err)
	}
	return &ValidatorInterceptor{
			validator: validator,
			observer:  observer,
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
		start := time.Now()
		requestID := fmt.Sprintf("%d", start.UnixNano())
		vi.observer.LogMetrics(ctx, requestID, info.FullMethod, time.Since(start), 0)
		message, ok := req.(proto.Message)
		if !ok {
			vi.observer.LogError(ctx, requestID, "invalid request type", "req is not proto.Message")
			return nil, status.Errorf(codes.InvalidArgument, "invalid request type")
		}
		if err := vi.validator.Validate(message); err != nil {
			vi.observer.LogError(ctx, requestID, fmt.Sprintf("validation failed: %v", err), "validation")
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %v", err))
		}
		resp, err := handler(ctx, req)
		if err != nil {
			vi.observer.LogError(ctx, requestID, err.Error(), fmt.Sprintf("request failed: %v", err))
		}
		vi.observer.LogMetrics(ctx, requestID, info.FullMethod, time.Since(start), 200)
		return resp, err
	}
}

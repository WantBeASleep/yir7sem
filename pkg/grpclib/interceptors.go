package grpclib

import (
	"context"
	"log/slog"

	"yir/pkg/ctxlib"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	requestIdKey = "x-request_id"
	methodKey    = "x-method"
)

func ServerCallLoggerInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		slog.WarnContext(ctx, "Server call w/o metadata aborted", slog.String(methodKey, info.FullMethod))
		return nil, status.Error(codes.Unauthenticated, "request id required")
	}

	requestIDArr := md.Get(requestIdKey)
	if len(requestIDArr) != 1 {
		slog.WarnContext(ctx, "Server call w/o request id aborted", slog.String(methodKey, info.FullMethod))
		return nil, status.Error(codes.Unauthenticated, "request id required")
	}

	requestID, err := uuid.Parse(requestIDArr[0])
	if err != nil {
		slog.WarnContext(ctx, "Server call w/o request id aborted", slog.String(methodKey, info.FullMethod))
		return nil, status.Error(codes.Unauthenticated, "request id invalid")
	}

	ctx = ctxlib.PublicSet(ctx, requestIdKey, requestID)
	ctx = ctxlib.PublicSet(ctx, methodKey, info.FullMethod)
	slog.InfoContext(ctx, "Server call")

	return handler(ctx, req)
}

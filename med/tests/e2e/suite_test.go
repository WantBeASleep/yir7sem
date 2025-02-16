//go:build e2e

package e2e_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	pb "med/internal/generated/grpc/service"

	"github.com/WantBeASleep/med_ml_lib/auth"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TestSuite struct {
	suite.Suite

	// обогащен аутентификацией
	ctx        context.Context
	grpcClient pb.MedSrvClient
}

func (suite *TestSuite) SetupSuite() {
	conn, err := grpc.NewClient(
		os.Getenv("APP_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(auth.AuthEnrichClientCall),
	)
	if err != nil {
		panic(fmt.Sprintf("grpc connection failed: %v", err))
	}
	suite.grpcClient = pb.NewMedSrvClient(conn)

	suite.ctx = auth.WithRequestID(context.Background(), uuid.New())
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

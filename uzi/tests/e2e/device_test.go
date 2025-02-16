//go:build e2e

package e2e_test

import (
	pb "uzi/internal/generated/grpc/service"

	"github.com/stretchr/testify/require"
	"github.com/thanhpk/randstr"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (suite *TestSuite) TestCreateGetDevice_Success() {
	type deviceTestElem struct {
		value   string
		founded bool
	}
	devices := map[int]deviceTestElem{}
	for range 5 {
		deviceName := randstr.String(5)
		resp, err := suite.grpcClient.CreateDevice(
			suite.ctx,
			&pb.CreateDeviceIn{Name: deviceName},
		)
		require.NoError(suite.T(), err)

		devices[int(resp.Id)] = deviceTestElem{value: deviceName}
	}

	resp, err := suite.grpcClient.GetDeviceList(suite.ctx, &emptypb.Empty{})
	require.NoError(suite.T(), err)

	for _, device := range resp.Devices {
		exceptedElem, ok := devices[int(device.Id)]
		if ok {
			require.Equal(suite.T(), exceptedElem.value, device.Name)
			exceptedElem.founded = true
			devices[int(device.Id)] = exceptedElem
		}
	}

	for _, v := range devices {
		require.True(suite.T(), v.founded)
	}
}

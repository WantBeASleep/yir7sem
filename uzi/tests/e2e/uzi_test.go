//go:build e2e

package e2e_test

import (
	pb "uzi/internal/generated/grpc/service"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/thanhpk/randstr"
)

func (suite *TestSuite) TestCreateGetUzi_Success() {
	// Создаем узи девайс
	deviceResp, err := suite.grpcClient.CreateDevice(
		suite.ctx,
		&pb.CreateDeviceIn{Name: randstr.String(5)},
	)
	require.NoError(suite.T(), err)

	patientID := uuid.New().String()
	projection := randstr.String(5)

	// Создаем УЗИ
	createResp, err := suite.grpcClient.CreateUzi(
		suite.ctx,
		&pb.CreateUziIn{
			PatientId:  patientID,
			DeviceId:   deviceResp.Id,
			Projection: projection,
		},
	)
	require.NoError(suite.T(), err)
	require.NotEmpty(suite.T(), createResp.Id)
	_, err = uuid.Parse(createResp.Id)
	require.NoError(suite.T(), err)

	// Получаем созданное УЗИ
	getResp, err := suite.grpcClient.GetUzi(
		suite.ctx,
		&pb.GetUziIn{Id: createResp.Id},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), patientID, getResp.Uzi.PatientId)
	require.Equal(suite.T(), deviceResp.Id, getResp.Uzi.DeviceId)
	require.Equal(suite.T(), projection, getResp.Uzi.Projection)
	require.False(suite.T(), getResp.Uzi.Checked)
	require.NotEmpty(suite.T(), getResp.Uzi.CreateAt)

	// Обновляем УЗИ
	newProjection := randstr.String(5)
	updateResp, err := suite.grpcClient.UpdateUzi(
		suite.ctx,
		&pb.UpdateUziIn{
			Id:         createResp.Id,
			Projection: &newProjection,
			Checked:    &[]bool{true}[0],
		},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), newProjection, updateResp.Uzi.Projection)
	require.True(suite.T(), updateResp.Uzi.Checked)

	// Получаем список УЗИ пациента
	listResp, err := suite.grpcClient.GetPatientUzis(
		suite.ctx,
		&pb.GetPatientUzisIn{PatientId: patientID},
	)
	require.NoError(suite.T(), err)
	require.Len(suite.T(), listResp.Uzis, 1)
	require.Equal(suite.T(), createResp.Id, listResp.Uzis[0].Id)
}

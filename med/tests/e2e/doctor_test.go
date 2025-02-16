//go:build e2e

package e2e_test

import (
	pb "med/internal/generated/grpc/service"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestDoctor_Success() {
	// регистрируем доктора
	doctorId := uuid.New().String()
	_, err := suite.grpcClient.RegisterDoctor(suite.ctx, &pb.RegisterDoctorIn{
		Doctor: &pb.Doctor{
			Id:       doctorId,
			Fullname: "Test Doctor",
			Org:      "Test Hospital",
			Job:      "Surgeon",
			Desc:     nil,
		},
	})
	require.NoError(suite.T(), err)

	// получаем информацию о докторе
	getDoctorResp, err := suite.grpcClient.GetDoctor(suite.ctx, &pb.GetDoctorIn{
		Id: doctorId,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), doctorId, getDoctorResp.Doctor.Id)
	require.Equal(suite.T(), "Test Doctor", getDoctorResp.Doctor.Fullname)
	require.Equal(suite.T(), "Test Hospital", getDoctorResp.Doctor.Org)
	require.Equal(suite.T(), "Surgeon", getDoctorResp.Doctor.Job)
	require.Nil(suite.T(), getDoctorResp.Doctor.Desc)

	// обновляем информацию о докторе
	newDesc := "Experienced surgeon"
	newOrg := "New Hospital"
	updateDoctorResp, err := suite.grpcClient.UpdateDoctor(suite.ctx, &pb.UpdateDoctorIn{
		Id:   doctorId,
		Org:  &newOrg,
		Desc: &newDesc,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), doctorId, updateDoctorResp.Doctor.Id)
	require.Equal(suite.T(), "Test Doctor", updateDoctorResp.Doctor.Fullname)
	require.Equal(suite.T(), newOrg, updateDoctorResp.Doctor.Org)
	require.Equal(suite.T(), "Surgeon", updateDoctorResp.Doctor.Job)
	require.Equal(suite.T(), newDesc, *updateDoctorResp.Doctor.Desc)

	// проверяем что изменения сохранились
	getDoctorResp, err = suite.grpcClient.GetDoctor(suite.ctx, &pb.GetDoctorIn{
		Id: doctorId,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), doctorId, getDoctorResp.Doctor.Id)
	require.Equal(suite.T(), "Test Doctor", getDoctorResp.Doctor.Fullname)
	require.Equal(suite.T(), newOrg, getDoctorResp.Doctor.Org)
	require.Equal(suite.T(), "Surgeon", getDoctorResp.Doctor.Job)
	require.Equal(suite.T(), newDesc, *getDoctorResp.Doctor.Desc)
}

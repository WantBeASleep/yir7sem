//go:build e2e

package e2e_test

import (
	pb "uzi/internal/generated/grpc/service"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/thanhpk/randstr"
)

func (suite *TestSuite) TestCreateGetUpdateEchographic_Success() {
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

	// Получаем эхографику
	getResp, err := suite.grpcClient.GetEchographic(
		suite.ctx,
		&pb.GetEchographicIn{Id: createResp.Id},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), createResp.Id, getResp.Echographic.Id)
	require.Empty(suite.T(), getResp.Echographic.Contors)
	require.Empty(suite.T(), getResp.Echographic.LeftLobeLength)
	require.Empty(suite.T(), getResp.Echographic.LeftLobeWidth)
	require.Empty(suite.T(), getResp.Echographic.LeftLobeThick)
	require.Empty(suite.T(), getResp.Echographic.LeftLobeVolum)
	require.Empty(suite.T(), getResp.Echographic.RightLobeLength)
	require.Empty(suite.T(), getResp.Echographic.RightLobeWidth)
	require.Empty(suite.T(), getResp.Echographic.RightLobeThick)
	require.Empty(suite.T(), getResp.Echographic.RightLobeVolum)
	require.Empty(suite.T(), getResp.Echographic.GlandVolum)
	require.Empty(suite.T(), getResp.Echographic.Isthmus)
	require.Empty(suite.T(), getResp.Echographic.Struct)
	require.Empty(suite.T(), getResp.Echographic.Echogenicity)
	require.Empty(suite.T(), getResp.Echographic.RegionalLymph)
	require.Empty(suite.T(), getResp.Echographic.Vascularization)
	require.Empty(suite.T(), getResp.Echographic.Location)
	require.Empty(suite.T(), getResp.Echographic.Additional)
	require.Empty(suite.T(), getResp.Echographic.Conclusion)

	// Обновляем эхографику
	contors := randstr.String(5)
	leftLobeLength := 1.1
	leftLobeWidth := 2.2
	leftLobeThick := 3.3
	leftLobeVolum := 4.4
	rightLobeLength := 5.5
	rightLobeWidth := 6.6
	rightLobeThick := 7.7
	rightLobeVolum := 8.8
	glandVolum := 9.9
	isthmus := 10.1
	structure := randstr.String(5)
	echogenicity := randstr.String(5)
	regionalLymph := randstr.String(5)
	vascularization := randstr.String(5)
	location := randstr.String(5)
	additional := randstr.String(5)
	conclusion := randstr.String(5)

	updateResp, err := suite.grpcClient.UpdateEchographic(
		suite.ctx,
		&pb.UpdateEchographicIn{
			Echographic: &pb.Echographic{
				Id:              createResp.Id,
				Contors:         &contors,
				LeftLobeLength:  &leftLobeLength,
				LeftLobeWidth:   &leftLobeWidth,
				LeftLobeThick:   &leftLobeThick,
				LeftLobeVolum:   &leftLobeVolum,
				RightLobeLength: &rightLobeLength,
				RightLobeWidth:  &rightLobeWidth,
				RightLobeThick:  &rightLobeThick,
				RightLobeVolum:  &rightLobeVolum,
				GlandVolum:      &glandVolum,
				Isthmus:         &isthmus,
				Struct:          &structure,
				Echogenicity:    &echogenicity,
				RegionalLymph:   &regionalLymph,
				Vascularization: &vascularization,
				Location:        &location,
				Additional:      &additional,
				Conclusion:      &conclusion,
			},
		},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), contors, *updateResp.Echographic.Contors)
	require.Equal(suite.T(), leftLobeLength, *updateResp.Echographic.LeftLobeLength)
	require.Equal(suite.T(), leftLobeWidth, *updateResp.Echographic.LeftLobeWidth)
	require.Equal(suite.T(), leftLobeThick, *updateResp.Echographic.LeftLobeThick)
	require.Equal(suite.T(), leftLobeVolum, *updateResp.Echographic.LeftLobeVolum)
	require.Equal(suite.T(), rightLobeLength, *updateResp.Echographic.RightLobeLength)
	require.Equal(suite.T(), rightLobeWidth, *updateResp.Echographic.RightLobeWidth)
	require.Equal(suite.T(), rightLobeThick, *updateResp.Echographic.RightLobeThick)
	require.Equal(suite.T(), rightLobeVolum, *updateResp.Echographic.RightLobeVolum)
	require.Equal(suite.T(), glandVolum, *updateResp.Echographic.GlandVolum)
	require.Equal(suite.T(), isthmus, *updateResp.Echographic.Isthmus)
	require.Equal(suite.T(), structure, *updateResp.Echographic.Struct)
	require.Equal(suite.T(), echogenicity, *updateResp.Echographic.Echogenicity)
	require.Equal(suite.T(), regionalLymph, *updateResp.Echographic.RegionalLymph)
	require.Equal(suite.T(), vascularization, *updateResp.Echographic.Vascularization)
	require.Equal(suite.T(), location, *updateResp.Echographic.Location)
	require.Equal(suite.T(), additional, *updateResp.Echographic.Additional)
	require.Equal(suite.T(), conclusion, *updateResp.Echographic.Conclusion)

	// Получаем обновленную эхографику
	getUpdatedResp, err := suite.grpcClient.GetEchographic(
		suite.ctx,
		&pb.GetEchographicIn{Id: createResp.Id},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), contors, *getUpdatedResp.Echographic.Contors)
	require.Equal(suite.T(), leftLobeLength, *getUpdatedResp.Echographic.LeftLobeLength)
	require.Equal(suite.T(), leftLobeWidth, *getUpdatedResp.Echographic.LeftLobeWidth)
	require.Equal(suite.T(), leftLobeThick, *getUpdatedResp.Echographic.LeftLobeThick)
	require.Equal(suite.T(), leftLobeVolum, *getUpdatedResp.Echographic.LeftLobeVolum)
	require.Equal(suite.T(), rightLobeLength, *getUpdatedResp.Echographic.RightLobeLength)
	require.Equal(suite.T(), rightLobeWidth, *getUpdatedResp.Echographic.RightLobeWidth)
	require.Equal(suite.T(), rightLobeThick, *getUpdatedResp.Echographic.RightLobeThick)
	require.Equal(suite.T(), rightLobeVolum, *getUpdatedResp.Echographic.RightLobeVolum)
	require.Equal(suite.T(), glandVolum, *getUpdatedResp.Echographic.GlandVolum)
	require.Equal(suite.T(), isthmus, *getUpdatedResp.Echographic.Isthmus)
	require.Equal(suite.T(), structure, *getUpdatedResp.Echographic.Struct)
	require.Equal(suite.T(), echogenicity, *getUpdatedResp.Echographic.Echogenicity)
	require.Equal(suite.T(), regionalLymph, *getUpdatedResp.Echographic.RegionalLymph)
	require.Equal(suite.T(), vascularization, *getUpdatedResp.Echographic.Vascularization)
	require.Equal(suite.T(), location, *getUpdatedResp.Echographic.Location)
	require.Equal(suite.T(), additional, *getUpdatedResp.Echographic.Additional)
	require.Equal(suite.T(), conclusion, *getUpdatedResp.Echographic.Conclusion)
}

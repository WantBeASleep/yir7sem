//go:build e2e

package e2e_test

import (
	pb "med/internal/generated/grpc/service"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestCard_Success() {
	// регистрируем доктора для теста
	doctorId := uuid.New().String()
	_, err := suite.grpcClient.RegisterDoctor(suite.ctx, &pb.RegisterDoctorIn{
		Doctor: &pb.Doctor{
			Id:       doctorId,
			Fullname: "Test Doctor",
			Org:      "Test Hospital",
			Job:      "Surgeon",
		},
	})
	require.NoError(suite.T(), err)

	// создаем пациента
	createPatientResp, err := suite.grpcClient.CreatePatient(suite.ctx, &pb.CreatePatientIn{
		Fullname:   "Test Patient",
		Email:      "test@example.com",
		Policy:     "1234567890",
		Active:     true,
		Malignancy: false,
	})
	require.NoError(suite.T(), err)

	// создаем карту пациента
	diagnosis := "Initial diagnosis"
	_, err = suite.grpcClient.CreateCard(suite.ctx, &pb.CreateCardIn{
		Card: &pb.Card{
			DoctorId:  doctorId,
			PatientId: createPatientResp.Id,
			Diagnosis: &diagnosis,
		},
	})
	require.NoError(suite.T(), err)

	// получаем карту пациента
	getCardResp, err := suite.grpcClient.GetCard(suite.ctx, &pb.GetCardIn{
		DoctorId:  doctorId,
		PatientId: createPatientResp.Id,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), doctorId, getCardResp.Card.DoctorId)
	require.Equal(suite.T(), createPatientResp.Id, getCardResp.Card.PatientId)
	require.NotNil(suite.T(), getCardResp.Card.Diagnosis)
	require.Equal(suite.T(), diagnosis, *getCardResp.Card.Diagnosis)

	// обновляем карту пациента
	newDiagnosis := "Updated diagnosis"
	updateCardResp, err := suite.grpcClient.UpdateCard(suite.ctx, &pb.UpdateCardIn{
		Card: &pb.Card{
			DoctorId:  doctorId,
			PatientId: createPatientResp.Id,
			Diagnosis: &newDiagnosis,
		},
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), doctorId, updateCardResp.Card.DoctorId)
	require.Equal(suite.T(), createPatientResp.Id, updateCardResp.Card.PatientId)
	require.NotNil(suite.T(), updateCardResp.Card.Diagnosis)
	require.Equal(suite.T(), newDiagnosis, *updateCardResp.Card.Diagnosis)

	// проверяем что изменения сохранились
	getCardResp, err = suite.grpcClient.GetCard(suite.ctx, &pb.GetCardIn{
		DoctorId:  doctorId,
		PatientId: createPatientResp.Id,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), doctorId, getCardResp.Card.DoctorId)
	require.Equal(suite.T(), createPatientResp.Id, getCardResp.Card.PatientId)
	require.NotNil(suite.T(), getCardResp.Card.Diagnosis)
	require.Equal(suite.T(), newDiagnosis, *getCardResp.Card.Diagnosis)
}

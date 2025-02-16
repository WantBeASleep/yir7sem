//go:build e2e

package e2e_test

import (
	"time"

	pb "med/internal/generated/grpc/service"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (suite *TestSuite) TestPatient_Success() {
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

	// получаем информацию о пациенте
	getPatientResp, err := suite.grpcClient.GetPatient(suite.ctx, &pb.GetPatientIn{
		Id: createPatientResp.Id,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), createPatientResp.Id, getPatientResp.Patient.Id)
	require.Equal(suite.T(), "Test Patient", getPatientResp.Patient.Fullname)
	require.Equal(suite.T(), "test@example.com", getPatientResp.Patient.Email)
	require.Equal(suite.T(), "1234567890", getPatientResp.Patient.Policy)
	require.True(suite.T(), getPatientResp.Patient.Active)
	require.False(suite.T(), getPatientResp.Patient.Malignancy)
	require.Nil(suite.T(), getPatientResp.Patient.LastUziDate)

	// проверяем что пациент не привязан к доктору
	getDoctorPatientsResp, err := suite.grpcClient.GetDoctorPatients(suite.ctx, &pb.GetDoctorPatientsIn{
		DoctorId: doctorId,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 0, len(getDoctorPatientsResp.Patients))

	// создаем карту пациента для связи с доктором
	diagnosis := "Initial diagnosis"
	_, err = suite.grpcClient.CreateCard(suite.ctx, &pb.CreateCardIn{
		Card: &pb.Card{
			DoctorId:  doctorId,
			PatientId: createPatientResp.Id,
			Diagnosis: &diagnosis,
		},
	})
	require.NoError(suite.T(), err)

	// обновляем информацию о пациенте
	lastUziDate := timestamppb.New(time.Now())
	malignancy := true
	active := false
	updatePatientResp, err := suite.grpcClient.UpdatePatient(suite.ctx, &pb.UpdatePatientIn{
		DoctorId:    doctorId,
		Id:          createPatientResp.Id,
		Active:      &active,
		Malignancy:  &malignancy,
		LastUziDate: lastUziDate,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), createPatientResp.Id, updatePatientResp.Patient.Id)
	require.Equal(suite.T(), "Test Patient", updatePatientResp.Patient.Fullname)
	require.Equal(suite.T(), "test@example.com", updatePatientResp.Patient.Email)
	require.Equal(suite.T(), "1234567890", updatePatientResp.Patient.Policy)
	require.False(suite.T(), updatePatientResp.Patient.Active)
	require.True(suite.T(), updatePatientResp.Patient.Malignancy)
	require.NotNil(suite.T(), updatePatientResp.Patient.LastUziDate)

	// проверяем что пациент теперь привязан к доктору
	getDoctorPatientsResp, err = suite.grpcClient.GetDoctorPatients(suite.ctx, &pb.GetDoctorPatientsIn{
		DoctorId: doctorId,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 1, len(getDoctorPatientsResp.Patients))
	require.Equal(suite.T(), createPatientResp.Id, getDoctorPatientsResp.Patients[0].Id)
	require.Equal(suite.T(), "Test Patient", getDoctorPatientsResp.Patients[0].Fullname)
	require.Equal(suite.T(), "test@example.com", getDoctorPatientsResp.Patients[0].Email)
	require.Equal(suite.T(), "1234567890", getDoctorPatientsResp.Patients[0].Policy)
	require.False(suite.T(), getDoctorPatientsResp.Patients[0].Active)
	require.True(suite.T(), getDoctorPatientsResp.Patients[0].Malignancy)
	require.NotNil(suite.T(), getDoctorPatientsResp.Patients[0].LastUziDate)
}

package controller

import (
	"net/http"
	"yir/api-gateway/internal/controller/auth"
	"yir/api-gateway/internal/controller/med/patient"
	"yir/api-gateway/internal/controller/med/worker"
	"yir/api-gateway/internal/controller/uzi"
	"yir/api-gateway/internal/logger"
	"yir/api-gateway/internal/middleware"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.VerifyToken)
	r.HandleFunc("/healthcheck", Healthcheck)

	authSubR := r.PathPrefix(authPrefix).Subrouter()
	authSubR.HandleFunc(authLogin, auth.Login)
	authSubR.HandleFunc(authRegister, auth.Login)
	authSubR.HandleFunc(authTokenRefresh, auth.Login)

	medSubR := r.PathPrefix(medPrefix).Subrouter()

	medPatientSubR := medSubR.PathPrefix(medPatientPrefix).Subrouter()
	medPatientSubR.HandleFunc(patientCreate, patient.Create)
	medPatientSubR.HandleFunc(patientInfo, patient.Info)
	medPatientSubR.HandleFunc(patientList, patient.List)
	//
	//medPatientSubR.HandleFunc(patientShots, patient.Shots) unimplemented in med patient service
	//
	medPatientSubR.HandleFunc(patientUpdate, patient.Update)

	medWorkerSubR := medSubR.PathPrefix("/worker").Subrouter()
	medWorkerSubR.HandleFunc(workerAdd, worker.Add)
	medWorkerSubR.HandleFunc(workerID, worker.ID)
	medWorkerSubR.HandleFunc(workerList, worker.List)
	medWorkerSubR.HandleFunc(workerPatients, worker.Patients)
	medWorkerSubR.HandleFunc(workerUpdate, worker.Update)

	uziSubR := r.PathPrefix("/uzi").Subrouter()
	uziSubR.HandleFunc(uziDeviceList, uzi.DeviceList)
	uziSubR.HandleFunc(uziFormationSegments_formation_id, uzi.FormationSegmID)
	uziSubR.HandleFunc(uziFormationSegments_uzi_id, uzi.FormationSegmUziID)
	uziSubR.HandleFunc(uziFormation, uzi.FormationID)
	uziSubR.HandleFunc(uziImageFormation, uzi.ImageFormation)
	uziSubR.HandleFunc(uziInfo, uzi.Info)
	uziSubR.HandleFunc(uziID, uzi.ID)

	return r
}

func Healthcheck(w http.ResponseWriter, req *http.Request) {
	logger.Logger.Info("Healthcheck",
		zap.String("client_ip", req.RemoteAddr),
	)
	w.Write([]byte("Hello!"))
}

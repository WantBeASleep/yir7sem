package controller

import (
	"net/http"
	"yir/gateway/internal/controller/auth"
	"yir/gateway/internal/controller/med"
	"yir/gateway/internal/controller/uzi"
	"yir/gateway/internal/logger"
	"yir/gateway/internal/middleware"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func InitRouter(ac *auth.AuthController, mc *med.MedController, uzic *uzi.UziController) *mux.Router {
	MainRouter := mux.NewRouter()
	MainRouter.Use(middleware.VerifyToken)
	MainRouter.HandleFunc("/healthcheck", Healthcheck)

	authSubR := MainRouter.PathPrefix(authPrefix).Subrouter()
	authSubR.HandleFunc(authLogin, ac.Login)               // токены не нужны
	authSubR.HandleFunc(authRegister, ac.Register)         // токены не нужны
	authSubR.HandleFunc(authTokenRefresh, ac.TokenRefresh) // нужен токен тип рефреш

	medSubR := MainRouter.PathPrefix(medPrefix).Subrouter()

	patientSubR := medSubR.PathPrefix(patientPrefix).Subrouter()
	patientSubR.HandleFunc(patientCreate, mc.PostPatient)
	patientSubR.HandleFunc(patientInfo, mc.GetPatientInfo)
	patientSubR.HandleFunc(patientList, mc.GetPatientList)
	//
	// medPatientSubR.HandleFunc(patientShots, mc.GetPatientShots) unimplemented in med patient service
	//
	patientSubR.HandleFunc(patientUpdate, mc.PutPatient)

	workerSubR := medSubR.PathPrefix(workerPrefix).Subrouter()
	workerSubR.HandleFunc(workerAdd, mc.PostWorker)
	workerSubR.HandleFunc(workerID, mc.GetWorkerID)
	workerSubR.HandleFunc(workerList, mc.GetWorkersList)
	workerSubR.HandleFunc(workerPatients, mc.GetWorkerPatients)
	workerSubR.HandleFunc(workerUpdate, mc.PutWorker)

	cardSubR := medSubR.PathPrefix(cardPrefix).Subrouter()
	cardSubR.HandleFunc(cardPrefix, mc.GetCards)
	cardSubR.HandleFunc(cardID, mc.PostCard)
	cardSubR.HandleFunc(cardID, mc.GetCardByID)
	cardSubR.HandleFunc(cardID, mc.DeleteCard)
	cardSubR.HandleFunc(cardID, mc.PutCard)

	uziSubR := MainRouter.PathPrefix("/uzi").Subrouter()
	uziSubR.HandleFunc(uziDeviceList, uzic.GetDeviceList)
	uziSubR.HandleFunc(uziFormationSegments_formation_id, uzic.GetFormationSegmID)
	uziSubR.HandleFunc(uziFormationSegmentsUziID, uzic.PostFormationSegmUziID)
	uziSubR.HandleFunc(uziFormation, uzic.GetFormationID)
	uziSubR.HandleFunc(uziImageID, uzic.GetImageID)
	uziSubR.HandleFunc(uziInfo, uzic.GetUziInfo)
	uziSubR.HandleFunc(uziID, uzic.GetUziID)

	return MainRouter
}

func Healthcheck(w http.ResponseWriter, req *http.Request) {
	logger.Logger.Info("Healthcheck",
		zap.String("client_ip", req.RemoteAddr),
	)
	w.Write([]byte("Hello!"))
}

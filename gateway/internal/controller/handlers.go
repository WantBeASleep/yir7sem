package controller

import (
	"net/http"
	"yir/gateway/internal/controller/auth"
	"yir/gateway/internal/controller/med"
	"yir/gateway/internal/controller/uzi"
	"yir/gateway/internal/custom"
	"yir/gateway/internal/middleware"
	"yir/gateway/internal/service"
	"yir/gateway/internal/service/medservice"
	"yir/gateway/repository"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func InitRouter(
	authSrv *service.AuthService,
	medSrv *medservice.MedService,
	uziSrv *service.UziService,
	s3Srv *repository.S3Repo,
	mdlware *middleware.AuthMiddleware,
	kafkaProd *repository.Producer,
) *mux.Router {
	ac := &auth.AuthController{
		Service:    authSrv,
		Middleware: mdlware,
	}
	mc := &med.MedController{
		Service: medSrv,
	}
	uzic := &uzi.UziController{
		Service: uziSrv,
		S3:      s3Srv,
		Kafka:   kafkaProd,
	}
	MainRouter := mux.NewRouter()
	MainRouter.HandleFunc("/healthcheck", Healthcheck).Methods("GET")
	authSubR := MainRouter.PathPrefix(authPrefix).Subrouter()
	authSubR.HandleFunc(authLogin, ac.Login).Methods("POST")                                                                // токены не нужны
	authSubR.HandleFunc(authRegister, ac.Register).Methods("POST")                                                          // токены не нужны
	authSubR.HandleFunc(authTokenRefresh, mdlware.VerifyToken(http.HandlerFunc(ac.TokenRefresh)).ServeHTTP).Methods("POST") // пздц слов нет

	medSubR := MainRouter.PathPrefix(medPrefix).Subrouter()
	medSubR.Use(ac.Middleware.VerifyToken)
	patientSubR := medSubR.PathPrefix(patientPrefix).Subrouter()
	patientSubR.HandleFunc(patientCreate, mc.PostPatient).Methods("POST")
	patientSubR.HandleFunc(patientInfo, mc.GetPatientInfo).Methods("GET")
	patientSubR.HandleFunc(patientList, mc.GetPatientList).Methods("GET")
	//
	// medPatientSubR.HandleFunc(patientShots, mc.GetPatientShots) unimplemented in med patient service
	//
	patientSubR.HandleFunc(patientUpdate, mc.PutPatient).Methods("PUT")

	workerSubR := medSubR.PathPrefix(workerPrefix).Subrouter()
	workerSubR.HandleFunc(workerAdd, mc.PostWorker).Methods("POST")
	workerSubR.HandleFunc(workerID, mc.GetWorkerID).Methods("GET")
	workerSubR.HandleFunc(workerList, mc.GetWorkersList).Methods("GET")
	workerSubR.HandleFunc(workerPatients, mc.GetWorkerPatients).Methods("GET")
	workerSubR.HandleFunc(workerUpdate, mc.PutWorker).Methods("PUT")

	cardSubR := medSubR.PathPrefix(cardPrefix).Subrouter()
	cardSubR.HandleFunc(cardPrefix, mc.GetCards).Methods("GET")
	cardSubR.HandleFunc(cardID, mc.PostCard).Methods("POST")
	cardSubR.HandleFunc(cardID, mc.GetCardByID).Methods("GET")
	cardSubR.HandleFunc(cardDelete, mc.DeleteCard).Methods("DELETE")
	cardSubR.HandleFunc(cardID, mc.PutCard).Methods("PUT")

	uziSubR := MainRouter.PathPrefix("/uzi").Subrouter()
	uziSubR.Use(ac.Middleware.VerifyToken)
	uziSubR.HandleFunc(uziDeviceList, uzic.GetDeviceList).Methods("GET")
	uziSubR.HandleFunc(uziFormationSegmentsFormationID, uzic.GetFormationWithSegments).Methods("GET")
	uziSubR.HandleFunc(uziFormationSegmentsPostByUziID, uzic.PostFormationWithSegments).Methods("POST")
	uziSubR.HandleFunc(uziFormationID, uzic.UpdateFormation).Methods("PUT")             // PATCH???
	uziSubR.HandleFunc(uziImageSegmentsID, uzic.GetUziImageWithSegments).Methods("GET") //
	uziSubR.HandleFunc(uziID, uzic.GetUziByID).Methods("GET")
	//uziSubR.HandleFunc(uziPost, uzic.PostUzi).Methods("POST") // Post Uzi нету
	uziSubR.HandleFunc(uziInfo, uzic.PostUzi).Methods("POST")
	uziSubR.HandleFunc(uziInfoID, uzic.GetUziByID).Methods("GET")
	uziSubR.HandleFunc(uziInfo, uzic.UpdateUzi).Methods("PUT")
	return MainRouter
}

func Healthcheck(w http.ResponseWriter, req *http.Request) {
	custom.Logger.Info("Healthcheck",
		zap.String("client_ip", req.RemoteAddr),
	)
	w.Write([]byte("Hello!"))
}

package controller

const (
	authPrefix       = "/auth"
	authLogin        = "/login"
	authRegister     = "/register"
	authTokenRefresh = "/token/refresh"

	medPrefix        = "/med"
	medPatientPrefix = "/patient"
	patientCreate    = "/create"
	patientInfo      = "/info"
	patientList      = "/list"
	patientShots     = "/shots"
	patientUpdate    = "/update"

	medWorkerPrefix = "/medworkers" // надо бы вовану сказать чтобы переделал пути под worker
	workerAdd       = "/add"
	workerID        = "/id/{id}"
	workerList      = "/list"
	workerPatients  = "/patients/{medWorkerId}"
	workerUpdate    = "/update/{id}" // есть методы PUT и PATCH, но пока у нас только PUT

	uziPrefix                         = "/uzi"
	uziDeviceList                     = "/device/list"
	uziFormationSegments_formation_id = "/formation/segments/{formation_id}"
	uziFormationSegments_uzi_id       = "/formation/segments/{uzi_id}"
	uziFormation                      = "/uzi/formation/{formation_id}"
	uziImageFormation                 = "/uzi/image/segments/{image_id}"
	uziInfo                           = "/uzi/info/{uzi_id}"
	uziID                             = "/uzi/{uzi_id}"
)

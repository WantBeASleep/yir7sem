package controller

const (
	authPrefix       = "/auth"
	authLogin        = "/login"
	authRegister     = "/register"
	authTokenRefresh = "/token/refresh"

	medPrefix     = "/med"
	patientPrefix = "/patient"
	patientCreate = "/create"
	patientInfo   = "/info"
	patientList   = "/list"
	patientShots  = "/shots"
	patientUpdate = "/update"

	workerPrefix   = "/medworkers" // надо бы вовану сказать чтобы переделал пути под worker
	workerAdd      = "/add"
	workerID       = "/id/{id}"
	workerList     = "/list"
	workerPatients = "/patients/{medWorkerId}"
	workerUpdate   = "/update/{id}" // есть методы PUT и PATCH, но пока у нас только PUT

	cardPrefix = "/cards"
	cardID     = "/{id}"

	uziPrefix                         = "/uzi"
	uziDeviceList                     = "/device/list"
	uziFormationSegments_formation_id = "/formation/segments/{formation_id}"
	uziFormationSegmentsUziID         = "/formation/segments/{uzi_id}"
	uziFormation                      = "/uzi/formation/{formation_id}"
	uziImageID                        = "/uzi/image/segments/{image_id}"
	uziInfo                           = "/uzi/info/{uzi_id}"
	uziID                             = "/uzi/{uzi_id}"
)

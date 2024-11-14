package controller

const (
	authPrefix = "/auth"
	//
	authLogin        = "/login"
	authRegister     = "/register"
	authTokenRefresh = "/token/refresh"

	medPrefix = "/med"
	//
	patientPrefix = "/patient"
	//
	patientCreate = "/create"
	patientInfo   = "/info/{id}"
	patientList   = "/list"
	patientShots  = "/shots/{id}"
	patientUpdate = "/update/{id}"

	workerPrefix = "/medworkers" // надо бы вовану сказать чтобы переделал пути под worker
	//
	workerAdd      = "/add"
	workerID       = "/{id}"
	workerList     = "/list"
	workerPatients = "/patients/{medWorkerId}"
	workerUpdate   = "/update/{id}" // есть методы PUT и PATCH, но пока у нас только PUT

	cardPrefix = "/card"
	//
	cardByID     = "/{id}"
	cardAdd    = "/add"
	cardList   = "/list"
	cardUpdate = "/update"
	cardDelete = "/delete/{id}"

	uziPrefix = "/uzi"
	//
	uziDeviceList                   = "/device/list"                       //GET
	uziFormationSegmentsFormationID = "/formation/segments/{formation_id}" // GET
	uziFormationSegmentsPostByUziID = "/formation/segments/{uzi_id}"       // POST
	uziFormationID                  = "/uzi/formation/{formation_id}"      // PUT/PATCH
	uziImageSegmentsID              = "/uzi/image/segments/{image_id}"     // GET
	uziInfo                         = "/uzi/info"                          // POST
	uziInfoID                       = "/uzi/info/{uzi_id}"                 // GET, PUT/PATCH
	uziPost                         = "/uzi/add"                           // POST
	uziID                           = "/uzi/{uzi_id}"                      // GET
)

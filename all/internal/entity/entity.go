package entity

type PatientCard struct {
	ID              uint64
	AppointmentTime string
	HasNodules      bool
	Diagnosis       string
	MedWorkerID     uint64
	PatientID       uint64
}

type MedicalWorker struct {
	ID              uint64
	FirstName       string
	MiddleName      string
	LastName        string
	MedOrganization string
	Job             string
	IsRemoteWorker  bool
	ExpertDetails   string
}

type Patient struct {
	ID            uint64
	FirstName     string
	LastName      string
	FatherName    string
	MedicalPolicy string
	Email         string
	IsActive      bool
}

type PatientInformation struct {
	Patient *Patient
	Card    *PatientCard
}

type PatientCardList struct {
	Cards []PatientInformation
	Count int //Кол-во карт
}

type MedicalWorkerUpdateRequest struct {
	FirstName       string
	LastName        string
	MiddleName      string
	MedOrganization string
	Job             string
	IsRemoteWorker  bool
	ExpertDetails   string
}

// Указатели для передачи тех полей, которые нужно обновить
type MedicalWorkerPartialUpdateRequest struct {
	FirstName       *string
	LastName        *string
	MiddleName      *string
	MedOrganization *string
	Job             *string
	IsRemoteWorker  *bool
	ExpertDetails   *string
}

type AddMedicalWorkerRequest struct {
	FirstName       string
	MiddleName      string
	LastName        string
	MedOrganization string
	Job             string
	IsRemoteWorker  bool
	ExpertDetails   string
}

type MedicalWorkerList struct {
	Workers []MedicalWorker
	Count   int //Кол-во работников
}

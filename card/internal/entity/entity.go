package entity

type PatientCard struct {
	ID              int
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
	Patient       *Patient
	Card          *PatientCard
	MedicalWorker *MedicalWorker
}

type PatientCardList struct {
	Cards []PatientCard
	Count int //Кол-во карт
}

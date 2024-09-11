package entity

type MedicalWorker struct {
	ID              int
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

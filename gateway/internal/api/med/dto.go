package med

import (
	"time"

	"github.com/google/uuid"
)

type Doctor struct {
	Id       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Org      string    `json:"org"`
	Job      string    `json:"job"`
	Desc     *string   `json:"desc"`
}

type Patient struct {
	Id          uuid.UUID  `json:"id"`
	FullName    string     `json:"full_name"`
	Email       string     `json:"email"`
	Policy      string     `json:"policy"`
	Active      bool       `json:"active"`
	Malignancy  bool       `json:"malignancy"`
	LastUziDate *time.Time `json:"last_uzi_date"`
}

type Card struct {
	DoctorID  uuid.UUID `json:"doctor_id"`
	PatientID uuid.UUID `json:"patient_id"`
	Diagnosis *string   `json:"diagnosis"`
}

type GetDoctorIn struct{}

type GetDoctorOut struct {
	Doctor
}

type UpdateDoctorIn struct {
	Org  *string `json:"org"`
	Job  *string `json:"job"`
	Desc *string `json:"desc"`
}

type UpdateDoctorOut struct {
	Doctor
}

type GetDoctorPatientsIn struct{}

type GetDoctorPatientsOut struct {
	Patients []Patient `json:"patients"`
}

type PostPatientIn struct {
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	Policy     string `json:"policy"`
	Active     bool   `json:"active"`
	Malignancy bool   `json:"malignancy"`
}

type PostPatientOut struct {
	Id uuid.UUID `json:"id"`
}

type GetPatientIn struct{}

type GetPatientOut struct {
	Patient
}

type UpdatePatientIn struct {
	Active      *bool      `json:"active"`
	Malignancy  *bool      `json:"malignancy"`
	LastUziDate *time.Time `json:"last_uzi_date"`
}

type UpdatePatientOut struct {
	Patient
}

type PostCardIn struct {
	PatientID uuid.UUID `json:"patient_id"`
	Diagnosis *string   `json:"diagnosis"`
}

type PostCardOut struct {
	DoctorID  uuid.UUID `json:"doctor_id"`
	PatientID uuid.UUID `json:"patient_id"`
}

type GetCardIn struct{}

type GetCardOut struct {
	Card
}

type UpdateCardIn struct {
	Diagnosis *string `json:"diagnosis"`
}

type UpdateCardOut struct {
	Card
}

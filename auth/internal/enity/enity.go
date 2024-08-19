package enity

type RequestLogin struct {
	Email string
	Password string
}

type DomainLogin struct {
	Email string
	PasswordHash string
}

type DomainUser struct {
	ID uint 
	Login string 
	PasswordHash string
	RefreshToken string
	MedWorkerID uint
}
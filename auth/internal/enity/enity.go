package enity

// делал из расчета: при реквесте и юзкейсках содержимое меняется?
// да (pass -> passhash), 2 структуры
// нет (token pair ac rt), 1 структура

type RequestLogin struct {
	Email    string
	Password string
}

type DomainLogin struct {
	Email        string
	PasswordHash string
}

type TokensPair struct {
	AccessToken  string
	RefreshToken string
}

type RequestRegister struct {
	Email       string
	LastName    string
	FirstName   string
	FathersName string
	MedOrg      string
	Password    string
}

type User struct {
	ID               int
	Login            string
	PasswordHash     string
	RefreshTokenWord string
	MedWorkerID      int
}

type UserTokenVerify struct {
	UserID           int
	RefreshTokenWord string
}

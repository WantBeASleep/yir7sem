package med

// TODO: uuid
type doctor struct {
	Id       string  `json:"id"`
	Fullname *string `json:"fullname"`
	Org      *string `json:"org"`
	Job      *string `json:"job"`
	Desc     *string `json:"desc"`
}

type getDoctorIn struct{}

type getDoctorOut struct {
	doctor
}

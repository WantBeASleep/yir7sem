package uzi

type UziUseCase struct {
	Service *UziService
}

func New(service *UziService) *UziUseCase {
	return &UziUseCase{
		Service: service,
	}
}

func (uuc *UziUseCase) PostUzi()                {}
func (uuc *UziUseCase) GetDeviceList()          {}
func (uuc *UziUseCase) GetFormationID()         {}
func (uuc *UziUseCase) GetFormationSegmID()     {}
func (uuc *UziUseCase) PostFormationSegmUziID() {}
func (uuc *UziUseCase) GetImageID()             {}
func (uuc *UziUseCase) GetUziInfo()             {}
func (uuc *UziUseCase) GetUziID()               {}

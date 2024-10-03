// https://popovza.kaiten.ru/space/420777/card/37360398
package services

import "fmt"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) AddMed() error {
	return fmt.Errorf("not implemented")
}

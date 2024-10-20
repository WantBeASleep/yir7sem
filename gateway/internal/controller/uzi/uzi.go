package uzi

import (
	"net/http"
	"yir/gateway/internal/usecase/uzi"
)

type UziController struct {
	uziuc *uzi.UziUseCase
}

func (uzic *UziController) PostUzi(w http.ResponseWriter, r *http.Request)                {}
func (uzic *UziController) GetDeviceList(w http.ResponseWriter, r *http.Request)          {}
func (uzic *UziController) GetFormationID(w http.ResponseWriter, r *http.Request)         {}
func (uzic *UziController) GetFormationSegmID(w http.ResponseWriter, r *http.Request)     {}
func (uzic *UziController) PostFormationSegmUziID(w http.ResponseWriter, r *http.Request) {}
func (uzic *UziController) GetImageID(w http.ResponseWriter, r *http.Request)             {}
func (uzic *UziController) GetUziInfo(w http.ResponseWriter, r *http.Request)             {}
func (uzic *UziController) GetUziID(w http.ResponseWriter, r *http.Request)               {}

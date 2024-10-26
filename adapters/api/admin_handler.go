package api

import (
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type AdminHandler struct {
	AdminUsecase usecases.AdminUsecase
}

func InitiateAdminHandler(usecase usecases.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		AdminUsecase: usecase,
	}
}

func (a *AdminHandler) InitializeAdmin() error {
	return a.AdminUsecase.InitializeAdmin()
}
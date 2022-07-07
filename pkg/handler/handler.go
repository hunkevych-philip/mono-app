package handler

import (
	"github.com/hunkevych-philip/mono-app/pkg/service"
)

type Handler struct {
	services *service.ServicesImpl
}

func NewHandler(services *service.ServicesImpl) *Handler {
	return &Handler{
		services: services,
	}
}

// Go performs high-level validations and calls Mono service
// to get a user's statement, then calls Excel service statement data to an Excel file
func (h *Handler) Go(xToken, account, fromStr string) error {
	statement, err := h.services.Mono.GetStatement(xToken, account, fromStr)
	if err != nil {
		return err
	}

	err = h.services.Excel.GenerateSheetForStatement(statement)
	if err != nil {
		return err
	}

	return nil
}

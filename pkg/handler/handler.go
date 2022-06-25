package handler

import (
	"github.com/hunkevych-philip/mono-app/pkg/service"
	"github.com/hunkevych-philip/mono-app/pkg/utils"
)

type Handler struct {
	services  *service.ServicesImpl
	utilities *utils.UtilsImpl
}

func NewHandler(services *service.ServicesImpl, utils *utils.UtilsImpl) *Handler {
	return &Handler{
		services:  services,
		utilities: utils,
	}
}

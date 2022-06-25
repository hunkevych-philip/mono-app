package handler

import (
	"github.com/HunkevychPhilip/todo/pkg/service"
	"github.com/HunkevychPhilip/todo/pkg/utils"
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
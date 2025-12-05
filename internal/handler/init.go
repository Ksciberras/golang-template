package handler

import (
	"golang-project-template/internal/config"
	"golang-project-template/internal/service"
)

type Handler struct {
	cfg *config.Config
	svc *service.Service
}

func Init(cfg *config.Config, svc *service.Service) *Handler {
	return &Handler{
		cfg,
		svc,
	}
}

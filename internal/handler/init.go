package handler

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

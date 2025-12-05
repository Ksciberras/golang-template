package service

import "golang-project-template/internal/config"

type Service struct {
	cfg *config.Config
}

func Init(cfg *config.Config) *Service {

	return &Service{
		cfg,
	}
}

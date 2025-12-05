package service

type Service struct {
	cfg *config.Config
}

func Init(cfg *config.Config) *Service {

	return &Service{
		cfg,
	}
}

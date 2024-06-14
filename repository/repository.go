package repository

import (
	"gRPC-basic/config"
)

type Repository struct {
	cfg *config.Config
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	r := &Repository{cfg: cfg}

	return r, nil
}

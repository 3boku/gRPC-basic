package network

import (
	"gRPC-basic/config"
	"gRPC-basic/service"
)

type Network struct {
	cfg *config.Config

	service *service.Service
}

func NewNetwork(cfg *config.Config, service *service.Service) (*Network, error) {
	r := &Network{cfg: cfg, service: service}

	return r, nil
}

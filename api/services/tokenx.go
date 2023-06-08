package services

import (
	"github.com/dvwzj/iam48-go/api/domain/tokenx"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type TokenXService struct {
	Client *resty.Client
	tokenx.TokenXRepository
}

type TokenXServiceConfiguration func(*TokenXService) error

func NewTokenXService(cfgs ...TokenXServiceConfiguration) (*TokenXService, error) {
	client := client.NewClient()
	tokenx, err := tokenx.New(tokenx.WithClient(client))
	if err != nil {
		return nil, err
	}
	service := &TokenXService{
		Client:           client,
		TokenXRepository: tokenx,
	}
	for _, cfg := range cfgs {
		if err := cfg(service); err != nil {
			return nil, err
		}
	}
	return service, nil
}

package services

import (
	"github.com/dvwzj/iam48-go/api/domain/pdpa"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type PdpaService struct {
	Client *resty.Client
	pdpa.PdpaRepository
}

type PdpaServiceConfiguration func(*PdpaService) error

func NewPdpaService(cfgs ...PdpaServiceConfiguration) (*PdpaService, error) {
	client := client.NewClient()
	pdpa, err := pdpa.New(pdpa.WithClient(client))
	if err != nil {
		return nil, err
	}
	service := &PdpaService{
		Client:         client,
		PdpaRepository: pdpa,
	}
	for _, cfg := range cfgs {
		if err := cfg(service); err != nil {
			return nil, err
		}
	}
	return service, nil
}

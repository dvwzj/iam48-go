package services

import (
	"github.com/dvwzj/iam48-go/api/domain/public"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type PublicService struct {
	Client *resty.Client
	public.PublicRepository
}

type PublicServiceConfiguration func(*PublicService) error

func NewPublicService(cfgs ...PublicServiceConfiguration) (*PublicService, error) {
	client := client.NewClient()
	public, err := public.New(public.WithClient(client))
	if err != nil {
		return nil, err
	}
	service := &PublicService{
		Client:           client,
		PublicRepository: public,
	}
	for _, cfg := range cfgs {
		if err := cfg(service); err != nil {
			return nil, err
		}
	}
	return service, nil
}

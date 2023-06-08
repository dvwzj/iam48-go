package services

import (
	"github.com/dvwzj/iam48-go/api/domain/theater"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type TheaterService struct {
	Client *resty.Client
	theater.TheaterRepository
}

type TheaterServiceConfiguration func(*TheaterService) error

func NewTheaterService(cfgs ...TheaterServiceConfiguration) (*TheaterService, error) {
	client := client.NewClient()
	theater, err := theater.New(theater.WithClient(client))
	if err != nil {
		return nil, err
	}
	service := &TheaterService{
		Client:            client,
		TheaterRepository: theater,
	}
	for _, cfg := range cfgs {
		if err := cfg(service); err != nil {
			return nil, err
		}
	}
	return service, nil
}

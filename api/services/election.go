package services

import (
	"github.com/dvwzj/iam48-go/api/domain/election"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type ElectionService struct {
	Client *resty.Client
	election.ElectionRepository
}

type ElectionServiceConfiguration func(*ElectionService) error

func NewElectionService(cfgs ...ElectionServiceConfiguration) (*ElectionService, error) {
	client := client.NewClient()
	election, err := election.New(election.WithClient(client))
	if err != nil {
		return nil, err
	}
	service := &ElectionService{
		Client:             client,
		ElectionRepository: election,
	}
	for _, cfg := range cfgs {
		if err := cfg(service); err != nil {
			return nil, err
		}
	}
	return service, nil
}

package services

import (
	userbalance "github.com/dvwzj/iam48-go/api/domain/user-balance"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type UserBalanceService struct {
	Client *resty.Client
	userbalance.UserBalanceRepository
}

type UserBalanceServiceConfiguration func(*UserBalanceService) error

func NewUserBalanceService(cfgs ...UserBalanceServiceConfiguration) (*UserBalanceService, error) {
	client := client.NewClient()
	userbalance, err := userbalance.New(userbalance.WithClient(client))
	if err != nil {
		return nil, err
	}
	service := &UserBalanceService{
		Client:                client,
		UserBalanceRepository: userbalance,
	}
	for _, cfg := range cfgs {
		if err := cfg(service); err != nil {
			return nil, err
		}
	}
	return service, nil
}

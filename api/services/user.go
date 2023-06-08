package services

import (
	"github.com/dvwzj/iam48-go/api/domain/user"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type UserService struct {
	Client *resty.Client
	user.UserRepository
}

type UserServiceConfiguration func(*UserService) error

func NewUserService(cfgs ...UserServiceConfiguration) (*UserService, error) {
	client := client.NewClient()
	user, err := user.New(user.WithClient(client))
	if err != nil {
		return nil, err
	}
	service := &UserService{
		Client:         client,
		UserRepository: user,
	}
	for _, cfg := range cfgs {
		if err := cfg(service); err != nil {
			return nil, err
		}
	}
	return service, nil
}

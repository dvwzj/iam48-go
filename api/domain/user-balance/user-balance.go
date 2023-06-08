package userbalance

import (
	"strconv"

	"github.com/dvwzj/iam48-go/api/entity"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type UserBalance struct {
	Client *resty.Client
}

func (s UserBalance) GetBalance(appCode string, userId int) (entity.CoreBalance, error) {
	var coreBalance entity.CoreBalance
	_, err := s.Client.R().
		SetResult(&coreBalance).
		SetPathParam("appCode", appCode).
		SetPathParam("userId", strconv.Itoa(userId)).
		Get("/app/{appCode}/user/{userId}/balance")
	if err != nil {
		return entity.CoreBalance{}, err
	}
	return coreBalance, nil
}

type UserBalanceConfiguration func(*UserBalance) error

func WithClient(client *resty.Client) UserBalanceConfiguration {
	return func(e *UserBalance) error {
		e.Client = client
		e.Client.SetBaseURL("https://coin-userbalance.bnk48.io")
		return nil
	}
}

func New(cfgs ...UserBalanceConfiguration) (UserBalance, error) {
	service := UserBalance{
		Client: client.NewClient().SetBaseURL("https://coin-userbalance.bnk48.io"),
	}
	for _, cfg := range cfgs {
		if err := cfg(&service); err != nil {
			return UserBalance{}, err
		}
	}
	return service, nil
}

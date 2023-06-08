package tokenx

import (
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type TokenX struct {
	Client *resty.Client
}

func (s TokenX) AccessDenied() (resty.Response, error) {
	res, err := s.Client.R().
		Get("/access-denied")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

type TokenXConfiguration func(*TokenX) error

func WithClient(client *resty.Client) TokenXConfiguration {
	return func(e *TokenX) error {
		e.Client = client
		e.Client.SetBaseURL("https://beta.tokenx-client2.com")
		return nil
	}
}

func New(cfgs ...TokenXConfiguration) (TokenX, error) {
	service := TokenX{
		Client: client.NewClient().SetBaseURL("https://beta.tokenx-client2.com"),
	}
	for _, cfg := range cfgs {
		if err := cfg(&service); err != nil {
			return TokenX{}, err
		}
	}
	return service, nil
}

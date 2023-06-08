package services

import (
	"github.com/dvwzj/iam48-go/api/domain/streaming"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type StreamingService struct {
	Client *resty.Client
	streaming.StreamingRepository
}

type StreamingServiceConfiguration func(*StreamingService) error

func NewStreamingService(cfgs ...StreamingServiceConfiguration) (*StreamingService, error) {
	client := client.NewClient()
	streaming, err := streaming.New(streaming.WithClient(client))
	if err != nil {
		return nil, err
	}
	service := &StreamingService{
		Client:              client,
		StreamingRepository: streaming,
	}
	for _, cfg := range cfgs {
		if err := cfg(service); err != nil {
			return nil, err
		}
	}
	return service, nil
}

package services

import (
	"github.com/dvwzj/iam48-go/api/domain/ticket"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type TicketService struct {
	Client *resty.Client
	ticket.TicketRepository
}

type TicketServiceConfiguration func(*TicketService) error

func NewTicketService(cfgs ...TicketServiceConfiguration) (*TicketService, error) {
	client := client.NewClient()
	ticket, err := ticket.New(ticket.WithClient(client))
	if err != nil {
		return nil, err
	}
	service := &TicketService{
		Client:           client,
		TicketRepository: ticket,
	}
	for _, cfg := range cfgs {
		if err := cfg(service); err != nil {
			return nil, err
		}
	}
	return service, nil
}

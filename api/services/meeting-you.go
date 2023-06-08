package services

import (
	meetingyou "github.com/dvwzj/iam48-go/api/domain/meeting-you"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type MeetingYouService struct {
	Client *resty.Client
	meetingyou.MeetingYouRepository
}

type MeetingYouServiceConfiguration func(*MeetingYouService) error

func NewMeetingYouService(cfgs ...MeetingYouServiceConfiguration) (*MeetingYouService, error) {
	client := client.NewClient()
	meetingyou, err := meetingyou.New(meetingyou.WithClient(client))
	if err != nil {
		return nil, err
	}
	service := &MeetingYouService{
		Client:               client,
		MeetingYouRepository: meetingyou,
	}
	for _, cfg := range cfgs {
		if err := cfg(service); err != nil {
			return nil, err
		}
	}
	return service, nil
}

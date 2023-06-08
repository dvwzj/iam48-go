package meetingyou

import (
	"strconv"

	"github.com/dvwzj/iam48-go/api/entity"
	"github.com/go-resty/resty/v2"
)

type MeetingYou struct {
	Client *resty.Client
}

func (s MeetingYou) JoinLobby(memberId int64) (entity.JoinLobbyInfo, error) {
	var joinLobbyInfo entity.JoinLobbyInfo
	_, err := s.Client.R().
		SetResult(&joinLobbyInfo).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Post("/channel/user/join/{memberId}")
	if err != nil {
		return entity.JoinLobbyInfo{}, err
	}
	return joinLobbyInfo, nil
}

func (s MeetingYou) SubmitReview(meetingReviewRequest entity.MeetingReviewRequest) (resty.Response, error) {
	res, err := s.Client.R().
		SetBody(meetingReviewRequest).
		Post("/review/submit")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

type MeetingYouConfiguration func(*MeetingYou) error

func WithClient(client *resty.Client) MeetingYouConfiguration {
	return func(e *MeetingYou) error {
		e.Client = client
		return nil
	}
}

func New(cfgs ...MeetingYouConfiguration) (MeetingYou, error) {
	service := MeetingYou{}
	for _, cfg := range cfgs {
		if err := cfg(&service); err != nil {
			return MeetingYou{}, err
		}
	}
	return service, nil
}

package theater

import (
	"strconv"

	"github.com/dvwzj/iam48-go/api/entity"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type Theater struct {
	Client *resty.Client
}

func (s Theater) GetStream(liveId int64, streamId int64) (entity.LiveStreamInfo, error) {
	var liveStreamInfo entity.LiveStreamInfo
	_, err := s.Client.R().
		SetResult(&liveStreamInfo).
		SetPathParam("liveId", strconv.FormatInt(liveId, 10)).
		SetPathParam("streamId", strconv.FormatInt(streamId, 10)).
		Get("/theater-live/{liveId}/stream/{streamId}")
	if err != nil {
		return entity.LiveStreamInfo{}, err
	}
	return liveStreamInfo, nil
}

func (s Theater) GetTheaterWatch(liveId int) (entity.TheaterWatchInfo, error) {
	var theaterWatchInfo entity.TheaterWatchInfo
	_, err := s.Client.R().
		SetResult(&theaterWatchInfo).
		SetPathParam("liveId", strconv.Itoa(liveId)).
		Get("/live/{liveId}/watch")
	if err != nil {
		return entity.TheaterWatchInfo{}, err
	}
	return theaterWatchInfo, nil
}

func (s Theater) SendTheaterGiftAllMembers(userId int64, liveId int64) (entity.TheaterGiftInfo, error) {
	var theaterGiftInfo entity.TheaterGiftInfo
	_, err := s.Client.R().
		SetResult(&theaterGiftInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("liveId", strconv.FormatInt(liveId, 10)).
		Post("/user/{userId}/theater-live/{liveId}/members/all")
	if err != nil {
		return entity.TheaterGiftInfo{}, err
	}
	return theaterGiftInfo, nil
}

func (s Theater) SendTheaterGiftChooseMembers(userId int64, liveId int64) (entity.TheaterGiftChooseMembersInfo, error) {
	var theaterGiftChooseMembersInfo entity.TheaterGiftChooseMembersInfo
	_, err := s.Client.R().
		SetResult(&theaterGiftChooseMembersInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("liveId", strconv.FormatInt(liveId, 10)).
		Post("/user/{userId}/theater-live/{liveId}/members")
	if err != nil {
		return entity.TheaterGiftChooseMembersInfo{}, err
	}
	return theaterGiftChooseMembersInfo, nil
}

type TheaterConfiguration func(*Theater) error

func WithClient(client *resty.Client) TheaterConfiguration {
	return func(e *Theater) error {
		e.Client = client
		e.Client.SetBaseURL("https://theater-live-api.bnk48.io")
		return nil
	}
}

func New(cfgs ...TheaterConfiguration) (Theater, error) {
	service := Theater{
		Client: client.NewClient().SetBaseURL("https://theater-live-api.bnk48.io"),
	}
	for _, cfg := range cfgs {
		if err := cfg(&service); err != nil {
			return Theater{}, err
		}
	}
	return service, nil
}

package streaming

import (
	"strconv"

	"github.com/dvwzj/iam48-go/api/entity"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type Streaming struct {
	Client *resty.Client
}

func (s Streaming) FloatAndGiftCommit(userId int64, transactionId int64) (entity.FloatMessageInfo, error) {
	var floatMessageInfo entity.FloatMessageInfo
	_, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("transactionId", strconv.FormatInt(transactionId, 10)).
		SetResult(&floatMessageInfo).
		Post("/user/{userId}/transaction/{transactionId}/commit")
	if err != nil {
		return entity.FloatMessageInfo{}, err
	}
	return floatMessageInfo, nil
}

func (s Streaming) FloatMessageInit(userId int64, liveId int64, floatMessageChatBody entity.FloatMessageChatBody) (entity.RequestGiftInfo, error) {
	var requestGiftInfo entity.RequestGiftInfo
	_, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("liveId", strconv.FormatInt(liveId, 10)).
		SetBody(floatMessageChatBody).
		SetResult(&requestGiftInfo).
		Post("/user/{userId}/floatmsg/member-live/{liveId}/request")
	if err != nil {
		return entity.RequestGiftInfo{}, err
	}
	return requestGiftInfo, nil
}

func (s Streaming) GetLiveUrlInfo(liveId int64) (entity.LiveUrlInfo, error) {
	var liveUrlInfo entity.LiveUrlInfo
	_, err := s.Client.R().
		SetPathParam("liveId", strconv.FormatInt(liveId, 10)).
		SetResult(&liveUrlInfo).
		Get("/live/{liveId}")
	if err != nil {
		return entity.LiveUrlInfo{}, err
	}
	return liveUrlInfo, nil
}

func (s Streaming) GetMemberLiveInfo(userId string, liveId string) (entity.LiveInfoData, error) {
	var liveInfoData entity.LiveInfoData
	_, err := s.Client.R().
		SetPathParam("userId", userId).
		SetPathParam("liveId", liveId).
		SetResult(&liveInfoData).
		Get("/user/{userId}/info/member-live/{liveId}")
	if err != nil {
		return entity.LiveInfoData{}, err
	}
	return liveInfoData, nil
}

func (s Streaming) GetMemberLiveUrlInfo(userId int64, liveId string) (entity.LiveUrlInfo, error) {
	var liveUrlInfo entity.LiveUrlInfo
	_, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("liveId", liveId).
		SetResult(&liveUrlInfo).
		Get("/user/{userId}/watch/member-live/{liveId}")
	if err != nil {
		return entity.LiveUrlInfo{}, err
	}
	return liveUrlInfo, nil
}

func (s Streaming) GiftInit(userId int64, liveId int64, giftBody entity.GiftBody) (entity.RequestGiftInfo, error) {
	var requestGiftInfo entity.RequestGiftInfo
	_, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("liveId", strconv.FormatInt(liveId, 10)).
		SetBody(giftBody).
		SetResult(&requestGiftInfo).
		Post("/user/{userId}/gifts/member-live/{liveId}/request")
	if err != nil {
		return entity.RequestGiftInfo{}, err
	}
	return requestGiftInfo, nil
}

func (s Streaming) SendMemberLiveGift(userId int64, liveId string, giftBody entity.GiftBody) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("liveId", liveId).
		SetBody(giftBody).
		Post("/user/{userId}/gifts/member-live/{liveId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s Streaming) UpdateUserFavorite(userId string, liveId string) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", userId).
		SetPathParam("liveId", liveId).
		Put("/user/{userId}/favorite/member-live/{liveId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s Streaming) UpdateUserUnFavorite(userId string, liveId string) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", userId).
		SetPathParam("liveId", liveId).
		Delete("/user/{userId}/favorite/member-live/{liveId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

type StreamingConfiguration func(*Streaming) error

func WithClient(client *resty.Client) StreamingConfiguration {
	return func(e *Streaming) error {
		e.Client = client
		e.Client.SetBaseURL("https://live-api.bnk48.io")
		return nil
	}
}

func New(cfgs ...StreamingConfiguration) (Streaming, error) {
	service := Streaming{
		Client: client.NewClient().SetBaseURL("https://live-api.bnk48.io"),
	}
	for _, cfg := range cfgs {
		if err := cfg(&service); err != nil {
			return Streaming{}, err
		}
	}
	return service, nil
}

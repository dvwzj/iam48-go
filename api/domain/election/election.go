package election

import (
	"strconv"

	"github.com/dvwzj/iam48-go/api/entity"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type Election struct {
	Client *resty.Client
}

func (s Election) ClaimShopee(electionClaimBodyInfo entity.ElectionClaimBodyInfo) (entity.ElectionClaimResponseInfo, error) {
	var electionClaimResponseInfo entity.ElectionClaimResponseInfo
	_, err := s.Client.R().
		SetBody(electionClaimBodyInfo).
		SetResult(&electionClaimResponseInfo).
		Post("/claim/v2/shopee")
	if err != nil {
		return entity.ElectionClaimResponseInfo{}, err
	}
	return electionClaimResponseInfo, nil
}

func (s Election) GetAllElectionCodeById(codeSetId int64) (entity.ElectionAllCodeSetResponseInfo, error) {
	var electionAllCodeSetResponseInfo entity.ElectionAllCodeSetResponseInfo
	_, err := s.Client.R().
		SetResult(&electionAllCodeSetResponseInfo).
		SetPathParam("codeSetId", strconv.FormatInt(codeSetId, 10)).
		Get("/code/{codeSetId}/all")
	if err != nil {
		return entity.ElectionAllCodeSetResponseInfo{}, err
	}
	return electionAllCodeSetResponseInfo, nil
}

func (s Election) GetElectionCode(skip int, take int) ([]entity.ElectionCodeResponseInfo, error) {
	var electionCodeResponseInfo []entity.ElectionCodeResponseInfo
	_, err := s.Client.R().
		SetResult(&electionCodeResponseInfo).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/code")
	if err != nil {
		return nil, err
	}
	return electionCodeResponseInfo, nil
}

func (s Election) GetElectionCodeById(codeSetId int64, latestCodeId int64) ([]entity.ElectionCodeSetByIdInfo, error) {
	var electionCodeSetByIdInfo []entity.ElectionCodeSetByIdInfo
	_, err := s.Client.R().
		SetResult(&electionCodeSetByIdInfo).
		SetPathParam("codeSetId", strconv.FormatInt(codeSetId, 10)).
		SetQueryParam("latestCodeId", strconv.FormatInt(latestCodeId, 10)).
		Get("/code/{codeSetId}")
	if err != nil {
		return nil, err
	}
	return electionCodeSetByIdInfo, nil
}

func (s Election) GetElectionHistory(skip int, take int) ([]entity.ElectionHistoryResponseInfo, error) {
	var electionHistoryResponseInfo []entity.ElectionHistoryResponseInfo
	_, err := s.Client.R().
		SetResult(&electionHistoryResponseInfo).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/history")
	if err != nil {
		return nil, err
	}
	return electionHistoryResponseInfo, nil
}

func (s Election) SentCodeToEmail(electionSentCodeToEmailBody entity.ElectionSentCodeToEmailBody) (resty.Response, error) {
	res, err := s.Client.R().
		SetBody(electionSentCodeToEmailBody).
		Post("/code/sent")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

type ElectionConfiguration func(*Election) error

func WithClient(client *resty.Client) ElectionConfiguration {
	return func(e *Election) error {
		e.Client = client
		e.Client.SetBaseURL("https://election-pool.bnk48.io")
		return nil
	}
}

func New(cfgs ...ElectionConfiguration) (Election, error) {
	election := Election{
		Client: client.NewClient().SetBaseURL("https://election-pool.bnk48.io"),
	}
	for _, cfg := range cfgs {
		if err := cfg(&election); err != nil {
			return Election{}, err
		}
	}
	return election, nil
}

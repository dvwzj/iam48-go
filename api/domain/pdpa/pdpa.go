package pdpa

import (
	"strconv"

	"github.com/dvwzj/iam48-go/api/entity"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type Pdpa struct {
	Client *resty.Client
}

func (s Pdpa) CheckConsentStatus(appCode string, ookbeeNumericId int64) (entity.PdpaConsentStatusResponseInfo, error) {
	var pdpaConsentStatusResponseInfo entity.PdpaConsentStatusResponseInfo
	_, err := s.Client.R().
		SetResult(&pdpaConsentStatusResponseInfo).
		SetPathParam("appCode", appCode).
		SetPathParam("ookbeeNumericId", strconv.FormatInt(ookbeeNumericId, 10)).
		Get("/{appCode}/consent/{ookbeeNumericId}/check")
	if err != nil {
		return entity.PdpaConsentStatusResponseInfo{}, err
	}
	return pdpaConsentStatusResponseInfo, nil
}

func (s Pdpa) CheckCurrentVersionPDPA(appCode string, ookbeeNumericId int64) (entity.PdpaCurrentVersionResponseInfo, error) {
	var pdpaCurrentVersionResponseInfo entity.PdpaCurrentVersionResponseInfo
	_, err := s.Client.R().
		SetResult(&pdpaCurrentVersionResponseInfo).
		SetPathParam("appCode", appCode).
		SetPathParam("ookbeeNumericId", strconv.FormatInt(ookbeeNumericId, 10)).
		Get("/{appCode}/consent/{ookbeeNumericId}/currentversion")
	if err != nil {
		return entity.PdpaCurrentVersionResponseInfo{}, err
	}
	return pdpaCurrentVersionResponseInfo, nil
}

type PdpaConfiguration func(*Pdpa) error

func WithClient(client *resty.Client) PdpaConfiguration {
	return func(e *Pdpa) error {
		e.Client = client
		e.Client.SetBaseURL("https://privacypolicy.bnk48.com")
		return nil
	}
}

func New(cfgs ...PdpaConfiguration) (Pdpa, error) {
	service := Pdpa{
		Client: client.NewClient().SetBaseURL("https://privacypolicy.bnk48.com"),
	}
	for _, cfg := range cfgs {
		if err := cfg(&service); err != nil {
			return Pdpa{}, err
		}
	}
	return service, nil
}

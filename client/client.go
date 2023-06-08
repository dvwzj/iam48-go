package client

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var (
	Environtment = "PROD"
	UserAgent    = "BNK48_102/1.2.89/Dalvik/2.1.0 (Linux; U; Android 5.1.1; HUAWEI MLA-L12 Build/HUAWEIMLA-L12)"
	AppId        = "BNK48_102"
	AppVersion   = "1.2.22"
	DeviceModel  = "HUAWEI HUAWEI MLA-L12"
	DeviceId     = "6802689067815614"
	DeviceName   = "HUAWEI MLA-L12"
)

var (
	DefaultHeaders = map[string]string{
		"environment":        Environtment,
		"user-agent":         UserAgent,
		"bnk48-appcode":      AppId,
		"bnk48-device-id":    DeviceId,
		"bnk48-device-model": DeviceModel,
	}
)

var (
	DefaultClient = NewClient()
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	if v := os.Getenv("IAM48_ENVIRONMENT"); v != "" {
		Environtment = v
	}
	if v := os.Getenv("IAM48_USER_AGENT"); v != "" {
		UserAgent = v
	}
	if v := os.Getenv("IAM48_APP_ID"); v != "" {
		AppId = v
	}
	if v := os.Getenv("IAM48_APP_VERSION"); v != "" {
		AppVersion = v
	}
	if v := os.Getenv("IAM48_DEVICE_MODEL"); v != "" {
		DeviceModel = v
	}
	if v := os.Getenv("IAM48_DEVICE_ID"); v != "" {
		DeviceId = v
	}
	if v := os.Getenv("IAM48_DEVICE_NAME"); v != "" {
		DeviceName = v
	}
	DefaultHeaders = map[string]string{
		"environment":        Environtment,
		"user-agent":         UserAgent,
		"bnk48-appcode":      AppId,
		"bnk48-device-id":    DeviceId,
		"bnk48-device-model": DeviceModel,
	}
	DefaultClient.SetHeaders(DefaultHeaders)
}

type ErrApiResponse struct {
	ApiVersion string `json:"apiVersion"`
	Error      struct {
		Code        int    `json:"code"`
		ErrorNumber int    `json:"errorNumber"`
		Message     string `json:"message"`
	} `json:"error"`
	TraceId string `json:"traceId"`
}

func NewClient() *resty.Client {
	client := resty.New().SetHeaders(DefaultHeaders)
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		if r.StatusCode() != 200 {
			body := r.Body()
			var errResp ErrApiResponse
			if err := json.Unmarshal(body, &errResp); err == nil {
				return errors.New(errResp.Error.Message)
			}
			if body != nil {
				return errors.New(string(body))
			}
		}
		return nil
	})
	return client
}

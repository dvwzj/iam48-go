package entity

type FacebookLinkPostInfo struct {
	// @SerializedName("accessToken")
	// @Nullable
	// private String accessToken;
	AccessToken *string `json:"accessToken"`
	// @SerializedName("appCode")
	// @Nullable
	// private String appCode;
	AppCode *string `json:"appCode"`
	// @SerializedName("deviceId")
	// @Nullable
	// private String deviceId;
	DeviceId *string `json:"deviceId"`
	// @SerializedName("facebookAppId")
	// @Nullable
	// private String facebookAppId;
	FacebookAppId *string `json:"facebookAppId"`
}

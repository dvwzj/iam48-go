package entity

type AdsItems struct {
	// @SerializedName(GlobalRedeemCodeActivity.KEY_CODE) // "code"
	// @Nullable
	// private String code;
	Code *string `json:"code"`
	// @SerializedName("displayPercentage")
	// @Nullable
	// private AdsDisplayPercentage displayPercentage;
	DisplayPercentage *AdsDisplayPercentage `json:"displayPercentage"`
	// @SerializedName("fileUrl")
	// @Nullable
	// private String fileUrl;
	FileUrl *string `json:"fileUrl"`
	// @SerializedName("imageUrl")
	// @Nullable
	// private String imageUrl;
	ImageUrl *string `json:"imageUrl"`
	// @SerializedName("interval")
	// @Nullable
	// private Long interval;
	Interval *int64 `json:"interval"`
	// @SerializedName("isAutoClose")
	// @Nullable
	// private Boolean isAutoClose;
	IsAutoClose *bool `json:"isAutoClose"`
	// @SerializedName("isDefault")
	// @Nullable
	// private Boolean isDefault;
	IsDefault *bool `json:"isDefault"`
	// @SerializedName("isDisplayCountdown")
	// @Nullable
	// private Boolean isDisplayCountdown;
	IsDisplayCountdown *bool `json:"isDisplayCountdown"`
	// @SerializedName("isUseWebView")
	// @Nullable
	// private Boolean isUseWebView;
	IsUseWebView *bool `json:"isUseWebView"`
	// @SerializedName("isVideo")
	// @Nullable
	// private Boolean isVideo;
	IsVideo *bool `json:"isVideo"`
	// @SerializedName("linkUrl")
	// @Nullable
	// private String linkUrl;
	LinkUrl *string `json:"linkUrl"`
	// @SerializedName("version")
	// @Nullable
	// private String version;
	Version *string `json:"version"`
}

package entity

type TheaterLiveInfo struct {
	// @SerializedName("token")
	// @Nullable
	// private String token;
	Token *string `json:"token"`
	// @SerializedName("watermarkConfig")
	// @Nullable
	// private WatermarkConfigInfo watermarkConfig;
	WatermarkConfig *WatermarkConfigInfo `json:"watermarkConfig"`
}

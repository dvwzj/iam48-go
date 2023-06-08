package entity

type SignedVideoContent struct {
	// @SerializedName("latestPlayInfo")
	// @Nullable
	// private LatestPlayInfo latestPlayInfo;
	LatestPlayInfo *LatestPlayInfo `json:"latestPlayInfo"`
	// @SerializedName("requestId")
	// @Nullable
	// private String requestId;
	RequestId *string `json:"requestId"`
	// @SerializedName("resourceUrl")
	// @Nullable
	// private String resourceUrl;
	ResourceUrl *string `json:"resourceUrl"`
	// @SerializedName("watermarkConfig")
	// @Nullable
	// private WatermarkConfigInfo watermarkConfig;
	WatermarkConfig *WatermarkConfigInfo `json:"watermarkConfig"`
}

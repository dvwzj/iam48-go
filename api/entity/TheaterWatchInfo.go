package entity

type TheaterWatchInfo struct {
	// @SerializedName("configuration")
	// @Nullable
	// private TheaterConfigurationInfo configuration;
	Configuration *TheaterConfigurationInfo `json:"configuration"`
	// @SerializedName("liveId")
	// @Nullable
	// private Long liveId;
	LiveId *int64 `json:"liveId"`
	// @SerializedName("liveStartedAt")
	// @Nullable
	// private String liveStartedAt;
	LiveStartedAt *string `json:"liveStartedAt"`
	// @SerializedName("memberId")
	// @Nullable
	// private List<Long> memberId;
	MemberId *[]int64 `json:"memberId"`
	// @SerializedName("requestId")
	// @Nullable
	// private String requestId;
	RequestId *string `json:"requestId"`
	// @SerializedName("watchCameraInfoList")
	// @Nullable
	// private List<TheaterWatchCameraInfo> watchCameraInfoList;
	WatchCameraInfoList *[]TheaterWatchCameraInfo `json:"watchCameraInfoList"`
}

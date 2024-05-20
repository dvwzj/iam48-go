package entity

type LiveUrlInfo struct {
	// @SerializedName("chatDelaySeconds")
	// @Nullable
	// private Long chatDelaySeconds;
	ChatDelaySeconds *float64 `json:"chatDelaySeconds,omitempty"`
	// @SerializedName("heartDelaySeconds")
	// @Nullable
	// private Long heartDelaySeconds;
	HeartDelaySeconds *float64 `json:"heartDelaySeconds,omitempty"`
	// @SerializedName("hlsUrl")
	// @Nullable
	// private String hlsUrl;
	HlsUrl *string `json:"hlsUrl,omitempty"`
	// @SerializedName("isFavorite")
	// @Nullable
	// private Boolean isFavorite;
	IsFavorite *bool `json:"isFavorite,omitempty"`
	// @SerializedName("liveStartedAt")
	// @Nullable
	// private String liveStartedAt;
	LiveStartedAt *string `json:"liveStartedAt,omitempty"`
	// @SerializedName("requestId")
	// @Nullable
	// private String requestId;
	RequestId *string `json:"requestId,omitempty"`
	// @SerializedName("signalrUrl")
	// @Nullable
	// private String signalrUrl;
	SignalrUrl *string `json:"signalrUrl,omitempty"`
	// @SerializedName("statistics")
	// @NotNull
	// private StatisticsInfo statistics;
	Statistics StatisticsInfo `json:"statistics"`
}

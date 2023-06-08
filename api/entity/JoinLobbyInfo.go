package entity

type JoinLobbyInfo struct {
	// @SerializedName("chatDelaySeconds")
	// @Nullable
	// private Long chatDelaySeconds;
	ChatDelaySeconds *int64 `json:"chatDelaySeconds"`
	// @SerializedName("signalrUrl")
	// @Nullable
	// private String signalrUrl;
	SignalrUrl *string `json:"signalrUrl"`
	// @SerializedName("videoUrls")
	// @Nullable
	// private List<String> videoUrls;
	VideoUrls *[]string `json:"videoUrls"`
}

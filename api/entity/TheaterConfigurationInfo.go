package entity

type TheaterConfigurationInfo struct {
	// @SerializedName("chatDelaySeconds")
	// @Nullable
	// private Long chatDelaySeconds;
	ChatDelaySeconds *int64 `json:"chatDelaySeconds"`
	// @SerializedName("heartDelaySeconds")
	// @Nullable
	// private Long heartDelaySeconds;
	HeartDelaySeconds *int64 `json:"heartDelaySeconds"`
	// @SerializedName("signalRUrl")
	// @Nullable
	// private String signalRUrl;
	SignalRUrl *string `json:"signalRUrl"`
}

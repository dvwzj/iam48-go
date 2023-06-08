package entity

type LiveStreamInfo struct {
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description,omitempty"`
	// @SerializedName("isStreaming")
	// @Nullable
	// private Boolean isStreaming;
	IsStreaming *bool `json:"isStreaming,omitempty"`
	// @SerializedName("liveId")
	// @Nullable
	// private String liveId;
	LiveId *string `json:"liveId,omitempty"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name,omitempty"`
	// @SerializedName("streamId")
	// @Nullable
	// private String streamId;
	StreamId *string `json:"streamId,omitempty"`
}

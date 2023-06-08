package entity

type VersionInfo struct {
	// @SerializedName("appCode")
	// @Nullable
	// private String appCode;
	AppCode *string `json:"appCode"`
	// @SerializedName("id")
	// @Nullable
	// private String id;
	Id *string `json:"id"`
	// @SerializedName("isForceUpdate")
	// @Nullable
	// private Boolean isForceUpdate;
	IsForceUpdate *bool `json:"isForceUpdate"`
	// @SerializedName("userMessage")
	// @Nullable
	// private String userMessage;
	UserMessage *string `json:"userMessage"`
	// @SerializedName("version")
	// @Nullable
	// private String version;
	Version *string `json:"version"`
}

package entity

type GreetingUserVideoInfo struct {
	// @SerializedName("thumbnailUrl")
	// @Nullable
	// private String thumbnailUrl;
	ThumbnailUrl *string `json:"thumbnailUrl"`
	// @SerializedName("totalSeconds")
	// @Nullable
	// private Double totalSeconds;
	TotalSeconds *float64 `json:"totalSeconds"`
	// @SerializedName("videoSubmittedAt")
	// @Nullable
	// private String videoSubmittedAt;
	VideoSubmittedAt *string `json:"videoSubmittedAt"`
	// @SerializedName("videoUrl")
	// @Nullable
	// private String videoUrl;
	VideoUrl *string `json:"videoUrl"`
}

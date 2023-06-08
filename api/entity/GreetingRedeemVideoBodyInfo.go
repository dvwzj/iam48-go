package entity

type GreetingRedeemVideoBodyInfo struct {
	// @SerializedName("totalVideoSeconds")
	// @Nullable
	// private Double totalVideoSeconds;
	TotalVideoSeconds *float64 `json:"totalVideoSeconds"`
	// @SerializedName("videoFileId")
	// @Nullable
	// private String videoFileId;
	VideoFileId *string `json:"videoFileId"`
	// @SerializedName("videoThumbnailFileId")
	// @Nullable
	// private String videoThumbnailFileId;
	VideoThumbnailFileId *string `json:"videoThumbnailFileId"`
}

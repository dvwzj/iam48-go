package entity

type LikeTimelineInfo struct {
	// @SerializedName("likeStatus")
	// @Nullable
	// private Boolean likeStatus = Boolean.FALSE;
	LikeStatus *bool `json:"likeStatus"`
}

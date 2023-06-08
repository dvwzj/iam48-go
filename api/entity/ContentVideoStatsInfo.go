package entity

type ContentVideoStatsInfo struct {
	// @SerializedName("totalGiftValue")
	// @Nullable
	// private Long totalGiftValue;
	TotalGiftValue *int64 `json:"totalGiftValue"`
	// @SerializedName("totalLikes")
	// @Nullable
	// private Long totalLikes;
	TotalLikes *int64 `json:"totalLikes"`
}

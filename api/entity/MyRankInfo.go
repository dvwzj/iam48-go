package entity

type MyRankInfo struct {
	// @SerializedName("rankNo")
	// @Nullable
	// private Long rankNo;
	RankNo *int64 `json:"rankNo"`
	// @SerializedName("score")
	// @Nullable
	// private Long score;
	Score *int64 `json:"score"`
}

package entity

type MyScoreInfo struct {
	// @SerializedName("currentScore")
	// @Nullable
	// private Long currentScore;
	CurrentScore *int64 `json:"currentScore"`
}

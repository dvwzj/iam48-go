package entity

type MissionRedeemInfo struct {
	// @SerializedName("rewardId")
	// @Nullable
	// private Long rewardId;
	RewardId *int64 `json:"rewardId"`
	// @SerializedName("score")
	// @Nullable
	// private Long score;
	Score *int64 `json:"score"`
}

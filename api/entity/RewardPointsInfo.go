package entity

type RewardPointsInfo struct {
	// @SerializedName("balance")
	// @Nullable
	// private Long balance = 0L;
	Balance *int64 `json:"balance"`
}

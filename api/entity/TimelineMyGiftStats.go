package entity

type TimelineMyGiftStats struct {
	// @SerializedName("amount")
	// @Nullable
	// private Long amount;
	Amount *int64 `json:"amount"`
	// @SerializedName("rank")
	// @Nullable
	// private Long rank;
	Rank *int64 `json:"rank"`
}

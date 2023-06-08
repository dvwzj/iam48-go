package entity

type RankStats struct {
	// @SerializedName("amount")
	// @Nullable
	// private Long amount;
	Amount *int64 `json:"amount"`
	// @SerializedName("rank")
	// @Nullable
	// private Integer rank;
	Rank *int `json:"rank"`
}

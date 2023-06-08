package entity

type AdsDisplayPercentage struct {
	// @SerializedName(ConstancesKt.KEY_FROM) // "from"
	// @Nullable
	// private Long from;
	From *int64 `json:"from"`
	// @SerializedName(ConstancesKt.KEY_TO) // "to"
	// @Nullable
	// private Long to;
	To *int64 `json:"to"`
}

package entity

type RankPeriodInfo struct {
	// @SerializedName("7days")
	// @Nullable
	// private String sevenDays;
	SevenDays *string `json:"7days"`
	// @SerializedName("30days")
	// @Nullable
	// private String thirtyDays;
	ThirtyDays *string `json:"30days"`
}

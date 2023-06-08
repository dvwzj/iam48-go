package entity

type MyStats struct {
	// @SerializedName("latestDate")
	// @Nullable
	// private String latestDate;
	LatestDate *string `json:"latestDate"`
	// @SerializedName("totalRedemptionItem")
	// @Nullable
	// private Long totalRedemptionItem;
	TotalRedemptionItem *int64 `json:"totalRedemptionItem"`
}

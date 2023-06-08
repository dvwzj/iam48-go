package entity

type StatisticsInfo struct {
	// @SerializedName("totalAudiences")
	// private long totalAudience;
	TotalAudience int64 `json:"totalAudience"`
	// @SerializedName("totalFavorites")
	// private long totalFavorite;
	TotalFavorite int64 `json:"totalFavorite"`
	// @SerializedName("totalGifts")
	// private long totalGift;
	TotalGift int64 `json:"totalGift"`
	// @SerializedName("totalHearts")
	// private long totalHearts;
	TotalHearts int64 `json:"totalHearts"`
}

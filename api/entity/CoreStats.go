package entity

type CoreStats struct {
	// @SerializedName("allTime")
	// @Nullable
	// private RankStats allTime;
	AllTime *RankStats `json:"allTime"`
	// @SerializedName("last30Days")
	// @Nullable
	// private RankStats last30Days;
	Last30Days *RankStats `json:"last30Days"`
	// @SerializedName("last7Days")
	// @Nullable
	// private RankStats last7Days;
	Last7Days *RankStats `json:"last7Days"`
	// @SerializedName("today")
	// @Nullable
	// private RankStats today;
	Today *RankStats `json:"today"`
}

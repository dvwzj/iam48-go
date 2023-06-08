package entity

type MembersStats struct {
	// @SerializedName("allTime")
	// @Nullable
	// private List<MembersStatsInfo> allTime;
	AllTime *[]MembersStatsInfo `json:"allTime"`
	// @SerializedName("last30Days")
	// @Nullable
	// private List<MembersStatsInfo> last30Days;
	Last30Days *[]MembersStatsInfo `json:"last30Days"`
	// @SerializedName("last7Days")
	// @Nullable
	// private List<MembersStatsInfo> last7Days;
	Last7Days *[]MembersStatsInfo `json:"last7Days"`
	// @SerializedName("today")
	// @Nullable
	// private List<MembersStatsInfo> today;
	Today *[]MembersStatsInfo `json:"today"`
}

package entity

type MemberRankStats struct {
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("stats")
	// @Nullable
	// private List<MemberRankStatsInfo> stats;
	Stats *[]MemberRankStatsInfo `json:"stats"`
}

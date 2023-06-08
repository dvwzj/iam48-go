package entity

type SpecialFansDayRankingList struct {
	// @SerializedName("fromDate")
	// @Nullable
	// private String fromDate;
	FromDate *string `json:"fromDate"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("ranking")
	// @Nullable
	// private List<SpecialFansDayRankingInfo> ranking;
	Ranking *[]SpecialFansDayRankingInfo `json:"ranking"`
	// @SerializedName("toDate")
	// @Nullable
	// private String toDate;
	ToDate *string `json:"toDate"`
	// @SerializedName("totalGiftAmount")
	// @Nullable
	// private Long totalGiftAmount;
	TotalGiftAmount *int64 `json:"totalGiftAmount"`
}

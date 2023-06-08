package entity

type SpecialFansDayRankingInfo struct {
	// @SerializedName("amount")
	// @Nullable
	// private Long amount;
	Amount *int64 `json:"amount"`
	// @Nullable
	// private String graduatedAt;
	GraduatedAt *string `json:"graduatedAt"`
	// @Nullable
	// private String imageUrl;
	ImageUrl *string `json:"imageUrl"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @Nullable
	// private String memberName;
	MemberName *string `json:"memberName"`
	// @SerializedName("rank")
	// @Nullable
	// private Long rank;
	Rank *int64 `json:"rank"`
	// @SerializedName("rankChange")
	// @Nullable
	// private Long rankChange;
	RankChange *int64 `json:"rankChange"`
	// private int rankHight;
	RankHight *int `json:"rankHight"`
}

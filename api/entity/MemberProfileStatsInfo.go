package entity

type MemberProfileStatsInfo struct {
	// @SerializedName("followers")
	// @Nullable
	// private Long followers;
	Followers *int64 `json:"followers"`
	// @SerializedName("gifts")
	// @Nullable
	// private Long gifts;
	Gifts *int64 `json:"gifts"`
	// @SerializedName("likes")
	// @Nullable
	// private Long likes;
	Likes *int64 `json:"likes"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("oshis")
	// @Nullable
	// private Long oshis;
	Oshis *int64 `json:"oshis"`
}

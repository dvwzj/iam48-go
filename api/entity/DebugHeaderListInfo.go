package entity

type DebugHeaderListInfo struct {
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("memberImageUrl")
	// @Nullable
	// private String memberImageUrl;
	MemberImageUrl *string `json:"memberImageUrl"`
}

package entity

type ActivitiesOshiMember struct {
	// @SerializedName("committedAt")
	// @Nullable
	// private String committedAt;
	CommittedAt *string `json:"committedAt"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
}

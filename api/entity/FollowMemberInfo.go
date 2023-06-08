package entity

type FollowMemberInfo struct {
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
}

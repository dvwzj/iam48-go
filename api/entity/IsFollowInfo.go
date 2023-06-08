package entity

type IsFollowInfo struct {
	// @SerializedName("followedAt")
	// @Nullable
	// private String followedAt;
	FollowedAt *string `json:"followedAt"`
}

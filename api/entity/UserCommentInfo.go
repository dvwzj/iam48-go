package entity

type UserCommentInfo struct {
	// @SerializedName("commentStats")
	// @Nullable
	// private List<UserCommentStatsInfo> commentStats;
	CommentStats *[]UserCommentStatsInfo `json:"commentStats"`
	// @SerializedName("displayItemBy")
	// @Nullable
	// private List<Long> displayItemBy;
	DisplayItemBy []*int64 `json:"displayItemBy"`
	// @SerializedName("isAllowUserComment")
	// @Nullable
	// private Boolean isAllowUserComment;
	IsAllowUserComment *bool `json:"isAllowUserComment"`
	// @SerializedName("totalCommentAmount")
	// @Nullable
	// private Integer totalCommentAmount;
	TotalCommentAmount *int `json:"totalCommentAmount"`
}

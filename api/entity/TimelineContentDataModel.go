package entity

type TimelineContentDataModel struct {
	// @SerializedName("isUserCommentAllow")
	// @Nullable
	// private Boolean isUserCommentAllow;
	IsUserCommentAllow *bool `json:"isUserCommentAllow"`
	// @SerializedName("latestCommentItem")
	// @Nullable
	// private List<CommentViewModel> latestCommentItem;
	LatestCommentItem *[]CommentViewModel `json:"latestCommentItem"`
	// @SerializedName("totalCommentAmount")
	// @Nullable
	// private Long totalCommentAmount;
	TotalCommentAmount *int64 `json:"totalCommentAmount"`
	// @SerializedName("totalGifts")
	// @Nullable
	// private Long totalGifts;
	TotalGifts *int64 `json:"totalGifts"`
	// @SerializedName("totalLikeAmount")
	// @Nullable
	// private Long totalLikeAmount;
	TotalLikeAmount *int64 `json:"totalLikeAmount"`
	// @SerializedName("totalViewAmount")
	// @Nullable
	// private Long totalViewAmount;
	TotalViewAmount *int64 `json:"totalViewAmount"`
	// @SerializedName("userImageUrlList")
	// @Nullable
	// private List<String> userImageUrlList;
	UserImageUrlList *[]string `json:"userImageUrlList"`
}

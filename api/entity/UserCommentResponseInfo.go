package entity

type UserCommentResponseInfo struct {
	// @SerializedName("commentId")
	// @Nullable
	// private Long commentId;
	CommentId *int64 `json:"commentId"`
	// @SerializedName("totalGift")
	// @Nullable
	// private Long totalGift;
	TotalGift *int64 `json:"totalGift"`
}

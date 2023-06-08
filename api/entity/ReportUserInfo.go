package entity

type ReportUserInfo struct {
	// @SerializedName("commentId")
	// @Nullable
	// private Long commentId;
	CommentId *int64 `json:"commentId"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("reporterId")
	// @Nullable
	// private Long reporterId;
	ReporterId *int64 `json:"reporterId"`
	// @SerializedName("topicId")
	// @Nullable
	// private Long topicId;
	TopicId *int64 `json:"topicId"`
	// @SerializedName("userId")
	// @Nullable
	// private Long userId;
	UserId *int64 `json:"userId"`
}

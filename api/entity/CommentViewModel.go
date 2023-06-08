package entity

type CommentViewModel struct {
	// @SerializedName("commentText")
	// @Nullable
	// private String commentText;
	CommentText *string `json:"commentText"`
	// @SerializedName("commenterInfo")
	// @Nullable
	// private CommenterInfo commenterInfo;
	CommenterInfo *CommenterInfo `json:"commenterInfo"`
	// @SerializedName("contentId")
	// @Nullable
	// private Long contentId;
	ContentId *int64 `json:"contentId"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("imageFileUrl")
	// @Nullable
	// private String imageFileUrl;
	ImageFileUrl *string `json:"imageFileUrl"`
	// @SerializedName("isLoved")
	// private boolean isLoved;
	IsLoved *bool `json:"isLoved"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("postedAt")
	// @Nullable
	// private String postedAt;
	PostedAt *string `json:"postedAt"`
	// @SerializedName("totalGifts")
	// @Nullable
	// private Long totalGifts;
	TotalGifts *int64 `json:"totalGifts"`
}

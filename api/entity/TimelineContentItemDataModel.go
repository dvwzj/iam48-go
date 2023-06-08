package entity

type TimelineContentItemDataModel struct {
	// @SerializedName("contentText")
	// @Nullable
	// private String contentText;
	ContentText *string `json:"contentText"`
	// @SerializedName("hashtags")
	// @Nullable
	// private List<String> hashtags;
	Hashtags *[]string `json:"hashtags"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("imageFileUrl")
	// @Nullable
	// private List<String> imageFileUrl;
	ImageFileUrl *[]string `json:"imageFileUrl"`
	// @SerializedName("memberId")
	// @Nullable
	// private List<Long> memberId;
	MemberId *[]int64 `json:"memberId"`
	// @SerializedName("postedAt")
	// @Nullable
	// private String postedAt;
	PostedAt *string `json:"postedAt"`
	// @SerializedName("thumbImageUrl")
	// @Nullable
	// private String thumbImageUrl;
	ThumbImageUrl *string `json:"thumbImageUrl"`
}

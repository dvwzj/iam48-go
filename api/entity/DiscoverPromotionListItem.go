package entity

type DiscoverPromotionListItem struct {
	// @SerializedName("bannerImageUrl")
	// @Nullable
	// private String bannerImageUrl;
	BannerImageUrl *string `json:"bannerImageUrl"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("linkUrl")
	// @Nullable
	// private String linkUrl;
	LinkUrl *string `json:"linkUrl"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
}

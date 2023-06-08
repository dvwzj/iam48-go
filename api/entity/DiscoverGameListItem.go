package entity

type DiscoverGameListItem struct {
	// @SerializedName("bannerFileUrl")
	// @Nullable
	// private String bannerFileUrl;
	BannerFileUrl *string `json:"bannerFileUrl"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("iconFileUrl")
	// @Nullable
	// private String iconFileUrl;
	IconFileUrl *string `json:"iconFileUrl"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("store")
	// @Nullable
	// private StoreInfo store;
	Store *StoreInfo `json:"store"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE)
	// @Nullable
	// private String title;
	Title *string `json:"title"`
}

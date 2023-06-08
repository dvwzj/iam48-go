package entity

type ElectionBannerInfo struct {
	// @SerializedName(ChooseMemberActivity.KEY_BANNER_TYPE) // "bannerType"
	// @Nullable
	// private String bannerType;
	BannerType *string `json:"bannerType"`
	// @SerializedName("externalUrl")
	// @Nullable
	// private String externalUrl;
	ExternalUrl *string `json:"externalUrl"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("imageFileUrl")
	// @Nullable
	// private String imageFileUrl;
	ImageFileUrl *string `json:"imageFileUrl"`
	// @SerializedName("orderIndex")
	// @Nullable
	// private Long orderIndex;
	OrderIndex *int64 `json:"orderIndex"`
}

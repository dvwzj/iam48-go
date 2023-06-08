package entity

type ContentRecommendInfo struct {
	// @SerializedName("coinAmount")
	// @Nullable
	// private Long coinAmount;
	CoinAmount *int64 `json:"coinAmount"`
	// @SerializedName("contentId")
	// @Nullable
	// private Long contentId;
	ContentId *int64 `json:"contentId"`
	// @SerializedName("displayIndex")
	// @Nullable
	// private Integer displayIndex;
	DisplayIndex *int `json:"displayIndex"`
	// @SerializedName("imageFileUrl")
	// @Nullable
	// private String imageFileUrl;
	ImageFileUrl *string `json:"imageFileUrl"`
	// @SerializedName("isPurchaseAllowance")
	// @Nullable
	// private Boolean isPurchaseAllowance;
	IsPurchaseAllowance *bool `json:"isPurchaseAllowance"`
	// @SerializedName("isSubscriptionAllowance")
	// @Nullable
	// private Boolean isSubscriptionAllowance;
	IsSubscriptionAllowance *bool `json:"isSubscriptionAllowance"`
	// @SerializedName("mediaType")
	// @Nullable
	// private String mediaType;
	MediaType *string `json:"mediaType"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME)
	// @Nullable
	// private String name;
	Name *string `json:"name"`
	// @SerializedName("publishedAt")
	// @Nullable
	// private String publishedAt;
	PublishedAt *string `json:"publishedAt"`
	// @SerializedName("teaserFileUrl")
	// @Nullable
	// private String teaserFileUrl;
	TeaserFileUrl *string `json:"teaserFileUrl"`
	// @SerializedName("totalWatchTime")
	// @Nullable
	// private Long totalWatchTime;
	TotalWatchTime *int64 `json:"totalWatchTime"`
	// @SerializedName("viewType")
	// @Nullable
	// private String viewType;
	ViewType *string `json:"viewType"`
}

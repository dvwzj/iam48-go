package entity

type Gifts struct {
	// @SerializedName("animateName")
	// @Nullable
	// private String animateName;
	AnimateName *string `json:"animateName"`
	// @SerializedName(ConstanceKt.CAMPAIGN_ID) // "campaignId"
	// @Nullable
	// private Long campaignId;
	CampaignId *int64 `json:"campaignId"`
	// @SerializedName("coinAmount")
	// @Nullable
	// private Long coinAmount;
	CoinAmount *int64 `json:"coinAmount"`
	// @SerializedName("coinPrice")
	// @Nullable
	// private Long coinPrice;
	CoinPrice *int64 `json:"coinPrice"`
	// @SerializedName("giftCategoryId")
	// @Nullable
	// private Long giftCategoryId;
	GiftCategoryId *int64 `json:"giftCategoryId"`
	// @SerializedName("giftValue")
	// @Nullable
	// private Long giftValue;
	GiftValue *int64 `json:"giftValue"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("imageFileUrl")
	// @Nullable
	// private String imageFileUrl;
	ImageFileUrl *string `json:"imageFileUrl"`
	// @SerializedName("imageUrl")
	// @Nullable
	// private String imageUrl;
	ImageUrl *string `json:"imageUrl"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
	// @SerializedName("value")
	// @Nullable
	// private Long value;
	Value *int64 `json:"value"`
}

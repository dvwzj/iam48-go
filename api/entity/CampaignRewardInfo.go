package entity

type CampaignRewardInfo struct {
	// @SerializedName("giftSkuId")
	// @Nullable
	// private Long giftSkuId;
	GiftSkuId *int64 `json:"giftSkuId"`
	// @SerializedName("isOnetimeReceive")
	// @Nullable
	// private Boolean isOnetimeReceive;
	IsOnetimeReceive *bool `json:"isOnetimeReceive"`
	// @SerializedName("minimumGiftRequire")
	// @Nullable
	// private Long minimumGiftRequire;
	MinimumGiftRequire *int64 `json:"minimumGiftRequire"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY) // "quantity"
	// @Nullable
	// private Long quantity;
	Quantity *int64 `json:"quantity"`
}

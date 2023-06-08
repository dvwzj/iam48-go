package entity

type CampaignDonateRewardInfo struct {
	// @SerializedName("amount")
	// @Nullable
	// private Integer amount;
	Amount *int `json:"amount"`
	// @SerializedName("giftSkuId")
	// @Nullable
	// private Long giftSkuId;
	GiftSkuId *int64 `json:"giftSkuId"`
}

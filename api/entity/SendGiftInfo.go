package entity

type SendGiftInfo struct {
	// @SerializedName("giftSkuId")
	// @Nullable
	// private Long giftSkuId;
	GiftSkuId *int64 `json:"giftSkuId"`
	// @SerializedName("pricePerUnit")
	// @Nullable
	// private Long pricePerUnit;
	PricePerUnit *int64 `json:"pricePerUnit"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY) // "quantity"
	// @Nullable
	// private Long quantity;
	Quantity *int64 `json:"quantity"`
}

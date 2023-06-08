package entity

type GiftBody struct {
	// @SerializedName("giftSkuId")
	// private long giftSkuId;
	GiftSkuId int64 `json:"giftSkuId"`
	// @SerializedName("pricePerUnit")
	// private long pricePerUnit;
	PricePerUnit int64 `json:"pricePerUnit"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY)
	// private long quantity;
	Quantity int64 `json:"quantity"`
}

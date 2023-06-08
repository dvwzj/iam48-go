package entity

type UserCommentBodyInfo struct {
	// @SerializedName("giftSkuId")
	// @Nullable
	// private Long giftSkuId;
	GiftSkuId *int64 `json:"giftSkuId"`
	// @SerializedName(ConstancesKt.KEY_MESSAGE) // "message"
	// @Nullable
	// private String message;
	Message *string `json:"message"`
	// @SerializedName("pricePerUnit")
	// @Nullable
	// private Long pricePerUnit;
	PricePerUnit *int64 `json:"pricePerUnit"`
}

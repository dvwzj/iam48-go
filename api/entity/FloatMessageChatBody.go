package entity

type FloatMessageChatBody struct {
	// @SerializedName("giftSkuId")
	// private final long giftSkuId;
	GiftSkuId int64 `json:"giftSkuId"`
	// @SerializedName(ConstancesKt.KEY_MESSAGE) // "message"
	// @NotNull
	// private final String message;
	Message string `json:"message"`
	// @SerializedName("pricePerUnit")
	// private final long pricePerUnit;
	PricePerUnit int64 `json:"pricePerUnit"`
}

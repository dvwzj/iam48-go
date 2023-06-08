package entity

type RequestGiftInfo struct {
	// @SerializedName("giftSkuId")
	// @Nullable
	// private Long giftSkuId;
	GiftSkuId *int64 `json:"giftSkuId"`
	// @SerializedName("liveId")
	// @Nullable
	// private Long liveId;
	LiveId *int64 `json:"liveId"`
	// @SerializedName("pricePerUnit")
	// @Nullable
	// private Long pricePerUnit;
	PricePerUnit *int64 `json:"pricePerUnit"`
	// @SerializedName("requestQuantity")
	// @Nullable
	// private Long requestQuantity;
	RequestQuantity *int64 `json:"requestQuantity"`
	// @SerializedName("requirePaymentQuantity")
	// @Nullable
	// private Long requirePaymentQuantity;
	RequirePaymentQuantity *int64 `json:"requirePaymentQuantity"`
	// @SerializedName("totalRequirePaymentCoin")
	// @Nullable
	// private Long totalRequirePaymentCoin;
	TotalRequirePaymentCoin *int64 `json:"totalRequirePaymentCoin"`
	// @SerializedName("transactionId")
	// @Nullable
	// private Long transactionId;
	TransactionId *int64 `json:"transactionId"`
	// @SerializedName("userId")
	// @Nullable
	// private Long userId;
	UserId *int64 `json:"userId"`
}

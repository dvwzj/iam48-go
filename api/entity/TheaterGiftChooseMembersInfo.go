package entity

type TheaterGiftChooseMembersInfo struct {
	// @SerializedName("autoPurchaseOrder")
	// @Nullable
	// private AutoPurchaseOrder autoPurchaseOrder;
	AutoPurchaseOrder *AutoPurchaseOrder `json:"autoPurchaseOrder"`
	// @SerializedName("giftSkuId")
	// @Nullable
	// private Long giftSkuId;
	GiftSkuId *int64 `json:"giftSkuId"`
	// @SerializedName("memberIds")
	// @Nullable
	// private List<Long> memberIds;
	MemberIds *[]int64 `json:"memberIds"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY) // "quantity"
	// @Nullable
	// private Long quantity;
	Quantity *int64 `json:"quantity"`
}

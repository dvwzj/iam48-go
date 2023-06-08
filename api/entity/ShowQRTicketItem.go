package entity

type ShowQRTicketItem struct {
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY) // "quantity"
	// @Nullable
	// private Long quantity;
	Quantity *int64 `json:"quantity"`
	// @SerializedName("ticketSkuId")
	// @Nullable
	// private Long ticketSkuId;
	TicketSkuId *int64 `json:"ticketSkuId"`
}

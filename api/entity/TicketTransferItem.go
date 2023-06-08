package entity

type TicketTransferItem struct {
	// @SerializedName("ticketSkuId")
	// @Nullable
	// private Long ticketSkuId;
	TicketSkuId *int64 `json:"ticketSkuId"`
	// @SerializedName("toWalletCode")
	// @Nullable
	// private String toWalletCode;
	ToWalletCode *string `json:"toWalletCode"`
}

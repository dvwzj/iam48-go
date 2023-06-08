package entity

type ExchangeDataPostBodyInfo struct {
	// @SerializedName("fromTicketSkuId")
	// @Nullable
	// private Long fromTicketSkuId;
	FromTicketSkuId *int64 `json:"fromTicketSkuId"`
	// @SerializedName("ownerPIN")
	// @Nullable
	// private String ownerPIN;
	OwnerPIN *string `json:"ownerPIN"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY) // "quantity"
	// @Nullable
	// private Long quantity;
	Quantity *int64 `json:"quantity"`
	// @SerializedName("toTicketSkuId")
	// @Nullable
	// private Long toTicketSkuId;
	ToTicketSkuId *int64 `json:"toTicketSkuId"`
}

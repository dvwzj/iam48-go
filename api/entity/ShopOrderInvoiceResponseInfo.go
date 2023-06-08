package entity

type ShopOrderInvoiceResponseInfo struct {
	// @SerializedName("invoiceNo")
	// @Nullable
	// private String invoiceNo;
	InvoiceNo *string `json:"invoiceNo"`
	// @SerializedName(ConstancesKt.KEY_REF_CODE) // "refCode"
	// @Nullable
	// private String refCode;
	RefCode *string `json:"refCode"`
	// @SerializedName("userId")
	// @Nullable
	// private Long userId;
	UserId *int64 `json:"userId"`
}

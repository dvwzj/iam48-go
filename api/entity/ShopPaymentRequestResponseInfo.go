package entity

type ShopPaymentRequestResponseInfo struct {
	// @SerializedName("invoiceNumber")
	// @Nullable
	// private String invoiceNumber;
	InvoiceNumber *string `json:"invoiceNumber"`
	// @SerializedName("isSuccess")
	// @Nullable
	// private Boolean isSuccess;
	IsSuccess *bool `json:"isSuccess"`
	// @SerializedName(ConstancesKt.KEY_MESSAGE) // "message"
	// @Nullable
	// private String message;
	Message *string `json:"message"`
	// @SerializedName(ImagesContract.URL)
	// @Nullable
	// private String url;
	Url *string `json:"url"`
}

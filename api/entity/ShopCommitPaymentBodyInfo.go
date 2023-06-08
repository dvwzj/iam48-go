package entity

type ShopCommitPaymentBodyInfo struct {
	// @SerializedName("gatewayCode")
	// @Nullable
	// private String gatewayCode;
	GatewayCode *string `json:"gatewayCode"`
	// @SerializedName("isSandbox")
	// @Nullable
	// private Boolean isSandbox;
	IsSandbox *bool `json:"isSandbox"`
	// @SerializedName("netPrice")
	// @Nullable
	// private Double netPrice;
	NetPrice *float64 `json:"netPrice"`
	// @SerializedName("originalInvoiceNo")
	// @Nullable
	// private String originalInvoiceNo;
	OriginalInvoiceNo *string `json:"originalInvoiceNo"`
	// @SerializedName(ConstancesKt.KEY_REF_CODE) // "refCode"
	// @Nullable
	// private String refCode;
	RefCode *string `json:"refCode"`
	// @SerializedName("userId")
	// @Nullable
	// private Long userId;
	UserId *int64 `json:"userId"`
}

package entity

type ShopCommitPaymentResponseInfo struct {
	// @SerializedName("isSuccess")
	// @Nullable
	// private Boolean isSuccess;
	IsSuccess *bool `json:"isSuccess"`
	// @SerializedName(ConstancesKt.KEY_MESSAGE) // "message"
	// @Nullable
	// private String message;
	Message *string `json:"message"`
	// @SerializedName("transactionId")
	// @Nullable
	// private Long transactionId;
	TransactionId *int64 `json:"transactionId"`
}

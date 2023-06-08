package entity

type UserTransactionInfo struct {
	// @SerializedName("coinAmount")
	// @Nullable
	// private Long coinAmount;
	CoinAmount *int64 `json:"coinAmount"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("productType")
	// @Nullable
	// private String productType;
	ProductType *string `json:"productType"`
	// @SerializedName("purchaseId")
	// @Nullable
	// private Long purchaseId;
	PurchaseId *int64 `json:"purchaseId"`
	// @SerializedName("purchasedAt")
	// @Nullable
	// private String purchasedAt;
	PurchasedAt *string `json:"purchasedAt"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
	// @SerializedName("transactionId")
	// @Nullable
	// private String transactionId;
	TransactionId *string `json:"transactionId"`
}

package entity

type TicketHistoryModel struct {
	// @SerializedName("claimReferenceOrderNo")
	// @Nullable
	// private String claimReferenceOrderNo;
	ClaimReferenceOrderNo *string `json:"claimReferenceOrderNo"`
	// @SerializedName("createdAt")
	// @Nullable
	// private String createdAt;
	CreatedAt *string `json:"createdAt"`
	// @SerializedName("itemName")
	// @Nullable
	// private String itemName;
	ItemName *string `json:"itemName"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY) // "quantity"
	// @Nullable
	// private String quantity;
	Quantity *string `json:"quantity"`
	// @SerializedName("referenceWalletInfo")
	// @Nullable
	// private WalletInfoModel referenceWalletInfo;
	ReferenceWalletInfo *WalletInfoModel `json:"referenceWalletInfo"`
	// @SerializedName("totalTransactionFee")
	// @Nullable
	// private Long totalTransactionFee;
	TotalTransactionFee *int64 `json:"totalTransactionFee"`
	// @SerializedName("transactionId")
	// @Nullable
	// private Long transactionId;
	TransactionId *int64 `json:"transactionId"`
	// @SerializedName("transactionType")
	// @Nullable
	// private String transactionType;
	TransactionType *string `json:"transactionType"`
}

package entity

type TicketTransferModel struct {
	// @SerializedName("maxQuantity")
	// @Nullable
	// private Integer maxQuantity;
	MaxQuantity *int `json:"maxQuantity"`
	// @SerializedName("ownerWallet")
	// @Nullable
	// private WalletInfoModel ownerWallet;
	OwnerWallet *WalletInfoModel `json:"ownerWallet"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY) // "quantity"
	// @Nullable
	// private Integer quantity;
	Quantity *int `json:"quantity"`
	// @SerializedName("receiverWallet")
	// @Nullable
	// private WalletInfoModel receiverWallet;
	ReceiverWallet *WalletInfoModel `json:"receiverWallet"`
	// @SerializedName(ConstancesKt.KEY_TICKET) // "ticket"
	// @Nullable
	// private TicketModel ticket;
	Ticket *TicketModel `json:"ticket"`
	// @SerializedName("totalTransactionFee")
	// @Nullable
	// private Integer totalTransactionFee;
	TotalTransactionFee *int `json:"totalTransactionFee"`
	// @SerializedName("transactionFeePerUnit")
	// @Nullable
	// private Integer transactionFeePerUnit;
	TransactionFeePerUnit *int `json:"transactionFeePerUnit"`
	// @SerializedName("transactionId")
	// @Nullable
	// private String transactionId;
	TransactionId *string `json:"transactionId"`
}

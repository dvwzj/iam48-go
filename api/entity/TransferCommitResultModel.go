package entity

type TransferCommitResultModel struct {
	// @SerializedName("createdAt")
	// @Nullable
	// private String createdAt;
	CreatedAt *string `json:"createdAt"`
	// @SerializedName("maxQuantity")
	// @Nullable
	// private Long maxQuantity;
	MaxQuantity *int64 `json:"maxQuantity"`
	// @SerializedName("memberId")
	// @Nullable
	// private List<Integer> memberId;
	MemberId []int `json:"memberId"`
	// @SerializedName("ownerPIN")
	// @Nullable
	// private Long ownerPIN;
	OwnerPIN *int64 `json:"ownerPIN"`
	// @SerializedName("ownerWallet")
	// @Nullable
	// private WalletInfoModel ownerWallet;
	OwnerWallet *WalletInfoModel `json:"ownerWallet"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY) // "quantity"
	// @Nullable
	// private Long quantity;
	Quantity *int64 `json:"quantity"`
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
	// private Long totalTransactionFee;
	TotalTransactionFee *int64 `json:"totalTransactionFee"`
	// @SerializedName("transactionFeePerUnit")
	// @Nullable
	// private Long transactionFeePerUnit;
	TransactionFeePerUnit *int64 `json:"transactionFeePerUnit"`
	// @SerializedName("transactionId")
	// @Nullable
	// private String transactionId;
	TransactionId *string `json:"transactionId"`
}

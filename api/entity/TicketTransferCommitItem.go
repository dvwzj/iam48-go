package entity

type TicketTransferCommitItem struct {
	// @SerializedName("memberId")
	// @Nullable
	// private List<Integer> memberId;
	MemberId *[]int `json:"memberId"`
	// @SerializedName("ownerPIN")
	// @Nullable
	// private String ownerPIN;
	OwnerPIN *string `json:"ownerPIN"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY) // "quantity"
	// @Nullable
	// private Long quantity;
	Quantity *int64 `json:"quantity"`
	// @SerializedName("transactionId")
	// @Nullable
	// private String transactionId;
	TransactionId *string `json:"transactionId"`
}

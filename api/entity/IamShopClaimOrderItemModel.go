package entity

type IamShopClaimOrderItemModel struct {
	// @SerializedName("createdAt")
	// @Nullable
	// private String createdAt;
	CreateAt *string `json:"createdAt"`
	// @SerializedName("createdAtAsString")
	// @Nullable
	// private String createdAtAsString;
	CreatedAtAsString *string `json:"createdAtAsString"`
	// @SerializedName("itemType")
	// @Nullable
	// private String itemType;
	ItemType *string `json:"itemType"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY) // "quantity"
	// @Nullable
	// private String quantity;
	Quantity *string `json:"quantity"`
	// @SerializedName(ConstancesKt.KEY_TICKET) // "ticket"
	// @Nullable
	// private IamShopTicketModel ticket;
	Ticket *IamShopTicketModel `json:"ticket"`
	// @SerializedName("transactionId")
	// @Nullable
	// private Long transactionId;
	TransactionId *int64 `json:"transactionId"`
}

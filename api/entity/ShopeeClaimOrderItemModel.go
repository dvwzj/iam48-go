package entity

type ShopeeClaimOrderItemModel struct {
	// @SerializedName("createdAt")
	// @Nullable
	// private String createdAt;
	CreatedAt *string `json:"createdAt"`
	// @SerializedName("itemType")
	// @Nullable
	// private String itemType;
	ItemType *string `json:"itemType"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY) // "quantity"
	// @Nullable
	// private String quantity;
	Quantity *string `json:"quantity"`
	// @SerializedName(ConstancesKt.KEY_TICKET)
	// @Nullable
	// private TicketModel ticket;
	Ticket *TicketModel `json:"ticket"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
	// @SerializedName("token")
	// @Nullable
	// private TokenModel token;
	Token *TokenModel `json:"token"`
	// @SerializedName("transactionId")
	// @Nullable
	// private Long transactionId;
	TransactionId *int64 `json:"transactionId"`
}

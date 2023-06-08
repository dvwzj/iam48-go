package entity

type ElectionClaimBodyInfo struct {
	// @SerializedName("buyerUsername")
	// @Nullable
	// private String buyerUsername;
	BuyerUsername *string `json:"buyerUsername"`
	// @SerializedName("orderId")
	// @Nullable
	// private String orderId;
	OrderId *string `json:"orderId"`
}

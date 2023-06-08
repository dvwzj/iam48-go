package entity

type IamShopClaimOrderModel struct {
	// @SerializedName("orderId")
	// @Nullable
	// private String orderId;
	OrderId *string `json:"orderId"`
}

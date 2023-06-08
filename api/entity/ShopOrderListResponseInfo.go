package entity

type ShopOrderListResponseInfo struct {
	// @SerializedName("expiredAt")
	// @Nullable
	// private String expiredAt;
	ExpiredAt *string `json:"expiredAt"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("orderItem")
	// @Nullable
	// private List<OrderItemInfo> orderItem;
	OrderItem *[]OrderItemInfo `json:"orderItem"`
	// @SerializedName("orderStatus")
	// @Nullable
	// private String orderStatus;
	OrderStatus *string `json:"orderStatus"`
	// @SerializedName("orderStatusCode")
	// @Nullable
	// private String orderStatusCode;
	OrderStatusCode *string `json:"orderStatusCode"`
	// @SerializedName("redeemStatus")
	// @Nullable
	// private String redeemStatus;
	RedeemStatus *string `json:"redeemStatus"`
	// @SerializedName(ConstancesKt.KEY_REF_CODE) // "refCode"
	// @Nullable
	// private String refCode;
	RefCode *string `json:"refCode"`
	// @SerializedName("totalAmount")
	// @Nullable
	// private Double totalAmount;
	TotalAmount *float64 `json:"totalAmount"`
}

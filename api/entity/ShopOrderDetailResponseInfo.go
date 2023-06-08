package entity

type ShopOrderDetailResponseInfo struct {
	// @SerializedName("address")
	// @Nullable
	// private String address;
	Address *string `json:"address"`
	// @SerializedName("amount")
	// @Nullable
	// private Double amount;
	Amount *float64 `json:"amount"`
	// @SerializedName("createdAt")
	// @Nullable
	// private String createdAt;
	CreatedAt *string `json:"createdAt"`
	// @SerializedName("expiredAt")
	// @Nullable
	// private String expiredAt;
	ExpiredAt *string `json:"expiredAt"`
	// @SerializedName("firstName")
	// @Nullable
	// private String firstName;
	FirstName *string `json:"firstName"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("lastName")
	// @Nullable
	// private String lastName;
	LastName *string `json:"lastName"`
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
	// @SerializedName("paymentAt")
	// @Nullable
	// private String paymentAt;
	PaymentAt *string `json:"paymentAt"`
	// @SerializedName("paymentChannel")
	// @Nullable
	// private String paymentChannel;
	PaymentChannel *string `json:"paymentChannel"`
	// @SerializedName("paymentImageUrl")
	// @Nullable
	// private String paymentImageUrl;
	PaymentImageUrl *string `json:"paymentImageUrl"`
	// @SerializedName("phone")
	// @Nullable
	// private String phone;
	Phone *string `json:"phone"`
	// @SerializedName("redeemStatus")
	// @Nullable
	// private String redeemStatus;
	RedeemStatus *string `json:"redeemStatus"`
	// @SerializedName(ConstancesKt.KEY_REF_CODE) // "refCode"
	// @Nullable
	// private String refCode;
	RefCode *string `json:"refCode"`
	// @SerializedName("shippingFee")
	// @Nullable
	// private Double shippingFee;
	ShippingFee *float64 `json:"shippingFee"`
	// @SerializedName("totalAmount")
	// @Nullable
	// private Double totalAmount;
	TotalAmount *float64 `json:"totalAmount"`
	// @SerializedName("trackingNo")
	// @Nullable
	// private String trackingNo;
	TrackingNo *string `json:"trackingNo"`
	// @SerializedName("trackingUrl")
	// @Nullable
	// private String trackingUrl;
	TrackingUrl *string `json:"trackingUrl"`
}

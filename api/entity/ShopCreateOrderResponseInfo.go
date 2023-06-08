package entity

type ShopCreateOrderResponseInfo struct {
	// @SerializedName("address")
	// @Nullable
	// private String address;
	Address *string `json:"address"`
	// @SerializedName("amount")
	// @Nullable
	// private Integer amount;
	Amount *int `json:"amount"`
	// @SerializedName("createdAt")
	// @Nullable
	// private String createdAt;
	CreatedAt *string `json:"createdAt"`
	// @SerializedName(Scopes.EMAIL) // "email"
	// @Nullable
	// private String email;
	Email *string `json:"email"`
	// @SerializedName("expiredAt")
	// @Nullable
	// private String expiredAt;
	ExpiredAt *string `json:"expiredAt"`
	// @SerializedName("firstName")
	// @Nullable
	// private String firstName;
	FirstName *string `json:"firstName"`
	// @SerializedName("invoiceNo")
	// @Nullable
	// private String invoiceNo;
	InvoiceNo *string `json:"invoiceNo"`
	// @SerializedName("lastName")
	// @Nullable
	// private String lastName;
	LastName *string `json:"lastName"`
	// @SerializedName("paymentChannel")
	// @Nullable
	// private String paymentChannel;
	PaymentChannel *string `json:"paymentChannel"`
	// @SerializedName("paymentUrl")
	// @Nullable
	// private String paymentUrl;
	PaymentUrl *string `json:"paymentUrl"`
	// @SerializedName("phone")
	// @Nullable
	// private String phone;
	Phone *string `json:"phone"`
	// @SerializedName(ConstancesKt.KEY_REF_CODE) // "refCode"
	// @Nullable
	// private String refCode;
	RefCode *string `json:"refCode"`
	// @SerializedName("shippingFee")
	// @Nullable
	// private Double shippingFee;
	ShippingFee *float64 `json:"shippingFee"`
	// @SerializedName("total")
	// @Nullable
	// private Double total;
	Total *float64 `json:"total"`
	// @SerializedName("userId")
	// @Nullable
	// private Long userId;
	UserId *int64 `json:"userId"`
}

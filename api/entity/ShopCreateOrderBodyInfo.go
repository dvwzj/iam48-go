package entity

type ShopCreateOrderBodyInfo struct {
	// @SerializedName("address")
	// @Nullable
	// private String address;
	Address *string `json:"address"`
	// @SerializedName("district")
	// @Nullable
	// private String district;
	District *string `json:"district"`
	// @SerializedName("firstName")
	// @Nullable
	// private String firstName;
	FirstName *string `json:"firstName"`
	// @SerializedName("lastName")
	// @Nullable
	// private String lastName;
	LastName *string `json:"lastName"`
	// @SerializedName("orderItem")
	// @Nullable
	// private List<OrderItemInfo> orderItem;
	OrderItem *[]OrderItemInfo `json:"orderItem"`
	// @SerializedName("paymentMethodId")
	// @Nullable
	// private Long paymentMethodId;
	PaymentMethodId *int64 `json:"paymentMethodId"`
	// @SerializedName("phone")
	// @Nullable
	// private String phone;
	Phone *string `json:"phone"`
	// @SerializedName("postalCode")
	// @Nullable
	// private Long postalCode;
	PostalCode *int64 `json:"postalCode"`
	// @SerializedName("province")
	// @Nullable
	// private String province;
	Province *string `json:"province"`
	// @SerializedName("subDistrict")
	// @Nullable
	// private String subDistrict;
	SubDistrict *string `json:"subDistrict"`
}

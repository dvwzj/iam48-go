package entity

type ShopShippingFeeBodyInfo struct {
	// @SerializedName("district")
	// @Nullable
	// private String district;
	District *string `json:"district"`
	// @SerializedName("orderItem")
	// @Nullable
	// private List<OrderItemInfo> orderItem;
	OrderItem *[]OrderItemInfo `json:"orderItem"`
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

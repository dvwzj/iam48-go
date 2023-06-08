package entity

type ShopShippingFeeResponseInfo struct {
	// @SerializedName("price")
	// @Nullable
	// private Double price;
	Price *float64 `json:"price"`
	// @SerializedName("totalWeight")
	// @Nullable
	// private Double totalWeight;
	TotalWeight *float64 `json:"totalWeight"`
}

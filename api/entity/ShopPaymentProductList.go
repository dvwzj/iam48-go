package entity

type ShopPaymentProductList struct {
	// @SerializedName("price")
	// @Nullable
	// private Double price;
	Price *float64 `json:"price"`
	// @SerializedName("productCode")
	// @Nullable
	// private String productCode;
	ProductCode *string `json:"productCode"`
	// @SerializedName("productDescription")
	// @Nullable
	// private String productDescription;
	ProductDescription *string `json:"productDescription"`
	// @SerializedName("productType")
	// @Nullable
	// private String productType;
	ProductType *string `json:"productType"`
	// @SerializedName("unit")
	// @Nullable
	// private Integer unit;
	Unit *int `json:"unit"`
}

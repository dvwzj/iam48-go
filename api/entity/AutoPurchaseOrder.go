package entity

type AutoPurchaseOrder struct {
	// @SerializedName("creditSource")
	// @Nullable
	// private String creditSource;
	CreditSource *string `json:"creditSource"`
	// @SerializedName("pricePerUnit")
	// @Nullable
	// private Long pricePerUnit;
	PricePerUnit *int64 `json:"pricePerUnit"`
}

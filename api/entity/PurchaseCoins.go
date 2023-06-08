package entity

type PurchaseCoins struct {
	// @SerializedName("appCode")
	// @Nullable
	// private String appCode;
	AppCode *string `json:"appCode"`
	// @SerializedName("coinAmount")
	// @Nullable
	// private Integer coinAmount;
	CoinAmount *int `json:"coinAmount"`
	// @SerializedName("currency")
	// @NotNull
	// private String currency;
	Currency string `json:"currency"`
	// @SerializedName("iconImageUrl")
	// @Nullable
	// private String iconImageUrl;
	IconImageUrl *string `json:"iconImageUrl"`
	// @SerializedName("price")
	// @NotNull
	// private String price;
	Price string `json:"price"`
	// @SerializedName("productId")
	// @Nullable
	// private String productId;
	ProductId *string `json:"productId"`
	// @SerializedName("purchaseStatus")
	// @NotNull
	// private Iam48BillingImpl.PurchaseStatus purchaseStatus;
	PurchaseStatus PurchaseStatus `json:"purchaseStatus"`
}

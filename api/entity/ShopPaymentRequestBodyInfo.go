package entity

type ShopPaymentRequestBodyInfo struct {
	// @SerializedName("appCode")
	// @Nullable
	// private String appCode;
	AppCode *string `json:"appCode"`
	// @SerializedName("customData")
	// @Nullable
	// private String customData;
	CustomData *string `json:"customData"`
	// @SerializedName("customerEmail")
	// @Nullable
	// private String customerEmail;
	CustomerEmail *string `json:"customerEmail"`
	// @SerializedName("customerName")
	// @Nullable
	// private String customerName;
	CustomerName *string `json:"customerName"`
	// @SerializedName("items")
	// @Nullable
	// private List<ShopPaymentProductList> items;
	Items *[]ShopPaymentProductList `json:"items"`
	// @SerializedName("price")
	// @Nullable
	// private Double price;
	Price *float64 `json:"price"`
	// @SerializedName("redirectUrl")
	// @Nullable
	// private String redirectUrl;
	RedirectUrl *string `json:"redirectUrl"`
	// @SerializedName("userId")
	// @Nullable
	// private Long userId;
	UserId *int64 `json:"userId"`
}

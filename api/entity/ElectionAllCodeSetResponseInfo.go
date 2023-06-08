package entity

type ElectionAllCodeSetResponseInfo struct {
	// @SerializedName("body")
	// @Nullable
	// private String body;
	Body *string `json:"body"`
	// @SerializedName("shopeeOrderId")
	// @Nullable
	// private String shopeeOrderId;
	ShopeeOrderId *string `json:"shopeeOrderId"`
	// @SerializedName("subject")
	// @Nullable
	// private String subject;
	Subject *string `json:"subject"`
	// @SerializedName("totalCodeAmount")
	// @Nullable
	// private Long totalCodeAmount;
	TotalCodeAmount *int64 `json:"totalCodeAmount"`
}

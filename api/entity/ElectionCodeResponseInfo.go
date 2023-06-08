package entity

type ElectionCodeResponseInfo struct {
	// @SerializedName("claimedAt")
	// @Nullable
	// private String claimedAt;
	ClaimedAt *string `json:"claimedAt"`
	// @SerializedName("displayCodePerPageAmount")
	// @Nullable
	// private Long displayCodePerPageAmount;
	DisplayCodePerPageAmount *int64 `json:"displayCodePerPageAmount"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("maxSelectedAmountForCopy")
	// @Nullable
	// private Long maxSelectedAmountForCopy;
	MaxSelectedAmountForCopy *int64 `json:"maxSelectedAmountForCopy"`
	// @SerializedName("shopeeOrderId")
	// @Nullable
	// private String shopeeOrderId;
	ShopeeOrderId *string `json:"shopeeOrderId"`
	// @SerializedName("totalCodeAmount")
	// @Nullable
	// private Long totalCodeAmount;
	TotalCodeAmount *int64 `json:"totalCodeAmount"`
}

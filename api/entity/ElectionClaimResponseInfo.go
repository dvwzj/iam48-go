package entity

type ElectionClaimResponseInfo struct {
	// @SerializedName("claimedAt")
	// @Nullable
	// private String claimedAt;
	ClaimedAt *string `json:"claimedAt"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("shopeeOrderId")
	// @Nullable
	// private String shopeeOrderId;
	ShopeeOrderId *string `json:"shopeeOrderId"`
	// @SerializedName("totalCodeAmount")
	// @Nullable
	// private Long totalCodeAmount;
	TotalCodeAmount *int64 `json:"totalCodeAmount"`
}

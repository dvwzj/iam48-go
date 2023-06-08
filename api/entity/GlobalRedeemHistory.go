package entity

type GlobalRedeemHistory struct {
	// @SerializedName("createdAt")
	// @Nullable
	// private String createdAt;
	CreatedAt *string `json:"createdAt"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("itemAmount")
	// @Nullable
	// private Long itemAmount;
	ItemAmount *int64 `json:"itemAmount"`
	// @SerializedName("redeemCode")
	// @Nullable
	// private String redeemCode;
	RedeemCode *string `json:"redeemCode"`
}

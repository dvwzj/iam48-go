package entity

type GlobalRedeemCodeInfo struct {
	// @SerializedName(GlobalRedeemCodeActivity.KEY_CODE) // "code"
	// @Nullable
	// private String code;
	Code *string `json:"code"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
	// @SerializedName("skus")
	// @Nullable
	// private List<GlobalRedeemProductInfo> redeemProductList;
	RedeemProductList *[]GlobalRedeemProductInfo `json:"skus"`
}

package entity

type AdsInfo struct {
	// @SerializedName(GlobalRedeemCodeActivity.KEY_CODE) // "code"
	// @Nullable
	// private String code;
	Code *string `json:"code"`
	// @SerializedName("items")
	// @Nullable
	// private List<AdsItems> items;
	Items *[]AdsItems `json:"items"`
	// @SerializedName("priority")
	// @Nullable
	// private Long priority;
	Priority *int64 `json:"priority"`
	// @SerializedName("version")
	// @Nullable
	// private String version;
	Version *string `json:"version"`
}

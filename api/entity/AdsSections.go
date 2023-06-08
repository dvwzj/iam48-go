package entity

type AdsSections struct {
	// @SerializedName("ads")
	// @Nullable
	// private List<AdsInfo> ads;
	Ads *[]AdsInfo `json:"ads"`
	// @SerializedName("type")
	// @Nullable
	// private String type;
	Type *string `json:"type"`
	// @SerializedName("version")
	// @Nullable
	// private String version;
	Version *string `json:"version"`
}

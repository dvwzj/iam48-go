package entity

type StoreInfo struct {
	// @SerializedName("android")
	// @Nullable
	// private String android;
	Android *string `json:"android"`
	// @SerializedName("iOS")
	// @Nullable
	// private String iOS;
	Ios *string `json:"iOS"`
}

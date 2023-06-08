package entity

type MyProvinceInfo struct {
	// @SerializedName("provinceId")
	// @Nullable
	// private Long provinceId;
	ProvinceId *int64 `json:"provinceId"`
}

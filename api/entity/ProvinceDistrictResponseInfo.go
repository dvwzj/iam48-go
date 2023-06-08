package entity

type ProvinceDistrictResponseInfo struct {
	// @SerializedName("districts")
	// @Nullable
	// private List<DistrictInfo> districts;
	Districts *[]DistrictInfo `json:"districts"`
	// @SerializedName("provinceId")
	// @Nullable
	// private Long provinceId;
	ProvinceId *int64 `json:"provinceId"`
	// @SerializedName("provinceNameEn")
	// @Nullable
	// private String provinceNameEn;
	ProvinceNameEn *string `json:"provinceNameEn"`
	// @SerializedName("provinceNameTh")
	// @Nullable
	// private String provinceNameTh;
	ProvinceNameTh *string `json:"provinceNameTh"`
}

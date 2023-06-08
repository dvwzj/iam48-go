package entity

type DistrictInfo struct {
	// @SerializedName("districtId")
	// @Nullable
	// private Long districtId;
	DistrictId *int64 `json:"districtId"`
	// @SerializedName("districtNameEn")
	// @Nullable
	// private String districtNameEn;
	DistrictNameEn *string `json:"districtNameEn"`
	// @SerializedName("districtNameTh")
	// @Nullable
	// private String districtNameTh;
	DistrictNameTh *string `json:"districtNameTh"`
}

package entity

type SubDistrictInfo struct {
	// @SerializedName("postalCodes")
	// @Nullable
	// private List<Long> postalCode;
	PostalCodes *[]int64 `json:"postalCodes"`
	// @SerializedName("subDistrictId")
	// @Nullable
	// private Long subDistrictId;
	SubDistrictId *int64 `json:"subDistrictId"`
	// @SerializedName("subDistrictNameEn")
	// @Nullable
	// private String subDistrictNameEn;
	SubDistrictNameEn *string `json:"subDistrictNameEn"`
	// @SerializedName("subDistrictNameTh")
	// @Nullable
	// private String subDistrictNameTh;
	SubDistrictNameTh *string `json:"subDistrictNameTh"`
}

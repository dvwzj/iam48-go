package entity

type AddUserAddressInfo struct {
	// @SerializedName("address")
	// @Nullable
	// private String address;
	Address *string `json:"address"`
	// @SerializedName("citizenId")
	// @Nullable
	// private String citizenId;
	CitizenId *string `json:"citizenId"`
	// @SerializedName("districtId")
	// @Nullable
	// private Long districtId;
	DistrictId *int64 `json:"districtId"`
	// @SerializedName("firstName")
	// @Nullable
	// private String firstName;
	FirstName *string `json:"firstName"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("isDefault")
	// @Nullable
	// private Boolean isDefault;
	IsDefault *bool `json:"isDefault"`
	// @SerializedName("lastName")
	// @Nullable
	// private String lastName;
	LastName *string `json:"lastName"`
	// @SerializedName("phone")
	// @Nullable
	// private String phone;
	Phone *string `json:"phone"`
	// @SerializedName("postalCode")
	// @Nullable
	// private Long postalCode;
	PostalCode *int64 `json:"postalCode"`
	// @SerializedName("provinceId")
	// @Nullable
	// private Long provinceId;
	ProvinceId *int64 `json:"provinceId"`
	// @SerializedName("subDistrictId")
	// @Nullable
	// private Long subDistrictId;
	SubDistrictId *int64 `json:"subDistrictId"`
}

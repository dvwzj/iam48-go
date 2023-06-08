package entity

type DebugTabbarInfo struct {
	// @SerializedName(Constants.DateOfBirthFlag.MONTH_YEAR) // "1"
	// @Nullable
	// private DebugTabbarImageInfo firstTabbar;
	FirstTabbar *DebugTabbarImageInfo `json:"1"`
	// @SerializedName("4")
	// @Nullable
	// private DebugTabbarImageInfo fourthTabbar;
	FourthTabbar *DebugTabbarImageInfo `json:"4"`
	// @SerializedName(Constants.DateOfBirthFlag.YEAR_ONLY) // "2"
	// @Nullable
	// private DebugTabbarImageInfo secondTabbar;
	SecondTabbar *DebugTabbarImageInfo `json:"2"`
	// @SerializedName("3")
	// @Nullable
	// private DebugTabbarImageInfo thirdTabbar;
	ThirdTabbar *DebugTabbarImageInfo `json:"3"`
}

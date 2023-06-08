package entity

type TabBarIconUrlListInfo struct {
	// @SerializedName(Constants.DateOfBirthFlag.MONTH_YEAR) // "1"
	// @Nullable
	// private TabBarIconUrlItemInfo firstTabBar;
	FirstTabBar *TabBarIconUrlItemInfo `json:"firstTabBar"`
	// @SerializedName("4")
	// @Nullable
	// private TabBarIconUrlItemInfo fourthTabBar;
	FourthTabBar *TabBarIconUrlItemInfo `json:"fourthTabBar"`
	// @SerializedName(Constants.DateOfBirthFlag.YEAR_ONLY) // "2"
	// @Nullable
	// private TabBarIconUrlItemInfo secondTabBar;
	SecondTabBar *TabBarIconUrlItemInfo `json:"secondTabBar"`
	// @SerializedName("3")
	// @Nullable
	// private TabBarIconUrlItemInfo thirdTabBar;
	ThirdTabBar *TabBarIconUrlItemInfo `json:"thirdTabBar"`
}

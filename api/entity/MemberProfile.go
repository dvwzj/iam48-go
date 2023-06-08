package entity

type MemberProfile struct {
	// @SerializedName("birthdate")
	// @Nullable
	// private String birthdate;
	Birthdate *string `json:"birthdate"`
	// @SerializedName(Brand.KEY_BRAND)
	// @Nullable
	// private String brand;
	Brand *string `json:"brand"`
	// @SerializedName("canRedeem")
	// @Nullable
	// private Boolean canRedeem;
	CanRedeem *bool `json:"canRedeem"`
	// @SerializedName("caption")
	// @Nullable
	// private String caption;
	Caption *string `json:"caption"`
	// @SerializedName("city")
	// @Nullable
	// private String city;
	City *string `json:"city"`
	// @SerializedName("cityEn")
	// @Nullable
	// private String cityEn;
	CityEn *string `json:"cityEn"`
	// @SerializedName("codeName")
	// @Nullable
	// private String codeName;
	CodeName *string `json:"codeName"`
	// @SerializedName("country")
	// @Nullable
	// private String country;
	Country *string `json:"country"`
	// @SerializedName("countryEn")
	// @Nullable
	// private String countryEn;
	CountryEn *string `json:"countryEn"`
	// @SerializedName("coverImageUrl")
	// @Nullable
	// private String coverImageUrl;
	CoverImageUrl *string `json:"coverImageUrl"`
	// @SerializedName("displayNameEn")
	// @Nullable
	// private String displayName;
	DisplayName *string `json:"displayName"`
	// @SerializedName("displayName")
	// @Nullable
	// private String displayNameTh;
	DisplayNameTh *string `json:"displayNameTh"`
	// private int filterScore;
	FilterScore int `json:"filterScore"`
	// @SerializedName("graduatedAt")
	// @Nullable
	// private String graduatedAt;
	GraduatedAt *string `json:"graduatedAt"`
	// @SerializedName("hashtags")
	// @Nullable
	// private List<String> hashtags;
	Hashtags *[]string `json:"hashtags"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// private boolean isChecked;
	IsChecked bool `json:"isChecked"`
	// @SerializedName("leftMessage")
	// @Nullable
	// private String leftMessage;
	LeftMessage *string `json:"leftMessage"`
	// @SerializedName("memberImageUrl")
	// @Nullable
	// private String memberImageUrl;
	MemberImageUrl *string `json:"memberImageUrl"`
	// @SerializedName("profileImageUrl")
	// @Nullable
	// private String profileImageUrl;
	ProfileImageUrl *string `json:"profileImageUrl"`
	// @SerializedName("subtitleEn")
	// @Nullable
	// private String subtitle;
	Subtitle *string `json:"subtitle"`
	// @SerializedName("subtitle")
	// @Nullable
	// private String subtitleTh;
	SubtitleTh *string `json:"subtitleTh"`
	// @SerializedName("tabbar")
	// @Nullable
	// private DebugTabbarInfo tabbar;
	Tabbar *DebugTabbarInfo `json:"tabbar"`
	// @SerializedName("version")
	// @Nullable
	// private Long version;
	Version *int64 `json:"version"`
}

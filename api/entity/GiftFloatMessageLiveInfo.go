package entity

type GiftFloatMessageLiveInfo struct {
	// @SerializedName("animationFileUrl")
	// @NotNull
	// private String animationFileUrl;
	AnimationFileUrl string `json:"animationFileUrl"`
	// @SerializedName("coinPrice")
	// private long coinPrice;
	CoinPrice int64 `json:"coinPrice"`
	// @SerializedName("iconImageUrl")
	// @NotNull
	// private String iconImageUrl;
	IconImageUrl string `json:"iconImageUrl"`
	// @SerializedName("id")
	// private long id;
	Id int64 `json:"id"`
	// @SerializedName("isSpecial")
	// private boolean isSpecial;
	IsSpecial bool `json:"isSpecial"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME)
	// @NotNull
	// private String name;
	Name string `json:"name"`
}

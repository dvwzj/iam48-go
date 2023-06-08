package entity

type Skus struct {
	// @SerializedName("animationFileUrl")
	// @Nullable
	// private String animationFileUrl;
	AnimationFileUrl *string `json:"animationFileUrl"`
	// @SerializedName("coinPrice")
	// @Nullable
	// private Long coinPrice;
	CoinPrice *int64 `json:"coinPrice"`
	// @SerializedName("iconImageUrl")
	// @Nullable
	// private String iconImageUrl;
	IconImageUrl *string `json:"iconImageUrl"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("isSpecial")
	// @Nullable
	// private Boolean isSpecial;
	IsSpecial *bool `json:"isSpecial"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "jsonConfig"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY) // "quantity"
	// @Nullable
	// private Long quantity;
	Quantity *int64 `json:"quantity"`
}

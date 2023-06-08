package entity

type GiftListInfo struct {
	// @SerializedName("animationFileUrl")
	// @Nullable
	// private String animationFileUrl;
	AnimationFileUrl *string `json:"animationFileUrl"`
	// @SerializedName("iconImageUrl")
	// @Nullable
	// private String iconImageUrl;
	IconImageUrl *string `json:"iconImageUrl"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
}

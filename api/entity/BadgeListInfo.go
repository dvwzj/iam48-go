package entity

type BadgeListInfo struct {
	// @SerializedName("achievementId")
	// @Nullable
	// private Long achievementId;
	AchievementId *int64 `json:"achievementId"`
	// @SerializedName("badgeImageUrl")
	// @Nullable
	// private String badgeImageUrl;
	BadgeImageUrl *string `json:"badgeImageUrl"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
	// @SerializedName("provinceId")
	// @Nullable
	// private Long provinceId;
	ProvinceId *int64 `json:"provinceId"`
}

package entity

type AchievementInfo struct {
	// @SerializedName("badgeImageUrl")
	// @Nullable
	// private String badgeImageUrl;
	BadgeImageUrl *string `json:"badgeImageUrl"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
	// @SerializedName("score")
	// @Nullable
	// private Long score;
	Score *int64 `json:"score"`
}

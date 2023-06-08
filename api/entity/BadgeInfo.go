package entity

type BadgeInfo struct {
	// @SerializedName("badgeList")
	// @Nullable
	// private List<BadgeListInfo> badgeList;
	BadgeList *[]BadgeListInfo `json:"badgeList"`
	// @SerializedName("isHorizontalDisplay")
	// @Nullable
	// private Boolean isHorizontalDisplay;
	IsHorizontalDisplay *bool `json:"isHorizontalDisplay"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME)
	// @Nullable
	// private String name;
	Name *string `json:"name"`
}

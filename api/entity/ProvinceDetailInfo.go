package entity

type ProvinceDetailInfo struct {
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
	// @SerializedName("latitude")
	// @Nullable
	// private Double latitude;
	Latitude *float64 `json:"latitude"`
	// @SerializedName("longitude")
	// @Nullable
	// private Double longitude;
	Longitude *float64 `json:"longitude"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
	// @SerializedName("recommend")
	// @Nullable
	// private String recommend;
	Recommend *string `json:"recommend"`
	// @SerializedName("slogan")
	// @Nullable
	// private String slogan;
	Slogan *string `json:"slogan"`
}

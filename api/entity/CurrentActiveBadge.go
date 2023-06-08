package entity

type CurrentActiveBadge struct {
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("isActive")
	// @Nullable
	// private Boolean isActive;
	IsActive *bool `json:"isActive"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME)
	// @Nullable
	// private String name;
	Name *string `json:"name"`
}

package entity

type GreetingRedeemTypeInfo struct {
	// @SerializedName("condition")
	// @Nullable
	// private String condition;
	Condition *string `json:"condition"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("iconName")
	// @Nullable
	// private String iconName;
	IconName *string `json:"iconName"`
	// @SerializedName("imageFileUrl")
	// @Nullable
	// private String imageFileUrl;
	ImageFileUrl *string `json:"imageFileUrl"`
	// @SerializedName("isAudioOnly")
	// @Nullable
	// private Boolean isAudioOnly;
	IsAudioOnly *bool `json:"isAudioOnly"`
	// @SerializedName("isRequireDialog")
	// @Nullable
	// private Boolean isRequireDialog;
	IsRequireDialog *bool `json:"isRequireDialog"`
	// @SerializedName("isRequireVideo")
	// @Nullable
	// private Boolean isRequireVideo;
	IsRequireVideo *bool `json:"isRequireVideo"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME)
	// @Nullable
	// private String name;
	Name *string `json:"name"`
	// @SerializedName("redeemTypeId")
	// @Nullable
	// private Long redeemTypeId;
	RedeemTypeId *int64 `json:"redeemTypeId"`
}

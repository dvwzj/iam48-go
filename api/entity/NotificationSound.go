package entity

type NotificationSound struct {
	// @Nullable
	// private String errorMessage;
	ErrorMessage *string `json:"errorMessage"`
	// @SerializedName("fileUrl")
	// @NotNull
	// private String fileUrl;
	FileUrl string `json:"fileUrl"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME)
	// @NotNull
	// private String name;
	Name string `json:"name"`
	// @SerializedName("reference")
	// @NotNull
	// private String reference;
	Reference string `json:"reference"`
	// @Nullable
	// private Uri soundUri;
	SoundUri *string `json:"soundUri"`
	// @SerializedName("tag")
	// @NotNull
	// private String tag;
	Tag string `json:"tag"`
	// @NotNull
	// private String version;
	Version string `json:"version"`
}

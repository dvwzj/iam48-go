package entity

type DeviceInfo struct {
	// @SerializedName("deviceId")
	// @Nullable
	// private String deviceId;
	DeviceId *string `json:"deviceId"`
	// @SerializedName("lastSeenAt")
	// @Nullable
	// private String lastSeenAt;
	LastSeenAt *string `json:"lastSeenAt"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
	// @SerializedName("type")
	// @Nullable
	// private String type;
	Type *string `json:"type"`
}

package entity

type LogoutDeviceBody struct {
	// @SerializedName("deviceId")
	// @Nullable
	// private String deviceId;
	DeviceId *string `json:"deviceId"`
}

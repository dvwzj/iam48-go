package entity

type UserDeviceInfo struct {
	// @SerializedName("deviceModel")
	// @Nullable
	// private String deviceModel;
	DeviceModel *string `json:"deviceModel"`
	// @SerializedName("deviceName")
	// @Nullable
	// private String deviceName;
	DeviceName *string `json:"deviceName"`
}

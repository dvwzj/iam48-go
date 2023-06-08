package entity

type TheaterWatchCameraInfo struct {
	// @SerializedName("cameraId")
	// @Nullable
	// private String cameraId;
	CameraId *string `json:"cameraId"`
	// @SerializedName("cameraName")
	// @Nullable
	// private String cameraName;
	CameraName *string `json:"cameraName"`
	// @SerializedName("hlsUrl")
	// @Nullable
	// private String hlsUrl;
	HlsUrl *string `json:"hlsUrl"`
	// @SerializedName("reference")
	// @Nullable
	// private String reference;
	Reference *string `json:"reference"`
	// @SerializedName("tag")
	// @Nullable
	// private String tag;
	Tag *string `json:"tag"`
}

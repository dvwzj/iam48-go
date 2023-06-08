package entity

type NotificationSoundList struct {
	// @SerializedName("sounds")
	// @Nullable
	// private ArrayList<NotificationSound> sounds;
	Sound *[]NotificationSound `json:"sounds"`
	// @SerializedName("version")
	// @Nullable
	// private String version;
	Version *string `json:"version"`
}

package entity

type PlayMusicInfo struct {
	// @SerializedName("fileUrl")
	// @Nullable
	// private String fileUrl;
	FileUrl *string `json:"fileUrl"`
}

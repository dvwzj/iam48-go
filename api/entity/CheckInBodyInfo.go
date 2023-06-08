package entity

type CheckInBodyInfo struct {
	// @SerializedName("content")
	// @Nullable
	// private String content;
	Content *string `json:"content"`
	// @SerializedName("hash")
	// @Nullable
	// private String hash;
	Hash *string `json:"hash"`
	// @SerializedName("key")
	// @Nullable
	// private String key;
	Key *string `json:"key"`
}

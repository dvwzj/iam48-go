package entity

type GreetingTempUrlResponseInfo struct {
	// @SerializedName("fileId")
	// @Nullable
	// private String fileId;
	FileId *string `json:"fileId"`
	// @SerializedName(ImagesContract.URL) // "url"
	// @Nullable
	// private String url;
	Url *string `json:"url"`
}

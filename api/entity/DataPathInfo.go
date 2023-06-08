package entity

type DataPathInfo struct {
	// @SerializedName("fileId")
	// @Nullable
	// private String fieldId;
	FieldId *string `json:"fileId"`
	// @SerializedName(ImagesContract.URL) // "url"
	// @Nullable
	// private String url;
	Url *string `json:"url"`
}

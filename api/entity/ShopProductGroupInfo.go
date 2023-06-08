package entity

type ShopProductGroupInfo struct {
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("displayIndex")
	// @Nullable
	// private Integer displayIndex;
	DisplayIndex *int `json:"displayIndex"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("isDisable")
	// @Nullable
	// private Boolean isDisable;
	IsDisable *bool `json:"isDisable"`
	// @SerializedName("startedPeriod")
	// @Nullable
	// private String startedPeriod;
	StartedPeriod *string `json:"startedPeriod"`
	// @SerializedName("thumbnailImageUrl")
	// @Nullable
	// private String thumbnailImageUrl;
	ThumbnailImageUrl *string `json:"thumbnailImageUrl"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
}

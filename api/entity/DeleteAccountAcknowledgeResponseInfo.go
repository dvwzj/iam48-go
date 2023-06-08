package entity

type DeleteAccountAcknowledgeResponseInfo struct {
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
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
}

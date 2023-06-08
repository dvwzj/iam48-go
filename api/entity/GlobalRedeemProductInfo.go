package entity

type GlobalRedeemProductInfo struct {
	// @SerializedName("amount")
	// @Nullable
	// private Long amount;
	Amount *int64 `json:"amount"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("displayIndex")
	// @Nullable
	// private Long displayIndex;
	DisplayIndex *int64 `json:"displayIndex"`
	// @SerializedName("expireInDay")
	// @Nullable
	// private String expireInDay;
	ExpireInDay *string `json:"expireInDay"`
	// @SerializedName("imageFileUrl")
	// @Nullable
	// private String imageFileUrl;
	ImageFileUrl *string `json:"imageFileUrl"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE)
	// @Nullable
	// private String title;
	Title *string `json:"title"`
}

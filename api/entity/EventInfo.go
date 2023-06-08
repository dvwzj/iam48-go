package entity

type EventInfo struct {
	// @SerializedName("datePeriod")
	// @Nullable
	// private String datePeriod;
	DatePeriod *string `json:"datePeriod"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("imageFileUrl")
	// @Nullable
	// private String imageFileUrl;
	ImageFileUrl *string `json:"imageFileUrl"`
	// @SerializedName("queueDistributeAt")
	// @Nullable
	// private String queueDistributeAt;
	QueueDistributeAt *string `json:"queueDistributeAt"`
	// @SerializedName("termsAndCondition")
	// @Nullable
	// private String termsAndCondition;
	TermsAndCondition *string `json:"termsAndCondition"`
	// @SerializedName("timePeriod")
	// @Nullable
	// private String timePeriod;
	TimePeriod *string `json:"timePeriod"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE)
	// @Nullable
	// private String title;
	Title *string `json:"title"`
}

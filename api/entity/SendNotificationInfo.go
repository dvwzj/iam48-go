package entity

type SendNotificationInfo struct {
	// @SerializedName("body")
	// @Nullable
	// private String body;
	Body *string `json:"body"`
	// @SerializedName(TPReportParams.PROP_KEY_DATA) // "data"
	// @Nullable
	// private NotificationDataInfo data;
	Data *NotificationDataInfo `json:"data"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
	// @SerializedName("topic")
	// @Nullable
	// private String topic;
	Topic *string `json:"topic"`
}

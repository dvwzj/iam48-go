package entity

type UserReportTopicInfo struct {
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @Nullable
	// private Boolean isSelected;
	IsSelected *bool `json:"isSelected"`
	// @SerializedName("topic")
	// @Nullable
	// private String topic;
	Topic *string `json:"topic"`
}

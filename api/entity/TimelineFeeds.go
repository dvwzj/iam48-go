package entity

type TimelineFeeds struct {
	// @SerializedName("content")
	// @Nullable
	// private TimelineContentItemDataModel content;
	Content *TimelineContentItemDataModel `json:"content"`
	// @SerializedName("discover")
	// @Nullable
	// private TimelineDiscoverDataModel discover;
	Discover *TimelineDiscoverDataModel `json:"discover"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("itemType")
	// @Nullable
	// private String itemType;
	ItemType *string `json:"itemType"`
	// @SerializedName("schedule")
	// @Nullable
	// private TimelineScheduleDataModel schedule;
	Schedule *TimelineScheduleDataModel `json:"schedule"`
}

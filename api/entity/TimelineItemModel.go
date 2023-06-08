package entity

type TimelineItemModel struct {
	// @SerializedName("feeds")
	// @NotNull
	// private List<TimelineFeeds> feeds;
	Feeds []TimelineFeeds `json:"feeds"`
}

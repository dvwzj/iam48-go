package entity

type ScheduleModelInfo struct {
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("endAt")
	// @NotNull
	// private String endAt;
	EndAt string `json:"endAt"`
	// @SerializedName("hashtags")
	// @Nullable
	// private List<String> hashtags;
	Hashtags *[]string `json:"hashtags"`
	// @SerializedName("id")
	// private long id;
	Id int64 `json:"id"`
	// @SerializedName("imageUrl")
	// @NotNull
	// private String imageUrl;
	ImageUrl string `json:"imageUrl"`
	// @SerializedName("isEnd")
	// private boolean isEnd;
	IsEnd bool `json:"isEnd"`
	// @SerializedName("isLive")
	// private boolean isLive;
	IsLive bool `json:"isLive"`
	// @SerializedName("memberId")
	// private long memberId;
	MemberId int64 `json:"memberId"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME)
	// @NotNull
	// private String name;
	Name string `json:"name"`
	// @SerializedName("placeName")
	// @NotNull
	// private String placeName;
	PlaceName string `json:"placeName"`
	// @SerializedName("startAt")
	// @NotNull
	// private String startAt;
	StartAt string `json:"startAt"`
}

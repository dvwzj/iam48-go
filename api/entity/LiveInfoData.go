package entity

type LiveInfoData struct {
	// @SerializedName("endAt")
	// @Nullable
	// private String endAt;
	EndAt *string `json:"endAt"`
	// @SerializedName("hashtags")
	// @Nullable
	// private List<String> hashtags;
	Hashtags *[]string `json:"hashtags"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("imageUrl")
	// @Nullable
	// private String imageUrl;
	ImageUrl *string `json:"imageUrl"`
	// @SerializedName("isEnd")
	// @Nullable
	// private Boolean isEnd;
	IsEnd *bool `json:"isEnd"`
	// @SerializedName("isLive")
	// @Nullable
	// private Boolean isLive;
	IsLive *bool `json:"isLive"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME)
	// @Nullable
	// private String name;
	Name *string `json:"name"`
	// @SerializedName("placeName")
	// @Nullable
	// private String placeName;
	PlaceName *string `json:"placeName"`
	// @SerializedName("ranks")
	// @Nullable
	// private List<LiveVipInfo> ranks;
	Ranks *[]LiveVipInfo `json:"ranks"`
	// @SerializedName("startAt")
	// @Nullable
	// private String startAt;
	StartAt *string `json:"startAt"`
	// @SerializedName("stats")
	// @Nullable
	// private StatisticsInfo stats;
	Stats *StatisticsInfo `json:"stats"`
}

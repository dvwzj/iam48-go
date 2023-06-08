package entity

type CampaignDetailInfo struct {
	// @SerializedName("contentId")
	// @Nullable
	// private Long contentId;
	ContentId *int64 `json:"contentId"`
	// @SerializedName("currentBackedCoinAmount")
	// @Nullable
	// private Long currentBackedCoinAmount;
	CurrentBackedCoinAmount *int64 `json:"currentBackedCoinAmount"`
	// @SerializedName("currentBackerCount")
	// @Nullable
	// private Long currentBackerCount;
	CurrentBackerCount *int64 `json:"currentBackerCount"`
	// @SerializedName("currentDonate")
	// @Nullable
	// private Long currentDonate;
	CurrentDonate *int64 `json:"currentDonate"`
	// @SerializedName("defaultMultiplier")
	// @Nullable
	// private Long defaultMultiplier;
	DefaultMultiplier *int64 `json:"defaultMultiplier"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("displayTextLine1")
	// @Nullable
	// private String displayTextLine1;
	DisplayTextLine1 *string `json:"displayTextLine1"`
	// @SerializedName("displayTextLine2")
	// @Nullable
	// private String displayTextLine2;
	DisplayTextLine2 *string `json:"displayTextLine2"`
	// @SerializedName("endAt")
	// @Nullable
	// private String endAt;
	EndAt *string `json:"endAt"`
	// @SerializedName("goalDescription")
	// @Nullable
	// private String goalDescription;
	GoalDescription *string `json:"goalDescription"`
	// @SerializedName("goalDonate")
	// @Nullable
	// private Long goalDonate;
	GoalDonate *int64 `json:"goalDonate"`
	// @SerializedName("goalTitle")
	// @Nullable
	// private String goalTitle;
	GoalTitle *string `json:"goalTitle"`
	// @SerializedName("hasDetails")
	// @Nullable
	// private Boolean hasDetails;
	HasDetails *bool `json:"hasDetails"`
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
	// @SerializedName("imageUrls")
	// @Nullable
	// private List<String> imageUrls;
	ImageUrls *[]string `json:"imageUrls"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("packages")
	// @Nullable
	// private List<CampaignGifts> packages;
	Packages *[]CampaignGifts `json:"packages"`
	// @SerializedName("progressPercentage")
	// @Nullable
	// private Float progressPercentage;
	ProgressPercentage *float32 `json:"progressPercentage"`
	// @SerializedName("rewards")
	// @Nullable
	// private List<CampaignRewardInfo> rewards;
	Rewards *[]CampaignRewardInfo `json:"rewards"`
	// @SerializedName("shareUrl")
	// @Nullable
	// private String shareUrl;
	ShareUrl *string `json:"shareUrl"`
	// @SerializedName("structureVersion")
	// @Nullable
	// private Integer structureVersion;
	StructureVersion *int `json:"structureVersion"`
	// @SerializedName("targetCoinAmount")
	// @Nullable
	// private Long targetCoinAmount;
	TargetCoinAmount *int64 `json:"targetCoinAmount"`
	// @SerializedName("tierList")
	// @Nullable
	// private List<Long> tierList;
	TierList *[]int64 `json:"tierList"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
	// @SerializedName("totalDonator")
	// @Nullable
	// private Long totalDonator;
	TotalDonator *int64 `json:"totalDonator"`
}

package entity

type CampaignMainItemInfo struct {
	// @SerializedName("currentBackedCoinAmount")
	// @Nullable
	// private Long currentBackedCoinAmount;
	CurrentBackedCoinAmount *int64 `json:"currentBackedCoinAmount"`
	// @SerializedName("currentBackerCount")
	// @Nullable
	// private Long currentBackerCount;
	CurrentBackerCount *int64 `json:"currentBackerCount"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("endAt")
	// @Nullable
	// private String endAt;
	EndAt *string `json:"endAt"`
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
	// @SerializedName("isEnded")
	// @Nullable
	// private Boolean isEnded;
	IsEnded *bool `json:"isEnded"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("progressPercentage")
	// @Nullable
	// private Float progressPercentage;
	ProgressPercentage *float64 `json:"progressPercentage"`
	// @SerializedName("startedAt")
	// @Nullable
	// private String startedAt;
	StartedAt *string `json:"startedAt"`
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
}

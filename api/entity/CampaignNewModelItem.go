package entity

type CampaignNewModelItem struct {
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
	// @SerializedName("targetCoinAmount")
	// @Nullable
	// private Long targetCoinAmount;
	TargetCoinAmount *int64 `json:"targetCoinAmount"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE)
	// @Nullable
	// private String title;
	Title *string `json:"title"`
}

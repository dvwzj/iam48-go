package entity

type TimelineRecommendMediaModel struct {
	// @SerializedName("coinAmount")
	// @Nullable
	// private Long coinAmount;
	CoinAmount *int64 `json:"coinAmount"`
	// @SerializedName("contentId")
	// @Nullable
	// private Long contentId;
	ContentId *int64 `json:"contentId"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("imageFileUrl")
	// @Nullable
	// private String imageFileUrl;
	ImageFileUrl *string `json:"imageFileUrl"`
	// @SerializedName("isPurchaseAllowance")
	// @Nullable
	// private Boolean isPurchaseAllowance;
	IsPurchaseAllowance *bool `json:"isPurchaseAllowance"`
	// @SerializedName("isSubscriptionAllowance")
	// @Nullable
	// private Boolean isSubscriptionAllowance;
	IsSubscriptionAllowance *bool `json:"isSubscriptionAllowance"`
	// @SerializedName("liveAt")
	// @Nullable
	// private String liveAt;
	LiveAt *string `json:"liveAt"`
	// @SerializedName("liveDate")
	// @Nullable
	// private String liveDate;
	LiveDate *string `json:"liveDate"`
	// @SerializedName("livePlace")
	// @Nullable
	// private String livePlace;
	LivePlace *string `json:"livePlace"`
	// @SerializedName("mediaType")
	// @Nullable
	// private String mediaType;
	MediaType *string `json:"mediaType"`
	// @SerializedName("memberIdList")
	// @Nullable
	// private List<Long> memberIdList;
	MemberIdList *[]int64 `json:"memberIdList"`
	// @SerializedName("publishedAt")
	// @Nullable
	// private String publishedAt;
	PublishedAt *string `json:"publishedAt"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
	// @SerializedName("totalWatchTime")
	// @Nullable
	// private Integer totalWatchTime;
	TotalWatchTime *int `json:"totalWatchTime"`
	// @SerializedName("viewType")
	// @Nullable
	// private String viewType;
	ViewType *string `json:"viewType"`
}

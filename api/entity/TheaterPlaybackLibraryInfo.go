package entity

type TheaterPlaybackLibraryInfo struct {
	// @SerializedName("coinAmount")
	// @Nullable
	// private Long coinAmount;
	CoinAmount *int64 `json:"coinAmount"`
	// @SerializedName("contentExpireAt")
	// @Nullable
	// private String contentExpireAt;
	ContentExpireAt *string `json:"contentExpireAt"`
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
	// @SerializedName("isPurchased")
	// @Nullable
	// private Boolean isPurchased;
	IsPurchased *bool `json:"isPurchased"`
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
	// @SerializedName("playableDurationMinute")
	// @Nullable
	// private Long playableDurationMinute;
	PlayableDurationMinute *int64 `json:"playableDurationMinute"`
	// @SerializedName("playableExpireAt")
	// @Nullable
	// private String playableExpireAt;
	PlayableExpireAt *string `json:"playableExpireAt"`
	// @SerializedName("publishedAt")
	// @Nullable
	// private String publishedAt;
	PublishedAt *string `json:"publishedAt"`
	// @SerializedName("purchaseableInDay")
	// @Nullable
	// private Integer purchaseableInDay;
	PurchaseableInDay *int `json:"purchaseableInDay"`
	// @SerializedName("subscriptionAllowance")
	// @Nullable
	// private Boolean subscriptionAllowance;
	SubscriptionAllowance *bool `json:"subscriptionAllowance"`
	// @SerializedName("theaterId")
	// @Nullable
	// private Long theaterId;
	TheaterId *int64 `json:"theaterId"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
	// @SerializedName("totalWatchTime")
	// @Nullable
	// private Long totalWatchTime;
	TotalWatchTime *int64 `json:"totalWatchTime"`
}

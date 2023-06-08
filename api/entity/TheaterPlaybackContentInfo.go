package entity

type TheaterPlaybackContentInfo struct {
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
	// @SerializedName("memberIdList")
	// @Nullable
	// private List<Long> memberIdList;
	MemberIdList *[]int64 `json:"memberIdList"`
	// @SerializedName("playableDurationMinute")
	// @Nullable
	// private Long playableDurationMinute;
	PlayableDurationMinute *int64 `json:"playableDurationMinute"`
	// @SerializedName("theaterId")
	// @Nullable
	// private Long theaterId;
	TheaterId *int64 `json:"theaterId"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE)
	// @Nullable
	// private String title;
	Title *string `json:"title"`
	// @SerializedName("totalWatchTime")
	// @Nullable
	// private Long totalWatchTime;
	TotalWatchTime *int64 `json:"totalWatchTime"`
}

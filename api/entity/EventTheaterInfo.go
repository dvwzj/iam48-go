package entity

type EventTheaterInfo struct {
	// @SerializedName("coinAmount")
	// @Nullable
	// private Integer coinAmount;
	CoinAmount *int `json:"coinAmount"`
	// @SerializedName("dateFrom")
	// @Nullable
	// private String dateFrom;
	DateFrom *string `json:"dateFrom"`
	// @SerializedName("dateTo")
	// @Nullable
	// private String dateTo;
	DateTo *string `json:"dateTo"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("imageFileUrl")
	// @Nullable
	// private String imageFileUrl;
	ImageFileUrl *string `json:"imageFileUrl"`
	// @SerializedName("isLive")
	// @Nullable
	// private Boolean isLive;
	IsLive *bool `json:"isLive"`
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
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
}

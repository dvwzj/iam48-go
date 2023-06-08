package entity

type ScheduleEvent struct {
	// @SerializedName("canPurchase")
	// @Nullable
	// private Boolean canPurchase;
	CanPurchase *bool `json:"canPurchase"`
	// @SerializedName("contentId")
	// @Nullable
	// private Long contentId;
	ContentId *int64 `json:"contentId"`
	// @SerializedName("date")
	// @Nullable
	// private String date;
	Date *string `json:"date"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("eventVideoId")
	// @Nullable
	// private String eventVideoId;
	EventVideoId *string `json:"eventVideoId"`
	// @SerializedName(ConstancesKt.KEY_FROM) // "from"
	// @Nullable
	// private String from;
	From *string `json:"from"`
	// @SerializedName("hashTags")
	// @Nullable
	// private ArrayList<String> hashTags;
	HashTags *[]string `json:"hashTags"`
	// @SerializedName("id")
	// private long id;
	Id *int64 `json:"id"`
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
	// private List<Integer> memberIdList;
	MemberIdList *[]int64 `json:"memberIdList"`
	// @SerializedName("placeName")
	// @Nullable
	// private String placeName;
	PlaceName *string `json:"placeName"`
	// @SerializedName("referenceItemId")
	// @Nullable
	// private Long referenceItemId;
	ReferenceItemId *int64 `json:"referenceItemId"`
	// @SerializedName("referenceUrl")
	// @Nullable
	// private String referenceUrl;
	ReferenceUrl *string `json:"referenceUrl"`
	// @SerializedName("shareUrl")
	// @Nullable
	// private String shareUrl;
	ShareUrl *string `json:"shareUrl"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
	// @SerializedName(ConstancesKt.KEY_TO) // "to"
	// @Nullable
	// private String to;
	To *string `json:"to"`
	// @SerializedName("type")
	// @Nullable
	// private String type;
	Type *string `json:"type"`
}

package entity

type NotificationInfo struct {
	// @SerializedName("caption")
	// @Nullable
	// private String caption;
	Caption *string `json:"caption"`
	// @SerializedName("content")
	// @Nullable
	// private String content;
	Content *string `json:"content"`
	// @Nullable
	// private String imageUrl = "";
	ImageUrl *string `json:"imageUrl"`
	// @Nullable
	// private Boolean isNew = Boolean.FALSE;
	IsNew *bool `json:"isNew"`
	// @SerializedName("isRead")
	// @Nullable
	// private Boolean isRead;
	IsRead *bool `json:"isRead"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("notificationId")
	// @Nullable
	// private Long notificationId;
	NotificationId *int64 `json:"notificationId"`
	// @SerializedName("publishedAt")
	// @Nullable
	// private String publishedAt;
	PublishedAt *string `json:"publishedAt"`
	// @SerializedName(SchemeController.FCM_PARAM_SCHEMELINK)
	// @Nullable
	// private String schemeLink;
	SchemeLink *string `json:"schemeLink"`
	// @SerializedName("topic")
	// @Nullable
	// private String topic;
	Topic *string `json:"topic"`
}

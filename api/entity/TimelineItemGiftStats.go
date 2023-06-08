package entity

type TimelineItemGiftStats struct {
	// @SerializedName("badgeId")
	// @Nullable
	// private Long badgeId;
	BadgeId *int64 `json:"badgeId"`
	// @SerializedName("giftAmount")
	// @Nullable
	// private Long giftAmount;
	GiftAmount *int64 `json:"giftAmount"`
	// @SerializedName("imageProfileUrl")
	// @Nullable
	// private String imageProfileUrl;
	ImageProfileUrl *string `json:"imageProfileUrl"`
	// @SerializedName("nullDisplayName")
	// @Nullable
	// private String nullDisplayName;
	NullDisplayName *string `json:"nullDisplayName"`
	// @SerializedName("rankNo")
	// @Nullable
	// private Long rankNo;
	RankNo *int64 `json:"rankNo"`
	// @SerializedName("userDisplayName")
	// @Nullable
	// private String userDisplayName;
	UserDisplayName *string `json:"userDisplayName"`
	// @SerializedName("userId")
	// @Nullable
	// private Long userId;
	UserId *int64 `json:"userId"`
}

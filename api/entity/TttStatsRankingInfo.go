package entity

type TttStatsRankingInfo struct {
	// @SerializedName("badgeId")
	// @Nullable
	// private Long badgeId;
	BadgeId *int64 `json:"badgeId"`
	// @SerializedName("displayName")
	// @Nullable
	// private String displayName;
	DisplayName *string `json:"displayName"`
	// @SerializedName("imageFileUrl")
	// @Nullable
	// private String imageFileUrl;
	ImageFileUrl *string `json:"imageFileUrl"`
	// @SerializedName("rankNo")
	// @Nullable
	// private Long rankNo;
	RankNo *int64 `json:"rankNo"`
	// @SerializedName("score")
	// @Nullable
	// private Long score;
	Score *int64 `json:"score"`
	// @SerializedName("userId")
	// @Nullable
	// private Long userId;
	UserId *int64 `json:"userId"`
}

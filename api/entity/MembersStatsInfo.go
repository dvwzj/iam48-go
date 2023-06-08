package entity

type MembersStatsInfo struct {
	// @SerializedName("badgeId")
	// @Nullable
	// private Long badgeId;
	BadgeId *int64 `json:"badgeId"`
	// @SerializedName("displayName")
	// @Nullable
	// private String displayName;
	DisplayName *string `json:"displayName"`
	// @SerializedName("fullName")
	// @Nullable
	// private String fullName;
	FullName *string `json:"fullName"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("stats")
	// @Nullable
	// private RankStats stats;
	Stats *RankStats `json:"stats"`
	// @SerializedName("thumbnailUrl")
	// @Nullable
	// private String thumbnailUrl;
	ThumbnailUrl *string `json:"thumbnailUrl"`
}

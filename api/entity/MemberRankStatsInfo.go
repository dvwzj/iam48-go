package entity

type MemberRankStatsInfo struct {
	// @SerializedName("displayName")
	// @Nullable
	// private String displayName;
	DisplayName *string `json:"displayName"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("thumbnailUrl")
	// @Nullable
	// private String thumbnailUrl;
	ThumbnailUrl *string `json:"thumbnailUrl"`
}

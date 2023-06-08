package entity

type PublicUserProfileInfo struct {
	// @SerializedName("badgeId")
	// @Nullable
	// private Long badgeId;
	BadgeId *int64 `json:"badgeId"`
	// @SerializedName("caption")
	// @Nullable
	// private String caption;
	Caption *string `json:"caption"`
	// @SerializedName("coverImageUrl")
	// @Nullable
	// private String coverImageUrl;
	CoverImageUrl *string `json:"coverImageUrl"`
	// @SerializedName("displayName")
	// @Nullable
	// private String displayName;
	DisplayName *string `json:"displayName"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("profileImageUrl")
	// @Nullable
	// private String profileImageUrl;
	ProfileImageUrl *string `json:"profileImageUrl"`
	// @SerializedName("totalFollowing")
	// @Nullable
	// private Long totalFollowing;
	TotalFollowing *int64 `json:"totalFollowing"`
	// @SerializedName("totalGiftAmount")
	// @Nullable
	// private Long totalGiftAmount;
	TotalGiftAmount *int64 `json:"totalGiftAmount"`
}

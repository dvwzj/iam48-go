package entity

type CommenterInfo struct {
	// @SerializedName("badgeId")
	// @Nullable
	// private Long badgeId;
	BadgeId *int64 `json:"badgeId"`
	// @SerializedName("displayImageUrl")
	// @Nullable
	// private String displayImageUrl;
	DisplayImageUrl *string `json:"displayImageUrl"`
	// @SerializedName("displayName")
	// @Nullable
	// private String displayName;
	DisplayName *string `json:"displayName"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("userId")
	// @Nullable
	// private Long userId;
	UserId *int64 `json:"userId"`
}

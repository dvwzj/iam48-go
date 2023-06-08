package entity

type EditUserProfileInfo struct {
	// @SerializedName("tempCoverImageId")
	// @Nullable
	// private String tempCoverImageId = "";
	TempCoverImageId *string `json:"tempCoverImageId"`
	// @SerializedName("tempProfileImageId")
	// @Nullable
	// private String tempProfileImageId = "";
	TempProfileImageId *string `json:"tempProfileImageId"`
	// @SerializedName("displayName")
	// @Nullable
	// private String displayName = "";
	DisplayName *string `json:"displayName"`
	// @SerializedName("profileImageUrl")
	// @Nullable
	// private String profileImageUrl = "";
	ProfileImageUrl *string `json:"profileImageUrl"`
	// @SerializedName("coverImageUrl")
	// @Nullable
	// private String coverImageUrl = "";
	CoverImageUrl *string `json:"coverImageUrl"`
	// @SerializedName("caption")
	// @Nullable
	// private String caption = "";
	Caption *string `json:"caption"`
}

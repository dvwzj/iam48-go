package entity

type UserProfileInfo struct {
	// @SerializedName("caption")
	// @Nullable
	// private String caption;
	Caption *string `json:"caption"`
	// @SerializedName("codeName")
	// @Nullable
	// private String codeName;
	CodeName *string `json:"codeName"`
	// @SerializedName("coverImageUrl")
	// @Nullable
	// private String coverImageUrl;
	CoverImageUrl *string `json:"coverImageUrl"`
	// @SerializedName("displayName")
	// @Nullable
	// private String displayName;
	DisplayName *string `json:"displayName"`
	// @SerializedName("firstName")
	// @Nullable
	// private String firstName;
	FirstName *string `json:"firstName"`
	// @SerializedName("fullName")
	// @Nullable
	// private String fullName;
	FullName *string `json:"fullName"`
	// @SerializedName("id")
	// private long id;
	Id int64 `json:"id"`
	// @SerializedName("imageCoverUrl")
	// @Nullable
	// private String imageCoverUrl;
	ImageCoverUrl *string `json:"imageCoverUrl"`
	// @SerializedName("imageUrl")
	// @Nullable
	// private String imageUrl;
	ImageUrl *string `json:"imageUrl"`
	// @SerializedName("nullDisplayName")
	// @Nullable
	// private String nullDisplayName;
	NullDisplayName *string `json:"nullDisplayName"`
	// @SerializedName("profileImageUrl")
	// @Nullable
	// private String profileImageUrl;
	ProfileImageUrl *string `json:"profileImageUrl"`
	// @SerializedName("subtitle")
	// @Nullable
	// private String subtitle;
	Subtitle *string `json:"subtitle"`
}

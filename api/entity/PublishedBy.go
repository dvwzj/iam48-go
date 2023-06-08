package entity

type PublishedBy struct {
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
}

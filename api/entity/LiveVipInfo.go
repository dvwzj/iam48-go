package entity

type LiveVipInfo struct {
	// @SerializedName("userId")
	// @Nullable
	// private Long userId;
	UserId *int64 `json:"userId"`
	// @SerializedName("userImageProfleUrl")
	// @Nullable
	// private String userImageProfleUrl;
	UserImageProfleUrl *string `json:"userImageProfleUrl"`
}

package entity

type VideoInfo struct {
	// @SerializedName("hashtags")
	// @Nullable
	// private String hashtags;
	Hashtags *string `json:"hashtags"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("isLive")
	// @Nullable
	// private Boolean isLive;
	IsLive *bool `json:"isLive"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
	// @SerializedName("publishedAt")
	// @Nullable
	// private String publishedAt;
	PublishedAt *string `json:"publishedAt"`
	// @SerializedName("thumbnailImageUrl")
	// @Nullable
	// private String thumbnailImageUrl;
	ThumbnailImageUrl *string `json:"thumbnailImageUrl"`
}

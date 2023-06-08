package entity

type ContentSocial struct {
	// @SerializedName("commentCount")
	// private long commentCount;
	CommentCount int64 `json:"commentCount"`
	// @SerializedName("comments")
	// @NotNull
	// private List<SocialCommentInfo> comments;
	Comments []SocialCommentInfo `json:"comments"`
	// @SerializedName("contentImageUrl")
	// @Nullable
	// private String contentImageUrl;
	ContentImageUrl *string `json:"contentImageUrl"`
	// @SerializedName("contentLabel")
	// @Nullable
	// private String contentLabel;
	ContentLabel *string `json:"contentLabel"`
	// @SerializedName("likeCount")
	// private long likeCount;
	LikeCount int64 `json:"likeCount"`
	// @NotNull
	// private String localId;
	LocalId string `json:"localId"`
	// @SerializedName("location")
	// @Nullable
	// private String location;
	Location *string `json:"location"`
	// @SerializedName("originalContentImageUrl")
	// @Nullable
	// private String originalContentImageUrl;
	OriginalContentImageUrl *string `json:"originalContentImageUrl"`
	// @SerializedName("originalUserImageUrl")
	// @Nullable
	// private String originalUserImageUrl;
	OriginalUserImageUrl *string `json:"originalUserImageUrl"`
	// @SerializedName("postDate")
	// @Nullable
	// private String postDate;
	PostDate *string `json:"postDate"`
	// @SerializedName("retweetCount")
	// private long retweetCount;
	RetweetCount int64 `json:"retweetCount"`
	// @SerializedName("userImageUrl")
	// @Nullable
	// private String userImageUrl;
	UserImageUrl *string `json:"userImageUrl"`
	// @SerializedName("username")
	// @Nullable
	// private String username;
	Username *string `json:"username"`
}

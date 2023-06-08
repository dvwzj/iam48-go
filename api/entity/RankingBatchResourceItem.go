package entity

type RankingBatchResourceItem struct {
	// @SerializedName("backgroundImageUrl")
	// @Nullable
	// private String backgroundImageUrl;
	BackgroundImageUrl *string `json:"backgroundImageUrl"`
	// @SerializedName("memberImageUrl")
	// @Nullable
	// private String memberImageUrl;
	MemberImageUrl *string `json:"memberImageUrl"`
}

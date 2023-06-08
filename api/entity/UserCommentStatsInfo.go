package entity

type UserCommentStatsInfo struct {
	// @SerializedName("canEdit")
	// @Nullable
	// private Boolean canEdit;
	CanEdit *bool `json:"canEdit"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("isHidden")
	// @Nullable
	// private Boolean isHidden;
	IsHidden *bool `json:"isHidden"`
	// @SerializedName("isLoved")
	// @Nullable
	// private Boolean isLoved;
	IsLoved *bool `json:"isLoved"`
	// @SerializedName("totalGifts")
	// @Nullable
	// private Long totalGifts;
	TotalGifts *int64 `json:"totalGifts"`
}

package entity

type SegmentInfo struct {
	// @SerializedName("alignment")
	// @Nullable
	// private String alignment;
	Alignment *string `json:"alignment"`
	// @SerializedName("contentSocial")
	// @Nullable
	// private ContentSocial contentSocial;
	ContentSocial *ContentSocial `json:"contentSocial"`
	// @SerializedName("createdAt")
	// @Nullable
	// private String createdAt;
	CreatedAt *string `json:"createdAt"`
	// @SerializedName("imageUrl")
	// @Nullable
	// private String imageUrl;
	ImageUrl *string `json:"imageUrl"`
	// @SerializedName("originalImageUrl")
	// @Nullable
	// private String originalImageUrl;
	OriginalImageUrl *string `json:"originalImageUrl"`
	// @SerializedName("segmentId")
	// @Nullable
	// private String segmentId;
	SegmentId *string `json:"segmentId"`
	// @SerializedName("styleFlags")
	// @Nullable
	// private List<? extends SegmentStyleFlag> styleFlags;
	StyleFlags *[]SegmentStyleFlag `json:"styleFlags"`
	// @SerializedName("textNodes")
	// @Nullable
	// private List<TextNodesInfo> textNodes;
	TextNodes *[]TextNodesInfo `json:"textNodes"`
	// @SerializedName("type")
	// @Nullable
	// private String type;
	Type *string `json:"type"`
	// @SerializedName("updatedAt")
	// @Nullable
	// private String updatedAt;
	UpdatedAt *string `json:"updatedAt"`
}

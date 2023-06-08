package entity

type ThankYouContentInfo struct {
	// @SerializedName("batchId")
	// @Nullable
	// private Long batchId;
	BatchId *int64 `json:"batchId"`
	// @SerializedName("contentText")
	// @Nullable
	// private String contentText;
	ContentText *string `json:"contentText"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("imageUrlList")
	// @NotNull
	// private ArrayList<String> imageUrlList;
	ImageUrlList []string `json:"imageUrlList"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("thumbnailImageUrl")
	// @Nullable
	// private String thumbnailImageUrl;
	ThumbnailImageUrl *string `json:"thumbnailImageUrl"`
}

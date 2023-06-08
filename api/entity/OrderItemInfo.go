package entity

type OrderItemInfo struct {
	// @SerializedName("price")
	// @Nullable
	// private Double price;
	Price *float64 `json:"price"`
	// @SerializedName("productId")
	// @Nullable
	// private Long productId;
	ProductId *int64 `json:"productId"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY) // "quantity"
	// @Nullable
	// private Integer quantity;
	Quantity *int64 `json:"quantity"`
	// @SerializedName("thumbnailImageUrl")
	// @Nullable
	// private String thumbnailImageUrl;
	ThumbnailImageUrl *string `json:"thumbnailImageUrl"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
}

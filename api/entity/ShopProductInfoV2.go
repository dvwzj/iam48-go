package entity

type ShopProductInfoV2 struct {
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("productList")
	// @Nullable
	// private List<ShopProductInfo> productList;
	ProductList *[]ShopProductInfo `json:"productList"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
}

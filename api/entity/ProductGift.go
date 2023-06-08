package entity

type ProductGift struct {
	// @SerializedName("categories")
	// @Nullable
	// private List<Category> categories;
	Categories *[]Category `json:"categories"`
}

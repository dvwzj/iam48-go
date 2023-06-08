package entity

type ShopProductInfo struct {
	// private int cartQuantity;
	CartQuantity int `json:"cartQuantity"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("endAt")
	// @Nullable
	// private String endAt;
	EndAt *string `json:"endAt"`
	// @SerializedName("estimatedDays")
	// @Nullable
	// private Integer estimatedDays;
	EstimatedDays *int `json:"estimatedDays"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("imageFileUrlList")
	// @Nullable
	// private List<String> imageFileUrlList;
	ImageFileUrlList *[]string `json:"imageFileUrlList"`
	// @SerializedName("isDisable")
	// @Nullable
	// private Boolean isDisable;
	IsDisable *bool `json:"isDisable"`
	// @SerializedName("maxQuantity")
	// @Nullable
	// private Integer maxQuantity;
	MaxQuantity *int `json:"maxQuantity"`
	// @SerializedName("packageGroupId")
	// @Nullable
	// private Long packageGroupId;
	PackageGroupId *int64 `json:"packageGroupId"`
	// @SerializedName("price")
	// @Nullable
	// private Double price;
	Price *float64 `json:"price"`
	// @SerializedName("productGroupId")
	// @Nullable
	// private Long productGroupId;
	ProductGroupId *int64 `json:"productGroupId"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY) // "quantity"
	// @Nullable
	// private Integer quantity;
	Quantity *int `json:"quantity"`
	// @SerializedName("startAt")
	// @Nullable
	// private String startAt;
	StartAt *string `json:"startAt"`
	// @SerializedName("thumbnailImageUrl")
	// @Nullable
	// private String thumbnailImageUrl;
	ThumbnailImageUrl *string `json:"thumbnailImageUrl"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
	// @SerializedName("weight")
	// @Nullable
	// private Double weight;
	Weight *float64 `json:"weight"`
}

package entity

type SortOrdersItem struct {
	// @SerializedName("indices")
	// @Nullable
	// private List<Long> indices;
	Indices *[]int64 `json:"indices"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
	// @SerializedName("nameEn")
	// @Nullable
	// private String nameEn;
	NameEn *string `json:"nameEn"`
}

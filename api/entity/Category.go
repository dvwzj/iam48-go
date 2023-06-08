package entity

type Category struct {
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
	// @SerializedName("skus")
	// @Nullable
	// private List<Skus> skus;
	Skus *[]Skus `json:"skus"`
}

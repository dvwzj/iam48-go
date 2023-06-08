package entity

type PurchaseGifts struct {
	// @SerializedName("gifts")
	// @Nullable
	// private List<Gifts> gifts;
	Gifts *[]Gifts `json:"gifts"`
	// @SerializedName("id")
	// @Nullable
	// private Integer id;
	Id *int `json:"id"`
	// @SerializedName("isSpecial")
	// @Nullable
	// private Boolean isSpecial;
	IsSpecial *bool `json:"isSpecial"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
}

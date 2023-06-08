package entity

type DiscoverLinks struct {
	// @SerializedName("seeAll")
	// @Nullable
	// private String seeAll;
	SeeAll *string `json:"seeAll"`
}

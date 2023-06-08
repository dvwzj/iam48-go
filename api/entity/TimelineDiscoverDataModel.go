package entity

type TimelineDiscoverDataModel struct {
	// @SerializedName("links")
	// @Nullable
	// private DiscoverLinks links;
	Links *DiscoverLinks `json:"links"`
	// @SerializedName("resource")
	// @Nullable
	// private ResourceItem resource;
	Resource *ResourceItem `json:"resource"`
}

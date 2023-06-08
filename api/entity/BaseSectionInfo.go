package entity

type BaseSectionInfo struct {
	// @SerializedName("items")
	// @NotNull
	// private List<? extends Object> items;
	Items []any `json:"items"`
	// @SerializedName("links")
	// @Nullable
	// private DiscoverSectionItemLink links;
	Links *DiscoverSectionItemLink `json:"links"`
	// @SerializedName("resource")
	// @Nullable
	// private DiscoverSelectionResourceItem resource;
	Resource *DiscoverSelectionResourceItem `json:"resource"`
	// @SerializedName("type")
	// @Nullable
	// private String type;
	Type *string `json:"type"`
}

package entity

// Moved from /models/theater to /models

type DiscoverSelectionResourceItem struct {
	// @SerializedName("imageurl")
	// @Nullable
	// private String imageurl;
	Imageurl *string `json:"imageurl"`
}

package entity

type TextNodesInfo struct {
	// @SerializedName("color")
	// @Nullable
	// private String color;
	Color *string `json:"color"`
	// @SerializedName("styleFlags")
	// @Nullable
	// private List<? extends SegmentStyleFlag> styleFlags;
	StyleFlags *[]SegmentStyleFlag `json:"styleFlags"`
	// @SerializedName("text")
	// @Nullable
	// private String text;
	Text *string `json:"text"`
}

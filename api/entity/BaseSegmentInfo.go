package entity

type BaseSegmentInfo struct {
	// @SerializedName("segments")
	// @Nullable
	// private List<SegmentInfo> segments;
	Segments *[]SegmentInfo `json:"segments"`
}

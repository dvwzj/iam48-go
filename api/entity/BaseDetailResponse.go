package entity

type BaseDetailResponse struct {
	// @SerializedName("details")
	// @Nullable
	// private BaseSegmentInfo details;
	Details *BaseSegmentInfo `json:"details"`
}

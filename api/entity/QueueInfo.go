package entity

type QueueInfo struct {
	// @SerializedName("totalQueueBeforeMe")
	// @Nullable
	// private Long totalQueueBeforeMe;
	TotalQueueBeforeMe *int64 `json:"totalQueueBeforeMe"`
	// @SerializedName("totalSecondsBeforeMe")
	// @Nullable
	// private Long totalSecondsBeforeMe;
	TotalSecondsBeforeMe *int64 `json:"totalSecondsBeforeMe"`
}

package entity

type FloatMessageInfo struct {
	// @SerializedName("currentRankNo")
	// @Nullable
	// private Long currentRankNo;
	CurrentRankNo *int64 `json:"currentRankNo"`
	// @SerializedName("displayPriority")
	// private long displayPriority;
	DisplayPriority int64 `json:"displayPriority"`
	// @SerializedName("id")
	// private long id;
	Id int64 `json:"id"`
}

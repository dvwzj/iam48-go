package entity

type MyScoreTransactionInfo struct {
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("label")
	// @Nullable
	// private String label;
	Label *string `json:"label"`
	// @SerializedName("score")
	// @Nullable
	// private Long score;
	Score *int64 `json:"score"`
	// @SerializedName("timestamp")
	// @Nullable
	// private String timestamp;
	Timestamp *string `json:"timestamp"`
}

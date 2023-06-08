package entity

type MeetingReviewRequest struct {
	// @SerializedName(ConstancesKt.KEY_MESSAGE) // "message"
	// @NotNull
	// private String message;
	Message string `json:"message"`
	// @SerializedName(ConstancesKt.KEY_REF_CODE) // "refCode"
	// @NotNull
	// private String refCode;
	RefCode string `json:"refCode"`
	// @SerializedName("score")
	// private long score;
	Score int64 `json:"score"`
}

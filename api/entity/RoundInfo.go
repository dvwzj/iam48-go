package entity

type RoundInfo struct {
	// @SerializedName("endEventDateTime")
	// @Nullable
	// private String endEventDateTime;
	EndEventDateTime *string `json:"endEventDateTime"`
	// @SerializedName("eventDateTime")
	// @Nullable
	// private String eventDateTime;
	EventDateTime *string `json:"eventDateTime"`
	// @SerializedName("isSpecialRound")
	// @Nullable
	// private Boolean isSpecialRound;
	IsSpecialRound *bool `json:"isSpecialRound"`
}

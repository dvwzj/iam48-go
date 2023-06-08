package entity

type GreetingUserDialogInfo struct {
	// @SerializedName("dialogSubmittedAt")
	// @Nullable
	// private String dialogSubmittedAt;
	DialogSubmittedAt *string `json:"dialogSubmittedAt"`
	// @SerializedName("dialogTemplate")
	// @Nullable
	// private String dialogTemplate;
	DialogTemplate *string `json:"dialogTemplate"`
	// @SerializedName("isAudioOnly")
	// @Nullable
	// private Boolean isAudioOnly;
	IsAudioOnly *bool `json:"isAudioOnly"`
	// @SerializedName("pronunciationName")
	// @Nullable
	// private String pronunciationName;
	PronunciationName *string `json:"pronunciationName"`
	// @SerializedName("userSpecificName")
	// @Nullable
	// private String userSpecificName;
	UserSpecificName *string `json:"userSpecificName"`
}

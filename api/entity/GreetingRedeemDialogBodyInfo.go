package entity

type GreetingRedeemDialogBodyInfo struct {
	// @SerializedName("dialogTemplateId")
	// @Nullable
	// private Long dialogTemplateId;
	DialogTemplateId *int64 `json:"dialogTemplateId"`
	// @SerializedName("displayName")
	// @Nullable
	// private String displayName;
	DisplayName *string `json:"displayName"`
	// @SerializedName("nameTemplateId")
	// @Nullable
	// private Long nameTemplateId;
	NameTemplateId *int64 `json:"nameTemplateId"`
	// @SerializedName("pronunciationName")
	// @Nullable
	// private String pronunciationName;
	PronunciationName *string `json:"pronunciationName"`
}

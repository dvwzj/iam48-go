package entity

type TokenXTermsResponseInfo struct {
	// @SerializedName("isSubmission")
	// @Nullable
	// private Boolean isSubmission;
	IsSubmission *bool `json:"isSubmission"`
	// @SerializedName("versionNo")
	// @Nullable
	// private String versionNo;
	VersionNo *string `json:"versionNo"`
}

package entity

type SecurityStatusResponseInfo struct {
	// @SerializedName("emailAddress")
	// @Nullable
	// private String emailAddress;
	EmailAddress *string `json:"emailAddress"`
	// @SerializedName("is2FAActive")
	// @Nullable
	// private Boolean is2FAActive;
	Is2FAActive *bool `json:"is2FAActive"`
}

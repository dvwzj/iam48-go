package entity

type VerifyFirstOTPEmailBodyInfo struct {
	// @SerializedName("otpPin")
	// @Nullable
	// private String otpPin;
	OtpPin *string `json:"otpPin"`
	// @SerializedName("referenceKey")
	// @Nullable
	// private String referenceKey;
	ReferenceKey *string `json:"referenceKey"`
}

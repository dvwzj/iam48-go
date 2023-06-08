package entity

type SendOTPEmailResponseInfo struct {
	// @SerializedName("emailOtpReference")
	// @Nullable
	// private final String emailOtpReference;
	EmailOtpReference *string `json:"emailOtpReference"`
	// @SerializedName("resendOtpAvailableInSeconds")
	// @Nullable
	// private final Integer resendOtpAvailableInSeconds;
	ResendOtpAvailableInSeconds *int `json:"resendOtpAvailableInSeconds"`
}

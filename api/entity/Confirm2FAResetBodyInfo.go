package entity

type Confirm2FAResetBodyInfo struct {
	// @SerializedName("emailOtpPin")
	// @Nullable
	// private String emailOtpPin;
	EmailOtpPin *string `json:"emailOtpPin"`
	// @SerializedName("emailOtpReference")
	// @Nullable
	// private String emailOtpReference;
	EmailOtpReference *string `json:"emailOtpReference"`
	// @SerializedName("twoFactorPin")
	// @Nullable
	// private String twoFactorPin;
	TwoFactorPin *string `json:"twoFactorPin"`
	// @SerializedName("twoFactorReference")
	// @Nullable
	// private String twoFactorReference;
	TwoFactorReference *string `json:"twoFactorReference"`
}

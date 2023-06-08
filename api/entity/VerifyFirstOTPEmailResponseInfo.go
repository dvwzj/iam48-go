package entity

type VerifyFirstOTPEmailResponseInfo struct {
	// @SerializedName("qrCode")
	// @NotNull
	// private final String qrCode;
	QrCode string `json:"qrCode"`
	// @SerializedName("setupKey")
	// @NotNull
	// private final String setupKey;
	SetupKey string `json:"setupKey"`
	// @SerializedName("twoFactorExpireAt")
	// @NotNull
	// private final String twoFactorExpireAt;
	TwoFactorExpireAt string `json:"twoFactorExpireAt"`
	// @SerializedName("twoFactorReference")
	// @NotNull
	// private final String twoFactorReference;
	TwoFactorReference string `json:"twoFactorReference"`
}

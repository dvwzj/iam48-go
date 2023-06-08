package entity

type Initial2FASetupResponseInfo struct {
	// @SerializedName("qrCode")
	// @Nullable
	// private String qrCode;
	QrCode *string `json:"qrCode"`
	// @SerializedName("setupKey")
	// @Nullable
	// private String setupKey;
	SetupKey *string `json:"setupKey"`
	// @SerializedName("twoFactorExpireAt")
	// @Nullable
	// private String twoFactorExpireAt;
	TwoFactorExpireAt *string `json:"twoFactorExpireAt"`
	// @SerializedName("twoFactorReference")
	// @Nullable
	// private String twoFactorReference;
	TwoFactorReference *string `json:"twoFactorReference"`
}

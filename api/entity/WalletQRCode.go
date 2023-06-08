package entity

type WalletQRCode struct {
	// @SerializedName("svgQrCode")
	// @Nullable
	// private String svgQrCode;
	SvgQrCode *string `json:"svgQrCode"`
}

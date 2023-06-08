package entity

type WalletInfoModel struct {
	// @SerializedName("walletCode")
	// @Nullable
	// private String walletCode;
	WalletCode *string `json:"walletCode"`
	// @SerializedName("walletImageUrl")
	// @Nullable
	// private String walletImageUrl;
	WalletImageUrl *string `json:"walletImageUrl"`
}

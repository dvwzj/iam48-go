package entity

type WalletPinModel struct {
	// @SerializedName("pin")
	// @Nullable
	// private String pin;
	Pin *string `json:"pin"`
}

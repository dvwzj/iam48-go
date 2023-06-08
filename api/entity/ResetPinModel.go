package entity

type ResetPinModel struct {
	// @SerializedName("pin")
	// @Nullable
	// private String pin;
	Pin *string `json:"pin"`
	// @SerializedName("pinResetCode")
	// @Nullable
	// private String pinResetCode;
	PinResetCode *string `json:"pinResetCode"`
}

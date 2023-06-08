package entity

type ChangePinModel struct {
	// @SerializedName("oldPin")
	// @Nullable
	// private String oldPin;
	OldPin *string `json:"oldPin"`
	// @SerializedName("pin")
	// @Nullable
	// private String pin;
	Pin *string `json:"pin"`
}

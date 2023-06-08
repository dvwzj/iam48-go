package entity

type TokenXAccessRequestBodyInfo struct {
	// @SerializedName("pin")
	// @Nullable
	// private String pin;
	Pin *string `json:"pin"`
}

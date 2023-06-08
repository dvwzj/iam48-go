package entity

type VerifyCodePinModel struct {
	// @SerializedName("pinResetCode")
	// @Nullable
	// private String pinResetCode;
	PinResetCode *string `json:"pinResetCode"`
}

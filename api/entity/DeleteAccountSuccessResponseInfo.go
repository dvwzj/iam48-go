package entity

type DeleteAccountSuccessResponseInfo struct {
	// @SerializedName("recoveryInfoMessage")
	// @Nullable
	// private final String recoveryInfoMessage;
	RecoveryInfoMessage *string `json:"recoveryInfoMessage"`
}

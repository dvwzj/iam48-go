package entity

type DeleteAccountSuccessRequestInfo struct {
	// @SerializedName("deleteReasonId")
	// @Nullable
	// private Long deleteReasonId;
	DeleteReasonId *int64 `json:"deleteReasonId"`
	// @SerializedName("emailOtpPin")
	// @Nullable
	// private String emailOtpPin;
	EmailOtpPin *string `json:"emailOtpPin"`
	// @SerializedName("emailOtpReference")
	// @Nullable
	// private String emailOtpReference;
	EmailOtpReference *string `json:"emailOtpReference"`
}

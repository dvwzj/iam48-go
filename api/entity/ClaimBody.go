package entity

type ClaimBody struct {
	// @SerializedName("billingSdkVersion")
	// @Nullable
	// private Long billingSdkVersion;
	BillingSdkVersion *int64 `json:"billingSdkVersion"`
	// @SerializedName("comment")
	// @Nullable
	// private Comment comment;
	Comment *Comment `json:"comment"`
	// @SerializedName("inAppDataSignature")
	// @Nullable
	// private String inAppDataSignature;
	InAppDataSignature *string `json:"inAppDataSignature"`
	// @SerializedName("inAppPurchaseData")
	// @Nullable
	// private String inAppPurchaseData;
	InAppPurchaseData *string `json:"inAppPurchaseData"`
}

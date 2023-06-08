package entity

type SubscriptionInfo struct {
	// @SerializedName("endAt")
	// @Nullable
	// private String endAt;
	EndAt *string `json:"endAt"`
	// @SerializedName("expireInDay")
	// @Nullable
	// private Integer expireInDay;
	ExpireInDay *int `json:"expireInDay"`
	// @SerializedName("originalInvoiceNo")
	// @Nullable
	// private String originalInvoiceNo;
	OriginalInvoiceNo *string `json:"originalInvoiceNo"`
	// @SerializedName("subscriptionName")
	// @Nullable
	// private String subscriptionName;
	SubscriptionName *string `json:"subscriptionName"`
	// @SerializedName("subscriptionPackageCode")
	// @Nullable
	// private String subscriptionPackageCode;
	SubscriptionPackageCode *string `json:"subscriptionPackageCode"`
}

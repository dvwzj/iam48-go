package entity

type VoteSenbatsuResponseInfo struct {
	// @SerializedName("thankYouImageUrl")
	// @Nullable
	// private String thankYouImageUrl;
	ThankYouImageUrl *string `json:"thankYouImageUrl"`
	// @SerializedName("transactionId")
	// @Nullable
	// private String transactionId;
	TransactionId *string `json:"transactionId"`
	// @SerializedName("userMessage")
	// @Nullable
	// private String userMessage;
	UserMessage *string `json:"userMessage"`
}

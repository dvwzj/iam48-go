package entity

type ClaimResponse struct {
	// @SerializedName("receivedAmount")
	// private final int receivedAmount;
	ReceivedAmount int `json:"receivedAmount"`
	// @SerializedName("userMessage")
	// @Nullable
	// private final String userMessage;
	UserMessage *string `json:"userMessage"`
}

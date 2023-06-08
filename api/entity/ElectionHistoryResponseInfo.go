package entity

type ElectionHistoryResponseInfo struct {
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("sentAt")
	// @Nullable
	// private String sentAt;
	SentAt *string `json:"sentAt"`
	// @SerializedName("sentToEmailAddress")
	// @Nullable
	// private String sentToEmailAddress;
	SentToEmailAddress *string `json:"sentToEmailAddress"`
	// @SerializedName("setId")
	// @Nullable
	// private Long setId;
	SetId *int64 `json:"setId"`
	// @SerializedName("shopeeOrderId")
	// @Nullable
	// private String shopeeOrderId;
	ShopeeOrderId *string `json:"shopeeOrderId"`
	// @SerializedName("totalCodeAmount")
	// @Nullable
	// private Integer totalCodeAmount;
	TotalCodeAmount *int64 `json:"totalCodeAmount"`
}

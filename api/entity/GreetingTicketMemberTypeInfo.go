package entity

type GreetingTicketMemberTypeInfo struct {
	// @SerializedName("availableRedeemType")
	// @Nullable
	// private List<GreetingAvailableRedeemTypeInfo> availableRedeemType;
	AvailableRedeemType *[]GreetingAvailableRedeemTypeInfo `json:"availableRedeemType"`
	// @SerializedName("endDate")
	// @Nullable
	// private String endDate;
	EndDate *string `json:"endDate"`
	// @SerializedName("expireDate")
	// @Nullable
	// private String expireDate;
	ExpireDate *string `json:"expireDate"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("startDate")
	// @Nullable
	// private String startDate;
	StartDate *string `json:"startDate"`
}

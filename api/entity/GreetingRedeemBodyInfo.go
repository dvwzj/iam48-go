package entity

type GreetingRedeemBodyInfo struct {
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("ownerPIN")
	// @Nullable
	// private String ownerPIN;
	OwnerPIN *string `json:"ownerPIN"`
	// @SerializedName("redeemTypeId")
	// @Nullable
	// private Long redeemTypeId;
	RedeemTypeId *int64 `json:"redeemTypeId"`
	// @SerializedName("ticketId")
	// @Nullable
	// private Long ticketId;
	TicketId *int64 `json:"ticketId"`
}

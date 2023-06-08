package entity

type RedeemMeetYouRequest struct {
	// @SerializedName("memberId")
	// private long memberId;
	MemberId int64 `json:"memberId"`
	// @SerializedName("ownerPIN")
	// @NotNull
	// private String ownerPin;
	OwnerPin string `json:"ownerPin"`
	// @SerializedName(ConstancesKt.KEY_TICKET_AMOUNT) // "ticketAmount"
	// private long ticketAmount;
	TicketAmount int64 `json:"ticketAmount"`
	// @SerializedName("ticketId")
	// private long ticketId;
	TicketId int64 `json:"ticketId"`
	// @SerializedName(SelectRoundActivity.KEY_TIME_SLOT_ID) // "timeSlotId"
	// private long timeSlotId;
	TimeSlotId int64 `json:"timeSlotId"`
}

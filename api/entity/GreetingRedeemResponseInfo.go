package entity

type GreetingRedeemResponseInfo struct {
	// @SerializedName("expireAt")
	// @Nullable
	// private String expireAt;
	ExpireAt *string `json:"expireAt"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("isAudioOnly")
	// @Nullable
	// private Boolean isAudioOnly;
	IsAudioOnly *bool `json:"isAudioOnly"`
	// @SerializedName("isRequireDialog")
	// @Nullable
	// private Boolean isRequireDialog;
	IsRequireDialog *bool `json:"isRequireDialog"`
	// @SerializedName("isRequireVideo")
	// @Nullable
	// private Boolean isRequireVideo;
	IsRequireVideo *bool `json:"isRequireVideo"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @Nullable
	// private Long refCode;
	RefCode *int64 `json:"refCode"`
	// @SerializedName(ConstancesKt.KEY_TICKET) // "ticket"
	// @Nullable
	// private GreetingTicketInfo ticket;
	Ticket *GreetingTicketInfo `json:"ticket"`
	// @SerializedName(ConstancesKt.KEY_TICKET_AMOUNT) // "ticketAmount"
	// @Nullable
	// private Long ticketAmount;
	TicketAmount *int64 `json:"ticketAmount"`
}

package entity

type RedeemInfo struct {
	// @SerializedName("eventInfo")
	// @Nullable
	// private EventInfo eventInfo;
	EventInfo *EventInfo `json:"eventInfo"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("isAllowToJoin")
	// @Nullable
	// private Boolean isAllowToJoin;
	IsAllowToJoin *bool `json:"isAllowToJoin"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("memberImageUrl")
	// @Nullable
	// private String memberImageUrl;
	MemberImageUrl *string `json:"memberImageUrl"`
	// @Nullable
	// private MemberProfile memberProfile;
	MemberProfile *MemberProfile `json:"memberProfile"`
	// @SerializedName("queueInfo")
	// @Nullable
	// private QueueInfo queueInfo;
	QueueInfo *QueueInfo `json:"queueInfo"`
	// @SerializedName(ConstancesKt.KEY_QUEUE_NO) // "queueNo"
	// @Nullable
	// private Long queueNo;
	QueueNo *int64 `json:"queueNo"`
	// @SerializedName("redeemAt")
	// @Nullable
	// private String redeemAt;
	RedeemAt *string `json:"redeemAt"`
	// @SerializedName(ConstancesKt.KEY_REF_CODE) // "refCode"
	// @Nullable
	// private String refCode;
	RefCode *string `json:"refCode"`
	// @SerializedName("roundInfo")
	// @Nullable
	// private RoundInfo roundInfo;
	RoundInfo *RoundInfo `json:"roundInfo"`
	// @SerializedName("scheduleInfo")
	// @Nullable
	// private ScheduleInfo scheduleInfo;
	ScheduleInfo *ScheduleInfo `json:"scheduleInfo"`
	// @SerializedName("statusLevel")
	// @Nullable
	// private Long statusLevel;
	StatusLevel *int64 `json:"statusLevel"`
	// @SerializedName("statusMessage")
	// @Nullable
	// private String statusMessage;
	StatusMessage *string `json:"statusMessage"`
	// @SerializedName("ticketInfo")
	// @Nullable
	// private TicketInfo ticketInfo;
	TicketInfo *TicketInfo `json:"ticketInfo"`
	// @SerializedName(ConstancesKt.KEY_TICKET_AMOUNT) // "ticketAmount"
	// @Nullable
	// private Long ticketQuantity;
	TicketQuantity *int64 `json:"ticketQuantity"`
	// @SerializedName("totalSeconds")
	// @Nullable
	// private String totalSeconds;
	TotalSeconds *string `json:"totalSeconds"`
}

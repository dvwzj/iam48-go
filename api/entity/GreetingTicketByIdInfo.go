package entity

type GreetingTicketByIdInfo struct {
	// @SerializedName("availableMemberList")
	// @NotNull
	// private List<GreetingAvailableMemberInfo> availableMemberList;
	AvailableMemberList []GreetingAvailableMemberInfo `json:"availableMemberList"`
	// @SerializedName(ConstancesKt.KEY_TICKET) // "ticket"
	// @Nullable
	// private GreetingTicketInfo ticket;
	Ticket *GreetingTicketInfo `json:"ticket"`
}

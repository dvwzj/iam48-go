package entity

type TimeSlot struct {
	// @SerializedName("canRedeem")
	// @Nullable
	// private Boolean canRedeem;
	CanRedeem *bool `json:"canRedeem"`
	// @SerializedName("maximumRedeemTicket")
	// @Nullable
	// private Long maximumRedeemTicket;
	MaximumRedeemTicket *int64 `json:"maximumRedeemTicket"`
	// @SerializedName("minimumRedeemTicket")
	// @Nullable
	// private Long minimumRedeemTicket;
	MinimumRedeemTicket *int64 `json:"minimumRedeemTicket"`
	// @SerializedName("secondsPerTicket")
	// @Nullable
	// private Long secondsPerTicket;
	SecondsPerTicket *int64 `json:"secondsPerTicket"`
	// @SerializedName("time")
	// @Nullable
	// private String time;
	Time *string `json:"time"`
	// @SerializedName(SelectRoundActivity.KEY_TIME_SLOT_ID) // "timeSlotId"
	// @Nullable
	// private Long timeSlotId;
	TimeSlotId *int64 `json:"timeSlotId"`
}

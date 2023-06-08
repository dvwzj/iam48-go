package entity

type ScheduleInfo struct {
	// @SerializedName("scheduleTimeFrom")
	// @Nullable
	// private String scheduleTimeFrom;
	ScheduleTimeFrom *string `json:"scheduleTimeFrom"`
	// @SerializedName("scheduleTimeTo")
	// @Nullable
	// private String scheduleTimeTo;
	ScheduleTimeTo *string `json:"scheduleTimeTo"`
}

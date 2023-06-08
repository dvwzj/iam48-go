package entity

type Round struct {
	// @Nullable
	// private String date;
	Date *string `json:"date"`
	// @Nullable
	// private List<TimeSlot> timeSlot;
	TimeSlot *[]TimeSlot `json:"timeSlot"`
}

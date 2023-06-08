package entity

type EkycVerifyUserBodyInfo struct {
	// @SerializedName("birthDate")
	// @Nullable
	// private String birthDate;
	BirthDate *string `json:"birthDate"`
	// @SerializedName("firstName")
	// @Nullable
	// private String firstName;
	FirstName *string `json:"firstName"`
	// @SerializedName("idcardNo")
	// @Nullable
	// private String idcardNo;
	IdcardNo *string `json:"idcardNo"`
	// @SerializedName("lastName")
	// @Nullable
	// private String lastName;
	LastName *string `json:"lastName"`
	// @SerializedName("sessionId")
	// @Nullable
	// private String sessionId;
	SessionId *string `json:"sessionId"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
}

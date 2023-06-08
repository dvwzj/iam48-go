package entity

type NotificationRegisterBody struct {
	// @SerializedName("fcmToken")
	// @NotNull
	// private String fcmToken;
	FcmToken string `json:"fcmToken"`
}

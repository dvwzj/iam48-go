package entity

type ElectionSentCodeToEmailBody struct {
	// @SerializedName("setId")
	// @Nullable
	// private Long setId;
	SetId *int64 `json:"setId"`
	// @SerializedName("toEmailAddress")
	// @Nullable
	// private String toEmailAddress;
	ToEmailAddress *string `json:"toEmailAddress"`
}

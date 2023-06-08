package entity

type SendGiftResponse struct {
	// @SerializedName("coin")
	// @Nullable
	// private Long coin;
	Coin *int64 `json:"coin"`
	// @SerializedName(ConstancesKt.KEY_MESSAGE) // "message"
	// @Nullable
	// private String message;
	Message *string `json:"message"`
	// @SerializedName("success")
	// @Nullable
	// private Boolean success;
	Success *bool `json:"success"`
}

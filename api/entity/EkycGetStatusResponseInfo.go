package entity

type EkycGetStatusResponseInfo struct {
	// @SerializedName("isSuccess")
	// @Nullable
	// private Boolean isSuccess;
	IsSuccess *bool `json:"isSuccess"`
	// @SerializedName(ConstancesKt.KEY_MESSAGE)
	// @Nullable
	// private String message;
	Message *string `json:"message"`
}

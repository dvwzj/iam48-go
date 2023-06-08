package entity

type NewCommentBody struct {
	// @SerializedName(ConstancesKt.KEY_MESSAGE)
	// @NotNull
	// private final String message;
	Message string `json:"message"`
}

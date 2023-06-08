package entity

type CommentInfo struct {
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName(ConstancesKt.KEY_MESSAGE) // "message"
	// @Nullable
	// private String message;
	Message *string `json:"message"`
	// @SerializedName("user")
	// @NotNull
	// private PublishedBy user;
	User PublishedBy `json:"user"`
}

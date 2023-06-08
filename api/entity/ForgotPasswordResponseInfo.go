package entity

type ForgotPasswordResponseInfo struct {
	// @SerializedName(Scopes.EMAIL) // "email"
	// @Nullable
	// private String email;
	Email *string `json:"email"`
}

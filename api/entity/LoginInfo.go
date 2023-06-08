package entity

type LoginInfo struct {
	// @SerializedName(Scopes.EMAIL)
	// @Nullable
	// private String email;
	Email *string `json:"email"`
	// @SerializedName("password")
	// @Nullable
	// private String password;
	Password *string `json:"password"`
}

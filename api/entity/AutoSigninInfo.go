package entity

type AutoSigninInfo struct {
	// @SerializedName("signinUrl")
	// @Nullable
	// private String signinUrl;
	SigninUrl *string `json:"signinUrl"`
}

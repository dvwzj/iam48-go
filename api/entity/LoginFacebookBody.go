package entity

type LoginFacebookBody struct {
	// @SerializedName("accessToken")
	// @Nullable
	// private String accessToken;
	AccessToken *string `json:"accessToken"`
	// @SerializedName(Scopes.EMAIL) // "email"
	// @Nullable
	// private String email;
	Email *string `json:"email"`
	// @SerializedName("facebookId")
	// @Nullable
	// private String facebookId;
	FacebookId *string `json:"facebookId"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
}

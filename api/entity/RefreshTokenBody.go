package entity

type RefreshTokenBody struct {
	// @SerializedName("refreshToken")
	// @NotNull
	// private final String refreshToken;
	RefreshToken string `json:"refreshToken"`
}

package entity

type RefreshTokenInfo struct {
	// @SerializedName("expireDate")
	// @Nullable
	// private String expireDate;
	ExpireDate *string `json:"expireDate"`
	// @SerializedName("id")
	// @Nullable
	// private Integer id;
	Id *int `json:"id"`
	// @SerializedName("refreshToken")
	// @Nullable
	// private String refreshToken;
	RefreshToken *string `json:"refreshToken"`
	// @SerializedName("token")
	// @Nullable
	// private String token;
	Token *string `json:"token"`
}

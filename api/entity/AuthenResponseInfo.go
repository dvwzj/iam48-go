package entity

type AuthenResponseInfo struct {
	// @SerializedName("expireDate")
	// @Nullable
	// private String expireDate;
	ExpireDate *string `json:"expireDate"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("refreshToken")
	// @Nullable
	// private String refreshToken;
	RefreshToken *string `json:"refreshToken"`
	// @SerializedName("token")
	// @Nullable
	// private String token;
	Token *string `json:"token"`
}

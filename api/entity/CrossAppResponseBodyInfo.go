package entity

type CrossAppResponseBodyInfo struct {
	// @NotNull
	// private String expireDate;
	ExpireDate string `json:"expireDate"`
	// private long id;
	Id int64 `json:"id"`
	// @NotNull
	// private String refreshToken;
	RefreshToken string `json:"refreshToken"`
	// @NotNull
	// private String token;
	Token string `json:"token"`
}

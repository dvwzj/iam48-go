package entity

type TokenModel struct {
	// @SerializedName("symbolUrl")
	// @Nullable
	// private String symbolUrl;
	SymbolUrl *string `json:"symbolUrl"`
	// @SerializedName("tokenCode")
	// @Nullable
	// private String tokenCode;
	TokenCode *string `json:"tokenCode"`
	// @SerializedName("tokenName")
	// @Nullable
	// private String tokenName;
	TokenName *string `json:"tokenName"`
}

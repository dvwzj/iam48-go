package entity

type FacebookLinkScopeResponseInfo struct {
	// @SerializedName("scope")
	// @Nullable
	// private String scope;
	Scope *string `json:"scope"`
}

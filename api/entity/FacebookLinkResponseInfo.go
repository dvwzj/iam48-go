package entity

type FacebookLinkResponseInfo struct {
	// @SerializedName("displayName")
	// @Nullable
	// private String displayName;
	DisplayName *string `json:"displayName"`
	// @SerializedName("emailAddress")
	// @Nullable
	// private String emailAddress;
	EmailAddress *string `json:"emailAddress"`
	// @SerializedName("userId")
	// @Nullable
	// private Long id;
	Id *int64 `json:"userId"`
}

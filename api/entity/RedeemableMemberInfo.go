package entity

type RedeemableMemberInfo struct {
	// @SerializedName("canRedeem")
	// @Nullable
	// private Boolean canRedeem;
	CanRedeem *bool `json:"canRedeem"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("leftMessage")
	// @Nullable
	// private String leftMessage;
	LeftMessage *string `json:"leftMessage"`
	// @SerializedName("memberImageUrl")
	// @Nullable
	// private String memberImageUrl;
	MemberImageUrl *string `json:"memberImageUrl"`
}

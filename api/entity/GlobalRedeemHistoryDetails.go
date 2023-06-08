package entity

type GlobalRedeemHistoryDetails struct {
	// @SerializedName(GlobalRedeemCodeActivity.KEY_CODE) // "code"
	// @Nullable
	// private String code;
	Code *string `json:"code"`
	// @SerializedName("createdAt")
	// @Nullable
	// private String createdAt;
	CreatedAt *string `json:"createdAt"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("skus")
	// @Nullable
	// private List<GlobalRedeemProductInfo> redeemProductList;
	RedeemProductList []*GlobalRedeemProductInfo `json:"skus"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE) // "title"
	// @Nullable
	// private String title;
	Title *string `json:"title"`
}

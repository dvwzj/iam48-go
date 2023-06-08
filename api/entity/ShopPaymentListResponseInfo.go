package entity

type ShopPaymentListResponseInfo struct {
	// @SerializedName("displayIndex")
	// @Nullable
	// private Integer displayIndex;
	DisplayIndex *int `json:"displayIndex"`
	// @SerializedName("gatewayCode")
	// @Nullable
	// private String gatewayCode;
	GatewayCode *string `json:"gatewayCode"`
	// @SerializedName("gatewayUrl")
	// @Nullable
	// private String gatewayUrl;
	GatewayUrl *string `json:"gatewayUrl"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("imageFileUrl")
	// @Nullable
	// private String imageFileUrl;
	ImageFileUrl *string `json:"imageFileUrl"`
	// @SerializedName(MeetingYouAlertDialog.KEY_TITLE)
	// @Nullable
	// private String title;
	Title *string `json:"title"`
}

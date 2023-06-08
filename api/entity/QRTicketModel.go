package entity

type QRTicketModel struct {
	// @SerializedName("endDate")
	// @Nullable
	// private String endDate;
	EndDate *string `json:"endDate"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY)
	// @Nullable
	// private Long quantity;
	Quantity *int64 `json:"quantity"`
	// @SerializedName("startDate")
	// @Nullable
	// private String startDate;
	StartDate *string `json:"startDate"`
	// @SerializedName("svgQrCode")
	// @Nullable
	// private String svgQrCode;
	SvgQrCode *string `json:"svgQrCode"`
}

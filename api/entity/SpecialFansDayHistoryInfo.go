package entity

type SpecialFansDayHistoryInfo struct {
	// @SerializedName("contentId")
	// @Nullable
	// private Long contentId;
	ContentId *int64 `json:"contentId"`
	// @SerializedName("fromDate")
	// @Nullable
	// private String fromDate;
	FromDate *string `json:"fromDate"`
	// @Nullable
	// private String graduatedAt;
	GraduatedAt *string `json:"graduatedAt"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @Nullable
	// private String imageUrl;
	ImageUrl *string `json:"imageUrl"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @Nullable
	// private String memberName;
	MemberName *string `json:"memberName"`
	// @SerializedName("toDate")
	// @Nullable
	// private String toDate;
	ToDate *string `json:"toDate"`
	// @SerializedName("totalGiftAmount")
	// @Nullable
	// private Long totalGiftAmount;
	TotalGiftAmount *int64 `json:"totalGiftAmount"`
}

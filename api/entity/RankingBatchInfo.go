package entity

type RankingBatchInfo struct {
	// @SerializedName("contentId")
	// @Nullable
	// private Long contentId;
	ContentId *int64 `json:"contentId"`
	// @SerializedName("fromDate")
	// @Nullable
	// private String fromDate;
	FromDate *string `json:"fromDate"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("resource")
	// @Nullable
	// private RankingBatchResourceItem resource;
	Resource *RankingBatchResourceItem `json:"resource"`
	// @SerializedName("tabbarIconUrl")
	// @Nullable
	// private TabBarIconUrlListInfo tabbarIconUrl;
	TabbarIconUrl *TabBarIconUrlListInfo `json:"tabbarIconUrl"`
	// @SerializedName("toDate")
	// @Nullable
	// private final String toDate;
	ToDate *string `json:"toDate"`
	// @SerializedName("totalGiftAmount")
	// @Nullable
	// private Long totalGiftAmount;
	TotalGiftAmount *int64 `json:"totalGiftAmount"`
}

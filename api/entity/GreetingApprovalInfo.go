package entity

type GreetingApprovalInfo struct {
	// @SerializedName("approveAt")
	// @Nullable
	// private String approveAt;
	ApproveAt *string `json:"approveAt"`
	// @SerializedName("isApprove")
	// @Nullable
	// private Boolean isApprove;
	IsApprove *bool `json:"isApprove"`
	// @SerializedName("rejectReason")
	// @Nullable
	// private String rejectReason;
	RejectReason *string `json:"rejectReason"`
}

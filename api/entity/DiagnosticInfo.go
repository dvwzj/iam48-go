package entity

type DiagnosticInfo struct {
	// @SerializedName("isSandbox")
	// @Nullable
	// private Boolean isSandbox;
	IsSandbox *bool `json:"isSandbox"`
	// @SerializedName("userId")
	// @Nullable
	// private Long userId;
	UserId *int64 `json:"userId"`
	// @SerializedName("walletGroupId")
	// @Nullable
	// private Long walletGroupId
	WalletGroupId *int64 `json:"walletGroupId"`
}

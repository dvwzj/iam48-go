package entity

type BalanceInfo struct {
	// @SerializedName("diagnosticInfo")
	// @Nullable
	// private DiagnosticInfo diagnosticInfo;
	DiagnosticInfo *DiagnosticInfo `json:"diagnosticInfo"`
	// @SerializedName("reservedBalance")
	// @Nullable
	// private Long reservedBalance;
	ReservedBalance *int64 `json:"reservedBalance"`
	// @SerializedName("spendableBalance")
	// @Nullable
	// private Long spendableBalance;
	SpendableBalance *int64 `json:"spendableBalance"`
}

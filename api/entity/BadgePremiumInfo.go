package entity

type BadgePremiumInfo struct {
	// @SerializedName("currentActiveBadge")
	// @Nullable
	// private CurrentActiveBadge currentActiveBadge;
	CurrentActiveBadge *CurrentActiveBadge `json:"currentActiveBadge"`
}

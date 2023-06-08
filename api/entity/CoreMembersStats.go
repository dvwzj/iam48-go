package entity

type CoreMembersStats struct {
	// @SerializedName("coin")
	// @Nullable
	// private MembersStats coin;
	Coin *MembersStats `json:"coin"`
	// @SerializedName("gift")
	// @Nullable
	// private MembersStats gift;
	Gift *MembersStats `json:"gift"`
	// @SerializedName("like")
	// @Nullable
	// private MembersStats like;
	Like *MembersStats `json:"like"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_PERIOD)
	// @Nullable
	// private RankPeriodInfo period;
	Period *RankPeriodInfo `json:"period"`
}

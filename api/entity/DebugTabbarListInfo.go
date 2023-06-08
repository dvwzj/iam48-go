package entity

type DebugTabbarListInfo struct {
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("tabbar")
	// @Nullable
	// private DebugTabbarInfo tabbar;
	Tabbar *DebugTabbarInfo `json:"tabbar"`
}

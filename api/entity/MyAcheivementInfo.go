package entity

type MyAcheivementInfo struct {
	// @SerializedName("achievementId")
	// @Nullable
	// private Long id;
	Id *int64 `json:"achievementId"`
}

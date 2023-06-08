package entity

type LatestPlayInfo struct {
	// @SerializedName("latestPlaySeconds")
	// @Nullable
	// private Integer latestPlaySeconds;
	LatestPlaySeconds *int `json:"latestPlaySeconds"`
}

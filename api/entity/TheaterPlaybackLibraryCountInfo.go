package entity

type TheaterPlaybackLibraryCountInfo struct {
	// @SerializedName("totalWatchablePlayback")
	// @Nullable
	// private Integer totalWatchablePlayback;
	TotalWatchablePlayback *int `json:"totalWatchablePlayback"`
}

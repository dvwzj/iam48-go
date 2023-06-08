package entity

type TheaterStatLiveInfo struct {
	// @SerializedName("canWatchTheaterLive")
	// @Nullable
	// private Boolean canWatchTheaterLive;
	CanWatchTheaterLive *bool `json:"canWatchTheaterLive"`
}

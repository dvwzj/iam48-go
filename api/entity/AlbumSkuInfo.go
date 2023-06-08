package entity

type AlbumSkuInfo struct {
	// @SerializedName("albumId")
	// @Nullable
	// private Long albumId;
	AlbumId *int64 `json:"albumId"`
	// @SerializedName("artistName")
	// @Nullable
	// private String artistName;
	ArtistName *string `json:"artistName"`
	// @SerializedName("durationSecond")
	// @Nullable
	// private Long durationSecond;
	DurationSecond *int64 `json:"durationSecond"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("songTitle")
	// @Nullable
	// private String songTitle;
	SongTitle *string `json:"songTitle"`
	// @SerializedName("trackNo")
	// @Nullable
	// private Long trackNo;
	TrackNo *int64 `json:"trackNo"`
}

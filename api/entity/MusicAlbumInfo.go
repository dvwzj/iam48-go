package entity

type MusicAlbumInfo struct {
	// @SerializedName("albumTitle")
	// @Nullable
	// private String albumTitle;
	AlbumTitle *string `json:"albumTitle"`
	// @SerializedName("artFileUrl")
	// @Nullable
	// private String artFileUrl;
	ArtFileUrl *string `json:"artFileUrl"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @Nullable
	// private Boolean isUnlocked;
	IsUnlocked *bool `json:"isUnlocked"`
	// @SerializedName("orderIndex")
	// @Nullable
	// private Long orderIndex;
	OrderIndex *int64 `json:"orderIndex"`
	// @SerializedName("skus")
	// @Nullable
	// private List<AlbumSkuInfo> skus;
	Skus *[]AlbumSkuInfo `json:"skus"`
}

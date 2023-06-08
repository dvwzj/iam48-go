package entity

type CheckInResponseInfo struct {
	// @SerializedName("badgeImageUrlList")
	// @Nullable
	// private List<String> badgeImageUrlList;
	BadgeImageUrlList *[]string `json:"badgeImageUrlList"`
	// @SerializedName("placeName")
	// @Nullable
	// private String placeName;
	PlaceName *string `json:"placeName"`
}

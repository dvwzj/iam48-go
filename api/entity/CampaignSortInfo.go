package entity

type CampaignSortInfo struct {
	// @SerializedName("displayIndex")
	// @Nullable
	// private Integer displayIndex;
	DisplayIndex *int `json:"displayIndex"`
	// @SerializedName("displayName")
	// @Nullable
	// private String displayName;
	DisplayName *string `json:"displayName"`
	// @SerializedName("requestName")
	// @Nullable
	// private String requestName;
	RequestName *string `json:"requestName"`
}

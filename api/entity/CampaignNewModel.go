package entity

type CampaignNewModel struct {
	// @SerializedName("items")
	// @Nullable
	// private List<CampaignNewModelItem> items;
	Items *[]CampaignNewModelItem `json:"items"`
	// @SerializedName("sortOrders")
	// @Nullable
	// private List<SortOrdersItem> sortOrders;
	SortOrders *[]SortOrdersItem `json:"sortOrders"`
}

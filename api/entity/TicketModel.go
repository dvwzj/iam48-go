package entity

type TicketModel struct {
	// @SerializedName("eventDate")
	// @Nullable
	// private String eventDate;
	EventDate *string `json:"eventDate"`
	// @SerializedName("eventPeriod")
	// @Nullable
	// private String eventPeriod;
	EventPeriod *string `json:"eventPeriod"`
	// @SerializedName("imageUrl")
	// @Nullable
	// private String imageUrl;
	ImageUrl *string `json:"imageUrl"`
	// @SerializedName("isAllowTransfer")
	// @Nullable
	// private Boolean isAllowTransfer;
	IsAllowTransfer *bool `json:"isAllowTransfer"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
	// @SerializedName("placeName")
	// @Nullable
	// private String placeName;
	PlaceName *string `json:"placeName"`
	// @SerializedName(ConstancesKt.KEY_QUANTITY) // "quantity"
	// @Nullable
	// private Integer quantity;
	Quantity *int `json:"quantity"`
	// @SerializedName("shortDescription")
	// @Nullable
	// private String shortDescription;
	ShortDescription *string `json:"shortDescription"`
	// @SerializedName("ticketPrefix")
	// @Nullable
	// private String ticketPrefix;
	TicketPrefix *string `json:"ticketPrefix"`
	// @SerializedName("ticketSkuId")
	// @Nullable
	// private Long ticketSkuId;
	TicketSkuId *int64 `json:"ticketSkuId"`
}

package entity

type ExchangeTicketResponseInfo struct {
	// @SerializedName("canRedeem")
	// @Nullable
	// private Boolean canRedeem;
	CanRedeem *bool `json:"canRedeem"`
	// @SerializedName("displayMessage")
	// @Nullable
	// private String displayMessage;
	DisplayMessage *string `json:"displayMessage"`
	// @SerializedName("displayMessageLevel")
	// @Nullable
	// private Long displayMessageLevel;
	DisplayMessageLevel *int64 `json:"displayMessageLevel"`
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
	// private Long quantity;
	Quantity *int64 `json:"quantity"`
	// @SerializedName("shortDescription")
	// @Nullable
	// private String shortDescription;
	ShortDescription *string `json:"shortDescription"`
	// @SerializedName("ticketSkuId")
	// @Nullable
	// private Long ticketSkuId;
	TicketSkuId *int64 `json:"ticketSkuId"`
}

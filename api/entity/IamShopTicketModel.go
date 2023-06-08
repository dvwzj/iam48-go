package entity

type IamShopTicketModel struct {
	// @SerializedName("eventDate")
	// @Nullable
	// private String eventDate;
	EventDate *string `json:"eventDate"`
	// @SerializedName("eventDateAsString")
	// @Nullable
	// private String eventDateAsString;
	EventDateAsString *string `json:"eventDateAsString"`
	// @SerializedName("eventPeriod")
	// @Nullable
	// private String eventPeriod;
	EventPeriod *string `json:"eventPeriod"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
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
	// @SerializedName("shortDescription")
	// @Nullable
	// private String shortDescription;
	ShortDescription *string `json:"shortDescription"`
}

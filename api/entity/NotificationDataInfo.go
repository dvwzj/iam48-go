package entity

type NotificationDataInfo struct {
	// @SerializedName("additionalProp1")
	// @Nullable
	// private String additionalProp1;
	AdditionalProp1 *string `json:"additionalProp1"`
	// @SerializedName("additionalProp2")
	// @Nullable
	// private String additionalProp2;
	AdditionalProp2 *string `json:"additionalProp2"`
	// @SerializedName("additionalProp3")
	// @Nullable
	// private String additionalProp3;
	AdditionalProp3 *string `json:"additionalProp3"`
}

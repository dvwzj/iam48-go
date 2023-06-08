package entity

type DeleteAccountReasonResponseInfo struct {
	// @SerializedName("displayIndex")
	// @Nullable
	// private Integer displayIndex;
	DisplayIndex *int `json:"displayIndex"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("topic")
	// @Nullable
	// private String topic;
	Topic *string `json:"topic"`
}

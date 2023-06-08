package entity

type PrefixSuffixTemplatesInfo struct {
	// @SerializedName("displayText")
	// @Nullable
	// private String displayText;
	DisplayText *string `json:"displayText"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// private boolean isSelected;
	IsSelected *bool `json:"isSelected"`
	// @SerializedName("templateText")
	// @Nullable
	// private String templateText;
	TemplateText *string `json:"templateText"`
}

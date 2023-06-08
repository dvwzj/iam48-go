package entity

type DiscoverSectionInfo struct {
	// @SerializedName("sections")
	// @NotNull
	// private List<? extends BaseSectionInfo> sections;
	Sections []BaseSectionInfo `json:"sections"`
}

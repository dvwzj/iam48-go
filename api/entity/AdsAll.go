package entity

type AdsAll struct {
	// @SerializedName("sections")
	// @Nullable
	// private List<AdsSections> sections;
	Sections *[]AdsSections `json:"sections"`
	// @SerializedName("version")
	// @Nullable
	// private String version;
	Version *string `json:"version"`
}

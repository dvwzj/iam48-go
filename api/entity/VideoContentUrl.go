package entity

type VideoContentUrl struct {
	// @SerializedName("resourceUrl")
	// @NotNull
	// private final String resourceUrl;
	ResourceUrl string `json:"resourceUrl"`
}

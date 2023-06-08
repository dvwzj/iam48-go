package entity

type ThemeInfo struct {
	// @SerializedName("themeName")
	// @Nullable
	// private String themeName;
	ThemeName *string `json:"themeName"`
}

package entity

type PdpaCurrentVersionResponseInfo struct {
	// @Nullable
	// private String apiVersion;
	ApiVersion *string `json:"apiVersion"`
	// @Nullable
	// private Long data;
	Data *int64 `json:"data"`
}

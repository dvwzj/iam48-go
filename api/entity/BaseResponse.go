package entity

type BaseResponse[T any] struct {
	// @SerializedName("apiVersion")
	// private String apiVersion;
	ApiVersion string `json:"apiVersion"`
	// @SerializedName(TPReportParams.PROP_KEY_DATA) // "data"
	// private T data;
	Data T `json:"data"`
}

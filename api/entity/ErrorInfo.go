package entity

type ErrorInfo struct {
	// private int code;
	Code int `json:"code"`
	// @Nullable
	// private Throwable error;
	Error *error `json:"error"`
	// @NotNull
	// private String message;
	Message string `json:"message"`
}

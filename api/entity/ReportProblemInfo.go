package entity

type ReportProblemInfo struct {
	// @SerializedName("appCode")
	// @Nullable
	// private String appCode;
	AppCode *string `json:"appCode"`
	// @SerializedName("appVersion")
	// @Nullable
	// private String appVersion;
	AppVersion *string `json:"appVersion"`
	// @SerializedName("contactEmail")
	// @Nullable
	// private String contactEmail;
	ContactEmail *string `json:"contactEmail"`
	// @SerializedName("contactName")
	// @Nullable
	// private String contactName;
	ContactName *string `json:"contactName"`
	// @SerializedName("deviceModel")
	// @Nullable
	// private String deviceModel;
	DeviceModel *string `json:"deviceModel"`
	// @SerializedName("fileId")
	// @Nullable
	// private List<String> fileId;
	FileId *[]string `json:"fileId"`
	// @SerializedName("os")
	// @Nullable
	// private String os;
	Os *string `json:"os"`
	// @SerializedName("osVersion")
	// @Nullable
	// private String osVersion;
	OsVersion *string `json:"osVersion"`
	// @SerializedName("phoneNumber")
	// @Nullable
	// private String phoneNumber;
	PhoneNumber *string `json:"phoneNumber"`
	// @SerializedName("problemDetail")
	// @Nullable
	// private String problemDetail;
	ProblemDetail *string `json:"problemDetail"`
	// @SerializedName("subject")
	// @Nullable
	// private String subject;
	Subject *string `json:"subject"`
	// @SerializedName("technicalData")
	// @Nullable
	// private String technicalData;
	TechnicalData *string `json:"technicalData"`
	// @SerializedName("userId")
	// @Nullable
	// private Long userId;
	UserId *int64 `json:"userId"`
}

package entity

type GreetingRedeemResponseByIdInfo struct {
	// @SerializedName("isAudioOnly")
	// @Nullable
	// private Boolean isAudioOnly;
	IsAudioOnly *bool `json:"isAudioOnly"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("videoFileUrl")
	// @Nullable
	// private String videoFileUrl;
	VideoFileUrl *string `json:"videoFileUrl"`
}

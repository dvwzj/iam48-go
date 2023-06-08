package entity

type ElectionCodeSetByIdInfo struct {
	// @SerializedName(GlobalRedeemCodeActivity.KEY_CODE) // "code"
	// @Nullable
	// private String code;
	Code *string `json:"code"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("sequence")
	// @Nullable
	// private Integer seq;
	Seq *int64 `json:"sequence"`
}

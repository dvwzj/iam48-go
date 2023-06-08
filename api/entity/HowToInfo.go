package entity

type HowToInfo struct {
	// @SerializedName("introContent")
	// @Nullable
	// private String introContent;
	IntroContent *string `json:"introContent"`
	// @SerializedName("introTitle")
	// @Nullable
	// private String introTitle;
	IntroTitle *string `json:"introTitle"`
	// @SerializedName("ruleContent")
	// @Nullable
	// private String ruleContent;
	RuleContent *string `json:"ruleContent"`
	// @SerializedName("ruleTitle")
	// @Nullable
	// private String ruleTitle;
	RuleTitle *string `json:"ruleTitle"`
}

package entity

type GreetingDialogInfo struct {
	// @SerializedName("dialogTemplates")
	// @Nullable
	// private List<ScriptTemplatesInfo> dialogTemplates;
	DialogTemplates *[]ScriptTemplatesInfo `json:"dialogTemplates"`
	// @SerializedName("namingTemplates")
	// @Nullable
	// private List<PrefixSuffixTemplatesInfo> namingTemplates;
	NamingTemplates *[]PrefixSuffixTemplatesInfo `json:"namingTemplates"`
}

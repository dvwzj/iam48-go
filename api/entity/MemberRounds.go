package entity

type MemberRounds struct {
	// @SerializedName("normalRound")
	// @Nullable
	// private List<Round> normalRound;
	NormalRound *[]Round `json:"normalRound"`
	// @SerializedName("specialRound")
	// @Nullable
	// private List<Round> specialRound;
	SpecialRound *[]Round `json:"specialRound"`
}

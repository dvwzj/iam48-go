package entity

type SocialCommentInfo struct {
	// @NotNull
	// private String localId;
	LocalId string `json:"localId"`
	// @NotNull
	// private String text;
	Text string `json:"text"`
	// @NotNull
	// private String username;
	Username string `json:"username"`
}

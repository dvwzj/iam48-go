package entity

type RegisterInfo struct {
	// @SerializedName(Scopes.EMAIL) // "email"
	// @Nullable
	// private String email;
	Email *string `json:"email"`
	// @SerializedName("firstname")
	// @Nullable
	// private String firstname;
	Firstname *string `json:"firstname"`
	// @SerializedName("idCard")
	// @Nullable
	// private String idCard;
	IdCard *string `json:"idCard"`
	// @SerializedName("lastname")
	// @Nullable
	// private String lastname;
	Lastname *string `json:"lastname"`
	// @SerializedName("password")
	// @Nullable
	// private String password;
	Password *string `json:"password"`
	// @SerializedName("phoneNo")
	// @Nullable
	// private String phoneNo;
	PhoneNo *string `json:"phoneNo"`
	// @SerializedName("placeOfIssue")
	// @Nullable
	// private String placeOfIssue;
	PlaceOfIssue *string `json:"placeOfIssue"`
	// @SerializedName("repeatPassword")
	// @Nullable
	// private String repeatPassword;
	RepeatPassword *string `json:"repeatPassword"`
}

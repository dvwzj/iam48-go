package entity

type GreetingMyInfo struct {
	// @SerializedName("approvalInfo")
	// @Nullable
	// private GreetingApprovalInfo approvalInfo;
	ApprovalInfo *GreetingApprovalInfo `json:"approvalInfo"`
	// @SerializedName("createdAt")
	// @Nullable
	// private String createdAt;
	CreatedAt *string `json:"createdAt"`
	// @SerializedName("expireAt")
	// @Nullable
	// private String expireAt;
	ExpireAt *string `json:"expireAt"`
	// @SerializedName("greetingTypeId")
	// @Nullable
	// private Long greetingTypeId;
	GreetingTypeId *int64 `json:"greetingTypeId"`
	// @SerializedName("greetingTypeName")
	// @Nullable
	// private String greetingTypeName;
	GreetingTypeName *string `json:"greetingTypeName"`
	// @SerializedName("iconName")
	// @Nullable
	// private String iconName;
	IconName *string `json:"iconName"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("isAudioOnly")
	// @Nullable
	// private Boolean isAudioOnly;
	IsAudioOnly *bool `json:"isAudioOnly"`
	// @SerializedName("isNeedSubmission")
	// @Nullable
	// private Boolean isNeedSubmission;
	IsNeedSubmission *bool `json:"isNeedSubmission"`
	// @SerializedName("isRequireDialog")
	// @Nullable
	// private Boolean isRequireDialog;
	IsRequireDialog *bool `json:"isRequireDialog"`
	// @SerializedName("isRequireVideo")
	// @Nullable
	// private Boolean isRequireVideo;
	IsRequireVideo *bool `json:"isRequireVideo"`
	// @SerializedName("memberId")
	// @Nullable
	// private Long memberId;
	MemberId *int64 `json:"memberId"`
	// @SerializedName("memberImageUrl")
	// @Nullable
	// private String memberImageUrl;
	MemberImageUrl *string `json:"memberImageUrl"`
	// @SerializedName("memberResponseAt")
	// @Nullable
	// private String memberResponseAt;
	MemberResponseAt *string `json:"memberResponseAt"`
	// @SerializedName("memberVideoLengthSeconds")
	// @Nullable
	// private String memberVideoLengthSeconds;
	MemberVideoLengthSeconds *string `json:"memberVideoLengthSeconds"`
	// @SerializedName("memberVideoThumbnailUrl")
	// @Nullable
	// private String memberVideoThumbnailUrl;
	MemberVideoThumbnailUrl *string `json:"memberVideoThumbnailUrl"`
	// @SerializedName("overallVideoSeconds")
	// @Nullable
	// private Long overallVideoSeconds;
	OverallVideoSeconds *int64 `json:"overallVideoSeconds"`
	// @SerializedName(ConstancesKt.KEY_REF_CODE) // "refCode"
	// @Nullable
	// private String refCode;
	RefCode *string `json:"refCode"`
	// @SerializedName("status")
	// @Nullable
	// private Long status;
	Status *int64 `json:"status"`
	// @SerializedName(ConstancesKt.KEY_TICKET_AMOUNT) // "ticketAmount"
	// @Nullable
	// private Long ticketAmount;
	TicketAmount *int64 `json:"ticketAmount"`
	// @SerializedName("ticketInfo")
	// @Nullable
	// private GreetingTicketInfo ticketInfo;
	TicketInfo *GreetingTicketInfo `json:"ticketInfo"`
	// @SerializedName("userDialogInfo")
	// @Nullable
	// private GreetingUserDialogInfo userDialogInfo;
	UserDialogInfo *GreetingUserDialogInfo `json:"userDialogInfo"`
	// @SerializedName("userLimitedVideoSeconds")
	// @Nullable
	// private Long userLimitedVideoSeconds;
	UserLimitedVideoSeconds *int64 `json:"userLimitedVideoSeconds"`
	// @SerializedName("userVideoInfo")
	// @Nullable
	// private GreetingUserVideoInfo userVideoInfo;
	UserVideoInfo *GreetingUserVideoInfo `json:"userVideoInfo"`
	// @SerializedName("userVideoMinimumSeconds")
	// @Nullable
	// private Long userVideoMinimumSeconds;
	UserVideoMinimumSeconds *int64 `json:"userVideoMinimumSeconds"`
}

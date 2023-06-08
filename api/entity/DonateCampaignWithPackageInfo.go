package entity

type DonateCampaignWithPackageInfo struct {
	// @SerializedName("rewardList")
	// @Nullable
	// private List<CampaignDonateRewardInfo> rewardList;
	RewardList *[]CampaignDonateRewardInfo `json:"rewardList"`
	// @SerializedName("transactionId")
	// @Nullable
	// private String transactionId;
	TransactionId *string `json:"transactionId"`
}

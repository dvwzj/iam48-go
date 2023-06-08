package entity

type MyRewardInfo struct {
	// @SerializedName("claimedAt")
	// @Nullable
	// private String claimedAt;
	ClaimedAt *string `json:"claimedAt"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("imageUrl")
	// @Nullable
	// private String imageUrl;
	ImageUrl *string `json:"imageUrl"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
	// @SerializedName("receivedAt")
	// @Nullable
	// private String receivedAt;
	ReceivedAt *string `json:"receivedAt"`
	// @SerializedName("requiredScore")
	// @Nullable
	// private Long requiredScore;
	RequiredScore *int64 `json:"requiredScore"`
	// @SerializedName("rewardId")
	// @Nullable
	// private Long rewardId;
	RewardId *int64 `json:"rewardId"`
}

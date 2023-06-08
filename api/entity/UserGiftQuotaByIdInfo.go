package entity

type UserGiftQuotaByIdInfo struct {
	// @SerializedName("amount")
	// @Nullable
	// private Long amount;
	Amount *int64 `json:"amount"`
	// @SerializedName("createdAt")
	// @Nullable
	// private String createdAt;
	CreatedAt *string `json:"createdAt"`
	// @SerializedName("description")
	// @Nullable
	// private String description;
	Description *string `json:"description"`
	// @SerializedName("expireAt")
	// @Nullable
	// private String expireAt;
	ExpireAt *string `json:"expireAt"`
	// @SerializedName("giftList")
	// @Nullable
	// private List<GiftListInfo> giftList;
	GiftList []*GiftListInfo `json:"giftList"`
	// @SerializedName("id")
	// @Nullable
	// private Long id;
	Id *int64 `json:"id"`
	// @SerializedName("imageFileUrl")
	// @Nullable
	// private String imageFileUrl;
	ImageFileUrl *string `json:"imageFileUrl"`
	// @SerializedName("memberIdList")
	// @Nullable
	// private List<Long> memberIdList;
	MemberIdList []*int64 `json:"memberIdList"`
	// @SerializedName(AnimatedPasterJsonConfig.CONFIG_NAME) // "name"
	// @Nullable
	// private String name;
	Name *string `json:"name"`
}

package streaming

import (
	"github.com/dvwzj/iam48-go/api/entity"
	"github.com/go-resty/resty/v2"
)

type StreamingRepository interface {
	// @POST("/user/{userId}/transaction/{transactionId}/commit")
	// g.b.l<FloatMessageInfo> floatAndGiftCommit(@Path("userId") long j2, @Path("transactionId") long j3);
	FloatAndGiftCommit(userId int64, transactionId int64) (entity.FloatMessageInfo, error)

	// @POST("/user/{userId}/floatmsg/member-live/{liveId}/request")
	// g.b.l<RequestGiftInfo> floatMessageInit(@Path("userId") long j2, @Path("liveId") long j3, @Body @NotNull FloatMessageChatBody floatMessageChatBody);
	FloatMessageInit(userId int64, liveId int64, floatMessageChatBody entity.FloatMessageChatBody) (entity.RequestGiftInfo, error)

	// @GET("/live/{liveId}")
	// g.b.l<LiveUrlInfo> getLiveUrlInfo(@Path("liveId") long j2);
	GetLiveUrlInfo(liveId int64) (entity.LiveUrlInfo, error)

	// @GET("/user/{userId}/info/member-live/{liveId}")
	// g.b.l<LiveInfoData> getMemberLiveInfo(@Path("userId") @NotNull String str, @Path("liveId") @NotNull String str2);
	GetMemberLiveInfo(userId string, liveId string) (entity.LiveInfoData, error)

	// @GET("/user/{userId}/watch/member-live/{liveId}")
	// g.b.l<LiveUrlInfo> getMemberLiveUrlInfo(@Path("userId") long j2, @Path("liveId") @NotNull String str);
	GetMemberLiveUrlInfo(userId int64, liveId string) (entity.LiveUrlInfo, error)

	// @POST("/user/{userId}/gifts/member-live/{liveId}/request")
	// g.b.l<RequestGiftInfo> giftInit(@Path("userId") long j2, @Path("liveId") long j3, @Body @NotNull GiftBody giftBody);
	GiftInit(userId int64, liveId int64, giftBody entity.GiftBody) (entity.RequestGiftInfo, error)

	// @POST("/user/{userId}/gifts/member-live/{liveId}")
	// g.b.l<f0> sendGift(@Path("userId") long j2, @Path("liveId") @NotNull String str, @Body @NotNull GiftBody giftBody);
	// Renamed to SendMemberLiveGift
	SendMemberLiveGift(userId int64, liveId string, giftBody entity.GiftBody) (resty.Response, error)

	// @PUT("/user/{userId}/favorite/member-live/{liveId}")
	// g.b.l<f0> updateUserFavorite(@Path("userId") @NotNull String str, @Path("liveId") @NotNull String str2);
	UpdateUserFavorite(userId string, liveId string) (resty.Response, error)

	// @DELETE("/user/{userId}/favorite/member-live/{liveId}")
	// g.b.l<f0> updateUserUnFavorite(@Path("userId") @NotNull String str, @Path("liveId") @NotNull String str2);
	UpdateUserUnFavorite(userId string, liveId string) (resty.Response, error)
}

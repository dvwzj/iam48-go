package ticket

import (
	"github.com/dvwzj/iam48-go/api/entity"
	"github.com/go-resty/resty/v2"
)

type TicketRepository interface {
	// @POST("/token-x/access")
	// g.b.l<TokenXAccessResponseBodyInfo> accessTokenX(@Body @NotNull TokenXAccessRequestBodyInfo tokenXAccessRequestBodyInfo);
	AccessTokenX(tokenXAccessRequestBodyInfo entity.TokenXAccessRequestBodyInfo) (entity.TokenXAccessResponseBodyInfo, error)

	// @POST("/token-x/access/fast")
	// g.b.l<TokenXAccessResponseBodyInfo> accessTokenXFast();
	AccessTokenXFast() (entity.TokenXAccessResponseBodyInfo, error)

	// @GET("/pin")
	// g.b.l<f0> checkWalletHasPin();
	CheckWalletHasPin() (resty.Response, error)

	// @DELETE("/wallet")
	// g.b.l<f0> deleteWallet();
	DeleteWallet() (resty.Response, error)

	// @POST("/exchange")
	// g.b.l<f0> exchange(@Body @NotNull ExchangeDataPostBodyInfo exchangeDataPostBodyInfo);
	Exchange(exchangeDataPostBodyInfo entity.ExchangeDataPostBodyInfo) (resty.Response, error)

	// @GET("/exchange/stats")
	// g.b.l<ExchangeStatsResponseInfo> exchangeStats();
	ExchangeStats() (entity.ExchangeStatsResponseInfo, error)

	// @GET("/exchange/ticket/{fromTicketSkuId}/to/{toTicketSkuId}")
	// g.b.l<ExchangeTicketToTicketResponseInfo> exchangeTicket(@Path("fromTicketSkuId") @NotNull String str, @Path("toTicketSkuId") @NotNull String str2);
	ExchangeTicket(fromTicketSkuId string, toTicketSkuId string) (entity.ExchangeTicketToTicketResponseInfo, error)

	// @GET("/exchange/ticket/available")
	// g.b.l<List<ExchangeTicketAvailableResponseInfo>> exchangeTicketAvailable();
	ExchangeTicketAvailable() ([]entity.ExchangeTicketAvailableResponseInfo, error)

	// @GET("/exchange/ticket/{ticketSkuId}/list")
	// g.b.l<List<ExchangeTicketResponseInfo>> exchangeTicketListById(@Path("ticketSkuId") @NotNull String str);
	ExchangeTicketListById(ticketSkuId string) ([]entity.ExchangeTicketResponseInfo, error)

	// @GET("/wallet")
	// g.b.l<WalletInfoModel> getMyWalletInfo();
	GetMyWalletInfo() (entity.WalletInfoModel, error)

	// @GET("/token-x/terms")
	// g.b.l<TokenXTermsResponseInfo> getTermsTokenX();
	GetTermsTokenX() (entity.TokenXTermsResponseInfo, error)

	// @GET("/inventory/{ticketSkuId}")
	// g.b.l<TicketDetailModel> getTicketDetail(@Path("ticketSkuId") long j2);
	GetTicketDetail(ticketSkuId int64) (entity.TicketDetailModel, error)

	// @GET("/history")
	// g.b.l<List<TicketHistoryModel>> getTicketHistory(@Query("skip") long j2, @Query("take") long j3);
	GetTicketHistory(skip int64, take int64) ([]entity.TicketHistoryModel, error)

	// @GET("/inventory")
	// g.b.l<List<TicketInventoryModel>> getTicketInventory(@Query("skip") long j2, @Query("take") long j3);
	GetTicketInventory(skip int64, take int64) ([]entity.TicketInventoryModel, error)

	// @GET("/wallet/{walletCode}")
	// g.b.l<WalletInfoModel> getWalletByWalletCode(@Path("walletCode") @NotNull String str);
	GetWalletByWalletCode(walletCode string) (entity.WalletInfoModel, error)

	// @GET("/qr/wallet")
	// g.b.l<WalletQRCode> getWalletQRCode();
	GetWalletQRCode() (entity.WalletQRCode, error)

	// @GET("/greeting/my/quota")
	// g.b.l<List<GreetingMyInfo>> greetingMyQuota(@Query("take") int i2, @Query("skip") long j2);
	GreetingMyQuota(take int, skip int64) ([]entity.GreetingMyInfo, error)

	// @GET("/greeting/my/replied")
	// g.b.l<List<GreetingMyInfo>> greetingMyReplied(@Query("take") int i2, @Query("skip") long j2);
	GreetingMyReplied(take int, skip int64) ([]entity.GreetingMyInfo, error)

	// @GET("/greeting/my/sent")
	// g.b.l<List<GreetingMyInfo>> greetingMySent(@Query("take") int i2, @Query("skip") long j2);
	GreetingMySent(take int, skip int64) ([]entity.GreetingMyInfo, error)

	// @GET("/greeting/my/stats")
	// g.b.l<GreetingMyStatsInfo> greetingMyStats();
	GreetingMyStats() (entity.GreetingMyStatsInfo, error)

	// @POST("/greeting/redeem")
	// g.b.l<GreetingRedeemResponseInfo> greetingRedeem(@Body @NotNull GreetingRedeemBodyInfo greetingRedeemBodyInfo);
	GreetingRedeem(greetingRedeemBodyInfo entity.GreetingRedeemBodyInfo) (entity.GreetingRedeemResponseInfo, error)

	// @POST("/greeting/redeem/{redeemId}/dialog")
	// g.b.l<GreetingMyInfo> greetingRedeemDialog(@Path("redeemId") long j2, @Body @NotNull GreetingRedeemDialogBodyInfo greetingRedeemDialogBodyInfo);
	GreetingRedeemDialog(redeemId int64, greetingRedeemDialogBodyInfo entity.GreetingRedeemDialogBodyInfo) (entity.GreetingMyInfo, error)

	// @GET("/greeting/redeem/{redeemId}/info")
	// g.b.l<GreetingMyInfo> greetingRedeemInfo(@Path("redeemId") long j2);
	GreetingRedeemInfo(redeemId int64) (entity.GreetingMyInfo, error)

	// @GET("/greeting/redeem/{redeemId}/response")
	// g.b.l<GreetingRedeemResponseByIdInfo> greetingRedeemResponse(@Path("redeemId") long j2);
	GreetingRedeemResponse(redeemId int64) (entity.GreetingRedeemResponseByIdInfo, error)

	// @GET("/greeting/redeem/{redeemId}/dialog")
	// g.b.l<GreetingDialogInfo> greetingRedeemScripDialog(@Path("redeemId") long j2);
	GreetingRedeemScripDialog(redeemId int64) (entity.GreetingDialogInfo, error)

	// @POST("/greeting/redeem/{redeemId}/thumbnail/tempurl")
	// g.b.l<GreetingTempUrlResponseInfo> greetingRedeemThumbnailTempUrl(@Path("redeemId") long j2);
	GreetingRedeemThumbnailTempUrl(redeemId int64) (entity.GreetingTempUrlResponseInfo, error)

	// @GET("/greeting/ticket/redeem-type/{redeemTypeId}")
	// g.b.l<GreetingRedeemTypeInfo> greetingRedeemTypeById(@Path("redeemTypeId") long j2);
	GreetingRedeemTypeById(redeemTypeId int64) (entity.GreetingRedeemTypeInfo, error)

	// @POST("/greeting/redeem/{redeemId}/video")
	// g.b.l<GreetingMyInfo> greetingRedeemVideo(@Path("redeemId") long j2, @Body @NotNull GreetingRedeemVideoBodyInfo greetingRedeemVideoBodyInfo);
	GreetingRedeemVideo(redeemId int64, greetingRedeemVideoBodyInfo entity.GreetingRedeemVideoBodyInfo) (entity.GreetingMyInfo, error)

	// @POST("/greeting/redeem/{redeemId}/video/tempurl")
	// g.b.l<GreetingTempUrlResponseInfo> greetingRedeemVideoTempUrl(@Path("redeemId") long j2);
	GreetingRedeemVideoTempUrl(redeemId int64) (entity.GreetingTempUrlResponseInfo, error)

	// @GET("/greeting/ticket/available")
	// g.b.l<List<GreetingTicketAvailableInfo>> greetingTicketAvailable();
	GreetingTicketAvailable() ([]entity.GreetingTicketAvailableInfo, error)

	// @GET("/greeting/ticket/{ticketId}")
	// g.b.l<GreetingTicketByIdInfo> greetingTicketById(@Path("ticketId") long j2);
	GreetingTicketById(ticketId int64) (entity.GreetingTicketByIdInfo, error)

	// @GET("/greeting/ticket/{ticketId}/member/{memberId}")
	// g.b.l<GreetingTicketMemberTypeInfo> greetingTicketByTicketAndMemberId(@Path("ticketId") long j2, @Path("memberId") long j3);
	GreetingTicketByTicketAndMemberId(ticketId int64, memberId int64) (entity.GreetingTicketMemberTypeInfo, error)

	// @DELETE("/token-x/logout")
	// g.b.l<f0> logoutTokenX(@NotNull @Query("t") String str);
	LogoutTokenX(t string) (resty.Response, error)

	// @GET("/meetingyou/ticket/{ticketId}/event/info")
	// g.b.l<EventInfo> meetYouEventInfo(@Path("ticketId") long j2);
	MeetYouEventInfo(ticketId int64) (entity.EventInfo, error)

	// @GET("/meetingyou/lastminute/member/{memberId}/event/info")
	// g.b.l<EventInfo> meetYouLastMinuteEventInfo(@Path("memberId") long j2);
	MeetYouLastMinuteEventInfo(memberId int64) (entity.EventInfo, error)

	// @GET("/meetingyou/lastminute/member/{memberId}")
	// g.b.l<MemberRounds> meetYouLastMinuteMemberRounds(@Path("memberId") long j2);
	MeetYouLastMinuteMemberRounds(memberId int64) (entity.MemberRounds, error)

	// @GET("/meetingyou/lastminute/members")
	// g.b.l<List<RedeemableMemberInfo>> meetYouLastMinuteRedeemableMemberList();
	MeetYouLastMinuteRedeemableMemberList() ([]entity.RedeemableMemberInfo, error)

	// @GET("meetingyou/lastminute/member/{memberId}/ticket/available")
	// g.b.l<List<TicketInfo>> meetYouLastMinuteTicketAvailable(@Path("memberId") long j2);
	MeetYouLastMinuteTicketAvailable(memberId int64) ([]entity.TicketInfo, error)

	// @GET("/meetingyou/ticket/{ticketId}/event/member/{memberId}")
	// g.b.l<MemberRounds> meetYouMemberRounds(@Path("ticketId") long j2, @Path("memberId") long j3);
	MeetYouMemberRounds(ticketId int64, memberId int64) (entity.MemberRounds, error)

	// @GET("/meetingyou/my/redeem/{refCode}")
	// g.b.l<RedeemInfo> meetYouMyMeetingDetail(@Path("refCode") @NotNull String str);
	MeetYouMyMeetingDetail(refCode string) (entity.RedeemInfo, error)

	// @GET("/meetingyou/my/history")
	// g.b.l<List<RedeemInfo>> meetYouMyMeetingHistory(@Query("take") int i2, @Query("skip") long j2);
	MeetYouMyMeetingHistory(take int, skip int64) ([]entity.RedeemInfo, error)

	// @GET("/meetingyou/my/schedule")
	// g.b.l<List<RedeemInfo>> meetYouMyMeetingSchedule(@Query("take") int i2, @Query("skip") long j2);
	MeetYouMyMeetingSchedule(take int, skip int64) ([]entity.RedeemInfo, error)

	// @GET("/meetingyou/my/stats")
	// g.b.l<MyStats> meetYouMyStats();
	MeetYouMyStats() (entity.MyStats, error)

	// @POST("/meetingyou/redeem")
	// g.b.l<RedeemInfo> meetYouRedeem(@Body @NotNull RedeemMeetYouRequest redeemMeetYouRequest);
	MeetYouRedeem(redeemMeetYouRequest entity.RedeemMeetYouRequest) (entity.RedeemInfo, error)

	// @GET("/meetingyou/ticket/{ticketId}/event/members")
	// g.b.l<List<RedeemableMemberInfo>> meetYouRedeemableMemberList(@Path("ticketId") long j2);
	MeetYouRedeemableMemberList(ticketId int64) ([]entity.RedeemableMemberInfo, error)

	// @GET("/meetingyou/ticket/available")
	// g.b.l<List<TicketInfo>> meetYouTicketAvailable();
	MeetYouTicketAvailable() ([]entity.TicketInfo, error)

	// @POST("/pin/change/confirm")
	// g.b.l<f0> onChangePin(@Body @NotNull ChangePinModel changePinModel);
	OnChangePin(changePinModel entity.ChangePinModel) (resty.Response, error)

	// @POST("/claim/v2/iamshop")
	// g.b.l<List<IamShopClaimOrderItemModel>> onClaimTicketIam(@Body @NotNull IamShopClaimOrderModel iamShopClaimOrderModel);
	OnClaimTicketIam(iamShopClaimOrderModel entity.IamShopClaimOrderModel) ([]entity.IamShopClaimOrderItemModel, error)

	// @POST("/claim/v2/shopee")
	// g.b.l<List<ShopeeClaimOrderItemModel>> onClaimTicketShopee(@Body @NotNull ShopeeClaimOrderModel shopeeClaimOrderModel);
	OnClaimTicketShopee(shopeeClaimOrderModel entity.ShopeeClaimOrderModel) ([]entity.ShopeeClaimOrderItemModel, error)

	// @GET("/claim/v2/shopee")
	// g.b.l<List<ShopeeClaimOrderItemModel>> onClaimTicketShopeeMock();
	OnClaimTicketShopeeMock() ([]entity.ShopeeClaimOrderItemModel, error)

	// @POST("/wallet/create")
	// g.b.l<WalletInfoModel> onCreateWallet(@Body @NotNull WalletPinModel walletPinModel);
	OnCreateWallet(walletPinModel entity.WalletPinModel) (entity.WalletInfoModel, error)

	// @POST("/pin/forgot")
	// g.b.l<f0> onForgotPin();
	OnForgotPin() (resty.Response, error)

	// @POST("/qr/ticket/offline")
	// g.b.l<QRTicketModel> onOfflineTicket(@Body @NotNull ShowQRTicketItem showQRTicketItem);
	OnOfflineTicket(showQRTicketItem entity.ShowQRTicketItem) (entity.QRTicketModel, error)

	// @POST("/qr/ticket")
	// g.b.l<QRTicketModel> onOnlineTicket(@Body @NotNull ShowQRTicketItem showQRTicketItem);
	OnOnlineTicket(showQRTicketItem entity.ShowQRTicketItem) (entity.QRTicketModel, error)

	// @POST("/pin/reset/confirm")
	// g.b.l<f0> onResetPin(@Body @NotNull ResetPinModel resetPinModel);
	OnResetPin(resetPinModel entity.ResetPinModel) (resty.Response, error)

	// @POST("/transfer")
	// g.b.l<TicketTransferModel> onTicketTransfer(@Body @NotNull TicketTransferItem ticketTransferItem);
	OnTicketTransfer(ticketTransferItem entity.TicketTransferItem) (entity.TicketTransferModel, error)

	// @POST("/transfer/commit")
	// g.b.l<TransferCommitResultModel> onTicketTransferCommit(@Body @NotNull TicketTransferCommitItem ticketTransferCommitItem);
	OnTicketTransferCommit(ticketTransferCommitItem entity.TicketTransferCommitItem) (entity.TransferCommitResultModel, error)

	// @POST("/pin/change/verify")
	// g.b.l<f0> onVerifyOldPin(@Body @NotNull WalletPinModel walletPinModel);
	OnVerifyOldPin(walletPinModel entity.WalletPinModel) (resty.Response, error)

	// @POST("/pin/reset/verify")
	// g.b.l<f0> onVerifyResetPin(@Body @NotNull VerifyCodePinModel verifyCodePinModel);
	OnVerifyResetPin(verifyCodePinModel entity.VerifyCodePinModel) (resty.Response, error)

	// @POST("/token-x/terms/{versionNo}")
	// g.b.l<TokenXTermsPostResponseInfo> postTermsTokenX(@Path("versionNo") @NotNull String str);
	PostTermsTokenX(versionNo string) (entity.TokenXTermsPostResponseInfo, error)

	// @PUT("/pin")
	// g.b.l<f0> putWalletPin(@NotNull @Query("pin") String str, @NotNull @Query("password") String str2);
	PutWalletPin(str string, str2 string) (resty.Response, error)
}

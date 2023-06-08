package user

import (
	"github.com/dvwzj/iam48-go/api/entity"
	"github.com/go-resty/resty/v2"
)

type UserRepository interface {
	// @POST("/user/{userId}/address/add")
	// g.b.l<UserAddressInfo> addAddress(@Path("userId") long j2, @Body @NotNull AddUserAddressInfo addUserAddressInfo);
	AddAddress(userId int64, addUserAddressInfo entity.AddUserAddressInfo) (entity.UserAddressInfo, error)

	// @POST("/user/{userId}/comments/mini-video/{contentId}")
	// g.b.l<CommentInfo> addComment(@Path("userId") long j2, @Path("contentId") long j3, @Body @NotNull NewCommentBody newCommentBody);
	AddComment(userId int64, contentId int64, newCommentBody entity.NewCommentBody) (entity.CommentInfo, error)

	// @POST("/user/{userId}/verifyredeem")
	// g.b.l<List<MusicAlbumInfo>> checkVerification(@Path("userId") long j2, @NotNull @Query("code") String str);
	CheckVerification(userId int64, str string) ([]entity.MusicAlbumInfo, error)

	// @POST("/payment/commit")
	// g.b.l<ShopCommitPaymentResponseInfo> commitPayment(@Body @NotNull ShopCommitPaymentBodyInfo shopCommitPaymentBodyInfo);
	CommitPayment(shopCommitPaymentBodyInfo entity.ShopCommitPaymentBodyInfo) (entity.ShopCommitPaymentResponseInfo, error)

	// @POST("/user/{userId}/security/2fa/reset/confirm")
	// g.b.l<SecurityStatusResponseInfo> confirm2FAReset(@Path("userId") long j2, @Body @NotNull Confirm2FAResetBodyInfo confirm2FAResetBodyInfo);
	Confirm2FAReset(userId int64, confirm2FAResetBodyInfo entity.Confirm2FAResetBodyInfo) (entity.SecurityStatusResponseInfo, error)

	// @POST("/user/{userId}/security/2fa/setup/confirm")
	// g.b.l<SecurityStatusResponseInfo> confirm2FASetup(@Path("userId") long j2, @Body @NotNull Confirm2FASetupBodyInfo confirm2FASetupBodyInfo);
	Confirm2FASetup(userId int64, confirm2FASetupBodyInfo entity.Confirm2FASetupBodyInfo) (entity.SecurityStatusResponseInfo, error)

	// @POST("/shop/{userId}/order/create")
	// g.b.l<ShopCreateOrderResponseInfo> createOrder(@Path("userId") long j2, @Body @NotNull ShopCreateOrderBodyInfo shopCreateOrderBodyInfo);
	CreateOrder(userId int64, shopCreateOrderBodyInfo entity.ShopCreateOrderBodyInfo) (entity.ShopCreateOrderResponseInfo, error)

	// @POST("/shop/{userId}/order/create/v2")
	// g.b.l<ShopCreateOrderResponseInfo> createOrderV2(@Path("userId") long j2, @Body @NotNull ShopCreateOrderBodyInfo shopCreateOrderBodyInfo);
	CreateOrderV2(userId int64, shopCreateOrderBodyInfo entity.ShopCreateOrderBodyInfo) (entity.ShopCreateOrderResponseInfo, error)

	// @POST("/crossapp/request")
	// g.b.l<CrossAppResponseBodyInfo> crossAppRequest(@Body @NotNull CrossAppRequestBodyInfo crossAppRequestBodyInfo);
	CrossAppRequest(crossAppRequestBodyInfo entity.CrossAppRequestBodyInfo) (entity.CrossAppResponseBodyInfo, error)

	// @POST("/user/{userId}/deletion")
	// g.b.l<DeleteAccountSuccessResponseInfo> deleteAccount(@Path("userId") long j2, @Body @NotNull DeleteAccountSuccessRequestInfo deleteAccountSuccessRequestInfo);
	DeleteAccount(userId int64, deleteAccountSuccessRequestInfo entity.DeleteAccountSuccessRequestInfo) (entity.DeleteAccountSuccessResponseInfo, error)

	// @DELETE("/user/{userId}/address/{addressId}/delete")
	// g.b.l<f0> deleteAddress(@Path("userId") long j2, @Path("addressId") long j3);
	DeleteAddress(userId int64, addressId int64) (resty.Response, error)

	// @DELETE("/user/{userId}/comments/{commentId}")
	// g.b.l<f0> deleteComment(@Path("userId") long j2, @Path("commentId") long j3);
	DeleteComment(userId int64, commentId int64) (resty.Response, error)

	// @DELETE("/ekyc/{userId}/delete")
	// g.b.l<EkycDeleteUserResponseInfo> deleteEkyc(@Path("userId") long j2);
	DeleteEkyc(userId int64) (entity.EkycDeleteUserResponseInfo, error)

	// @DELETE("/user/{userId}/link")
	// g.b.l<f0> deleteUserLinkFacebook(@Path("userId") long j2);
	DeleteUserLinkFacebook(userId int64) (resty.Response, error)

	// @POST("/user/{userId}/campaigns/package/{packageId}/donate")
	// g.b.l<DonateCampaignWithPackageInfo> donateCampaignWithPackage(@Path("userId") long j2, @Path("packageId") long j3, @Body @NotNull CampaignDonateOrder campaignDonateOrder);
	DonateCampaignWithPackage(userId int64, packageId int64, campaignDonateOrder entity.CampaignDonateOrder) (entity.DonateCampaignWithPackageInfo, error)

	// @PUT("/user/{userId}/address/{addressId}/edit")
	// g.b.l<UserAddressInfo> editAddress(@Path("userId") long j2, @Path("addressId") long j3, @Body @NotNull AddUserAddressInfo addUserAddressInfo);
	EditAddress(userId int64, addressId int64, addUserAddressInfo entity.AddUserAddressInfo) (entity.UserAddressInfo, error)

	// @PUT("/user/{userId}/profile")
	// g.b.l<f0> editUserProfile(@Path("userId") long j2, @Body @NotNull EditUserProfileInfo editUserProfileInfo);
	EditUserProfile(userId int64, editUserProfileInfo entity.EditUserProfileInfo) (resty.Response, error)

	// @POST("/accounts/forgot")
	// g.b.l<f0> forgotPassword(@Body @NotNull ForgotPasswordResponseInfo forgotPasswordResponseInfo);
	ForgotPassword(forgotPasswordResponseInfo entity.ForgotPasswordResponseInfo) (resty.Response, error)

	// @GET("/user/{userId}/deletion/acknowledge")
	// g.b.l<List<DeleteAccountAcknowledgeResponseInfo>> getAccountDeletionAcknowledge(@Path("userId") long j2);
	GetAccountDeletionAcknowledge(userId int64) ([]entity.DeleteAccountAcknowledgeResponseInfo, error)

	// @GET("/user/{userId}/deletion/reason")
	// g.b.l<List<DeleteAccountReasonResponseInfo>> getAccountDeletionReason(@Path("userId") long j2);
	GetAccountDeletionReason(userId int64) ([]entity.DeleteAccountReasonResponseInfo, error)

	// @GET("/user/{userId}/activities/oshi-member")
	// g.b.l<List<ActivitiesOshiMember>> getActivitiesOshiMember(@Path("userId") long j2);
	GetActivitiesOshiMember(userId int64) ([]entity.ActivitiesOshiMember, error)

	// @GET("/user/{userId}/campaigns/participated/campaign/{campaignId}/rank")
	// g.b.l<CampaignTopSupporterMyRank> getCampaignTopSupporterMyRank(@Path("userId") long j2, @Path("campaignId") long j3);
	GetCampaignTopSupporterMyRank(userId int64, campaignId int64) (entity.CampaignTopSupporterMyRank, error)

	// @GET("/user/{userId}/purchase/theater-live/{liveId}")
	// g.b.l<f0> getCanPurchaseEventTheaterLive(@Path("userId") long j2, @Path("liveId") long j3);
	GetCanPurchaseEventTheaterLive(userId int64, liveId int64) (resty.Response, error)

	// @GET("/report/comment/topic")
	// g.b.l<List<UserReportTopicInfo>> getCommentReportTopic();
	GetCommentReportTopic() ([]entity.UserReportTopicInfo, error)

	// @GET("/user/{userId}/badge/current")
	// g.b.l<BadgePremiumInfo> getCurrentBadgePremium(@Path("userId") long j2);
	GetCurrentBadgePremium(userId int64) (entity.BadgePremiumInfo, error)

	// @GET("/ekyc/{userId}/status")
	// g.b.l<EkycGetStatusResponseInfo> getEkycStatus(@Path("userId") long j2);
	GetEkycStatus(userId int64) (entity.EkycGetStatusResponseInfo, error)

	// @GET("/ekyc/token")
	// g.b.l<EkycGetTokenResponseInfo> getEkycToken();
	GetEkycToken() (entity.EkycGetTokenResponseInfo, error)

	// @GET("/auth/facebook/app/{facebookAppId}/parameters")
	// g.b.l<FacebookLinkScopeResponseInfo> getFacebookLinkParameters(@Path("facebookAppId") @NotNull String str);
	GetFacebookLinkParameters(facebookAppId string) (entity.FacebookLinkScopeResponseInfo, error)

	// @GET("/user/{userId}/follows/members")
	// g.b.l<List<FollowMemberInfo>> getFollowMembers(@Path("userId") long j2);
	GetFollowMembers(userId int64) ([]entity.FollowMemberInfo, error)

	// @GET("/user/{userId}/redeemcode")
	// g.b.l<GlobalRedeemCodeInfo> getGlobalRedeemCode(@Path("userId") long j2, @NotNull @Query("code") String str);
	GetGlobalRedeemCode(userId int64, code string) (entity.GlobalRedeemCodeInfo, error)

	// @GET("/user/{userId}/redeemcode/history")
	// g.b.l<List<GlobalRedeemHistory>> getGlobalRedeemHistories(@Path("userId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetGlobalRedeemHistories(userId int64, skip int, take int) ([]entity.GlobalRedeemHistory, error)

	// @GET("/user/{userId}/redeemcode/history/info")
	// g.b.l<GlobalRedeemHistoryDetails> getGlobalRedeemHistoryDetail(@Path("userId") long j2, @NotNull @Query("code") String str);
	GetGlobalRedeemHistoryDetail(userId int64, code string) (entity.GlobalRedeemHistoryDetails, error)

	// @GET("/user/{userId}/follows/member/{memberId}")
	// g.b.l<IsFollowInfo> getIsFollow(@Path("userId") long j2, @Path("memberId") long j3);
	GetIsFollow(userId int64, memberId int64) (entity.IsFollowInfo, error)

	// @HEAD("/user/{userId}/follows/member/{memberId}")
	// g.b.l<IsFollowInfo> getIsFollowByHead(@Path("userId") long j2, @Path("memberId") long j3);
	GetIsFollowByHead(userId int64, memberId int64) (entity.IsFollowInfo, error)

	// @GET("/user/{userId}/likes/member-lives")
	// g.b.l<List<LiveVideoData>> getLikeMemberLives(@Path("userId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetLikeMemberLives(userId int64, skip int, take int) ([]entity.LiveVideoData, error)

	// @GET("/user/{userId}/likes/mini-videos")
	// g.b.l<List<VideoInfo>> getLikeMiniVideos(@Path("userId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetLikeMiniVideos(userId int64, skip int, take int) ([]entity.VideoInfo, error)

	// @GET("/member-live/{contentId}")
	// g.b.l<VideoContentUrl> getLiveVideoByContentId(@Path("contentId") long j2);
	GetLiveVideoByContentId(contentId int64) (entity.VideoContentUrl, error)

	// @GET("/user/{userId}/comments/content/{contentId}/member")
	// g.b.l<UserCommentInfo> getMemberComment(@Path("userId") long j2, @Path("contentId") long j3, @Query("take") int i2, @Nullable @Query("lastCommentId") Long l2);
	GetMemberComment(userId int64, contentId int64, take int, lastCommentId *int64) (entity.UserCommentInfo, error)

	// @GET("/mini-video/{contentId}")
	// g.b.l<VideoContentUrl> getMiniVideoByContentId(@Path("contentId") long j2);
	GetMiniVideoByContentId(contentId int64) (entity.VideoContentUrl, error)

	// @GET("/user/{userId}/comments/content/{contentId}/my")
	// g.b.l<UserCommentInfo> getMyComment(@Path("userId") long j2, @Path("contentId") long j3, @Query("take") int i2, @Nullable @Query("lastCommentId") Long l2);
	GetMyComment(userId int64, contentId int64, take int, lastCommentId *int64) (entity.UserCommentInfo, error)

	// @GET("/user/{userId}/notification")
	// g.b.l<List<NotificationInfo>> getNotification(@Path("userId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetNotification(userId int64, skip int, take int) ([]entity.NotificationInfo, error)

	// @GET("/user/{userId}/notification/sound")
	// g.b.l<NotificationSoundList> getNotificationSoundList(@Path("userId") long j2, @Header("If-None-Match") @NotNull String str);
	GetNotificationSoundList(userId int64, ifNoneMatch string) (entity.NotificationSoundList, error)

	// @GET("/shop/{userId}/order/detail/{orderNo}")
	// g.b.l<ShopOrderDetailResponseInfo> getOrderById(@Path("userId") int i2, @Path("orderNo") @NotNull String str);
	GetOrderById(userId int, orderNo string) (entity.ShopOrderDetailResponseInfo, error)

	// @GET("/shop/{userId}/order/invoice")
	// g.b.l<ShopOrderInvoiceResponseInfo> getOrderInvoice(@Path("userId") long j2, @NotNull @Query("RefCode") String str);
	GetOrderInvoice(userId int64, refCode string) (entity.ShopOrderInvoiceResponseInfo, error)

	// @GET("/shop/{userId}/order/list")
	// g.b.l<List<ShopOrderListResponseInfo>> getOrderList(@Path("userId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetOrderList(userId int64, skip int, take int) ([]entity.ShopOrderListResponseInfo, error)

	// @GET("/shop/{userId}/order/payment/status/{orderNo}")
	// g.b.l<String> getOrderPaymentStatus(@Path("userId") long j2, @Path("orderNo") long j3);
	GetOrderPaymentStatus(userId int64, orderNo int64) (string, error)

	// @GET("/user/{userId}/oshi-member")
	// g.b.l<ActivitiesOshiMember> getOshiMember(@Path("userId") long j2);
	GetOshiMember(userId int64) (entity.ActivitiesOshiMember, error)

	// @GET("/user/{userId}/music/{musicSkuId}")
	// g.b.l<retrofit2.Response<PlayMusicInfo>> getPlayMusic(@Path("userId") long j2, @Path("musicSkuId") long j3);
	GetPlayMusic(userId int64, musicSkuId int64) (entity.PlayMusicInfo, error)

	// @GET("/user/{userId}/comments/content/{contentId}/recently")
	// g.b.l<UserCommentInfo> getRecentlyComment(@Path("userId") long j2, @Path("contentId") long j3, @Query("take") int i2, @Nullable @Query("lastCommentId") Long l2);
	GetRecentlyComment(userId int64, contentId int64, take int, lastCommentId *int64) (entity.UserCommentInfo, error)

	// @GET("/user/{userId}/content/{contentId}/recommend")
	// g.b.l<List<TheaterRecommendedPlaybackInfo>> getRecommendedVideo(@Path("userId") long j2, @Path("contentId") long j3, @Query("take") int i2, @Query("skip") int i3);
	GetRecommendedVideo(userId int64, contentId int64, take int, skip int) ([]entity.TheaterRecommendedPlaybackInfo, error)

	// @POST("/report/problem/generatetempurl")
	// g.b.l<DataPathInfo> getReportProblemImageUrl();
	GetReportProblemImageUrl() (entity.DataPathInfo, error)

	// @GET("/report/user/topic")
	// g.b.l<List<UserReportTopicInfo>> getReportTopic();
	GetReportTopic() ([]entity.UserReportTopicInfo, error)

	// @GET("/user/{userId}/reward-points/balance")
	// g.b.l<RewardPointsInfo> getRewardPoints(@Path("userId") long j2);
	GetRewardPoints(userId int64) (entity.RewardPointsInfo, error)

	// @GET("/user/{userId}/security")
	// g.b.l<SecurityStatusResponseInfo> getSecurityStatus(@Path("userId") long j2);
	GetSecurityStatus(userId int64) (entity.SecurityStatusResponseInfo, error)

	// @POST("/shop/{userId}/order/delivery-rate")
	// g.b.l<ShopShippingFeeResponseInfo> getShippingFee(@Path("userId") long j2, @Body @NotNull ShopShippingFeeBodyInfo shopShippingFeeBodyInfo);
	GetShippingFee(userId int64, shopShippingFeeBodyInfo entity.ShopShippingFeeBodyInfo) (entity.ShopShippingFeeResponseInfo, error)

	// @GET("/subscription/user/{userId}")
	// g.b.l<SubscriptionInfo> getSubscriptionPackage(@Path("userId") long j2);
	GetSubscriptionPackage(userId int64) (entity.SubscriptionInfo, error)

	// @GET("/user/{userId}/campaigns/participated/campaign/{campaignId}")
	// g.b.l<CampaignParticipatedItemInfo> getSupportedCampaignInfo(@Path("userId") long j2, @Path("campaignId") long j3);
	GetSupportedCampaignInfo(userId int64, campaignId int64) (entity.CampaignParticipatedItemInfo, error)

	// @GET("/batch-thankyou-video/{contentId}")
	// g.b.l<ThankYouContentVideoInfo> getThankYouContentVideo(@Path("contentId") long j2);
	GetThankYouContentVideo(contentId int64) (entity.ThankYouContentVideoInfo, error)

	// @GET("/batch-thankyou-video/{contentId}")
	// g.b.l<SignedVideoContent> getThankYouVideoUrl(@Path("contentId") long j2);
	GetThankYouVideoUrl(contentId int64) (entity.SignedVideoContent, error)

	// @GET("/user/{userId}/theater-playback/archive")
	// g.b.l<List<TheaterBoxInfo>> getTheaterBox(@Path("userId") long j2, @Query("take") int i2, @Query("skip") int i3);
	GetTheaterBox(userId int64, take int, skip int) ([]entity.TheaterBoxInfo, error)

	// @POST("/user/{userId}/request/theater-live/{liveId}")
	// g.b.l<TheaterLiveInfo> getTheaterLiveInfo(@Path("userId") long j2, @Path("liveId") long j3);
	GetTheaterLiveInfo(userId int64, liveId int64) (entity.TheaterLiveInfo, error)

	// @GET("/user/{userId}/request/theater-live/{liveId}")
	// g.b.l<TheaterStatLiveInfo> getTheaterLiveStatInfo(@Path("userId") long j2, @Path("liveId") long j3);
	GetTheaterLiveStatInfo(userId int64, liveId int64) (entity.TheaterStatLiveInfo, error)

	// @GET("/user/{userId}/theater-playback/library/list")
	// g.b.l<List<TheaterPlaybackLibraryInfo>> getTheaterPlaybackLibrary(@Path("userId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetTheaterPlaybackLibrary(userId int64, skip int, take int) ([]entity.TheaterPlaybackLibraryInfo, error)

	// @GET("/user/{userId}/theater-playback/library")
	// g.b.l<TheaterPlaybackLibraryCountInfo> getTheaterPlaybackLibraryAmount(@Path("userId") long j2);
	GetTheaterPlaybackLibraryAmount(userId int64) (entity.TheaterPlaybackLibraryCountInfo, error)

	// @GET("/user/{userId}/theater-playback/{contentId}/recommend")
	// g.b.l<List<TheaterRecommendedPlaybackInfo>> getTheaterPlaybackRecommendedVideo(@Path("userId") long j2, @Path("contentId") long j3, @Query("skip") int i2, @Query("take") int i3);
	GetTheaterPlaybackRecommendedVideo(userId int64, contentId int64, skip int, take int) ([]entity.TheaterRecommendedPlaybackInfo, error)

	// @GET("/user/{userId}/watch/theater-playback/{contentId}")
	// g.b.l<SignedVideoContent> getTheaterPlaybackWatch(@Path("userId") long j2, @Path("contentId") long j3);
	GetTheaterPlaybackWatch(userId int64, contentId int64) (entity.SignedVideoContent, error)

	// @GET("/user/{userId}/theaters/{theaterId}/signin")
	// g.b.l<AutoSigninInfo> getTheaterSignin(@Path("userId") long j2, @Path("theaterId") long j3);
	GetTheaterSignin(userId int64, theaterId int64) (entity.AutoSigninInfo, error)

	// @GET("/campaign-result-video/{contentId}")
	// g.b.l<SignedVideoContent> getTimelineCampaignResultVideoUrl(@Path("contentId") long j2);
	GetTimelineCampaignResultVideoUrl(contentId int64) (entity.SignedVideoContent, error)

	// @GET("/user/{userId}/stats/timeline/{contentId}")
	// g.b.l<TimelineMyGiftStats> getTimelineMyGiftStats(@Path("userId") long j2, @Path("contentId") long j3);
	GetTimelineMyGiftStats(userId int64, contentId int64) (entity.TimelineMyGiftStats, error)

	// @GET("/timeline-video/{contentId}")
	// g.b.l<SignedVideoContent> getTimelineVideoUrl(@Path("contentId") long j2);
	GetTimelineVideoUrl(contentId int64) (entity.SignedVideoContent, error)

	// @GET("/user/{userId}/comments/content/{contentId}/top100")
	// g.b.l<UserCommentInfo> getTopComment(@Path("userId") long j2, @Path("contentId") long j3);
	GetTopComment(userId int64, contentId int64) (entity.UserCommentInfo, error)

	// @GET("/user/{userId}/address")
	// g.b.l<List<UserAddressInfo>> getUserAddress(@Path("userId") long j2);
	GetUserAddress(userId int64) ([]entity.UserAddressInfo, error)

	// @GET("/user/{userId}/campaigns/participated")
	// g.b.l<List<CampaignParticipatedItemInfo>> getUserCampaign(@Path("userId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetUserCampaign(userId int64, skip int, take int) ([]entity.CampaignParticipatedItemInfo, error)

	// @GET("/user/{userId}/device")
	// g.b.l<List<DeviceInfo>> getUserDevices(@Path("userId") long j2);
	GetUserDevices(userId int64) ([]entity.DeviceInfo, error)

	// @GET("/user/{userId}/gifts")
	// g.b.l<List<UserGiftsInfo>> getUserGifts(@Path("userId") long j2);
	GetUserGifts(userId int64) ([]entity.UserGiftsInfo, error)

	// @GET("/user/{userId}/gifts/quota")
	// g.b.l<List<UserGiftQuotaInfo>> getUserGiftsQuota(@Path("userId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetUserGiftsQuota(userId int64, skip int, take int) ([]entity.UserGiftQuotaInfo, error)

	// @GET("/user/{userId}/gifts/quota/{quotaId}")
	// g.b.l<UserGiftQuotaByIdInfo> getUserGiftsQuotaById(@Path("userId") long j2, @Path("quotaId") long j3);
	GetUserGiftsQuotaById(userId int64, quotaId int64) (entity.UserGiftQuotaByIdInfo, error)

	// @GET("/user/{userId}/likes/mini-videos")
	// g.b.l<List<LiveVideoData>> getUserLikedVideos(@Path("userId") long j2, @Query("skip") int i2, @Query("take") int i3);
	// Renamed to GetUserLikedMiniVideos
	GetUserLikedMiniVideos(userId int64, skip int, take int) ([]entity.LiveVideoData, error)

	// @GET("/user/{userId}/link/{facebookAppId}")
	// g.b.l<FacebookLinkResponseInfo> getUserLinkFacebook(@Path("userId") long j2, @Path("facebookAppId") @NotNull String str);
	GetUserLinkFacebook(userId int64, facebookAppId string) (entity.FacebookLinkResponseInfo, error)

	// @GET("/user/{userId}/music")
	// g.b.l<retrofit2.Response<List<MusicAlbumInfo>>> getUserMusic(@Path("userId") long j2);
	GetUserMusic(userId int64) ([]entity.MusicAlbumInfo, error)

	// @GET("/user/{userId}/profile")
	// g.b.l<UserProfileInfo> getUserProfile(@Path("userId") long j2);
	GetUserProfile(userId int64) (entity.UserProfileInfo, error)

	// @GET("/user/{userId}/stats/coin")
	// g.b.l<CoreStats> getUserRank(@Path("userId") long j2);
	GetUserRank(userId int64) (entity.CoreStats, error)

	// @GET("/user/{userId}/transactions")
	// g.b.l<List<UserTransactionInfo>> getUserTransaction(@Path("userId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetUserTransaction(userId int64, skip int, take int) ([]entity.UserTransactionInfo, error)

	// @GET("/user/{userId}/content/purchased/list")
	// g.b.l<List<TheaterPlaybackLibraryInfo>> getVideoCollectionList(@Path("userId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetVideoCollectionList(userId int64, skip int, take int) ([]entity.TheaterPlaybackLibraryInfo, error)

	// @GET("/user/{userId}/content/purchased")
	// g.b.l<TheaterPlaybackLibraryCountInfo> getVideoLibraryAmount(@Path("userId") long j2);
	GetVideoLibraryAmount(userId int64) (entity.TheaterPlaybackLibraryCountInfo, error)

	// @POST("/user/{userId}/security/2fa/setup/initial")
	// g.b.l<Initial2FASetupResponseInfo> initial2FASetup(@Path("userId") long j2);
	Initial2FASetup(userId int64) (entity.Initial2FASetupResponseInfo, error)

	// @GET("/user/{userId}/likes/member-live/{memberLiveId}")
	// g.b.l<f0> isLikeMemberLive(@Path("userId") long j2, @Path("memberLiveId") long j3);
	IsLikeMemberLive(userId int64, memberLiveId int64) (resty.Response, error)

	// @GET("/user/{userId}/likes/mini-video/{videoId}")
	// g.b.l<f0> isLikeMiniVideo(@Path("userId") long j2, @Path("videoId") long j3);
	IsLikeMiniVideo(userId int64, videoId int64) (resty.Response, error)

	// @GET("/user/{userId}/likes/like-status/{memberTimelineId}")
	// g.b.l<LikeTimelineInfo> isLikeTimeline(@Path("userId") long j2, @Path("memberTimelineId") long j3);
	IsLikeTimeline(userId int64, memberTimelineId int64) (entity.LikeTimelineInfo, error)

	// @PUT("/user/{userId}/oshi-member")
	// g.b.l<f0> isOshiMember(@Path("userId") long j2, @Body @NotNull OshiMember oshiMember);
	IsOshiMember(userId int64, oshiMember entity.OshiMember) (resty.Response, error)

	// @GET("/user/{userId}/push-notification/subscribe/member-live-all/status")
	// g.b.l<Boolean> isSubscribeMemberLive(@Path("userId") long j2);
	IsSubscribeMemberLive(userId int64) (bool, error)

	// @GET("/user/{userId}/push-notification/subscribe/member/{memberId}/status")
	// g.b.l<Boolean> isSubscribeNotification(@Path("userId") long j2, @Path("memberId") long j3);
	IsSubscribeNotification(userId int64, memberId int64) (bool, error)

	// @GET("/user/{userId}/push-notification/subscribe/theater-live/status")
	// g.b.l<Boolean> isSubscribeTheaterLiveNotification(@Path("userId") long j2);
	IsSubscribeTheaterLiveNotification(userId int64) (bool, error)

	// @PUT("/user/{userId}/likes/member-live/{memberLiveId}")
	// g.b.l<Boolean> likeMemberLive(@Path("userId") long j2, @Path("memberLiveId") long j3);
	LikeMemberLive(userId int64, memberLiveId int64) (bool, error)

	// @PUT("/user/{userId}/likes/mini-video/{videoId}")
	// g.b.l<Boolean> likeMiniVideo(@Path("userId") long j2, @Path("videoId") long j3);
	LikeMiniVideo(userId int64, videoId int64) (bool, error)

	// @PUT("/user/{userId}/likes/member-timeline/{memberTimelineId}")
	// g.b.l<f0> likeTimeline(@Path("userId") long j2, @Path("memberTimelineId") long j3);
	LikeTimeline(userId int64, memberTimelineId int64) (resty.Response, error)

	// @POST("/auth/facebook")
	// g.b.l<AuthenResponseInfo> loginFacebook(@Body @NotNull LoginFacebookBody loginFacebookBody);
	LoginFacebook(loginFacebookBody entity.LoginFacebookBody) (entity.AuthenResponseInfo, error)

	// @POST("/auth/email")
	// g.b.l<AuthenResponseInfo> loginWithEmail(@Body @NotNull LoginInfo loginInfo);
	LoginWithEmail(loginInfo entity.LoginInfo) (entity.AuthenResponseInfo, error)

	// @HTTP(hasBody = true, method = "DELETE", path = "/user/{userId}/device")
	// g.b.l<f0> logoutDevice(@Path("userId") long j2, @Body @NotNull LogoutDeviceBody logoutDeviceBody);
	LogoutDevice(userId int64, logoutDeviceBody entity.LogoutDeviceBody) (resty.Response, error)

	// @GET("/user/{userId}/watch/media/{contentId}")
	// g.b.l<SignedVideoContent> onCheckMediaCanWatch(@Path("userId") long j2, @Path("contentId") long j3);
	OnCheckMediaCanWatch(userId int64, contentId int64) (entity.SignedVideoContent, error)

	// @POST("/user/{userId}/purchase/theater-live/{liveId}")
	// g.b.l<f0> onPurchaseEventTheaterLive(@Path("userId") long j2, @Path("liveId") long j3);
	OnPurchaseEventTheaterLive(userId int64, liveId int64) (resty.Response, error)

	// @POST("/user/{userId}/purchase/media/{contentId}")
	// g.b.l<f0> onPurchaseMedia(@Path("userId") long j2, @Path("contentId") long j3);
	OnPurchaseMedia(userId int64, contentId int64) (resty.Response, error)

	// @GET("/user/{userId}/shop/signin")
	// g.b.l<AutoSigninInfo> onShopAutoSignIn(@Path("userId") long j2);
	OnShopAutoSignIn(userId int64) (entity.AutoSigninInfo, error)

	// @POST("/user/{userId}/push-notification/subscribe/theater-live")
	// g.b.l<Boolean> onSubscribeTheaterLiveNotification(@Path("userId") long j2);
	OnSubscribeTheaterLiveNotification(userId int64) (bool, error)

	// @DELETE("/user/{userId}/push-notification/subscribe/theater-live")
	// g.b.l<Boolean> onUnSubscribeTheaterLiveNotification(@Path("userId") long j2);
	OnUnSubscribeTheaterLiveNotification(userId int64) (bool, error)

	// @POST("/user/{userId}/profile/cover/generateTempUrl")
	// g.b.l<DataPathInfo> postGeneratedCoverUrl(@Path("userId") long j2);
	PostGeneratedCoverUrl(userId int64) (entity.DataPathInfo, error)

	// @POST("/user/{userId}/profile/avatar/generateTempUrl")
	// g.b.l<DataPathInfo> postGeneratedProfileUrl(@Path("userId") long j2);
	PostGeneratedProfileUrl(userId int64) (entity.DataPathInfo, error)

	// @POST("/user/{userId}/redeemcode")
	// g.b.l<GlobalRedeemCodeInfo> postGlobalRedeemCode(@Path("userId") long j2, @NotNull @Query("code") String str);
	PostGlobalRedeemCode(userId int64, str string) (entity.GlobalRedeemCodeInfo, error)

	// @POST("/user/{userId}/purchase/theater-live/{liveId}")
	// g.b.l<f0> postTheaterLivePurchase(@Path("userId") long j2, @Path("liveId") long j3);
	PostTheaterLivePurchase(userId int64, liveId int64) (resty.Response, error)

	// @POST("/user/{userId}/purchase/theater-playback/{contentId}")
	// g.b.l<f0> postTheaterPlaybackPurchase(@Path("userId") long j2, @Path("contentId") long j3);
	PostTheaterPlaybackPurchase(userId int64, contentId int64) (resty.Response, error)

	// @POST("/user/{userId}/device")
	// g.b.l<f0> postUserDevice(@Path("userId") long j2, @Body @NotNull UserDeviceInfo userDeviceInfo);
	PostUserDevice(userId int64, userDeviceInfo entity.UserDeviceInfo) (resty.Response, error)

	// @POST("/user/{userId}/link/facebook")
	// g.b.l<FacebookLinkResponseInfo> postUserLinkFacebook(@Path("userId") long j2, @Body @NotNull FacebookLinkPostInfo facebookLinkPostInfo);
	PostUserLinkFacebook(userId int64, facebookLinkPostInfo entity.FacebookLinkPostInfo) (entity.FacebookLinkResponseInfo, error)

	// @PUT("/user/{userId}/notification/all")
	// g.b.l<Boolean> putAllNotification(@Path("userId") long j2);
	PutAllNotification(userId int64) (bool, error)

	// @PUT("/user/{userId}/follows/member/{memberId}")
	// g.b.u<Object> putFollow(@Path("userId") long j2, @Path("memberId") long j3);
	PutFollow(userId int64, memberId int64) (bool, error)

	// @PUT("/shop/{userId}/order/invoice")
	// g.b.l<ShopCreateOrderResponseInfo> putOrderInvoice(@Path("userId") long j2, @NotNull @Query("RefCode") String str, @NotNull @Query("InvNo") String str2);
	PutOrderInvoice(userId int64, str string, str2 string) (entity.ShopCreateOrderResponseInfo, error)

	// @PUT("/user/{userId}/notification/{notificationId}/read")
	// g.b.l<NotificationInfo> putReadNotification(@Path("userId") long j2, @Path("notificationId") long j3);
	PutReadNotification(userId int64, notificationId int64) (entity.NotificationInfo, error)

	// @DELETE("/user/{userId}/follows/member/{memberId}")
	// g.b.l<Object> putUnfollow(@Path("userId") long j2, @Path("memberId") long j3);
	PutUnfollow(userId int64, memberId int64) (bool, error)

	// @POST("/user/{userId}/redeem")
	// g.b.l<f0> redeemCode(@Path("userId") long j2, @NotNull @Query("code") String str);
	RedeemCode(userId int64, str string) (resty.Response, error)

	// @POST("/auth/refresh")
	// g.b.l<RefreshTokenInfo> refreshToken(@Body @NotNull RefreshTokenBody refreshTokenBody);
	RefreshToken(refreshTokenBody entity.RefreshTokenBody) (entity.RefreshTokenInfo, error)

	// @POST("/accounts/register")
	// g.b.l<RegisResponse> register(@Body @NotNull RegisterInfo registerInfo);
	Register(registerInfo entity.RegisterInfo) (entity.RegisResponse, error)

	// @POST("/user/{userId}/push-notification/register")
	// g.b.l<f0> registerNotification(@Path("userId") long j2, @Body @NotNull NotificationRegisterBody notificationRegisterBody);
	RegisterNotification(userId int64, notificationRegisterBody entity.NotificationRegisterBody) (resty.Response, error)

	// @POST("/report/comment")
	// g.b.l<Boolean> reportComment(@Body @NotNull ReportUserInfo reportUserInfo);
	ReportComment(reportUserInfo entity.ReportUserInfo) (bool, error)

	// @POST("/report/problem")
	// g.b.l<f0> reportProblem(@Body @NotNull ReportProblemInfo reportProblemInfo);
	ReportProblem(reportProblemInfo entity.ReportProblemInfo) (resty.Response, error)

	// @POST("/report/user")
	// g.b.l<Boolean> reportUser(@Body @NotNull ReportUserInfo reportUserInfo);
	ReportUser(reportUserInfo entity.ReportUserInfo) (bool, error)

	// @POST("{fullUrl}")
	// g.b.l<CoreShopPaymentInfo> requestPaymentVendor(@Path(encoded = true, value = "fullUrl") @NotNull String str, @Body @NotNull ShopPaymentRequestBodyInfo shopPaymentRequestBodyInfo);
	RequestPaymentVendor(str string, shopPaymentRequestBodyInfo entity.ShopPaymentRequestBodyInfo) (entity.CoreShopPaymentInfo, error)

	// @POST("/user/{userId}/media/{contentId}/play")
	// g.b.l<f0> saveMediaPlayTime(@Path("userId") long j2, @Path("contentId") long j3, @Query("latestSeconds") int i2);
	SaveMediaPlayTime(userId int64, contentId int64, i2 int) (resty.Response, error)

	// @POST("/user/{userId}/theater-playback/{contentId}/play")
	// g.b.l<f0> savePlaybackPlayTime(@Path("userId") long j2, @Path("contentId") long j3, @Query("latestSeconds") int i2);
	SavePlaybackPlayTime(userId int64, contentId int64, i2 int) (resty.Response, error)

	// @POST("/user/{userId}/security/2fa/setup/email-otp/send")
	// g.b.l<SendEmailOTPResponseInfo> sendEmailOTP2FASetup(@Path("userId") long j2);
	SendEmailOTP2FASetup(userId int64) (entity.SendEmailOTPResponseInfo, error)

	// @POST("/user/{userId}/deletion/email-otp/send")
	// g.b.l<SendEmailOTPResponseInfo> sendEmailOTPAccountDeletion(@Path("userId") long j2, @NotNull @Query("ref") String str);
	SendEmailOTPAccountDeletion(userId int64, str string) (entity.SendEmailOTPResponseInfo, error)

	// @POST("/user/{userId}/security/2fa/reset/email-otp-1st/send")
	// g.b.l<SendOTPEmailResponseInfo> sendFirstOTPEmail2FAReset(@Path("userId") long j2);
	SendFirstOTPEmail2FAReset(userId int64) (entity.SendOTPEmailResponseInfo, error)

	// @POST("/user/{userId}/gifts/member/{memberId}")
	// g.b.l<f0> sendGift(@Path("userId") long j2, @Path("memberId") long j3, @Body @NotNull SendGiftInfo sendGiftInfo);
	SendGift(userId int64, memberId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error)

	// @POST("/user/{userId}/gifts/campaign-result/{contentId}")
	// g.b.l<f0> sendGiftCampaignResult(@Path("userId") long j2, @Path("contentId") long j3, @Body @NotNull SendGiftInfo sendGiftInfo);
	SendGiftCampaignResult(userId int64, contentId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error)

	// @POST("/user/{userId}/gifts/comment/{commentId}")
	// g.b.l<f0> sendGiftComment(@Path("userId") long j2, @Path("commentId") long j3, @Body @NotNull SendGiftInfo sendGiftInfo);
	SendGiftComment(userId int64, commentId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error)

	// @POST("/user/{userId}/gifts/digital-studio/{contentId}")
	// g.b.l<f0> sendGiftDigitalStudio(@Path("userId") long j2, @Path("contentId") long j3, @Body @NotNull SendGiftInfo sendGiftInfo);
	SendGiftDigitalStudio(userId int64, contentId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error)

	// @POST("/user/{userId}/gifts/meeting-you/{refCode}/lobby")
	// g.b.l<f0> sendGiftMeetingLobbyRoom(@Path("userId") long j2, @Path("refCode") @NotNull String str, @Body @NotNull SendGiftInfo sendGiftInfo);
	SendGiftMeetingLobbyRoom(userId int64, str string, sendGiftInfo entity.SendGiftInfo) (resty.Response, error)

	// @POST("/user/{userId}/gifts/meeting-you/{refCode}/thankyou")
	// g.b.l<SendMeetingGiftThankYouResponse> sendGiftMeetingThankYou(@Path("userId") long j2, @Path("refCode") @NotNull String str, @Body @NotNull SendGiftInfo sendGiftInfo);
	SendGiftMeetingThankYou(userId int64, str string, sendGiftInfo entity.SendGiftInfo) (entity.SendMeetingGiftThankYouResponse, error)

	// @POST("/user/{userId}/gifts/member-live/{contentId}")
	// g.b.l<f0> sendGiftMemberLive(@Path("userId") long j2, @Path("contentId") long j3, @Body @NotNull SendGiftInfo sendGiftInfo);
	SendGiftMemberLive(userId int64, contentId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error)

	// @POST("/user/{userId}/gifts/mini-video/{contentId}")
	// g.b.l<f0> sendGiftMiniVideo(@Path("userId") long j2, @Path("contentId") long j3, @Body @NotNull SendGiftInfo sendGiftInfo);
	SendGiftMiniVideo(userId int64, contentId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error)

	// @POST("/user/{userId}/gifts/batch-ranking/{memberId}")
	// g.b.l<f0> sendGiftThankyou(@Path("userId") long j2, @Path("memberId") long j3, @Body @NotNull SendGiftInfo sendGiftInfo);
	SendGiftThankyou(userId int64, memberId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error)

	// @POST("/user/{userId}/gifts/theater/{contentId}")
	// g.b.l<f0> sendGiftTheater(@Path("userId") long j2, @Path("contentId") long j3, @Body @NotNull SendGiftInfo sendGiftInfo);
	SendGiftTheater(userId int64, contentId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error)

	// @POST("/user/{userId}/gifts/timeline/{contentId}")
	// g.b.l<f0> sendGiftTimeline(@Path("userId") long j2, @Path("contentId") long j3, @Body @NotNull SendGiftInfo sendGiftInfo);
	SendGiftTimeline(userId int64, contentId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error)

	// @POST("/push-notification/message/send")
	// g.b.l<f0> sendNotification(@Body @NotNull SendNotificationInfo sendNotificationInfo);
	SendNotification(sendNotificationInfo entity.SendNotificationInfo) (resty.Response, error)

	// @POST("/user/{userId}/security/2fa/reset/email-otp-2nd/send")
	// g.b.l<SendOTPEmailResponseInfo> sendSecondOTPEmail2FAReset(@Path("userId") long j2);
	SendSecondOTPEmail2FAReset(userId int64) (entity.SendOTPEmailResponseInfo, error)

	// @PUT("/user/{userId}/comments/content/{contentId}")
	// g.b.l<UserCommentResponseInfo> sendUserComment(@Path("userId") long j2, @Path("contentId") long j3, @Body @NotNull UserCommentBodyInfo userCommentBodyInfo);
	SendUserComment(userId int64, contentId int64, userCommentBodyInfo entity.UserCommentBodyInfo) (entity.UserCommentResponseInfo, error)

	// @POST("/user/{userId}/push-notification/subscribe/member-live-all")
	// g.b.l<Boolean> subScribeMeberLive(@Path("userId") long j2);
	SubScribeMeberLive(userId int64) (bool, error)

	// @POST("/user/{userId}/push-notification/subscribe/member/{memberId}")
	// g.b.l<Boolean> subscribeNotification(@Path("userId") long j2, @Path("memberId") long j3);
	SubscribeNotification(userId int64, memberId int64) (bool, error)

	// @DELETE("/user/{userId}/likes/member-live/{memberLiveId}")
	// g.b.l<Boolean> unLikeMemberLive(@Path("userId") long j2, @Path("memberLiveId") long j3);
	UnLikeMemberLive(userId int64, memberLiveId int64) (bool, error)

	// @DELETE("/user/{userId}/likes/mini-video/{videoId}")
	// g.b.l<Boolean> unLikeMiniVideo(@Path("userId") long j2, @Path("videoId") long j3);
	UnLikeMiniVideo(userId int64, videoId int64) (bool, error)

	// @DELETE("/user/{userId}/oshi-member")
	// g.b.l<f0> unOshiMember(@Path("userId") long j2);
	UnOshiMember(userId int64) (resty.Response, error)

	// @DELETE("/user/{userId}/push-notification/subscribe/member-live-all")
	// g.b.l<Boolean> unSubscribeMemberLive(@Path("userId") long j2);
	UnSubscribeMemberLive(userId int64) (bool, error)

	// @DELETE("/user/{userId}/push-notification/subscribe/member/{memberId}")
	// g.b.l<Boolean> unSubscribeNotification(@Path("userId") long j2, @Path("memberId") long j3);
	UnSubscribeNotification(userId int64, memberId int64) (bool, error)

	// @DELETE("/user/{userId}/likes/member-timeline/{memberTimelineId}")
	// g.b.l<f0> unlikeTimeline(@Path("userId") long j2, @Path("memberTimelineId") long j3);
	UnlikeTimeline(userId int64, memberTimelineId int64) (resty.Response, error)

	// @HTTP(hasBody = true, method = "DELETE", path = "/push-notification/register")
	// g.b.l<f0> unregisterNotification(@Body @NotNull NotificationRegisterBody notificationRegisterBody);
	UnregisterNotification(notificationRegisterBody entity.NotificationRegisterBody) (resty.Response, error)

	// @POST("/ekyc/{userId}/verify")
	// g.b.l<EkycVerifyUserResponseInfo> verifyEkyc(@Path("userId") long j2, @Body @NotNull EkycVerifyUserBodyInfo ekycVerifyUserBodyInfo);
	VerifyEkyc(userId int64, ekycVerifyUserBodyInfo entity.EkycVerifyUserBodyInfo) (entity.EkycVerifyUserResponseInfo, error)

	// @POST("/user/{userId}/security/2fa/reset/email-otp-1st/verify")
	// g.b.l<VerifyFirstOTPEmailResponseInfo> verifyFirstOTPEmail2FAReset(@Path("userId") long j2, @Body @NotNull VerifyFirstOTPEmailBodyInfo verifyFirstOTPEmailBodyInfo);
	VerifyFirstOTPEmail2FAReset(userId int64, verifyFirstOTPEmailBodyInfo entity.VerifyFirstOTPEmailBodyInfo) (entity.VerifyFirstOTPEmailResponseInfo, error)

	// @GET("/user/{userId}/verify/founder")
	// g.b.l<f0> verifyFounder(@Path("userId") long j2);
	VerifyFounder(userId int64) (resty.Response, error)

	// @POST("/user/{userId}/senbatsu-votes/member/{memberId}")
	// g.b.l<VoteSenbatsuResponseInfo> voteSenbatsu(@Path("userId") long j2, @Path("memberId") long j3, @NotNull @Query("voteCode") String str);
	VoteSenbatsu(userId int64, memberId int64, str string) (entity.VoteSenbatsuResponseInfo, error)
}

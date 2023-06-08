package user

import (
	"strconv"

	"github.com/dvwzj/iam48-go/api/entity"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type User struct {
	Client *resty.Client
}

func (s User) AddAddress(userId int64, addUserAddressInfo entity.AddUserAddressInfo) (entity.UserAddressInfo, error) {
	var userAddressInfo entity.UserAddressInfo
	_, err := s.Client.R().
		SetBody(addUserAddressInfo).
		SetResult(&userAddressInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Post("/user/{userId}/address/add")
	if err != nil {
		return entity.UserAddressInfo{}, err
	}
	return userAddressInfo, nil
}

func (s User) AddComment(userId int64, contentId int64, newCommentBody entity.NewCommentBody) (entity.CommentInfo, error) {
	var commentInfo entity.CommentInfo
	_, err := s.Client.R().
		SetBody(newCommentBody).
		SetResult(&commentInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Post("/user/{userId}/comments/mini-video/{contentId}")
	if err != nil {
		return entity.CommentInfo{}, err
	}
	return commentInfo, nil
}

func (s User) CheckVerification(userId int64, str string) ([]entity.MusicAlbumInfo, error) {
	var musicAlbumInfo []entity.MusicAlbumInfo
	_, err := s.Client.R().
		SetResult(&musicAlbumInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("code", str).
		Post("/user/{userId}/verifyredeem")
	if err != nil {
		return nil, err
	}
	return musicAlbumInfo, nil
}

func (s User) CommitPayment(shopCommitPaymentBodyInfo entity.ShopCommitPaymentBodyInfo) (entity.ShopCommitPaymentResponseInfo, error) {
	var shopCommitPaymentResponseInfo entity.ShopCommitPaymentResponseInfo
	_, err := s.Client.R().
		SetBody(shopCommitPaymentBodyInfo).
		SetResult(&shopCommitPaymentResponseInfo).
		Post("/payment/commit")
	if err != nil {
		return entity.ShopCommitPaymentResponseInfo{}, err
	}
	return shopCommitPaymentResponseInfo, nil
}

func (s User) Confirm2FAReset(userId int64, confirm2FAResetBodyInfo entity.Confirm2FAResetBodyInfo) (entity.SecurityStatusResponseInfo, error) {
	var securityStatusResponseInfo entity.SecurityStatusResponseInfo
	_, err := s.Client.R().
		SetBody(confirm2FAResetBodyInfo).
		SetResult(&securityStatusResponseInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Post("/user/{userId}/security/2fa/reset/confirm")
	if err != nil {
		return entity.SecurityStatusResponseInfo{}, err
	}
	return securityStatusResponseInfo, nil
}

func (s User) Confirm2FASetup(userId int64, confirm2FASetupBodyInfo entity.Confirm2FASetupBodyInfo) (entity.SecurityStatusResponseInfo, error) {
	var securityStatusResponseInfo entity.SecurityStatusResponseInfo
	_, err := s.Client.R().
		SetBody(confirm2FASetupBodyInfo).
		SetResult(&securityStatusResponseInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Post("/user/{userId}/security/2fa/setup/confirm")
	if err != nil {
		return entity.SecurityStatusResponseInfo{}, err
	}
	return securityStatusResponseInfo, nil
}

func (s User) CreateOrder(userId int64, shopCreateOrderBodyInfo entity.ShopCreateOrderBodyInfo) (entity.ShopCreateOrderResponseInfo, error) {
	var shopCreateOrderResponseInfo entity.ShopCreateOrderResponseInfo
	_, err := s.Client.R().
		SetBody(shopCreateOrderBodyInfo).
		SetResult(&shopCreateOrderResponseInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Post("/shop/{userId}/order/create")
	if err != nil {
		return entity.ShopCreateOrderResponseInfo{}, err
	}
	return shopCreateOrderResponseInfo, nil
}

func (s User) CreateOrderV2(userId int64, shopCreateOrderBodyInfo entity.ShopCreateOrderBodyInfo) (entity.ShopCreateOrderResponseInfo, error) {
	var shopCreateOrderResponseInfo entity.ShopCreateOrderResponseInfo
	_, err := s.Client.R().
		SetBody(shopCreateOrderBodyInfo).
		SetResult(&shopCreateOrderResponseInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Post("/shop/{userId}/order/create/v2")
	if err != nil {
		return entity.ShopCreateOrderResponseInfo{}, err
	}
	return shopCreateOrderResponseInfo, nil
}

func (s User) CrossAppRequest(crossAppRequestBodyInfo entity.CrossAppRequestBodyInfo) (entity.CrossAppResponseBodyInfo, error) {
	var crossAppResponseBodyInfo entity.CrossAppResponseBodyInfo
	_, err := s.Client.R().
		SetBody(crossAppRequestBodyInfo).
		SetResult(&crossAppResponseBodyInfo).
		Post("/crossapp/request")
	if err != nil {
		return entity.CrossAppResponseBodyInfo{}, err
	}
	return crossAppResponseBodyInfo, nil
}

func (s User) DeleteAccount(userId int64, deleteAccountSuccessRequestInfo entity.DeleteAccountSuccessRequestInfo) (entity.DeleteAccountSuccessResponseInfo, error) {
	var deleteAccountSuccessResponseInfo entity.DeleteAccountSuccessResponseInfo
	_, err := s.Client.R().
		SetBody(deleteAccountSuccessRequestInfo).
		SetResult(&deleteAccountSuccessResponseInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Post("/user/{userId}/deletion")
	if err != nil {
		return entity.DeleteAccountSuccessResponseInfo{}, err
	}
	return deleteAccountSuccessResponseInfo, nil
}

func (s User) DeleteAddress(userId int64, addressId int64) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("addressId", strconv.FormatInt(addressId, 10)).
		Delete("/user/{userId}/address/{addressId}/delete")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) DeleteComment(userId int64, commentId int64) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("commentId", strconv.FormatInt(commentId, 10)).
		Delete("/user/{userId}/comments/{commentId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) DeleteEkyc(userId int64) (entity.EkycDeleteUserResponseInfo, error) {
	var ekycDeleteUserResponseInfo entity.EkycDeleteUserResponseInfo
	_, err := s.Client.R().
		SetResult(&ekycDeleteUserResponseInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Delete("/ekyc/{userId}/delete")
	if err != nil {
		return entity.EkycDeleteUserResponseInfo{}, err
	}
	return ekycDeleteUserResponseInfo, nil
}

func (s User) DeleteUserLinkFacebook(userId int64) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Delete("/user/{userId}/link")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) DonateCampaignWithPackage(userId int64, packageId int64, campaignDonateOrder entity.CampaignDonateOrder) (entity.DonateCampaignWithPackageInfo, error) {
	var donateCampaignWithPackageInfo entity.DonateCampaignWithPackageInfo
	_, err := s.Client.R().
		SetBody(campaignDonateOrder).
		SetResult(&donateCampaignWithPackageInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("packageId", strconv.FormatInt(packageId, 10)).
		Post("/user/{userId}/campaigns/package/{packageId}/donate")
	if err != nil {
		return entity.DonateCampaignWithPackageInfo{}, err
	}
	return donateCampaignWithPackageInfo, nil
}

func (s User) EditAddress(userId int64, addressId int64, addUserAddressInfo entity.AddUserAddressInfo) (entity.UserAddressInfo, error) {
	var userAddressInfo entity.UserAddressInfo
	_, err := s.Client.R().
		SetBody(addUserAddressInfo).
		SetResult(&userAddressInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("addressId", strconv.FormatInt(addressId, 10)).
		Put("/user/{userId}/address/{addressId}/edit")
	if err != nil {
		return entity.UserAddressInfo{}, err
	}
	return userAddressInfo, nil
}

func (s User) EditUserProfile(userId int64, editUserProfileInfo entity.EditUserProfileInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetBody(editUserProfileInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Put("/user/{userId}/profile")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) ForgotPassword(forgotPasswordResponseInfo entity.ForgotPasswordResponseInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetBody(forgotPasswordResponseInfo).
		Post("/accounts/forgot")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) GetAccountDeletionAcknowledge(userId int64) ([]entity.DeleteAccountAcknowledgeResponseInfo, error) {
	var deleteAccountAcknowledgeResponseInfo []entity.DeleteAccountAcknowledgeResponseInfo
	_, err := s.Client.R().
		SetResult(&deleteAccountAcknowledgeResponseInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/deletion/acknowledge")
	if err != nil {
		return nil, err
	}
	return deleteAccountAcknowledgeResponseInfo, nil
}

func (s User) GetAccountDeletionReason(userId int64) ([]entity.DeleteAccountReasonResponseInfo, error) {
	var deleteAccountReasonResponseInfo []entity.DeleteAccountReasonResponseInfo
	_, err := s.Client.R().
		SetResult(&deleteAccountReasonResponseInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/deletion/reason")
	if err != nil {
		return nil, err
	}
	return deleteAccountReasonResponseInfo, nil
}

func (s User) GetActivitiesOshiMember(userId int64) ([]entity.ActivitiesOshiMember, error) {
	var activitiesOshiMember []entity.ActivitiesOshiMember
	_, err := s.Client.R().
		SetResult(&activitiesOshiMember).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/activities/oshi-member")
	if err != nil {
		return nil, err
	}
	return activitiesOshiMember, nil
}

func (s User) GetCampaignTopSupporterMyRank(userId int64, campaignId int64) (entity.CampaignTopSupporterMyRank, error) {
	var campaignTopSupporterMyRank entity.CampaignTopSupporterMyRank
	_, err := s.Client.R().
		SetResult(&campaignTopSupporterMyRank).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("campaignId", strconv.FormatInt(campaignId, 10)).
		Get("/user/{userId}/campaigns/participated/campaign/{campaignId}/rank")
	if err != nil {
		return entity.CampaignTopSupporterMyRank{}, err
	}
	return campaignTopSupporterMyRank, nil
}

func (s User) GetCanPurchaseEventTheaterLive(userId int64, liveId int64) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("liveId", strconv.FormatInt(liveId, 10)).
		Get("/user/{userId}/purchase/theater-live/{liveId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) GetCommentReportTopic() ([]entity.UserReportTopicInfo, error) {
	var userReportTopicInfo []entity.UserReportTopicInfo
	_, err := s.Client.R().
		SetResult(&userReportTopicInfo).
		Get("/report/comment/topic")
	if err != nil {
		return nil, err
	}
	return userReportTopicInfo, nil
}

func (s User) GetCurrentBadgePremium(userId int64) (entity.BadgePremiumInfo, error) {
	var badgePremiumInfo entity.BadgePremiumInfo
	_, err := s.Client.R().
		SetResult(&badgePremiumInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/badge/current")
	if err != nil {
		return entity.BadgePremiumInfo{}, err
	}
	return badgePremiumInfo, nil
}

func (s User) GetEkycStatus(userId int64) (entity.EkycGetStatusResponseInfo, error) {
	var ekycGetStatusResponseInfo entity.EkycGetStatusResponseInfo
	_, err := s.Client.R().
		SetResult(&ekycGetStatusResponseInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/ekyc/{userId}/status")
	if err != nil {
		return entity.EkycGetStatusResponseInfo{}, err
	}
	return ekycGetStatusResponseInfo, nil
}

func (s User) GetEkycToken() (entity.EkycGetTokenResponseInfo, error) {
	var ekycGetTokenResponseInfo entity.EkycGetTokenResponseInfo
	_, err := s.Client.R().
		SetResult(&ekycGetTokenResponseInfo).
		Get("/ekyc/token")
	if err != nil {
		return entity.EkycGetTokenResponseInfo{}, err
	}
	return ekycGetTokenResponseInfo, nil
}

func (s User) GetFacebookLinkParameters(facebookAppId string) (entity.FacebookLinkScopeResponseInfo, error) {
	var facebookLinkScopeResponseInfo entity.FacebookLinkScopeResponseInfo
	_, err := s.Client.R().
		SetResult(&facebookLinkScopeResponseInfo).
		SetPathParam("facebookAppId", facebookAppId).
		Get("/auth/facebook/app/{facebookAppId}/parameters")
	if err != nil {
		return entity.FacebookLinkScopeResponseInfo{}, err
	}
	return facebookLinkScopeResponseInfo, nil
}

func (s User) GetFollowMembers(userId int64) ([]entity.FollowMemberInfo, error) {
	var followMemberInfo []entity.FollowMemberInfo
	_, err := s.Client.R().
		SetResult(&followMemberInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/follows/members")
	if err != nil {
		return nil, err
	}
	return followMemberInfo, nil
}

func (s User) GetGlobalRedeemCode(userId int64, code string) (entity.GlobalRedeemCodeInfo, error) {
	var globalRedeemCodeInfo entity.GlobalRedeemCodeInfo
	_, err := s.Client.R().
		SetResult(&globalRedeemCodeInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("code", code).
		Get("/user/{userId}/redeemcode")
	if err != nil {
		return entity.GlobalRedeemCodeInfo{}, err
	}
	return globalRedeemCodeInfo, nil
}

func (s User) GetGlobalRedeemHistories(userId int64, skip int, take int) ([]entity.GlobalRedeemHistory, error) {
	var globalRedeemHistory []entity.GlobalRedeemHistory
	_, err := s.Client.R().
		SetResult(&globalRedeemHistory).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/user/{userId}/redeemcode/history")
	if err != nil {
		return nil, err
	}
	return globalRedeemHistory, nil
}

func (s User) GetGlobalRedeemHistoryDetail(userId int64, code string) (entity.GlobalRedeemHistoryDetails, error) {
	var globalRedeemHistoryDetails entity.GlobalRedeemHistoryDetails
	_, err := s.Client.R().
		SetResult(&globalRedeemHistoryDetails).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("code", code).
		Get("/user/{userId}/redeemcode/history/info")
	if err != nil {
		return entity.GlobalRedeemHistoryDetails{}, err
	}
	return globalRedeemHistoryDetails, nil
}

func (s User) GetIsFollow(userId int64, memberId int64) (entity.IsFollowInfo, error) {
	var isFollowInfo entity.IsFollowInfo
	_, err := s.Client.R().
		SetResult(&isFollowInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Get("/user/{userId}/follows/member/{memberId}")
	if err != nil {
		return entity.IsFollowInfo{}, err
	}
	return isFollowInfo, nil
}

func (s User) GetIsFollowByHead(userId int64, memberId int64) (entity.IsFollowInfo, error) {
	var isFollowInfo entity.IsFollowInfo
	_, err := s.Client.R().
		SetResult(&isFollowInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Head("/user/{userId}/follows/member/{memberId}")
	if err != nil {
		return entity.IsFollowInfo{}, err
	}
	return isFollowInfo, nil
}

func (s User) GetLikeMemberLives(userId int64, skip int, take int) ([]entity.LiveVideoData, error) {
	var liveVideoData []entity.LiveVideoData
	_, err := s.Client.R().
		SetResult(&liveVideoData).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/user/{userId}/likes/member-lives")
	if err != nil {
		return nil, err
	}
	return liveVideoData, nil
}

func (s User) GetLikeMiniVideos(userId int64, skip int, take int) ([]entity.VideoInfo, error) {
	var videoInfo []entity.VideoInfo
	_, err := s.Client.R().
		SetResult(&videoInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/user/{userId}/likes/mini-videos")
	if err != nil {
		return nil, err
	}
	return videoInfo, nil
}

func (s User) GetLiveVideoByContentId(contentId int64) (entity.VideoContentUrl, error) {
	var videoContentUrl entity.VideoContentUrl
	_, err := s.Client.R().
		SetResult(&videoContentUrl).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Get("/member-live/{contentId}")
	if err != nil {
		return entity.VideoContentUrl{}, err
	}
	return videoContentUrl, nil
}

func (s User) GetMemberComment(userId int64, contentId int64, take int, lastCommentId *int64) (entity.UserCommentInfo, error) {
	var userCommentInfo entity.UserCommentInfo
	r := s.Client.R().
		SetResult(&userCommentInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetQueryParam("take", strconv.Itoa(take))
	if lastCommentId != nil {
		r.SetQueryParam("lastCommentId", strconv.FormatInt(*lastCommentId, 10))
	}
	_, err := r.Get("/user/{userId}/comments/content/{contentId}/member")
	if err != nil {
		return entity.UserCommentInfo{}, err
	}
	return userCommentInfo, nil
}

func (s User) GetMiniVideoByContentId(contentId int64) (entity.VideoContentUrl, error) {
	var videoContentUrl entity.VideoContentUrl
	_, err := s.Client.R().
		SetResult(&videoContentUrl).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Get("/mini-video/{contentId}")
	if err != nil {
		return entity.VideoContentUrl{}, err
	}
	return videoContentUrl, nil
}

func (s User) GetMyComment(userId int64, contentId int64, take int, lastCommentId *int64) (entity.UserCommentInfo, error) {
	var userCommentInfo entity.UserCommentInfo
	r := s.Client.R().
		SetResult(&userCommentInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetQueryParam("take", strconv.Itoa(take))
	if lastCommentId != nil {
		r.SetQueryParam("lastCommentId", strconv.FormatInt(*lastCommentId, 10))
	}
	_, err := r.Get("/user/{userId}/comments/content/{contentId}/my")
	if err != nil {
		return entity.UserCommentInfo{}, err
	}
	return userCommentInfo, nil
}

func (s User) GetNotification(userId int64, skip int, take int) ([]entity.NotificationInfo, error) {
	var notificationInfo []entity.NotificationInfo
	_, err := s.Client.R().
		SetResult(&notificationInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/user/{userId}/notification")
	if err != nil {
		return nil, err
	}
	return notificationInfo, nil
}

func (s User) GetNotificationSoundList(userId int64, ifNoneMatch string) (entity.NotificationSoundList, error) {
	var notificationSoundList entity.NotificationSoundList
	_, err := s.Client.R().
		SetResult(&notificationSoundList).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetHeader("If-None-Match", ifNoneMatch).
		Get("/user/{userId}/notification/sound")
	if err != nil {
		return entity.NotificationSoundList{}, err
	}
	return notificationSoundList, nil
}

func (s User) GetOrderById(userId int, orderNo string) (entity.ShopOrderDetailResponseInfo, error) {
	var shopOrderDetailResponseInfo entity.ShopOrderDetailResponseInfo
	_, err := s.Client.R().
		SetResult(&shopOrderDetailResponseInfo).
		SetPathParam("userId", strconv.Itoa(userId)).
		SetPathParam("orderNo", orderNo).
		Get("/shop/{userId}/order/detail/{orderNo}")
	if err != nil {
		return entity.ShopOrderDetailResponseInfo{}, err
	}
	return shopOrderDetailResponseInfo, nil
}

func (s User) GetOrderInvoice(userId int64, refCode string) (entity.ShopOrderInvoiceResponseInfo, error) {
	var shopOrderInvoiceResponseInfo entity.ShopOrderInvoiceResponseInfo
	_, err := s.Client.R().
		SetResult(&shopOrderInvoiceResponseInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("RefCode", refCode).
		Get("/shop/{userId}/order/invoice")
	if err != nil {
		return entity.ShopOrderInvoiceResponseInfo{}, err
	}
	return shopOrderInvoiceResponseInfo, nil
}

func (s User) GetOrderList(userId int64, skip int, take int) ([]entity.ShopOrderListResponseInfo, error) {
	var shopOrderListResponseInfo []entity.ShopOrderListResponseInfo
	_, err := s.Client.R().
		SetResult(&shopOrderListResponseInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/shop/{userId}/order/list")
	if err != nil {
		return nil, err
	}
	return shopOrderListResponseInfo, nil
}

func (s User) GetOrderPaymentStatus(userId int64, orderNo int64) (string, error) {
	var str string
	_, err := s.Client.R().
		SetResult(&str).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("orderNo", strconv.FormatInt(orderNo, 10)).
		Get("/shop/{userId}/order/payment/status/{orderNo}")
	if err != nil {
		return "", err
	}
	return str, nil
}

func (s User) GetOshiMember(userId int64) (entity.ActivitiesOshiMember, error) {
	var activitiesOshiMember entity.ActivitiesOshiMember
	_, err := s.Client.R().
		SetResult(&activitiesOshiMember).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/oshi-member")
	if err != nil {
		return entity.ActivitiesOshiMember{}, err
	}
	return activitiesOshiMember, nil
}

func (s User) GetPlayMusic(userId int64, musicSkuId int64) (entity.PlayMusicInfo, error) {
	var playMusicInfo entity.PlayMusicInfo
	_, err := s.Client.R().
		SetResult(&playMusicInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("musicSkuId", strconv.FormatInt(musicSkuId, 10)).
		Get("/user/{userId}/music/{musicSkuId}")
	if err != nil {
		return entity.PlayMusicInfo{}, err
	}
	return playMusicInfo, nil
}

func (s User) GetRecentlyComment(userId int64, contentId int64, take int, lastCommentId *int64) (entity.UserCommentInfo, error) {
	var userCommentInfo entity.UserCommentInfo
	r := s.Client.R().
		SetResult(&userCommentInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetQueryParam("take", strconv.Itoa(take))
	if lastCommentId != nil {
		r.SetQueryParam("lastCommentId", strconv.FormatInt(*lastCommentId, 10))
	}
	_, err := r.Get("/user/{userId}/comments/content/{contentId}/recently")
	if err != nil {
		return entity.UserCommentInfo{}, err
	}
	return userCommentInfo, nil
}

func (s User) GetRecommendedVideo(userId int64, contentId int64, take int, skip int) ([]entity.TheaterRecommendedPlaybackInfo, error) {
	var theaterRecommendedPlaybackInfo []entity.TheaterRecommendedPlaybackInfo
	_, err := s.Client.R().
		SetResult(&theaterRecommendedPlaybackInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetQueryParam("take", strconv.Itoa(take)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		Get("/user/{userId}/content/{contentId}/recommend")
	if err != nil {
		return nil, err
	}
	return theaterRecommendedPlaybackInfo, nil
}

func (s User) GetReportProblemImageUrl() (entity.DataPathInfo, error) {
	var dataPathInfo entity.DataPathInfo
	_, err := s.Client.R().
		SetResult(&dataPathInfo).
		Post("/report/problem/generatetempurl")
	if err != nil {
		return entity.DataPathInfo{}, err
	}
	return dataPathInfo, nil
}

func (s User) GetReportTopic() ([]entity.UserReportTopicInfo, error) {
	var userReportTopicInfo []entity.UserReportTopicInfo
	_, err := s.Client.R().
		SetResult(&userReportTopicInfo).
		Get("/report/user/topic")
	if err != nil {
		return nil, err
	}
	return userReportTopicInfo, nil
}

func (s User) GetRewardPoints(userId int64) (entity.RewardPointsInfo, error) {
	var rewardPointsInfo entity.RewardPointsInfo
	_, err := s.Client.R().
		SetResult(&rewardPointsInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/reward-points/balance")
	if err != nil {
		return entity.RewardPointsInfo{}, err
	}
	return rewardPointsInfo, nil
}

func (s User) GetSecurityStatus(userId int64) (entity.SecurityStatusResponseInfo, error) {
	var securityStatusResponseInfo entity.SecurityStatusResponseInfo
	_, err := s.Client.R().
		SetResult(&securityStatusResponseInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/security")
	if err != nil {
		return entity.SecurityStatusResponseInfo{}, err
	}
	return securityStatusResponseInfo, nil
}

func (s User) GetShippingFee(userId int64, shopShippingFeeBodyInfo entity.ShopShippingFeeBodyInfo) (entity.ShopShippingFeeResponseInfo, error) {
	var shopShippingFeeResponseInfo entity.ShopShippingFeeResponseInfo
	_, err := s.Client.R().
		SetResult(&shopShippingFeeResponseInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetBody(shopShippingFeeBodyInfo).
		Post("/shop/{userId}/order/delivery-rate")
	if err != nil {
		return entity.ShopShippingFeeResponseInfo{}, err
	}
	return shopShippingFeeResponseInfo, nil
}

func (s User) GetSubscriptionPackage(userId int64) (entity.SubscriptionInfo, error) {
	var subscriptionInfo entity.SubscriptionInfo
	_, err := s.Client.R().
		SetResult(&subscriptionInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/subscription/user/{userId}")
	if err != nil {
		return entity.SubscriptionInfo{}, err
	}
	return subscriptionInfo, nil
}

func (s User) GetSupportedCampaignInfo(userId int64, campaignId int64) (entity.CampaignParticipatedItemInfo, error) {
	var campaignParticipatedItemInfo entity.CampaignParticipatedItemInfo
	_, err := s.Client.R().
		SetResult(&campaignParticipatedItemInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("campaignId", strconv.FormatInt(campaignId, 10)).
		Get("/user/{userId}/campaigns/participated/campaign/{campaignId}")
	if err != nil {
		return entity.CampaignParticipatedItemInfo{}, err
	}
	return campaignParticipatedItemInfo, nil
}

func (s User) GetThankYouContentVideo(contentId int64) (entity.ThankYouContentVideoInfo, error) {
	var thankYouContentVideoInfo entity.ThankYouContentVideoInfo
	_, err := s.Client.R().
		SetResult(&thankYouContentVideoInfo).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Get("/batch-thankyou-video/{contentId}")
	if err != nil {
		return entity.ThankYouContentVideoInfo{}, err
	}
	return thankYouContentVideoInfo, nil
}

func (s User) GetThankYouVideoUrl(contentId int64) (entity.SignedVideoContent, error) {
	var signedVideoContent entity.SignedVideoContent
	_, err := s.Client.R().
		SetResult(&signedVideoContent).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Get("/batch-thankyou-video/{contentId}")
	if err != nil {
		return entity.SignedVideoContent{}, err
	}
	return signedVideoContent, nil
}

func (s User) GetTheaterBox(userId int64, take int, skip int) ([]entity.TheaterBoxInfo, error) {
	var theaterBoxInfo []entity.TheaterBoxInfo
	_, err := s.Client.R().
		SetResult(&theaterBoxInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("take", strconv.Itoa(take)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		Get("/user/{userId}/theater-playback/archive")
	if err != nil {
		return nil, err
	}
	return theaterBoxInfo, nil
}

func (s User) GetTheaterLiveInfo(userId int64, liveId int64) (entity.TheaterLiveInfo, error) {
	var theaterLiveInfo entity.TheaterLiveInfo
	_, err := s.Client.R().
		SetResult(&theaterLiveInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("liveId", strconv.FormatInt(liveId, 10)).
		Get("/user/{userId}/request/theater-live/{liveId}")
	if err != nil {
		return entity.TheaterLiveInfo{}, err
	}
	return theaterLiveInfo, nil
}

func (s User) GetTheaterLiveStatInfo(userId int64, liveId int64) (entity.TheaterStatLiveInfo, error) {
	var theaterStatLiveInfo entity.TheaterStatLiveInfo
	_, err := s.Client.R().
		SetResult(&theaterStatLiveInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("liveId", strconv.FormatInt(liveId, 10)).
		Get("/user/{userId}/request/theater-live/{liveId}")
	if err != nil {
		return entity.TheaterStatLiveInfo{}, err
	}
	return theaterStatLiveInfo, nil
}

func (s User) GetTheaterPlaybackLibrary(userId int64, skip int, take int) ([]entity.TheaterPlaybackLibraryInfo, error) {
	var theaterPlaybackLibraryInfo []entity.TheaterPlaybackLibraryInfo
	_, err := s.Client.R().
		SetResult(&theaterPlaybackLibraryInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/user/{userId}/theater-playback/library/list")
	if err != nil {
		return nil, err
	}
	return theaterPlaybackLibraryInfo, nil
}

func (s User) GetTheaterPlaybackLibraryAmount(userId int64) (entity.TheaterPlaybackLibraryCountInfo, error) {
	var theaterPlaybackLibraryCountInfo entity.TheaterPlaybackLibraryCountInfo
	_, err := s.Client.R().
		SetResult(&theaterPlaybackLibraryCountInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/theater-playback/library")
	if err != nil {
		return entity.TheaterPlaybackLibraryCountInfo{}, err
	}
	return theaterPlaybackLibraryCountInfo, nil
}

func (s User) GetTheaterPlaybackRecommendedVideo(userId int64, contentId int64, skip int, take int) ([]entity.TheaterRecommendedPlaybackInfo, error) {
	var theaterRecommendedPlaybackInfo []entity.TheaterRecommendedPlaybackInfo
	_, err := s.Client.R().
		SetResult(&theaterRecommendedPlaybackInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/user/{userId}/theater-playback/{contentId}/recommend")
	if err != nil {
		return nil, err
	}
	return theaterRecommendedPlaybackInfo, nil
}

func (s User) GetTheaterPlaybackWatch(userId int64, contentId int64) (entity.SignedVideoContent, error) {
	var signedVideoContent entity.SignedVideoContent
	_, err := s.Client.R().
		SetResult(&signedVideoContent).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Get("/user/{userId}/watch/theater-playback/{contentId}")
	if err != nil {
		return entity.SignedVideoContent{}, err
	}
	return signedVideoContent, nil
}

func (s User) GetTheaterSignin(userId int64, theaterId int64) (entity.AutoSigninInfo, error) {
	var autoSigninInfo entity.AutoSigninInfo
	_, err := s.Client.R().
		SetResult(&autoSigninInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("theaterId", strconv.FormatInt(theaterId, 10)).
		Get("/user/{userId}/theaters/{theaterId}/signin")
	if err != nil {
		return entity.AutoSigninInfo{}, err
	}
	return autoSigninInfo, nil
}

func (s User) GetTimelineCampaignResultVideoUrl(contentId int64) (entity.SignedVideoContent, error) {
	var signedVideoContent entity.SignedVideoContent
	_, err := s.Client.R().
		SetResult(&signedVideoContent).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Get("/campaign-result-video/{contentId}")
	if err != nil {
		return entity.SignedVideoContent{}, err
	}
	return signedVideoContent, nil
}

func (s User) GetTimelineMyGiftStats(userId int64, contentId int64) (entity.TimelineMyGiftStats, error) {
	var timelineMyGiftStats entity.TimelineMyGiftStats
	_, err := s.Client.R().
		SetResult(&timelineMyGiftStats).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Get("/user/{userId}/stats/timeline/{contentId}")
	if err != nil {
		return entity.TimelineMyGiftStats{}, err
	}
	return timelineMyGiftStats, nil
}

func (s User) GetTimelineVideoUrl(contentId int64) (entity.SignedVideoContent, error) {
	var signedVideoContent entity.SignedVideoContent
	_, err := s.Client.R().
		SetResult(&signedVideoContent).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Get("/timeline-video/{contentId}")
	if err != nil {
		return entity.SignedVideoContent{}, err
	}
	return signedVideoContent, nil
}

func (s User) GetTopComment(userId int64, contentId int64) (entity.UserCommentInfo, error) {
	var userCommentInfo entity.UserCommentInfo
	_, err := s.Client.R().
		SetResult(&userCommentInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Get("/user/{userId}/comments/content/{contentId}/top100")
	if err != nil {
		return entity.UserCommentInfo{}, err
	}
	return userCommentInfo, nil
}

func (s User) GetUserAddress(userId int64) ([]entity.UserAddressInfo, error) {
	var userAddressInfo []entity.UserAddressInfo
	_, err := s.Client.R().
		SetResult(&userAddressInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/address")
	if err != nil {
		return nil, err
	}
	return userAddressInfo, nil
}

func (s User) GetUserCampaign(userId int64, skip int, take int) ([]entity.CampaignParticipatedItemInfo, error) {
	var campaignParticipatedItemInfo []entity.CampaignParticipatedItemInfo
	_, err := s.Client.R().
		SetResult(&campaignParticipatedItemInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/user/{userId}/campaigns/participated")
	if err != nil {
		return nil, err
	}
	return campaignParticipatedItemInfo, nil
}

func (s User) GetUserDevices(userId int64) ([]entity.DeviceInfo, error) {
	var deviceInfo []entity.DeviceInfo
	_, err := s.Client.R().
		SetResult(&deviceInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/device")
	if err != nil {
		return nil, err
	}
	return deviceInfo, nil
}

func (s User) GetUserGifts(userId int64) ([]entity.UserGiftsInfo, error) {
	var userGiftsInfo []entity.UserGiftsInfo
	_, err := s.Client.R().
		SetResult(&userGiftsInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/gifts")
	if err != nil {
		return nil, err
	}
	return userGiftsInfo, nil
}

func (s User) GetUserGiftsQuota(userId int64, skip int, take int) ([]entity.UserGiftQuotaInfo, error) {
	var userGiftQuotaInfo []entity.UserGiftQuotaInfo
	_, err := s.Client.R().
		SetResult(&userGiftQuotaInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/user/{userId}/gifts/quota")
	if err != nil {
		return nil, err
	}
	return userGiftQuotaInfo, nil
}

func (s User) GetUserGiftsQuotaById(userId int64, quotaId int64) (entity.UserGiftQuotaByIdInfo, error) {
	var userGiftQuotaByIdInfo entity.UserGiftQuotaByIdInfo
	_, err := s.Client.R().
		SetResult(&userGiftQuotaByIdInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("quotaId", strconv.FormatInt(quotaId, 10)).
		Get("/user/{userId}/gifts/quota/{quotaId}")
	if err != nil {
		return entity.UserGiftQuotaByIdInfo{}, err
	}
	return userGiftQuotaByIdInfo, nil
}

func (s User) GetUserLikedMiniVideos(userId int64, skip int, take int) ([]entity.LiveVideoData, error) {
	var liveVideoData []entity.LiveVideoData
	_, err := s.Client.R().
		SetResult(&liveVideoData).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/user/{userId}/likes/mini-videos")
	if err != nil {
		return nil, err
	}
	return liveVideoData, nil
}

func (s User) GetUserLinkFacebook(userId int64, facebookAppId string) (entity.FacebookLinkResponseInfo, error) {
	var facebookLinkResponseInfo entity.FacebookLinkResponseInfo
	_, err := s.Client.R().
		SetResult(&facebookLinkResponseInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("facebookAppId", facebookAppId).
		Get("/user/{userId}/link/{facebookAppId}")
	if err != nil {
		return entity.FacebookLinkResponseInfo{}, err
	}
	return facebookLinkResponseInfo, nil
}

func (s User) GetUserMusic(userId int64) ([]entity.MusicAlbumInfo, error) {
	var musicAlbumInfo []entity.MusicAlbumInfo
	_, err := s.Client.R().
		SetResult(&musicAlbumInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/music")
	if err != nil {
		return nil, err
	}
	return musicAlbumInfo, nil
}

func (s User) GetUserProfile(userId int64) (entity.UserProfileInfo, error) {
	var userProfileInfo entity.UserProfileInfo
	_, err := s.Client.R().
		SetResult(&userProfileInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/profile")
	if err != nil {
		return entity.UserProfileInfo{}, err
	}
	return userProfileInfo, nil
}

func (s User) GetUserRank(userId int64) (entity.CoreStats, error) {
	var coreStats entity.CoreStats
	_, err := s.Client.R().
		SetResult(&coreStats).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/stats/coin")
	if err != nil {
		return entity.CoreStats{}, err
	}
	return coreStats, nil
}

func (s User) GetUserTransaction(userId int64, skip int, take int) ([]entity.UserTransactionInfo, error) {
	var userTransactionInfo []entity.UserTransactionInfo
	_, err := s.Client.R().
		SetResult(&userTransactionInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/user/{userId}/transactions")
	if err != nil {
		return nil, err
	}
	return userTransactionInfo, nil
}

func (s User) GetVideoCollectionList(userId int64, skip int, take int) ([]entity.TheaterPlaybackLibraryInfo, error) {
	var theaterPlaybackLibraryInfo []entity.TheaterPlaybackLibraryInfo
	_, err := s.Client.R().
		SetResult(&theaterPlaybackLibraryInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/user/{userId}/content/purchased/list")
	if err != nil {
		return nil, err
	}
	return theaterPlaybackLibraryInfo, nil
}

func (s User) GetVideoLibraryAmount(userId int64) (entity.TheaterPlaybackLibraryCountInfo, error) {
	var theaterPlaybackLibraryCountInfo entity.TheaterPlaybackLibraryCountInfo
	_, err := s.Client.R().
		SetResult(&theaterPlaybackLibraryCountInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/content/purchased")
	if err != nil {
		return entity.TheaterPlaybackLibraryCountInfo{}, err
	}
	return theaterPlaybackLibraryCountInfo, nil
}

func (s User) Initial2FASetup(userId int64) (entity.Initial2FASetupResponseInfo, error) {
	var initial2FASetupResponseInfo entity.Initial2FASetupResponseInfo
	_, err := s.Client.R().
		SetResult(&initial2FASetupResponseInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Post("/user/{userId}/security/2fa/setup/initial")
	if err != nil {
		return entity.Initial2FASetupResponseInfo{}, err
	}
	return initial2FASetupResponseInfo, nil
}

func (s User) IsLikeMemberLive(userId int64, memberLiveId int64) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberLiveId", strconv.FormatInt(memberLiveId, 10)).
		Get("/user/{userId}/likes/member-live/{memberLiveId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) IsLikeMiniVideo(userId int64, videoId int64) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("videoId", strconv.FormatInt(videoId, 10)).
		Get("/user/{userId}/likes/mini-video/{videoId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) IsLikeTimeline(userId int64, memberTimelineId int64) (entity.LikeTimelineInfo, error) {
	var likeTimelineInfo entity.LikeTimelineInfo
	_, err := s.Client.R().
		SetResult(&likeTimelineInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberTimelineId", strconv.FormatInt(memberTimelineId, 10)).
		Get("/user/{userId}/likes/like-status/{memberTimelineId}")
	if err != nil {
		return entity.LikeTimelineInfo{}, err
	}
	return likeTimelineInfo, nil
}

func (s User) IsOshiMember(userId int64, oshiMember entity.OshiMember) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetBody(oshiMember).
		Put("/user/{userId}/oshi-member")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) IsSubscribeMemberLive(userId int64) (bool, error) {
	var isSubscribeMemberLive bool
	_, err := s.Client.R().
		SetResult(&isSubscribeMemberLive).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/push-notification/subscribe/member-live-all/status")
	if err != nil {
		return false, err
	}
	return isSubscribeMemberLive, nil
}

func (s User) IsSubscribeNotification(userId int64, memberId int64) (bool, error) {
	var isSubscribeNotification bool
	_, err := s.Client.R().
		SetResult(&isSubscribeNotification).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Get("/user/{userId}/push-notification/subscribe/member/{memberId}/status")
	if err != nil {
		return false, err
	}
	return isSubscribeNotification, nil
}

func (s User) IsSubscribeTheaterLiveNotification(userId int64) (bool, error) {
	var isSubscribeTheaterLiveNotification bool
	_, err := s.Client.R().
		SetResult(&isSubscribeTheaterLiveNotification).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/push-notification/subscribe/theater-live/status")
	if err != nil {
		return false, err
	}
	return isSubscribeTheaterLiveNotification, nil
}

func (s User) LikeMemberLive(userId int64, memberLiveId int64) (bool, error) {
	var likeMemberLive bool
	_, err := s.Client.R().
		SetResult(&likeMemberLive).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberLiveId", strconv.FormatInt(memberLiveId, 10)).
		Put("/user/{userId}/likes/member-live/{memberLiveId}")
	if err != nil {
		return false, err
	}
	return likeMemberLive, nil
}

func (s User) LikeMiniVideo(userId int64, videoId int64) (bool, error) {
	var likeMiniVideo bool
	_, err := s.Client.R().
		SetResult(&likeMiniVideo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("videoId", strconv.FormatInt(videoId, 10)).
		Put("/user/{userId}/likes/mini-video/{videoId}")
	if err != nil {
		return false, err
	}
	return likeMiniVideo, nil
}

func (s User) LikeTimeline(userId int64, memberTimelineId int64) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberTimelineId", strconv.FormatInt(memberTimelineId, 10)).
		Put("/user/{userId}/likes/member-timeline/{memberTimelineId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) LoginFacebook(loginFacebookBody entity.LoginFacebookBody) (entity.AuthenResponseInfo, error) {
	var loginFacebook entity.AuthenResponseInfo
	_, err := s.Client.R().
		SetResult(&loginFacebook).
		SetBody(loginFacebookBody).
		Post("/auth/facebook")
	if err != nil {
		return entity.AuthenResponseInfo{}, err
	}
	return loginFacebook, nil
}

func (s User) LoginWithEmail(loginInfo entity.LoginInfo) (entity.AuthenResponseInfo, error) {
	var loginWithEmail entity.AuthenResponseInfo
	_, err := s.Client.R().
		SetResult(&loginWithEmail).
		SetBody(loginInfo).
		Post("/auth/email")
	if err != nil {
		return entity.AuthenResponseInfo{}, err
	}
	return loginWithEmail, nil
}

func (s User) LogoutDevice(userId int64, logoutDeviceBody entity.LogoutDeviceBody) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetBody(logoutDeviceBody).
		Delete("/user/{userId}/device")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) OnCheckMediaCanWatch(userId int64, contentId int64) (entity.SignedVideoContent, error) {
	var onCheckMediaCanWatch entity.SignedVideoContent
	_, err := s.Client.R().
		SetResult(&onCheckMediaCanWatch).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Get("/user/{userId}/watch/media/{contentId}")
	if err != nil {
		return entity.SignedVideoContent{}, err
	}
	return onCheckMediaCanWatch, nil
}

func (s User) OnPurchaseEventTheaterLive(userId int64, liveId int64) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("liveId", strconv.FormatInt(liveId, 10)).
		Post("/user/{userId}/purchase/theater-live/{liveId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) OnPurchaseMedia(userId int64, contentId int64) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Post("/user/{userId}/purchase/media/{contentId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) OnShopAutoSignIn(userId int64) (entity.AutoSigninInfo, error) {
	var onShopAutoSignIn entity.AutoSigninInfo
	_, err := s.Client.R().
		SetResult(&onShopAutoSignIn).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/shop/signin")
	if err != nil {
		return entity.AutoSigninInfo{}, err
	}
	return onShopAutoSignIn, nil
}

func (s User) OnSubscribeTheaterLiveNotification(userId int64) (bool, error) {
	var onSubscribeTheaterLiveNotification bool
	_, err := s.Client.R().
		SetResult(&onSubscribeTheaterLiveNotification).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Post("/user/{userId}/push-notification/subscribe/theater-live")
	if err != nil {
		return false, err
	}
	return onSubscribeTheaterLiveNotification, nil
}

func (s User) OnUnSubscribeTheaterLiveNotification(userId int64) (bool, error) {
	var onUnSubscribeTheaterLiveNotification bool
	_, err := s.Client.R().
		SetResult(&onUnSubscribeTheaterLiveNotification).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Delete("/user/{userId}/push-notification/subscribe/theater-live")
	if err != nil {
		return false, err
	}
	return onUnSubscribeTheaterLiveNotification, nil
}

func (s User) PostGeneratedCoverUrl(userId int64) (entity.DataPathInfo, error) {
	var postGeneratedCoverUrl entity.DataPathInfo
	_, err := s.Client.R().
		SetResult(&postGeneratedCoverUrl).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Post("/user/{userId}/profile/cover/generateTempUrl")
	if err != nil {
		return entity.DataPathInfo{}, err
	}
	return postGeneratedCoverUrl, nil
}

func (s User) PostGeneratedProfileUrl(userId int64) (entity.DataPathInfo, error) {
	var postGeneratedProfileUrl entity.DataPathInfo
	_, err := s.Client.R().
		SetResult(&postGeneratedProfileUrl).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Post("/user/{userId}/profile/avatar/generateTempUrl")
	if err != nil {
		return entity.DataPathInfo{}, err
	}
	return postGeneratedProfileUrl, nil
}

func (s User) PostGlobalRedeemCode(userId int64, str string) (entity.GlobalRedeemCodeInfo, error) {
	var postGlobalRedeemCode entity.GlobalRedeemCodeInfo
	_, err := s.Client.R().
		SetResult(&postGlobalRedeemCode).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("code", str).
		Post("/user/{userId}/redeemcode")
	if err != nil {
		return entity.GlobalRedeemCodeInfo{}, err
	}
	return postGlobalRedeemCode, nil
}

func (s User) PostTheaterLivePurchase(userId int64, liveId int64) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("liveId", strconv.FormatInt(liveId, 10)).
		Post("/user/{userId}/purchase/theater-live/{liveId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) PostTheaterPlaybackPurchase(userId int64, contentId int64) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Post("/user/{userId}/purchase/theater-playback/{contentId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) PostUserDevice(userId int64, userDeviceInfo entity.UserDeviceInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetBody(userDeviceInfo).
		Post("/user/{userId}/device")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) PostUserLinkFacebook(userId int64, facebookLinkPostInfo entity.FacebookLinkPostInfo) (entity.FacebookLinkResponseInfo, error) {
	var postUserLinkFacebook entity.FacebookLinkResponseInfo
	_, err := s.Client.R().
		SetResult(&postUserLinkFacebook).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetBody(facebookLinkPostInfo).
		Post("/user/{userId}/link/facebook")
	if err != nil {
		return entity.FacebookLinkResponseInfo{}, err
	}
	return postUserLinkFacebook, nil
}

func (s User) PutAllNotification(userId int64) (bool, error) {
	var putAllNotification bool
	_, err := s.Client.R().
		SetResult(&putAllNotification).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Put("/user/{userId}/notification/all")
	if err != nil {
		return false, err
	}
	return putAllNotification, nil
}

func (s User) PutFollow(userId int64, memberId int64) (bool, error) {
	var putFollow bool
	_, err := s.Client.R().
		SetResult(&putFollow).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Put("/user/{userId}/follows/member/{memberId}")
	if err != nil {
		return false, err
	}
	return putFollow, nil
}

func (s User) PutOrderInvoice(userId int64, str string, str2 string) (entity.ShopCreateOrderResponseInfo, error) {
	var putOrderInvoice entity.ShopCreateOrderResponseInfo
	_, err := s.Client.R().
		SetResult(&putOrderInvoice).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("RefCode", str).
		SetQueryParam("InvNo", str2).
		Put("/shop/{userId}/order/invoice")
	if err != nil {
		return entity.ShopCreateOrderResponseInfo{}, err
	}
	return putOrderInvoice, nil
}

func (s User) PutReadNotification(userId int64, notificationId int64) (entity.NotificationInfo, error) {
	var putReadNotification entity.NotificationInfo
	_, err := s.Client.R().
		SetResult(&putReadNotification).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("notificationId", strconv.FormatInt(notificationId, 10)).
		Put("/user/{userId}/notification/{notificationId}/read")
	if err != nil {
		return entity.NotificationInfo{}, err
	}
	return putReadNotification, nil
}

func (s User) PutUnfollow(userId int64, memberId int64) (bool, error) {
	var putUnfollow bool
	_, err := s.Client.R().
		SetResult(&putUnfollow).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Delete("/user/{userId}/follows/member/{memberId}")
	if err != nil {
		return false, err
	}
	return putUnfollow, nil
}

func (s User) RedeemCode(userId int64, str string) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("code", str).
		Post("/user/{userId}/redeem")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) RefreshToken(refreshTokenBody entity.RefreshTokenBody) (entity.RefreshTokenInfo, error) {
	var refreshToken entity.RefreshTokenInfo
	_, err := s.Client.R().
		SetResult(&refreshToken).
		SetBody(refreshTokenBody).
		Post("/auth/refresh")
	if err != nil {
		return entity.RefreshTokenInfo{}, err
	}
	return refreshToken, nil
}

func (s User) Register(registerInfo entity.RegisterInfo) (entity.RegisResponse, error) {
	var register entity.RegisResponse
	_, err := s.Client.R().
		SetResult(&register).
		SetBody(registerInfo).
		Post("/accounts/register")
	if err != nil {
		return entity.RegisResponse{}, err
	}
	return register, nil
}

func (s User) RegisterNotification(userId int64, notificationRegisterBody entity.NotificationRegisterBody) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetBody(notificationRegisterBody).
		Post("/user/{userId}/push-notification/register")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) ReportComment(reportUserInfo entity.ReportUserInfo) (bool, error) {
	var reportComment bool
	_, err := s.Client.R().
		SetResult(&reportComment).
		SetBody(reportUserInfo).
		Post("/report/comment")
	if err != nil {
		return false, err
	}
	return reportComment, nil
}

func (s User) ReportProblem(reportProblemInfo entity.ReportProblemInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetBody(reportProblemInfo).
		Post("/report/problem")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) ReportUser(reportUserInfo entity.ReportUserInfo) (bool, error) {
	var reportUser bool
	_, err := s.Client.R().
		SetResult(&reportUser).
		SetBody(reportUserInfo).
		Post("/report/user")
	if err != nil {
		return false, err
	}
	return reportUser, nil
}

func (s User) RequestPaymentVendor(str string, shopPaymentRequestBodyInfo entity.ShopPaymentRequestBodyInfo) (entity.CoreShopPaymentInfo, error) {
	var requestPaymentVendor entity.CoreShopPaymentInfo
	_, err := s.Client.R().
		SetResult(&requestPaymentVendor).
		SetPathParam("fullUrl", str).
		SetBody(shopPaymentRequestBodyInfo).
		Post("{fullUrl}")
	if err != nil {
		return entity.CoreShopPaymentInfo{}, err
	}
	return requestPaymentVendor, nil
}

func (s User) SaveMediaPlayTime(userId int64, contentId int64, i2 int) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetQueryParam("latestSeconds", strconv.Itoa(i2)).
		Post("/user/{userId}/media/{contentId}/play")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) SavePlaybackPlayTime(userId int64, contentId int64, i2 int) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetQueryParam("latestSeconds", strconv.Itoa(i2)).
		Post("/user/{userId}/theater-playback/{contentId}/play")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) SendEmailOTP2FASetup(userId int64) (entity.SendEmailOTPResponseInfo, error) {
	var sendEmailOTP2FASetup entity.SendEmailOTPResponseInfo
	_, err := s.Client.R().
		SetResult(&sendEmailOTP2FASetup).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Post("/user/{userId}/security/2fa/setup/email-otp/send")
	if err != nil {
		return entity.SendEmailOTPResponseInfo{}, err
	}
	return sendEmailOTP2FASetup, nil
}

func (s User) SendEmailOTPAccountDeletion(userId int64, str string) (entity.SendEmailOTPResponseInfo, error) {
	var sendEmailOTPAccountDeletion entity.SendEmailOTPResponseInfo
	_, err := s.Client.R().
		SetResult(&sendEmailOTPAccountDeletion).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("ref", str).
		Post("/user/{userId}/deletion/email-otp/send")
	if err != nil {
		return entity.SendEmailOTPResponseInfo{}, err
	}
	return sendEmailOTPAccountDeletion, nil
}

func (s User) SendFirstOTPEmail2FAReset(userId int64) (entity.SendOTPEmailResponseInfo, error) {
	var sendFirstOTPEmail2FAReset entity.SendOTPEmailResponseInfo
	_, err := s.Client.R().
		SetResult(&sendFirstOTPEmail2FAReset).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Post("/user/{userId}/security/2fa/reset/email-otp-1st/send")
	if err != nil {
		return entity.SendOTPEmailResponseInfo{}, err
	}
	return sendFirstOTPEmail2FAReset, nil
}

func (s User) SendGift(userId int64, memberId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		SetBody(sendGiftInfo).
		Post("/user/{userId}/gifts/member/{memberId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) SendGiftCampaignResult(userId int64, contentId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetBody(sendGiftInfo).
		Post("/user/{userId}/gifts/campaign-result/{contentId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) SendGiftComment(userId int64, commentId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("commentId", strconv.FormatInt(commentId, 10)).
		SetBody(sendGiftInfo).
		Post("/user/{userId}/gifts/comment/{commentId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) SendGiftDigitalStudio(userId int64, contentId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetBody(sendGiftInfo).
		Post("/user/{userId}/gifts/digital-studio/{contentId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) SendGiftMeetingLobbyRoom(userId int64, str string, sendGiftInfo entity.SendGiftInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("refCode", str).
		SetBody(sendGiftInfo).
		Post("/user/{userId}/gifts/meeting-you/{refCode}/lobby")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) SendGiftMeetingThankYou(userId int64, str string, sendGiftInfo entity.SendGiftInfo) (entity.SendMeetingGiftThankYouResponse, error) {
	var sendMeetingGiftThankYouResponse entity.SendMeetingGiftThankYouResponse
	_, err := s.Client.R().
		SetResult(&sendMeetingGiftThankYouResponse).
		SetBody(sendGiftInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("refCode", str).
		Post("/user/{userId}/gifts/meeting-you/{refCode}/thankyou")
	if err != nil {
		return entity.SendMeetingGiftThankYouResponse{}, err
	}
	return sendMeetingGiftThankYouResponse, nil
}

func (s User) SendGiftMemberLive(userId int64, contentId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetBody(sendGiftInfo).
		Post("/user/{userId}/gifts/member-live/{contentId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) SendGiftMiniVideo(userId int64, contentId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetBody(sendGiftInfo).
		Post("/user/{userId}/gifts/mini-video/{contentId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) SendGiftThankyou(userId int64, memberId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		SetBody(sendGiftInfo).
		Post("/user/{userId}/gifts/batch-ranking/{memberId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) SendGiftTheater(userId int64, contentId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetBody(sendGiftInfo).
		Post("/user/{userId}/gifts/theater/{contentId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) SendGiftTimeline(userId int64, contentId int64, sendGiftInfo entity.SendGiftInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetBody(sendGiftInfo).
		Post("/user/{userId}/gifts/timeline/{contentId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) SendNotification(sendNotificationInfo entity.SendNotificationInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetBody(sendNotificationInfo).
		Post("/push-notification/message/send")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) SendSecondOTPEmail2FAReset(userId int64) (entity.SendOTPEmailResponseInfo, error) {
	var sendOTPEmailResponseInfo entity.SendOTPEmailResponseInfo
	_, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetResult(&sendOTPEmailResponseInfo).
		Post("/user/{userId}/security/2fa/reset/email-otp-2nd/send")
	if err != nil {
		return entity.SendOTPEmailResponseInfo{}, err
	}
	return sendOTPEmailResponseInfo, nil
}

func (s User) SendUserComment(userId int64, contentId int64, userCommentBodyInfo entity.UserCommentBodyInfo) (entity.UserCommentResponseInfo, error) {
	var userCommentResponseInfo entity.UserCommentResponseInfo
	_, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetBody(userCommentBodyInfo).
		SetResult(&userCommentResponseInfo).
		Put("/user/{userId}/comments/content/{contentId}")
	if err != nil {
		return entity.UserCommentResponseInfo{}, err
	}
	return userCommentResponseInfo, nil
}

func (s User) SubScribeMeberLive(userId int64) (bool, error) {
	var subScribeMeberLive bool
	_, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetResult(&subScribeMeberLive).
		Post("/user/{userId}/push-notification/subscribe/member-live-all")
	if err != nil {
		return false, err
	}
	return subScribeMeberLive, nil
}

func (s User) SubscribeNotification(userId int64, memberId int64) (bool, error) {
	var subscribeNotification bool
	_, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		SetResult(&subscribeNotification).
		Post("/user/{userId}/push-notification/subscribe/member/{memberId}")
	if err != nil {
		return false, err
	}
	return subscribeNotification, nil
}

func (s User) UnLikeMemberLive(userId int64, memberLiveId int64) (bool, error) {
	var unLikeMemberLive bool
	_, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberLiveId", strconv.FormatInt(memberLiveId, 10)).
		SetResult(&unLikeMemberLive).
		Delete("/user/{userId}/likes/member-live/{memberLiveId}")
	if err != nil {
		return false, err
	}
	return unLikeMemberLive, nil
}

func (s User) UnLikeMiniVideo(userId int64, videoId int64) (bool, error) {
	var unLikeMiniVideo bool
	_, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("videoId", strconv.FormatInt(videoId, 10)).
		SetResult(&unLikeMiniVideo).
		Delete("/user/{userId}/likes/mini-video/{videoId}")
	if err != nil {
		return false, err
	}
	return unLikeMiniVideo, nil
}

func (s User) UnOshiMember(userId int64) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Delete("/user/{userId}/oshi-member")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) UnSubscribeMemberLive(userId int64) (bool, error) {
	var unSubscribeMemberLive bool
	_, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetResult(&unSubscribeMemberLive).
		Delete("/user/{userId}/push-notification/subscribe/member-live-all")
	if err != nil {
		return false, err
	}
	return unSubscribeMemberLive, nil
}

func (s User) UnSubscribeNotification(userId int64, memberId int64) (bool, error) {
	var unSubscribeNotification bool
	_, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		SetResult(&unSubscribeNotification).
		Delete("/user/{userId}/push-notification/subscribe/member/{memberId}")
	if err != nil {
		return false, err
	}
	return unSubscribeNotification, nil
}

func (s User) UnlikeTimeline(userId int64, memberTimelineId int64) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberTimelineId", strconv.FormatInt(memberTimelineId, 10)).
		Delete("/user/{userId}/likes/member-timeline/{memberTimelineId}")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) UnregisterNotification(notificationRegisterBody entity.NotificationRegisterBody) (resty.Response, error) {
	res, err := s.Client.R().
		SetBody(notificationRegisterBody).
		Delete("/push-notification/register")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) VerifyEkyc(userId int64, ekycVerifyUserBodyInfo entity.EkycVerifyUserBodyInfo) (entity.EkycVerifyUserResponseInfo, error) {
	var ekycVerifyUserResponseInfo entity.EkycVerifyUserResponseInfo
	_, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetBody(ekycVerifyUserBodyInfo).
		SetResult(&ekycVerifyUserResponseInfo).
		Post("/ekyc/{userId}/verify")
	if err != nil {
		return entity.EkycVerifyUserResponseInfo{}, err
	}
	return ekycVerifyUserResponseInfo, nil
}

func (s User) VerifyFirstOTPEmail2FAReset(userId int64, verifyFirstOTPEmailBodyInfo entity.VerifyFirstOTPEmailBodyInfo) (entity.VerifyFirstOTPEmailResponseInfo, error) {
	var verifyFirstOTPEmailResponseInfo entity.VerifyFirstOTPEmailResponseInfo
	_, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetBody(verifyFirstOTPEmailBodyInfo).
		SetResult(&verifyFirstOTPEmailResponseInfo).
		Post("/user/{userId}/security/2fa/reset/email-otp-1st/verify")
	if err != nil {
		return entity.VerifyFirstOTPEmailResponseInfo{}, err
	}
	return verifyFirstOTPEmailResponseInfo, nil
}

func (s User) VerifyFounder(userId int64) (resty.Response, error) {
	res, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/verify/founder")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s User) VoteSenbatsu(userId int64, memberId int64, str string) (entity.VoteSenbatsuResponseInfo, error) {
	var voteSenbatsuResponseInfo entity.VoteSenbatsuResponseInfo
	_, err := s.Client.R().
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		SetQueryParam("voteCode", str).
		SetResult(&voteSenbatsuResponseInfo).
		Post("/user/{userId}/senbatsu-votes/member/{memberId}")
	if err != nil {
		return entity.VoteSenbatsuResponseInfo{}, err
	}
	return voteSenbatsuResponseInfo, nil
}

type UserConfiguration func(*User) error

func WithClient(client *resty.Client) UserConfiguration {
	return func(e *User) error {
		e.Client = client
		e.Client.SetBaseURL("https://user.bnk48.io")
		return nil
	}
}

func New(cfgs ...UserConfiguration) (User, error) {
	service := User{
		Client: client.NewClient().SetBaseURL("https://user.bnk48.io"),
	}
	for _, cfg := range cfgs {
		if err := cfg(&service); err != nil {
			return User{}, err
		}
	}
	return service, nil
}

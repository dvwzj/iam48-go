package ticket

import (
	"strconv"

	"github.com/dvwzj/iam48-go/api/entity"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type Ticket struct {
	Client *resty.Client
}

func (s Ticket) AccessTokenX(tokenXAccessRequestBodyInfo entity.TokenXAccessRequestBodyInfo) (entity.TokenXAccessResponseBodyInfo, error) {
	var tokenXAccessResponseBodyInfo entity.TokenXAccessResponseBodyInfo
	_, err := s.Client.R().
		SetBody(tokenXAccessRequestBodyInfo).
		SetResult(&tokenXAccessResponseBodyInfo).
		Post("/token-x/access")
	if err != nil {
		return entity.TokenXAccessResponseBodyInfo{}, err
	}
	return tokenXAccessResponseBodyInfo, nil
}

func (s Ticket) AccessTokenXFast() (entity.TokenXAccessResponseBodyInfo, error) {
	var tokenXAccessResponseBodyInfo entity.TokenXAccessResponseBodyInfo
	_, err := s.Client.R().
		SetResult(&tokenXAccessResponseBodyInfo).
		Post("/token-x/access/fast")
	if err != nil {
		return entity.TokenXAccessResponseBodyInfo{}, err
	}
	return tokenXAccessResponseBodyInfo, nil
}

func (s Ticket) CheckWalletHasPin() (resty.Response, error) {
	res, err := s.Client.R().
		Get("/pin")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s Ticket) DeleteWallet() (resty.Response, error) {
	res, err := s.Client.R().
		Delete("/wallet")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s Ticket) Exchange(exchangeDataPostBodyInfo entity.ExchangeDataPostBodyInfo) (resty.Response, error) {
	res, err := s.Client.R().
		SetBody(exchangeDataPostBodyInfo).
		Post("/exchange")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s Ticket) ExchangeStats() (entity.ExchangeStatsResponseInfo, error) {
	var exchangeStatsResponseInfo entity.ExchangeStatsResponseInfo
	_, err := s.Client.R().
		SetResult(&exchangeStatsResponseInfo).
		Get("/exchange/stats")
	if err != nil {
		return entity.ExchangeStatsResponseInfo{}, err
	}
	return exchangeStatsResponseInfo, nil
}

func (s Ticket) ExchangeTicket(fromTicketSkuId string, toTicketSkuId string) (entity.ExchangeTicketToTicketResponseInfo, error) {
	var exchangeTicketToTicketResponseInfo entity.ExchangeTicketToTicketResponseInfo
	_, err := s.Client.R().
		SetResult(&exchangeTicketToTicketResponseInfo).
		Get("/exchange/ticket/" + fromTicketSkuId + "/to/" + toTicketSkuId)
	if err != nil {
		return entity.ExchangeTicketToTicketResponseInfo{}, err
	}
	return exchangeTicketToTicketResponseInfo, nil
}

func (s Ticket) ExchangeTicketAvailable() ([]entity.ExchangeTicketAvailableResponseInfo, error) {
	var exchangeTicketAvailableResponseInfo []entity.ExchangeTicketAvailableResponseInfo
	_, err := s.Client.R().
		SetResult(&exchangeTicketAvailableResponseInfo).
		Get("/exchange/ticket/available")
	if err != nil {
		return nil, err
	}
	return exchangeTicketAvailableResponseInfo, nil
}

func (s Ticket) ExchangeTicketListById(ticketSkuId string) ([]entity.ExchangeTicketResponseInfo, error) {
	var exchangeTicketResponseInfo []entity.ExchangeTicketResponseInfo
	_, err := s.Client.R().
		SetResult(&exchangeTicketResponseInfo).
		SetPathParam("ticketSkuId", ticketSkuId).
		Get("/exchange/ticket/{ticketSkuId}/list")
	if err != nil {
		return nil, err
	}
	return exchangeTicketResponseInfo, nil
}

func (s Ticket) GetMyWalletInfo() (entity.WalletInfoModel, error) {
	var walletInfoModel entity.WalletInfoModel
	_, err := s.Client.R().
		SetResult(&walletInfoModel).
		Get("/wallet")
	if err != nil {
		return entity.WalletInfoModel{}, err
	}
	return walletInfoModel, nil
}

func (s Ticket) GetTermsTokenX() (entity.TokenXTermsResponseInfo, error) {
	var tokenXTermsResponseInfo entity.TokenXTermsResponseInfo
	_, err := s.Client.R().
		SetResult(&tokenXTermsResponseInfo).
		Get("/token-x/terms")
	if err != nil {
		return entity.TokenXTermsResponseInfo{}, err
	}
	return tokenXTermsResponseInfo, nil
}

func (s Ticket) GetTicketDetail(ticketSkuId int64) (entity.TicketDetailModel, error) {
	var ticketDetailModel entity.TicketDetailModel
	_, err := s.Client.R().
		SetResult(&ticketDetailModel).
		SetPathParam("ticketSkuId", strconv.FormatInt(ticketSkuId, 10)).
		Get("/inventory/{ticketSkuId}")
	if err != nil {
		return entity.TicketDetailModel{}, err
	}
	return ticketDetailModel, nil
}

func (s Ticket) GetTicketHistory(skip int64, take int64) ([]entity.TicketHistoryModel, error) {
	var ticketHistoryModel []entity.TicketHistoryModel
	_, err := s.Client.R().
		SetResult(&ticketHistoryModel).
		SetQueryParam("skip", strconv.FormatInt(skip, 10)).
		SetQueryParam("take", strconv.FormatInt(take, 10)).
		Get("/history")
	if err != nil {
		return nil, err
	}
	return ticketHistoryModel, nil
}

func (s Ticket) GetTicketInventory(skip int64, take int64) ([]entity.TicketInventoryModel, error) {
	var ticketInventoryModel []entity.TicketInventoryModel
	_, err := s.Client.R().
		SetResult(&ticketInventoryModel).
		SetQueryParam("skip", strconv.FormatInt(skip, 10)).
		SetQueryParam("take", strconv.FormatInt(take, 10)).
		Get("/inventory")
	if err != nil {
		return nil, err
	}
	return ticketInventoryModel, nil
}

func (s Ticket) GetWalletByWalletCode(walletCode string) (entity.WalletInfoModel, error) {
	var walletInfoModel entity.WalletInfoModel
	_, err := s.Client.R().
		SetResult(&walletInfoModel).
		SetPathParam("walletCode", walletCode).
		Get("/wallet/{walletCode}")
	if err != nil {
		return entity.WalletInfoModel{}, err
	}
	return walletInfoModel, nil
}

func (s Ticket) GetWalletQRCode() (entity.WalletQRCode, error) {
	var walletQRCode entity.WalletQRCode
	_, err := s.Client.R().
		SetResult(&walletQRCode).
		Get("/qr/wallet")
	if err != nil {
		return entity.WalletQRCode{}, err
	}
	return walletQRCode, nil
}

func (s Ticket) GreetingMyQuota(take int, skip int64) ([]entity.GreetingMyInfo, error) {
	var greetingMyInfo []entity.GreetingMyInfo
	_, err := s.Client.R().
		SetResult(&greetingMyInfo).
		SetQueryParam("take", strconv.Itoa(take)).
		SetQueryParam("skip", strconv.FormatInt(skip, 10)).
		Get("/greeting/my/quota")
	if err != nil {
		return nil, err
	}
	return greetingMyInfo, nil
}

func (s Ticket) GreetingMyReplied(take int, skip int64) ([]entity.GreetingMyInfo, error) {
	var greetingMyInfo []entity.GreetingMyInfo
	_, err := s.Client.R().
		SetResult(&greetingMyInfo).
		SetQueryParam("take", strconv.Itoa(take)).
		SetQueryParam("skip", strconv.FormatInt(skip, 10)).
		Get("/greeting/my/replied")
	if err != nil {
		return nil, err
	}
	return greetingMyInfo, nil
}

func (s Ticket) GreetingMySent(take int, skip int64) ([]entity.GreetingMyInfo, error) {
	var greetingMyInfo []entity.GreetingMyInfo
	_, err := s.Client.R().
		SetResult(&greetingMyInfo).
		SetQueryParam("take", strconv.Itoa(take)).
		SetQueryParam("skip", strconv.FormatInt(skip, 10)).
		Get("/greeting/my/sent")
	if err != nil {
		return nil, err
	}
	return greetingMyInfo, nil
}

func (s Ticket) GreetingMyStats() (entity.GreetingMyStatsInfo, error) {
	var greetingMyStatsInfo entity.GreetingMyStatsInfo
	_, err := s.Client.R().
		SetResult(&greetingMyStatsInfo).
		Get("/greeting/my/stats")
	if err != nil {
		return entity.GreetingMyStatsInfo{}, err
	}
	return greetingMyStatsInfo, nil
}

func (s Ticket) GreetingRedeem(greetingRedeemBodyInfo entity.GreetingRedeemBodyInfo) (entity.GreetingRedeemResponseInfo, error) {
	var greetingRedeemResponseInfo entity.GreetingRedeemResponseInfo
	_, err := s.Client.R().
		SetResult(&greetingRedeemResponseInfo).
		SetBody(greetingRedeemBodyInfo).
		Post("/greeting/redeem")
	if err != nil {
		return entity.GreetingRedeemResponseInfo{}, err
	}
	return greetingRedeemResponseInfo, nil
}

func (s Ticket) GreetingRedeemDialog(redeemId int64, greetingRedeemDialogBodyInfo entity.GreetingRedeemDialogBodyInfo) (entity.GreetingMyInfo, error) {
	var greetingMyInfo entity.GreetingMyInfo
	_, err := s.Client.R().
		SetResult(&greetingMyInfo).
		SetBody(greetingRedeemDialogBodyInfo).
		SetPathParam("redeemId", strconv.FormatInt(redeemId, 10)).
		Post("/greeting/redeem/{redeemId}/dialog")
	if err != nil {
		return entity.GreetingMyInfo{}, err
	}
	return greetingMyInfo, nil
}

func (s Ticket) GreetingRedeemInfo(redeemId int64) (entity.GreetingMyInfo, error) {
	var greetingMyInfo entity.GreetingMyInfo
	_, err := s.Client.R().
		SetResult(&greetingMyInfo).
		SetPathParam("redeemId", strconv.FormatInt(redeemId, 10)).
		Get("/greeting/redeem/{redeemId}/info")
	if err != nil {
		return entity.GreetingMyInfo{}, err
	}
	return greetingMyInfo, nil
}

func (s Ticket) GreetingRedeemResponse(redeemId int64) (entity.GreetingRedeemResponseByIdInfo, error) {
	var greetingRedeemResponseByIdInfo entity.GreetingRedeemResponseByIdInfo
	_, err := s.Client.R().
		SetResult(&greetingRedeemResponseByIdInfo).
		SetPathParam("redeemId", strconv.FormatInt(redeemId, 10)).
		Get("/greeting/redeem/{redeemId}/response")
	if err != nil {
		return entity.GreetingRedeemResponseByIdInfo{}, err
	}
	return greetingRedeemResponseByIdInfo, nil
}

func (s Ticket) GreetingRedeemScripDialog(redeemId int64) (entity.GreetingDialogInfo, error) {
	var greetingDialogInfo entity.GreetingDialogInfo
	_, err := s.Client.R().
		SetResult(&greetingDialogInfo).
		SetPathParam("redeemId", strconv.FormatInt(redeemId, 10)).
		Get("/greeting/redeem/{redeemId}/dialog")
	if err != nil {
		return entity.GreetingDialogInfo{}, err
	}
	return greetingDialogInfo, nil
}

func (s Ticket) GreetingRedeemThumbnailTempUrl(redeemId int64) (entity.GreetingTempUrlResponseInfo, error) {
	var greetingTempUrlResponseInfo entity.GreetingTempUrlResponseInfo
	_, err := s.Client.R().
		SetResult(&greetingTempUrlResponseInfo).
		SetPathParam("redeemId", strconv.FormatInt(redeemId, 10)).
		Get("/greeting/redeem/{redeemId}/thumbnail/tempurl")
	if err != nil {
		return entity.GreetingTempUrlResponseInfo{}, err
	}
	return greetingTempUrlResponseInfo, nil
}

func (s Ticket) GreetingRedeemTypeById(redeemTypeId int64) (entity.GreetingRedeemTypeInfo, error) {
	var greetingRedeemTypeInfo entity.GreetingRedeemTypeInfo
	_, err := s.Client.R().
		SetResult(&greetingRedeemTypeInfo).
		SetPathParam("redeemTypeId", strconv.FormatInt(redeemTypeId, 10)).
		Get("/greeting/ticket/redeem-type/{redeemTypeId}")
	if err != nil {
		return entity.GreetingRedeemTypeInfo{}, err
	}
	return greetingRedeemTypeInfo, nil
}

func (s Ticket) GreetingRedeemVideo(redeemId int64, greetingRedeemVideoBodyInfo entity.GreetingRedeemVideoBodyInfo) (entity.GreetingMyInfo, error) {
	var greetingMyInfo entity.GreetingMyInfo
	_, err := s.Client.R().
		SetResult(&greetingMyInfo).
		SetBody(greetingRedeemVideoBodyInfo).
		SetPathParam("redeemId", strconv.FormatInt(redeemId, 10)).
		Post("/greeting/redeem/{redeemId}/video")
	if err != nil {
		return entity.GreetingMyInfo{}, err
	}
	return greetingMyInfo, nil
}

func (s Ticket) GreetingRedeemVideoTempUrl(redeemId int64) (entity.GreetingTempUrlResponseInfo, error) {
	var greetingTempUrlResponseInfo entity.GreetingTempUrlResponseInfo
	_, err := s.Client.R().
		SetResult(&greetingTempUrlResponseInfo).
		SetPathParam("redeemId", strconv.FormatInt(redeemId, 10)).
		Post("/greeting/redeem/{redeemId}/video/tempurl")
	if err != nil {
		return entity.GreetingTempUrlResponseInfo{}, err
	}
	return greetingTempUrlResponseInfo, nil
}

func (s Ticket) GreetingTicketAvailable() ([]entity.GreetingTicketAvailableInfo, error) {
	var greetingTicketAvailableInfo []entity.GreetingTicketAvailableInfo
	_, err := s.Client.R().
		SetResult(&greetingTicketAvailableInfo).
		Get("/greeting/ticket/available")
	if err != nil {
		return nil, err
	}
	return greetingTicketAvailableInfo, nil
}

func (s Ticket) GreetingTicketById(ticketId int64) (entity.GreetingTicketByIdInfo, error) {
	var greetingTicketByIdInfo entity.GreetingTicketByIdInfo
	_, err := s.Client.R().
		SetResult(&greetingTicketByIdInfo).
		SetPathParam("ticketId", strconv.FormatInt(ticketId, 10)).
		Get("/greeting/ticket/{ticketId}")
	if err != nil {
		return entity.GreetingTicketByIdInfo{}, err
	}
	return greetingTicketByIdInfo, nil
}

func (s Ticket) GreetingTicketByTicketAndMemberId(ticketId int64, memberId int64) (entity.GreetingTicketMemberTypeInfo, error) {
	var greetingTicketMemberTypeInfo entity.GreetingTicketMemberTypeInfo
	_, err := s.Client.R().
		SetResult(&greetingTicketMemberTypeInfo).
		SetPathParam("ticketId", strconv.FormatInt(ticketId, 10)).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Get("/greeting/ticket/{ticketId}/member/{memberId}")
	if err != nil {
		return entity.GreetingTicketMemberTypeInfo{}, err
	}
	return greetingTicketMemberTypeInfo, nil
}

func (s Ticket) LogoutTokenX(t string) (resty.Response, error) {
	res, err := s.Client.R().
		SetQueryParam("t", t).
		Delete("/token-x/logout")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s Ticket) MeetYouEventInfo(ticketId int64) (entity.EventInfo, error) {
	var eventInfo entity.EventInfo
	_, err := s.Client.R().
		SetResult(&eventInfo).
		SetPathParam("ticketId", strconv.FormatInt(ticketId, 10)).
		Get("/meetingyou/ticket/{ticketId}/event/info")
	if err != nil {
		return entity.EventInfo{}, err
	}
	return eventInfo, nil
}

func (s Ticket) MeetYouLastMinuteEventInfo(memberId int64) (entity.EventInfo, error) {
	var eventInfo entity.EventInfo
	_, err := s.Client.R().
		SetResult(&eventInfo).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Get("/meetingyou/lastminute/member/{memberId}/event/info")
	if err != nil {
		return entity.EventInfo{}, err
	}
	return eventInfo, nil
}

func (s Ticket) MeetYouLastMinuteMemberRounds(memberId int64) (entity.MemberRounds, error) {
	var memberRounds entity.MemberRounds
	_, err := s.Client.R().
		SetResult(&memberRounds).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Get("/meetingyou/lastminute/member/{memberId}")
	if err != nil {
		return entity.MemberRounds{}, err
	}
	return memberRounds, nil
}

func (s Ticket) MeetYouLastMinuteRedeemableMemberList() ([]entity.RedeemableMemberInfo, error) {
	var redeemableMemberInfo []entity.RedeemableMemberInfo
	_, err := s.Client.R().
		SetResult(&redeemableMemberInfo).
		Get("/meetingyou/lastminute/members")
	if err != nil {
		return nil, err
	}
	return redeemableMemberInfo, nil
}

func (s Ticket) MeetYouLastMinuteTicketAvailable(memberId int64) ([]entity.TicketInfo, error) {
	var ticketInfo []entity.TicketInfo
	_, err := s.Client.R().
		SetResult(&ticketInfo).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Get("meetingyou/lastminute/member/{memberId}/ticket/available")
	if err != nil {
		return nil, err
	}
	return ticketInfo, nil
}

func (s Ticket) MeetYouMemberRounds(ticketId int64, memberId int64) (entity.MemberRounds, error) {
	var memberRounds entity.MemberRounds
	_, err := s.Client.R().
		SetResult(&memberRounds).
		SetPathParam("ticketId", strconv.FormatInt(ticketId, 10)).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Get("/meetingyou/ticket/{ticketId}/event/member/{memberId}")
	if err != nil {
		return entity.MemberRounds{}, err
	}
	return memberRounds, nil
}

func (s Ticket) MeetYouMyMeetingDetail(refCode string) (entity.RedeemInfo, error) {
	var redeemInfo entity.RedeemInfo
	_, err := s.Client.R().
		SetResult(&redeemInfo).
		SetPathParam("refCode", refCode).
		Get("/meetingyou/my/redeem/{refCode}")
	if err != nil {
		return entity.RedeemInfo{}, err
	}
	return redeemInfo, nil
}

func (s Ticket) MeetYouMyMeetingHistory(take int, skip int64) ([]entity.RedeemInfo, error) {
	var redeemInfo []entity.RedeemInfo
	_, err := s.Client.R().
		SetResult(&redeemInfo).
		SetQueryParam("take", strconv.Itoa(take)).
		SetQueryParam("skip", strconv.FormatInt(skip, 10)).
		Get("/meetingyou/my/history")
	if err != nil {
		return nil, err
	}
	return redeemInfo, nil
}

func (s Ticket) MeetYouMyMeetingSchedule(take int, skip int64) ([]entity.RedeemInfo, error) {
	var redeemInfo []entity.RedeemInfo
	_, err := s.Client.R().
		SetResult(&redeemInfo).
		SetQueryParam("take", strconv.Itoa(take)).
		SetQueryParam("skip", strconv.FormatInt(skip, 10)).
		Get("/meetingyou/my/schedule")
	if err != nil {
		return nil, err
	}
	return redeemInfo, nil
}

func (s Ticket) MeetYouMyStats() (entity.MyStats, error) {
	var myStats entity.MyStats
	_, err := s.Client.R().
		SetResult(&myStats).
		Get("/meetingyou/my/stats")
	if err != nil {
		return entity.MyStats{}, err
	}
	return myStats, nil
}

func (s Ticket) MeetYouRedeem(redeemMeetYouRequest entity.RedeemMeetYouRequest) (entity.RedeemInfo, error) {
	var redeemInfo entity.RedeemInfo
	_, err := s.Client.R().
		SetResult(&redeemInfo).
		SetBody(redeemMeetYouRequest).
		Post("/meetingyou/redeem")
	if err != nil {
		return entity.RedeemInfo{}, err
	}
	return redeemInfo, nil
}

func (s Ticket) MeetYouRedeemableMemberList(ticketId int64) ([]entity.RedeemableMemberInfo, error) {
	var redeemableMemberInfo []entity.RedeemableMemberInfo
	_, err := s.Client.R().
		SetResult(&redeemableMemberInfo).
		SetPathParam("ticketId", strconv.FormatInt(ticketId, 10)).
		Get("/meetingyou/ticket/{ticketId}/event/members")
	if err != nil {
		return nil, err
	}
	return redeemableMemberInfo, nil
}

func (s Ticket) MeetYouTicketAvailable() ([]entity.TicketInfo, error) {
	var ticketInfo []entity.TicketInfo
	_, err := s.Client.R().
		SetResult(&ticketInfo).
		Get("/meetingyou/ticket/available")
	if err != nil {
		return nil, err
	}
	return ticketInfo, nil
}

func (s Ticket) OnChangePin(changePinModel entity.ChangePinModel) (resty.Response, error) {
	res, err := s.Client.R().
		SetBody(changePinModel).
		Post("/pin/change/confirm")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s Ticket) OnClaimTicketIam(iamShopClaimOrderModel entity.IamShopClaimOrderModel) ([]entity.IamShopClaimOrderItemModel, error) {
	var iamShopClaimOrderItemModel []entity.IamShopClaimOrderItemModel
	_, err := s.Client.R().
		SetResult(&iamShopClaimOrderItemModel).
		SetBody(iamShopClaimOrderModel).
		Post("/claim/v2/iamshop")
	if err != nil {
		return nil, err
	}
	return iamShopClaimOrderItemModel, nil
}

func (s Ticket) OnClaimTicketShopee(shopeeClaimOrderModel entity.ShopeeClaimOrderModel) ([]entity.ShopeeClaimOrderItemModel, error) {
	var shopeeClaimOrderItemModel []entity.ShopeeClaimOrderItemModel
	_, err := s.Client.R().
		SetResult(&shopeeClaimOrderItemModel).
		SetBody(shopeeClaimOrderModel).
		Post("/claim/v2/shopee")
	if err != nil {
		return nil, err
	}
	return shopeeClaimOrderItemModel, nil
}

func (s Ticket) OnClaimTicketShopeeMock() ([]entity.ShopeeClaimOrderItemModel, error) {
	var shopeeClaimOrderItemModel []entity.ShopeeClaimOrderItemModel
	_, err := s.Client.R().
		SetResult(&shopeeClaimOrderItemModel).
		Get("/claim/v2/shopee")
	if err != nil {
		return nil, err
	}
	return shopeeClaimOrderItemModel, nil
}

func (s Ticket) OnCreateWallet(walletPinModel entity.WalletPinModel) (entity.WalletInfoModel, error) {
	var walletInfoModel entity.WalletInfoModel
	_, err := s.Client.R().
		SetResult(&walletInfoModel).
		SetBody(walletPinModel).
		Post("/wallet/create")
	if err != nil {
		return entity.WalletInfoModel{}, err
	}
	return walletInfoModel, nil
}

func (s Ticket) OnForgotPin() (resty.Response, error) {
	res, err := s.Client.R().
		Post("/pin/forgot")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s Ticket) OnOfflineTicket(showQRTicketItem entity.ShowQRTicketItem) (entity.QRTicketModel, error) {
	var qrTicketModel entity.QRTicketModel
	_, err := s.Client.R().
		SetResult(&qrTicketModel).
		SetBody(showQRTicketItem).
		Post("/qr/ticket/offline")
	if err != nil {
		return entity.QRTicketModel{}, err
	}
	return qrTicketModel, nil
}

func (s Ticket) OnOnlineTicket(showQRTicketItem entity.ShowQRTicketItem) (entity.QRTicketModel, error) {
	var qrTicketModel entity.QRTicketModel
	_, err := s.Client.R().
		SetResult(&qrTicketModel).
		SetBody(showQRTicketItem).
		Post("/qr/ticket")
	if err != nil {
		return entity.QRTicketModel{}, err
	}
	return qrTicketModel, nil
}

func (s Ticket) OnResetPin(resetPinModel entity.ResetPinModel) (resty.Response, error) {
	res, err := s.Client.R().
		SetBody(resetPinModel).
		Post("/pin/reset/confirm")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s Ticket) OnTicketTransfer(ticketTransferItem entity.TicketTransferItem) (entity.TicketTransferModel, error) {
	var ticketTransferModel entity.TicketTransferModel
	_, err := s.Client.R().
		SetResult(&ticketTransferModel).
		SetBody(ticketTransferItem).
		Post("/transfer")
	if err != nil {
		return entity.TicketTransferModel{}, err
	}
	return ticketTransferModel, nil
}

func (s Ticket) OnTicketTransferCommit(ticketTransferCommitItem entity.TicketTransferCommitItem) (entity.TransferCommitResultModel, error) {
	var transferCommitResultModel entity.TransferCommitResultModel
	_, err := s.Client.R().
		SetResult(&transferCommitResultModel).
		SetBody(ticketTransferCommitItem).
		Post("/transfer/commit")
	if err != nil {
		return entity.TransferCommitResultModel{}, err
	}
	return transferCommitResultModel, nil
}

func (s Ticket) OnVerifyOldPin(walletPinModel entity.WalletPinModel) (resty.Response, error) {
	res, err := s.Client.R().
		SetBody(walletPinModel).
		Post("/pin/change/verify")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s Ticket) OnVerifyResetPin(verifyCodePinModel entity.VerifyCodePinModel) (resty.Response, error) {
	res, err := s.Client.R().
		SetBody(verifyCodePinModel).
		Post("/pin/reset/verify")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

func (s Ticket) PostTermsTokenX(versionNo string) (entity.TokenXTermsPostResponseInfo, error) {
	var tokenXTermsPostResponseInfo entity.TokenXTermsPostResponseInfo
	_, err := s.Client.R().
		SetResult(&tokenXTermsPostResponseInfo).
		Post("/token-x/terms/" + versionNo)
	if err != nil {
		return entity.TokenXTermsPostResponseInfo{}, err
	}
	return tokenXTermsPostResponseInfo, nil
}

func (s Ticket) PutWalletPin(str string, str2 string) (resty.Response, error) {
	res, err := s.Client.R().
		SetQueryParam("pin", str).
		SetQueryParam("password", str2).
		Put("/pin")
	if err != nil {
		return resty.Response{}, err
	}
	return *res, nil
}

type TicketConfiguration func(*Ticket) error

func WithClient(client *resty.Client) TicketConfiguration {
	return func(e *Ticket) error {
		e.Client = client
		e.Client.SetBaseURL("https://ticket.bnk48.io")
		return nil
	}
}

func New(cfgs ...TicketConfiguration) (Ticket, error) {
	service := Ticket{
		Client: client.NewClient().SetBaseURL("https://ticket.bnk48.io"),
	}
	for _, cfg := range cfgs {
		if err := cfg(&service); err != nil {
			return Ticket{}, err
		}
	}
	return service, nil
}

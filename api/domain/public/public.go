package public

import (
	"strconv"

	"github.com/dvwzj/iam48-go/api/entity"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type Public struct {
	Client *resty.Client
}

func (s Public) GetAds() (entity.AdsAll, error) {
	var ads entity.AdsAll
	_, err := s.Client.R().
		SetResult(&ads).
		Get("/ads")
	if err != nil {
		return entity.AdsAll{}, err
	}
	return ads, nil
}

func (s Public) GetAllMembers() ([]entity.MemberProfile, error) {
	var members []entity.MemberProfile
	_, err := s.Client.R().
		SetResult(&members).
		Get("/members/all")
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (s Public) GetAllMembersFansTopRank() ([]entity.MemberRankStats, error) {
	var members []entity.MemberRankStats
	_, err := s.Client.R().
		SetResult(&members).
		Get("/members/all/fans/toprank")
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (s Public) GetAllMembersFansTopRankSummary() ([]entity.MemberRankStats, error) {
	var members []entity.MemberRankStats
	_, err := s.Client.R().
		SetResult(&members).
		Get("/members/all/fans/toprank/summary")
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (s Public) GetAllMembersFansTopRankToday() ([]entity.MemberRankStats, error) {
	var members []entity.MemberRankStats
	_, err := s.Client.R().
		SetResult(&members).
		Get("/members/all/fans/toprank/today")
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (s Public) GetAllMembersFansTopRankTodaySummary() ([]entity.MemberRankStats, error) {
	var members []entity.MemberRankStats
	_, err := s.Client.R().
		SetResult(&members).
		Get("/members/all/fans/toprank/today/summary")
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (s Public) GetAllMembersStats() (entity.CoreMembersStats, error) {
	var members entity.CoreMembersStats
	_, err := s.Client.R().
		SetResult(&members).
		Get("/members/stats")
	if err != nil {
		return entity.CoreMembersStats{}, err
	}
	return members, nil
}

func (s Public) GetAllMiniVideo(skip int, take int) ([]entity.LiveVideoData, error) {
	var videos []entity.LiveVideoData
	_, err := s.Client.R().
		SetResult(&videos).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/content/mini-video/videos")
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (s Public) GetAllTimelineFeed(amount int64, lastId *int64) (entity.TimelineItemModel, error) {
	var timeline entity.TimelineItemModel
	r := s.Client.R().
		SetResult(&timeline).
		SetQueryParam("amount", strconv.FormatInt(amount, 10))
	if lastId != nil {
		r.SetQueryParam("last_id", strconv.FormatInt(*lastId, 10))
	}
	_, err := r.Get("/timeline/all")
	if err != nil {
		return entity.TimelineItemModel{}, err
	}
	return timeline, nil
}

func (s Public) GetAllTopRang() ([]entity.MembersStatsInfo, error) {
	var members []entity.MembersStatsInfo
	_, err := s.Client.R().
		SetResult(&members).
		Get("/users/all/toprank")
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (s Public) GetAllTopRankToday() ([]entity.MembersStatsInfo, error) {
	var members []entity.MembersStatsInfo
	_, err := s.Client.R().
		SetResult(&members).
		Get("/users/all/toprank/today")
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (s Public) GetCampaign(skip int, take int) (entity.CampaignNewModel, error) {
	var campaign entity.CampaignNewModel
	_, err := s.Client.R().
		SetResult(&campaign).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/campaigns/running")
	if err != nil {
		return entity.CampaignNewModel{}, err
	}
	return campaign, nil
}

func (s Public) GetCampaignAll(skip int, take int) ([]entity.CampaignMainItemInfo, error) {
	var campaign []entity.CampaignMainItemInfo
	_, err := s.Client.R().
		SetResult(&campaign).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/campaigns/all")
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s Public) GetCampaignById(campaignId int64) (entity.CampaignDetailInfo, error) {
	var campaign entity.CampaignDetailInfo
	_, err := s.Client.R().
		SetResult(&campaign).
		SetPathParam("campaignId", strconv.FormatInt(campaignId, 10)).
		Get("/campaign/{campaignId}")
	if err != nil {
		return entity.CampaignDetailInfo{}, err
	}
	return campaign, nil
}

func (s Public) GetCampaignEnd(skip int, take int) ([]entity.CampaignMainItemInfo, error) {
	var campaign []entity.CampaignMainItemInfo
	_, err := s.Client.R().
		SetResult(&campaign).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/campaigns/ended")
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s Public) GetCampaignEndByMemberSet(memberIdSet string, skip int, take int) ([]entity.CampaignMainItemInfo, error) {
	var campaign []entity.CampaignMainItemInfo
	_, err := s.Client.R().
		SetResult(&campaign).
		SetPathParam("memberIdSet", memberIdSet).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/campaigns/ended/members/{memberIdSet}")
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s Public) GetCampaignMoreInfo(campaignId int64) (entity.BaseDetailResponse, error) {
	var campaign entity.BaseDetailResponse
	_, err := s.Client.R().
		SetResult(&campaign).
		SetPathParam("campaignId", strconv.FormatInt(campaignId, 10)).
		Get("/campaign/{campaignId}/details")
	if err != nil {
		return entity.BaseDetailResponse{}, err
	}
	return campaign, nil
}

func (s Public) GetCampaignNearEnd(skip int, take int) ([]entity.CampaignMainItemInfo, error) {
	var campaign []entity.CampaignMainItemInfo
	_, err := s.Client.R().
		SetResult(&campaign).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/campaigns/nearend")
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s Public) GetCampaignNearEndByMemberSet(memberIdSet string, skip int, take int) ([]entity.CampaignMainItemInfo, error) {
	var campaign []entity.CampaignMainItemInfo
	_, err := s.Client.R().
		SetResult(&campaign).
		SetPathParam("memberIdSet", memberIdSet).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/campaigns/nearend/members/{memberIdSet}")
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s Public) GetCampaignNearSuccess(skip int, take int) ([]entity.CampaignMainItemInfo, error) {
	var campaign []entity.CampaignMainItemInfo
	_, err := s.Client.R().
		SetResult(&campaign).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/campaigns/nearsuccess")
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s Public) GetCampaignNearSuccessByMemberSet(memberIdSet string, skip int, take int) ([]entity.CampaignMainItemInfo, error) {
	var campaign []entity.CampaignMainItemInfo
	_, err := s.Client.R().
		SetResult(&campaign).
		SetPathParam("memberIdSet", memberIdSet).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/campaigns/nearsuccess/members/{memberIdSet}")
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s Public) GetCampaignNewest(skip int, take int) ([]entity.CampaignMainItemInfo, error) {
	var campaign []entity.CampaignMainItemInfo
	_, err := s.Client.R().
		SetResult(&campaign).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/campaigns/newest")
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s Public) GetCampaignNewestByMemberSet(memberIdSet string, skip int, take int) ([]entity.CampaignMainItemInfo, error) {
	var campaign []entity.CampaignMainItemInfo
	_, err := s.Client.R().
		SetResult(&campaign).
		SetPathParam("memberIdSet", memberIdSet).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/campaigns/newest/members/{memberIdSet}")
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s Public) GetCampaignPopular(skip int, take int) ([]entity.CampaignMainItemInfo, error) {
	var campaign []entity.CampaignMainItemInfo
	_, err := s.Client.R().
		SetResult(&campaign).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/campaigns/popular")
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s Public) GetCampaignPopularByMemberSet(memberIdSet string, skip int, take int) ([]entity.CampaignMainItemInfo, error) {
	var campaign []entity.CampaignMainItemInfo
	_, err := s.Client.R().
		SetResult(&campaign).
		SetPathParam("memberIdSet", memberIdSet).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/campaigns/popular/members/{memberIdSet}")
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s Public) GetCampaignRank(campaignId int64) ([]entity.MembersStatsInfo, error) {
	var membersStats []entity.MembersStatsInfo
	_, err := s.Client.R().
		SetResult(&membersStats).
		SetPathParam("campaignId", strconv.FormatInt(campaignId, 10)).
		Get("/campaign/{campaignId}/backers/byrank")
	if err != nil {
		return nil, err
	}
	return membersStats, nil
}

func (s Public) GetCampaignSort() ([]entity.CampaignSortInfo, error) {
	var campaignSort []entity.CampaignSortInfo
	_, err := s.Client.R().
		SetResult(&campaignSort).
		Get("/campaigns/sort")
	if err != nil {
		return nil, err
	}
	return campaignSort, nil
}

func (s Public) GetCommentList(contentId int64, take int, lastCommentId int) ([]entity.CommentViewModel, error) {
	var comment []entity.CommentViewModel
	_, err := s.Client.R().
		SetResult(&comment).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetQueryParam("take", strconv.Itoa(take)).
		SetQueryParam("lastCommentId", strconv.Itoa(lastCommentId)).
		Get("/comment/content/{contentId}")
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s Public) GetCommentMultipleList(contentId int64, commentIdList string) ([]entity.CommentViewModel, error) {
	var comment []entity.CommentViewModel
	_, err := s.Client.R().
		SetResult(&comment).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetQueryParam("commentIdList", commentIdList).
		Get("/comment/content/{contentId}/multiple")
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s Public) GetContentRecommend() ([]entity.ContentRecommendInfo, error) {
	var contentRecommend []entity.ContentRecommendInfo
	_, err := s.Client.R().
		SetResult(&contentRecommend).
		Get("/content/purchasable/recommend")
	if err != nil {
		return nil, err
	}
	return contentRecommend, nil
}

func (s Public) GetDebugFooterList(dk string) ([]entity.DebugTabbarListInfo, error) {
	var debugFooter []entity.DebugTabbarListInfo
	_, err := s.Client.R().
		SetResult(&debugFooter).
		SetQueryParam("dk", dk).
		Get("/_debug/tabbar")
	if err != nil {
		return nil, err
	}
	return debugFooter, nil
}

func (s Public) GetDebugHeaderList(dk string) ([]entity.DebugHeaderListInfo, error) {
	var debugHeader []entity.DebugHeaderListInfo
	_, err := s.Client.R().
		SetResult(&debugHeader).
		SetQueryParam("dk", dk).
		Get("/_debug/champ-banner")
	if err != nil {
		return nil, err
	}
	return debugHeader, nil
}

func (s Public) GetDigitalStudioStats(videoContentId int64, skip int, take int) (entity.ContentVideoStatsInfo, error) {
	var digitalStudioStats entity.ContentVideoStatsInfo
	_, err := s.Client.R().
		SetResult(&digitalStudioStats).
		SetPathParam("videoContentId", strconv.FormatInt(videoContentId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/content/digital-studio/video/{videoContentId}/stats")
	if err != nil {
		return entity.ContentVideoStatsInfo{}, err
	}
	return digitalStudioStats, nil
}

func (s Public) GetDigitalStudioVideos(skip int, take int) ([]entity.LiveVideoData, error) {
	var digitalStudioVideos []entity.LiveVideoData
	_, err := s.Client.R().
		SetResult(&digitalStudioVideos).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/content/digital-studio/videos")
	if err != nil {
		return nil, err
	}
	return digitalStudioVideos, nil
}

func (s Public) GetDigitalStudioVideosById(videoId int64) (entity.LiveVideoData, error) {
	var digitalStudioVideos entity.LiveVideoData
	_, err := s.Client.R().
		SetResult(&digitalStudioVideos).
		SetPathParam("videoId", strconv.FormatInt(videoId, 10)).
		Get("/content/digital-studio/video/{videoId}")
	if err != nil {
		return entity.LiveVideoData{}, err
	}
	return digitalStudioVideos, nil
}

func (s Public) GetDigitalstudioComments(contentId int64, skip int, take int) ([]entity.CommentInfo, error) {
	var digitalstudioComments []entity.CommentInfo
	_, err := s.Client.R().
		SetResult(&digitalstudioComments).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/content/digital-studio/video/{contentId}/comments")
	if err != nil {
		return nil, err
	}
	return digitalstudioComments, nil
}

func (s Public) GetDiscover() (entity.DiscoverSectionInfo, error) {
	var discover entity.DiscoverSectionInfo
	_, err := s.Client.R().
		SetResult(&discover).
		Get("/discover")
	if err != nil {
		return entity.DiscoverSectionInfo{}, err
	}
	return discover, nil
}

func (s Public) GetDiscoverTimelineGift(contentId int64) (entity.TimelineContentDataModel, error) {
	var discoverTimelineGift entity.TimelineContentDataModel
	_, err := s.Client.R().
		SetResult(&discoverTimelineGift).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Get("/timeline/{contentId}/info/v2")
	if err != nil {
		return entity.TimelineContentDataModel{}, err
	}
	return discoverTimelineGift, nil
}

func (s Public) GetElectionBanner() ([]entity.ElectionBannerInfo, error) {
	var electionBanner []entity.ElectionBannerInfo
	_, err := s.Client.R().
		SetResult(&electionBanner).
		Get("/banner/election")
	if err != nil {
		return nil, err
	}
	return electionBanner, nil
}

func (s Public) GetEvent(year int, month int) ([]entity.ScheduleEvent, error) {
	var event []entity.ScheduleEvent
	_, err := s.Client.R().
		SetResult(&event).
		SetPathParam("year", strconv.Itoa(year)).
		SetPathParam("month", strconv.Itoa(month)).
		Get("/events/{year}/{month}")
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (s Public) GetEventDetails(eventId int64) (entity.ScheduleEvent, error) {
	var eventDetails entity.ScheduleEvent
	_, err := s.Client.R().
		SetResult(&eventDetails).
		SetPathParam("eventId", strconv.FormatInt(eventId, 10)).
		Get("/event/{eventId}")
	if err != nil {
		return entity.ScheduleEvent{}, err
	}
	return eventDetails, nil
}

func (s Public) GetEventTheater() ([]entity.EventTheaterInfo, error) {
	var eventTheater []entity.EventTheaterInfo
	_, err := s.Client.R().
		SetResult(&eventTheater).
		Get("/theater")
	if err != nil {
		return nil, err
	}
	return eventTheater, nil
}

func (s Public) GetEventTheaterInfo(id int64) (entity.EventTheaterInfo, error) {
	var eventTheaterInfo entity.EventTheaterInfo
	_, err := s.Client.R().
		SetResult(&eventTheaterInfo).
		SetPathParam("id", strconv.FormatInt(id, 10)).
		Get("/theater/{id}")
	if err != nil {
		return entity.EventTheaterInfo{}, err
	}
	return eventTheaterInfo, nil
}

func (s Public) GetGames() ([]entity.DiscoverGameListItem, error) {
	var games []entity.DiscoverGameListItem
	_, err := s.Client.R().
		SetResult(&games).
		Get("/games")
	if err != nil {
		return nil, err
	}
	return games, nil
}

func (s Public) GetGiftFloatMessageLiveChat(app string) ([]entity.GiftFloatMessageLiveInfo, error) {
	var giftFloatMessageLiveChat []entity.GiftFloatMessageLiveInfo
	_, err := s.Client.R().
		SetResult(&giftFloatMessageLiveChat).
		SetQueryParam("app", app).
		Get("/products/gift/section/member-live-chat")
	if err != nil {
		return nil, err
	}
	return giftFloatMessageLiveChat, nil
}

func (s Public) GetGiftFloatMessageLiveChatById(memberId int64, app string) ([]entity.GiftFloatMessageLiveInfo, error) {
	var giftFloatMessageLiveChatById []entity.GiftFloatMessageLiveInfo
	_, err := s.Client.R().
		SetResult(&giftFloatMessageLiveChatById).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		SetQueryParam("app", app).
		Get("/products/gift/section/member-live-chat/{memberId}")
	if err != nil {
		return nil, err
	}
	return giftFloatMessageLiveChatById, nil
}

func (s Public) GetHistoryRanking(batchId int64) (entity.SpecialFansDayRankingList, error) {
	var historyRanking entity.SpecialFansDayRankingList
	_, err := s.Client.R().
		SetResult(&historyRanking).
		SetPathParam("batchId", strconv.FormatInt(batchId, 10)).
		Get("/batch/{batchId}")
	if err != nil {
		return entity.SpecialFansDayRankingList{}, err
	}
	return historyRanking, nil
}

func (s Public) GetLivePlayback(skip int, take int) ([]entity.LiveVideoData, error) {
	var livePlayback []entity.LiveVideoData
	_, err := s.Client.R().
		SetResult(&livePlayback).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/content/member-live/videos")
	if err != nil {
		return nil, err
	}
	return livePlayback, nil
}

func (s Public) GetLivePlaybackById(videoId int64) (entity.LiveVideoData, error) {
	var livePlaybackById entity.LiveVideoData
	_, err := s.Client.R().
		SetResult(&livePlaybackById).
		SetPathParam("videoId", strconv.FormatInt(videoId, 10)).
		Get("/content/member-live/video/{videoId}")
	if err != nil {
		return entity.LiveVideoData{}, err
	}
	return livePlaybackById, nil
}

func (s Public) GetLiveVideoStats(videoContentId int64) (entity.ContentVideoStatsInfo, error) {
	var liveVideoStats entity.ContentVideoStatsInfo
	_, err := s.Client.R().
		SetResult(&liveVideoStats).
		SetPathParam("videoContentId", strconv.FormatInt(videoContentId, 10)).
		Get("/content/member-live/video/{videoContentId}/stats")
	if err != nil {
		return entity.ContentVideoStatsInfo{}, err
	}
	return liveVideoStats, nil
}

func (s Public) GetMediaInfo(contentId int64) (entity.TimelineRecommendMediaModel, error) {
	var mediaInfo entity.TimelineRecommendMediaModel
	_, err := s.Client.R().
		SetResult(&mediaInfo).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Get("/media/{contentId}")
	if err != nil {
		return entity.TimelineRecommendMediaModel{}, err
	}
	return mediaInfo, nil
}

func (s Public) GetMeetYouBannerList() ([]entity.Banner, error) {
	var meetYouBannerList []entity.Banner
	_, err := s.Client.R().
		SetResult(&meetYouBannerList).
		Get("/banner/meeting-you")
	if err != nil {
		return nil, err
	}
	return meetYouBannerList, nil
}

func (s Public) GetBannerList(bannerType string) ([]entity.Banner, error) {
	var bannerList []entity.Banner
	_, err := s.Client.R().
		SetResult(&bannerList).
		SetPathParam("bannerType", bannerType).
		Get("/banner/{bannerType}")
	if err != nil {
		return nil, err
	}
	return bannerList, nil
}

func (s Public) GetMeetYouBannerListLastMinute() ([]entity.Banner, error) {
	var meetYouBannerListLastMinute []entity.Banner
	_, err := s.Client.R().
		SetResult(&meetYouBannerListLastMinute).
		Get("/banner/meeting-you-last-minute-full")
	if err != nil {
		return nil, err
	}
	return meetYouBannerListLastMinute, nil
}

func (s Public) GetMemberCatchupRecommend(memberId int64, contentId int64, skip int, take int) ([]entity.LiveVideoData, error) {
	var memberCatchupRecommend []entity.LiveVideoData
	_, err := s.Client.R().
		SetResult(&memberCatchupRecommend).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/member/{memberId}/member-live/{contentId}/recommend")
	if err != nil {
		return nil, err
	}
	return memberCatchupRecommend, nil
}

func (s Public) GetMemberFansTopRank(memberId int64) ([]entity.MembersStatsInfo, error) {
	var memberFansTopRank []entity.MembersStatsInfo
	_, err := s.Client.R().
		SetResult(&memberFansTopRank).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Get("/member/{memberId}/fans/toprank")
	if err != nil {
		return nil, err
	}
	return memberFansTopRank, nil
}

func (s Public) GetMemberFansTopRankSummary(memberId int64) ([]entity.MembersStatsInfo, error) {
	var memberFansTopRankSummary []entity.MembersStatsInfo
	_, err := s.Client.R().
		SetResult(&memberFansTopRankSummary).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Get("/member/{memberId}/fans/toprank/summary")
	if err != nil {
		return nil, err
	}
	return memberFansTopRankSummary, nil
}

func (s Public) GetMemberLiveSchedule() ([]entity.ScheduleModelInfo, error) {
	var memberLiveSchedule []entity.ScheduleModelInfo
	_, err := s.Client.R().
		SetResult(&memberLiveSchedule).
		Get("/schedules/member-live")
	if err != nil {
		return nil, err
	}
	return memberLiveSchedule, nil
}

func (s Public) GetMemberLiveVideo(memberId int64, skip int, take int) ([]entity.LiveVideoData, error) {
	var memberLiveVideo []entity.LiveVideoData
	_, err := s.Client.R().
		SetResult(&memberLiveVideo).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/member/{memberId}/member-lives")
	if err != nil {
		return nil, err
	}
	return memberLiveVideo, nil
}

func (s Public) GetMemberMiniVideo(memberId int64, skip int, take int) ([]entity.LiveVideoData, error) {
	var memberMiniVideo []entity.LiveVideoData
	_, err := s.Client.R().
		SetResult(&memberMiniVideo).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/member/{memberId}/mini-videos")
	if err != nil {
		return nil, err
	}
	return memberMiniVideo, nil
}

func (s Public) GetMemberProfile(memberId int64) (entity.MemberProfile, error) {
	var memberProfile entity.MemberProfile
	_, err := s.Client.R().
		SetResult(&memberProfile).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Get("/member/{memberId}/profile")
	if err != nil {
		return entity.MemberProfile{}, err
	}
	return memberProfile, nil
}

func (s Public) GetMemberStats(memberId int64) (entity.MemberProfileStatsInfo, error) {
	var memberStats entity.MemberProfileStatsInfo
	_, err := s.Client.R().
		SetResult(&memberStats).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Get("/member/{memberId}/stats")
	if err != nil {
		return entity.MemberProfileStatsInfo{}, err
	}
	return memberStats, nil
}

func (s Public) GetMembersFansTopRankToday(memberId int64) ([]entity.MembersStatsInfo, error) {
	var membersFansTopRankToday []entity.MembersStatsInfo
	_, err := s.Client.R().
		SetResult(&membersFansTopRankToday).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Get("/member/{memberId}/fans/toprank/today")
	if err != nil {
		return nil, err
	}
	return membersFansTopRankToday, nil
}

func (s Public) GetMembersFansTopRankTodaySummary(memberId int64) ([]entity.MembersStatsInfo, error) {
	var membersFansTopRankTodaySummary []entity.MembersStatsInfo
	_, err := s.Client.R().
		SetResult(&membersFansTopRankTodaySummary).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		Get("/member/{memberId}/fans/toprank/today/summary")
	if err != nil {
		return nil, err
	}
	return membersFansTopRankTodaySummary, nil
}

func (s Public) GetMiniVideoById(videoId int64) (entity.LiveVideoData, error) {
	var miniVideoById entity.LiveVideoData
	_, err := s.Client.R().
		SetResult(&miniVideoById).
		SetPathParam("videoId", strconv.FormatInt(videoId, 10)).
		Get("/content/mini-video/video/{videoId}")
	if err != nil {
		return entity.LiveVideoData{}, err
	}
	return miniVideoById, nil
}

func (s Public) GetMiniVideoComments(contentId int64, skip int, take int) ([]entity.CommentInfo, error) {
	var miniVideoComments []entity.CommentInfo
	_, err := s.Client.R().
		SetResult(&miniVideoComments).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/content/mini-video/video/{contentId}/comments")
	if err != nil {
		return nil, err
	}
	return miniVideoComments, nil
}

func (s Public) GetMiniVideoStats(videoContentId int64) (entity.ContentVideoStatsInfo, error) {
	var miniVideoStats entity.ContentVideoStatsInfo
	_, err := s.Client.R().
		SetResult(&miniVideoStats).
		SetPathParam("videoContentId", strconv.FormatInt(videoContentId, 10)).
		Get("/content/mini-video/video/{videoContentId}/stats")
	if err != nil {
		return entity.ContentVideoStatsInfo{}, err
	}
	return miniVideoStats, nil
}

func (s Public) GetNineCookiesGift() (entity.Skus, error) {
	var nineCookiesGift entity.Skus
	_, err := s.Client.R().
		SetResult(&nineCookiesGift).
		Get("/products/gift/section/timeline-feed")
	if err != nil {
		return entity.Skus{}, err
	}
	return nineCookiesGift, nil
}

func (s Public) GetOnlyTimeline(memberId int64, lastId *int64, amount int64) (entity.TimelineItemModel, error) {
	var onlyTimeline entity.TimelineItemModel
	r := s.Client.R().
		SetResult(&onlyTimeline).
		SetPathParam("memberId", strconv.FormatInt(memberId, 10)).
		SetQueryParam("amount", strconv.FormatInt(amount, 10))
	if lastId != nil {
		r.SetQueryParam("last_id", strconv.FormatInt(*lastId, 10))
	}
	_, err := r.Get("/timeline/only/member/{memberId}")
	if err != nil {
		return entity.TimelineItemModel{}, err
	}
	return onlyTimeline, nil
}

func (s Public) GetProductCoin(app string) ([]entity.PurchaseCoins, error) {
	var productCoin []entity.PurchaseCoins
	_, err := s.Client.R().
		SetResult(&productCoin).
		SetQueryParam("app", app).
		Get("/products/coin")
	if err != nil {
		return nil, err
	}
	return productCoin, nil
}

func (s Public) GetProductGift() (entity.ProductGift, error) {
	var productGift entity.ProductGift
	_, err := s.Client.R().
		SetResult(&productGift).
		Get("/products/gift")
	if err != nil {
		return entity.ProductGift{}, err
	}
	return productGift, nil
}

func (s Public) GetProductMusic() ([]entity.MusicAlbumInfo, error) {
	var productMusic []entity.MusicAlbumInfo
	_, err := s.Client.R().
		SetResult(&productMusic).
		Get("/products/music")
	if err != nil {
		return nil, err
	}
	return productMusic, nil
}

func (s Public) GetPromotion() ([]entity.DiscoverPromotionListItem, error) {
	var promotion []entity.DiscoverPromotionListItem
	_, err := s.Client.R().
		SetResult(&promotion).
		Get("/promotion")
	if err != nil {
		return nil, err
	}
	return promotion, nil
}

func (s Public) GetProvinceDistrict() ([]entity.ProvinceDistrictResponseInfo, error) {
	var provinceDistrict []entity.ProvinceDistrictResponseInfo
	_, err := s.Client.R().
		SetResult(&provinceDistrict).
		Get("/addressinfo/province-district")
	if err != nil {
		return nil, err
	}
	return provinceDistrict, nil
}

func (s Public) GetPurchaseGifts() ([]entity.PurchaseGifts, error) {
	var purchaseGifts []entity.PurchaseGifts
	_, err := s.Client.R().
		SetResult(&purchaseGifts).
		Get("/api/purchase/gifts")
	if err != nil {
		return nil, err
	}
	return purchaseGifts, nil
}

func (s Public) GetRecommendMedia(skip int, take int) ([]entity.TimelineRecommendMediaModel, error) {
	var recommendMedia []entity.TimelineRecommendMediaModel
	_, err := s.Client.R().
		SetResult(&recommendMedia).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/media")
	if err != nil {
		return nil, err
	}
	return recommendMedia, nil
}

func (s Public) GetShopPaymentList() ([]entity.ShopPaymentListResponseInfo, error) {
	var shopPaymentList []entity.ShopPaymentListResponseInfo
	_, err := s.Client.R().
		SetResult(&shopPaymentList).
		Get("/shop/payment/list")
	if err != nil {
		return nil, err
	}
	return shopPaymentList, nil
}

func (s Public) GetShopProductById(productId int64) (entity.ShopProductInfo, error) {
	var shopProduct entity.ShopProductInfo
	_, err := s.Client.R().
		SetResult(&shopProduct).
		Get("/shop/product/{productId}")
	if err != nil {
		return entity.ShopProductInfo{}, err
	}
	return shopProduct, nil
}

func (s Public) GetShopProductGroup(skip int, take int) ([]entity.ShopProductGroupInfo, error) {
	var shopProductGroup []entity.ShopProductGroupInfo
	_, err := s.Client.R().
		SetResult(&shopProductGroup).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/shop/product/group")
	if err != nil {
		return nil, err
	}
	return shopProductGroup, nil
}

func (s Public) GetShopProductGroupDebug(dk string, skip int, take int) ([]entity.ShopProductGroupInfo, error) {
	var shopProductGroup []entity.ShopProductGroupInfo
	_, err := s.Client.R().
		SetResult(&shopProductGroup).
		SetQueryParam("dk", dk).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/_debug/product/group")
	if err != nil {
		return nil, err
	}
	return shopProductGroup, nil
}

func (s Public) GetShopProductList(skip int, take int) ([]entity.ShopProductInfo, error) {
	var shopProduct []entity.ShopProductInfo
	_, err := s.Client.R().
		SetResult(&shopProduct).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/shop/product/list")
	if err != nil {
		return nil, err
	}
	return shopProduct, nil
}

func (s Public) GetShopProductListV2(groupId int64, skip int, take int) (entity.ShopProductInfoV2, error) {
	var shopProduct entity.ShopProductInfoV2
	_, err := s.Client.R().
		SetResult(&shopProduct).
		SetQueryParam("groupId", strconv.FormatInt(groupId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/shop/product/list/v2")
	if err != nil {
		return entity.ShopProductInfoV2{}, err
	}
	return shopProduct, nil
}

func (s Public) GetShopProductListV2Debug(dk string, groupId int64, skip int, take int) (entity.ShopProductInfoV2, error) {
	var shopProduct entity.ShopProductInfoV2
	_, err := s.Client.R().
		SetResult(&shopProduct).
		SetQueryParam("dk", dk).
		SetQueryParam("groupId", strconv.FormatInt(groupId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/_debug/product/list/v2")
	if err != nil {
		return entity.ShopProductInfoV2{}, err
	}
	return shopProduct, nil
}

func (s Public) GetSpecialFanDayWinner() (entity.RankingBatchInfo, error) {
	var rankingBatch entity.RankingBatchInfo
	_, err := s.Client.R().
		SetResult(&rankingBatch).
		Get("/batch/latest")
	if err != nil {
		return entity.RankingBatchInfo{}, err
	}
	return rankingBatch, nil
}

func (s Public) GetSpecialFansDayByDate(date string) ([]entity.SpecialFansDayRankingInfo, error) {
	var specialFansDay []entity.SpecialFansDayRankingInfo
	_, err := s.Client.R().
		SetResult(&specialFansDay).
		SetPathParam("date", date).
		Get("/batch/ranking/{date}")
	if err != nil {
		return nil, err
	}
	return specialFansDay, nil
}

func (s Public) GetSpecialFansDayHistory() ([]entity.SpecialFansDayHistoryInfo, error) {
	var specialFansDay []entity.SpecialFansDayHistoryInfo
	_, err := s.Client.R().
		SetResult(&specialFansDay).
		Get("/batch")
	if err != nil {
		return nil, err
	}
	return specialFansDay, nil
}

func (s Public) GetSpecialFansDayRanking() (entity.SpecialFansDayRankingList, error) {
	var specialFansDay entity.SpecialFansDayRankingList
	_, err := s.Client.R().
		SetResult(&specialFansDay).
		Get("/batch/current")
	if err != nil {
		return entity.SpecialFansDayRankingList{}, err
	}
	return specialFansDay, nil
}

func (s Public) GetSubDistrict(districtId int64) ([]entity.SubDistrictInfo, error) {
	var subDistrict []entity.SubDistrictInfo
	_, err := s.Client.R().
		SetResult(&subDistrict).
		SetPathParam("districtId", strconv.FormatInt(districtId, 10)).
		Get("/addressinfo/province-district/{districtId}/sub-district")
	if err != nil {
		return nil, err
	}
	return subDistrict, nil
}

func (s Public) GetThankYouContent(batchId int64) (entity.ThankYouContentInfo, error) {
	var thankYou entity.ThankYouContentInfo
	_, err := s.Client.R().
		SetResult(&thankYou).
		SetPathParam("batchId", strconv.FormatInt(batchId, 10)).
		Get("/content/batch-thankyou/content/{batchId}")
	if err != nil {
		return entity.ThankYouContentInfo{}, err
	}
	return thankYou, nil
}

func (s Public) GetTheaterComments(contentId int64, skip int, take int) ([]entity.CommentInfo, error) {
	var comment []entity.CommentInfo
	_, err := s.Client.R().
		SetResult(&comment).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/content/theater-show/video/{contentId}/comments")
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s Public) GetTheaterPlaybackByContentId(contentId int64) (entity.TheaterPlaybackContentInfo, error) {
	var theaterPlayback entity.TheaterPlaybackContentInfo
	_, err := s.Client.R().
		SetResult(&theaterPlayback).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Get("/theater/playback/{contentId}")
	if err != nil {
		return entity.TheaterPlaybackContentInfo{}, err
	}
	return theaterPlayback, nil
}

func (s Public) GetTheaterPlaybackList(skip int, take int) ([]entity.TheaterPlaybackInfo, error) {
	var theaterPlayback []entity.TheaterPlaybackInfo
	_, err := s.Client.R().
		SetResult(&theaterPlayback).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/theater/playback")
	if err != nil {
		return nil, err
	}
	return theaterPlayback, nil
}

func (s Public) GetTheaterVideoStats(skip int, take int) (entity.ContentVideoStatsInfo, error) {
	var contentVideoStats entity.ContentVideoStatsInfo
	_, err := s.Client.R().
		SetResult(&contentVideoStats).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/content/theater-show/video/stats")
	if err != nil {
		return entity.ContentVideoStatsInfo{}, err
	}
	return contentVideoStats, nil
}

func (s Public) GetTheaterVideos(skip int, take int) ([]entity.LiveVideoData, error) {
	var liveVideoData []entity.LiveVideoData
	_, err := s.Client.R().
		SetResult(&liveVideoData).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/content/theater/videos")
	if err != nil {
		return nil, err
	}
	return liveVideoData, nil
}

func (s Public) GetTheaterVideosById(videoId int64) (entity.LiveVideoData, error) {
	var liveVideoData entity.LiveVideoData
	_, err := s.Client.R().
		SetResult(&liveVideoData).
		SetPathParam("videoId", strconv.FormatInt(videoId, 10)).
		Get("/content/theater/video/{videoId}")
	if err != nil {
		return entity.LiveVideoData{}, err
	}
	return liveVideoData, nil
}

func (s Public) GetTheme() (entity.ThemeInfo, error) {
	var theme entity.ThemeInfo
	_, err := s.Client.R().
		SetResult(&theme).
		Get("/theme")
	if err != nil {
		return entity.ThemeInfo{}, err
	}
	return theme, nil
}

func (s Public) GetTimelineFeedById(memberId string, amount int64, lastId *int64) (entity.TimelineItemModel, error) {
	var timelineItem entity.TimelineItemModel
	r := s.Client.R().
		SetResult(&timelineItem).
		SetPathParam("memberId", memberId).
		SetQueryParam("amount", strconv.FormatInt(amount, 10))
	if lastId != nil {
		r.SetQueryParam("last_id", strconv.FormatInt(*lastId, 10))
	}
	_, err := r.Get("/timeline/multi/member/{memberId}")
	if err != nil {
		return entity.TimelineItemModel{}, err
	}
	return timelineItem, nil
}

func (s Public) GetTimelineGiftStats(contentId int64) ([]entity.TimelineItemGiftStats, error) {
	var timelineItemGiftStats []entity.TimelineItemGiftStats
	_, err := s.Client.R().
		SetResult(&timelineItemGiftStats).
		SetPathParam("contentId", strconv.FormatInt(contentId, 10)).
		Get("/timeline/{contentId}/rank/gifts")
	if err != nil {
		return nil, err
	}
	return timelineItemGiftStats, nil
}

func (s Public) GetTimelineItem(itemType string, itemId int64) (entity.TimelineFeeds, error) {
	var timelineFeeds entity.TimelineFeeds
	_, err := s.Client.R().
		SetResult(&timelineFeeds).
		SetPathParam("itemType", itemType).
		SetPathParam("itemId", strconv.FormatInt(itemId, 10)).
		Get("/timeline/{itemType}/{itemId}")
	if err != nil {
		return entity.TimelineFeeds{}, err
	}
	return timelineFeeds, nil
}

func (s Public) GetUserLikedVideos(userId int64, skip int, take int) ([]entity.LiveVideoData, error) {
	var liveVideoData []entity.LiveVideoData
	_, err := s.Client.R().
		SetResult(&liveVideoData).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("take", strconv.Itoa(take)).
		Get("/user/{userId}/liked/videos")
	if err != nil {
		return nil, err
	}
	return liveVideoData, nil
}

func (s Public) GetUserMemeberFollow(userId int64) ([]entity.FollowMemberInfo, error) {
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

func (s Public) GetPublicUserProfile(userId int64) (entity.PublicUserProfileInfo, error) {
	var publicUserProfileInfo entity.PublicUserProfileInfo
	_, err := s.Client.R().
		SetResult(&publicUserProfileInfo).
		SetPathParam("userId", strconv.FormatInt(userId, 10)).
		Get("/user/{userId}/profile")
	if err != nil {
		return entity.PublicUserProfileInfo{}, err
	}
	return publicUserProfileInfo, nil
}

func (s Public) GetWaitingForLiveSchedule(scheduleId *string) (entity.ScheduleModelInfo, error) {
	var scheduleModelInfo entity.ScheduleModelInfo
	_, err := s.Client.R().
		SetResult(&scheduleModelInfo).
		SetPathParam("scheduleId", *scheduleId).
		Get("/schedules/member-live/{scheduleId}")
	if err != nil {
		return entity.ScheduleModelInfo{}, err
	}
	return scheduleModelInfo, nil
}

func (s Public) Version() (entity.VersionInfo, error) {
	var versionInfo entity.VersionInfo
	_, err := s.Client.R().
		SetResult(&versionInfo).
		Get("/version")
	if err != nil {
		return entity.VersionInfo{}, err
	}
	return versionInfo, nil
}

type PublicConfiguration func(*Public) error

func WithClient(client *resty.Client) PublicConfiguration {
	return func(e *Public) error {
		e.Client = client
		e.Client.SetBaseURL("https://public.bnk48.io")
		return nil
	}
}

func New(cfgs ...PublicConfiguration) (Public, error) {
	service := Public{
		Client: client.NewClient().SetBaseURL("https://public.bnk48.io"),
	}
	for _, cfg := range cfgs {
		if err := cfg(&service); err != nil {
			return Public{}, err
		}
	}
	return service, nil
}

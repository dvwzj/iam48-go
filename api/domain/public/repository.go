package public

import "github.com/dvwzj/iam48-go/api/entity"

type PublicRepository interface {
	// @GET("/ads")
	// g.b.l<AdsAll> getAds();
	GetAds() (entity.AdsAll, error)

	// @GET("/members/all")
	// g.b.l<List<MemberProfile>> getAllMembers();
	GetAllMembers() ([]entity.MemberProfile, error)

	// @GET("/members/all/fans/toprank")
	// g.b.l<List<MemberRankStats>> getAllMembersFansTopRank();
	GetAllMembersFansTopRank() ([]entity.MemberRankStats, error)

	// @GET("/members/all/fans/toprank/summary")
	// g.b.l<List<MemberRankStats>> getAllMembersFansTopRankSummary();
	GetAllMembersFansTopRankSummary() ([]entity.MemberRankStats, error)

	// @GET("/members/all/fans/toprank/today")
	// g.b.l<List<MemberRankStats>> getAllMembersFansTopRankToday();
	GetAllMembersFansTopRankToday() ([]entity.MemberRankStats, error)

	// @GET("/members/all/fans/toprank/today/summary")
	// g.b.l<List<MemberRankStats>> getAllMembersFansTopRankTodaySummary();
	GetAllMembersFansTopRankTodaySummary() ([]entity.MemberRankStats, error)

	// @GET("/members/stats")
	// g.b.l<CoreMembersStats> getAllMembersStats();
	GetAllMembersStats() (entity.CoreMembersStats, error)

	// @GET("/content/mini-video/videos")
	// g.b.l<List<LiveVideoData>> getAllMiniVideo(@Query("skip") int i2, @Query("take") int i3);
	GetAllMiniVideo(skip int, take int) ([]entity.LiveVideoData, error)

	// @GET("/timeline/all")
	// g.b.l<TimelineItemModel> getAllTimelineFeed(@Query("amount") long j2, @Nullable @Query("last_id") Long l2);
	GetAllTimelineFeed(amount int64, lastId *int64) (entity.TimelineItemModel, error)

	// @GET("/users/all/toprank")
	// g.b.l<List<MembersStatsInfo>> getAllTopRank();
	GetAllTopRang() ([]entity.MembersStatsInfo, error)

	// @GET("/users/all/toprank/today")
	// g.b.l<List<MembersStatsInfo>> getAllTopRankToday();
	GetAllTopRankToday() ([]entity.MembersStatsInfo, error)

	// @GET("/campaigns/running")
	// g.b.l<CampaignNewModel> getCampaign(@Query("skip") int i2, @Query("take") int i3);
	GetCampaign(skip int, take int) (entity.CampaignNewModel, error)

	// @GET("/campaigns/all")
	// g.b.l<List<CampaignMainItemInfo>> getCampaignAll(@Query("skip") int i2, @Query("take") int i3);
	GetCampaignAll(skip int, take int) ([]entity.CampaignMainItemInfo, error)

	// @GET("/campaign/{campaignId}")
	// g.b.l<CampaignDetailInfo> getCampaignById(@Path("campaignId") long j2);
	GetCampaignById(campaignId int64) (entity.CampaignDetailInfo, error)

	// @GET("/campaigns/ended")
	// g.b.l<List<CampaignMainItemInfo>> getCampaignEnd(@Query("skip") int i2, @Query("take") int i3);
	GetCampaignEnd(skip int, take int) ([]entity.CampaignMainItemInfo, error)

	// @GET("/campaigns/ended/members/{memberIdSet}")
	// g.b.l<List<CampaignMainItemInfo>> getCampaignEndByMemberSet(@Path("memberIdSet") @NotNull String str, @Query("skip") int i2, @Query("take") int i3);
	GetCampaignEndByMemberSet(memberIdSet string, skip int, take int) ([]entity.CampaignMainItemInfo, error)

	// @GET("/campaign/{campaignId}/details")
	// g.b.l<BaseDetailResponse> getCampaignMoreInfo(@Path("campaignId") long j2);
	GetCampaignMoreInfo(campaignId int64) (entity.BaseDetailResponse, error)

	// @GET("/campaigns/nearend")
	// g.b.l<List<CampaignMainItemInfo>> getCampaignNearEnd(@Query("skip") int i2, @Query("take") int i3);
	GetCampaignNearEnd(skip int, take int) ([]entity.CampaignMainItemInfo, error)

	// @GET("/campaigns/nearend/members/{memberIdSet}")
	// g.b.l<List<CampaignMainItemInfo>> getCampaignNearEndByMemberSet(@Path("memberIdSet") @NotNull String str, @Query("skip") int i2, @Query("take") int i3);
	GetCampaignNearEndByMemberSet(memberIdSet string, skip int, take int) ([]entity.CampaignMainItemInfo, error)

	// @GET("/campaigns/nearsuccess")
	// g.b.l<List<CampaignMainItemInfo>> getCampaignNearSuccess(@Query("skip") int i2, @Query("take") int i3);
	GetCampaignNearSuccess(skip int, take int) ([]entity.CampaignMainItemInfo, error)

	// @GET("/campaigns/nearsuccess/members/{memberIdSet}")
	// g.b.l<List<CampaignMainItemInfo>> getCampaignNearSuccessByMemberSet(@Path("memberIdSet") @NotNull String str, @Query("skip") int i2, @Query("take") int i3);
	GetCampaignNearSuccessByMemberSet(memberIdSet string, skip int, take int) ([]entity.CampaignMainItemInfo, error)

	// @GET("/campaigns/newest")
	// g.b.l<List<CampaignMainItemInfo>> getCampaignNewest(@Query("skip") int i2, @Query("take") int i3);
	GetCampaignNewest(skip int, take int) ([]entity.CampaignMainItemInfo, error)

	// @GET("/campaigns/newest/members/{memberIdSet}")
	// g.b.l<List<CampaignMainItemInfo>> getCampaignNewestByMemberSet(@Path("memberIdSet") @NotNull String str, @Query("skip") int i2, @Query("take") int i3);
	GetCampaignNewestByMemberSet(memberIdSet string, skip int, take int) ([]entity.CampaignMainItemInfo, error)

	// @GET("/campaigns/popular")
	// g.b.l<List<CampaignMainItemInfo>> getCampaignPopular(@Query("skip") int i2, @Query("take") int i3);
	GetCampaignPopular(skip int, take int) ([]entity.CampaignMainItemInfo, error)

	// @GET("/campaigns/popular/members/{memberIdSet}")
	// g.b.l<List<CampaignMainItemInfo>> getCampaignPopularByMemberSet(@Path("memberIdSet") @NotNull String str, @Query("skip") int i2, @Query("take") int i3);
	GetCampaignPopularByMemberSet(memberIdSet string, skip int, take int) ([]entity.CampaignMainItemInfo, error)

	// @GET("/campaign/{campaignId}/backers/byrank")
	// g.b.l<List<MembersStatsInfo>> getCampaignRank(@Path("campaignId") long j2);
	GetCampaignRank(campaignId int64) ([]entity.MembersStatsInfo, error)

	// @GET("/campaigns/sort")
	// g.b.l<List<CampaignSortInfo>> getCampaignSort();
	GetCampaignSort() ([]entity.CampaignSortInfo, error)

	// @GET("/comment/content/{contentId}")
	// g.b.l<List<CommentViewModel>> getCommentList(@Path("contentId") long j2, @Query("take") int i2, @Query("lastCommentId") int i3);
	GetCommentList(contentId int64, take int, lastCommentId int) ([]entity.CommentViewModel, error)

	// @GET("/comment/content/{contentId}/multiple")
	// g.b.l<List<CommentViewModel>> getCommentMultipleList(@Path("contentId") long j2, @NotNull @Query("commentIdList") String str);
	GetCommentMultipleList(contentId int64, commentIdList string) ([]entity.CommentViewModel, error)

	// @GET("/content/purchasable/recommend")
	// g.b.l<List<ContentRecommendInfo>> getContentRecommend();
	GetContentRecommend() ([]entity.ContentRecommendInfo, error)

	// @GET("/_debug/tabbar")
	// g.b.l<List<DebugTabbarListInfo>> getDebugFooterList(@NotNull @Query("dk") String str);
	GetDebugFooterList(dk string) ([]entity.DebugTabbarListInfo, error)

	// @GET("/_debug/champ-banner")
	// g.b.l<List<DebugHeaderListInfo>> getDebugHeaderList(@NotNull @Query("dk") String str);
	GetDebugHeaderList(dk string) ([]entity.DebugHeaderListInfo, error)

	// @GET("/content/digital-studio/video/{videoContentId}/stats")
	// g.b.l<ContentVideoStatsInfo> getDigitalStudioStats(@Query("skip") int i2, @Query("take") int i3);
	GetDigitalStudioStats(videoContentId int64, skip int, take int) (entity.ContentVideoStatsInfo, error)

	// @GET("/content/digital-studio/videos")
	// g.b.l<List<LiveVideoData>> getDigitalStudioVideos(@Query("skip") int i2, @Query("take") int i3);
	GetDigitalStudioVideos(skip int, take int) ([]entity.LiveVideoData, error)

	// @GET("/content/digital-studio/video/{videoId}")
	// g.b.l<LiveVideoData> getDigitalStudioVideosById(@Path("videoId") long j2);
	GetDigitalStudioVideosById(videoId int64) (entity.LiveVideoData, error)

	// @GET("/content/digital-studio/video/{contentId}/comments")
	// g.b.l<List<CommentInfo>> getDigitalstudioComments(@Path("contentId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetDigitalstudioComments(contentId int64, skip int, take int) ([]entity.CommentInfo, error)

	// @GET("/discover")
	// g.b.l<retrofit2.Response<DiscoverSectionInfo>> getDiscover();
	GetDiscover() (entity.DiscoverSectionInfo, error)

	// @GET("/timeline/{contentId}/info/v2")
	// g.b.l<TimelineContentDataModel> getDiscoverTimelineGift(@Path("contentId") long j2);
	GetDiscoverTimelineGift(contentId int64) (entity.TimelineContentDataModel, error)

	// @GET("/banner/election")
	// g.b.l<List<ElectionBannerInfo>> getElectionBanner();
	GetElectionBanner() ([]entity.ElectionBannerInfo, error)

	// @GET("/events/{year}/{month}")
	// g.b.l<List<ScheduleEvent>> getEvent(@Path("year") int i2, @Path("month") int i3);
	GetEvent(year int, month int) ([]entity.ScheduleEvent, error)

	// @GET("/event/{eventId}")
	// g.b.l<ScheduleEvent> getEventDetails(@Path("eventId") long j2);
	GetEventDetails(eventId int64) (entity.ScheduleEvent, error)

	// @GET("/theater")
	// g.b.l<List<EventTheaterInfo>> getEventTheater();
	GetEventTheater() ([]entity.EventTheaterInfo, error)

	// @GET("/theater/{id}")
	// g.b.l<EventTheaterInfo> getEventTheaterInfo(@Path("id") long j2);
	GetEventTheaterInfo(id int64) (entity.EventTheaterInfo, error)

	// @GET("/games")
	// g.b.l<retrofit2.Response<List<DiscoverGameListItem>>> getGames();
	GetGames() ([]entity.DiscoverGameListItem, error)

	// @GET("/products/gift/section/member-live-chat")
	// g.b.l<List<GiftFloatMessageLiveInfo>> getGiftFloatMessageLiveChat(@NotNull @Query("app") String str);
	GetGiftFloatMessageLiveChat(app string) ([]entity.GiftFloatMessageLiveInfo, error)

	// @GET("/products/gift/section/member-live-chat/{memberId}")
	// g.b.l<List<GiftFloatMessageLiveInfo>> getGiftFloatMessageLiveChatById(@Path("memberId") long j2, @NotNull @Query("app") String str);
	GetGiftFloatMessageLiveChatById(memberId int64, app string) ([]entity.GiftFloatMessageLiveInfo, error)

	// @GET("/batch/{batchId}")
	// g.b.l<SpecialFansDayRankingList> getHistoryRanking(@Path("batchId") long j2);
	GetHistoryRanking(batchId int64) (entity.SpecialFansDayRankingList, error)

	// @GET("/content/member-live/videos")
	// g.b.l<List<LiveVideoData>> getLivePlayback(@Query("skip") int i2, @Query("take") int i3);
	GetLivePlayback(skip int, take int) ([]entity.LiveVideoData, error)

	// @GET("/content/member-live/video/{videoId}")
	// g.b.l<LiveVideoData> getLivePlaybackById(@Path("videoId") long j2);
	GetLivePlaybackById(videoId int64) (entity.LiveVideoData, error)

	// @GET("/content/member-live/video/{videoContentId}/stats")
	// g.b.l<ContentVideoStatsInfo> getLiveVideoStats(@Path("videoContentId") long j2);
	GetLiveVideoStats(videoContentId int64) (entity.ContentVideoStatsInfo, error)

	// @GET("/media/{contentId}")
	// g.b.l<TimelineRecommendMediaModel> getMediaInfo(@Path("contentId") long j2);
	GetMediaInfo(contentId int64) (entity.TimelineRecommendMediaModel, error)

	// @GET("/banner/meeting-you")
	// g.b.l<List<Banner>> getMeetYouBannerList();
	GetMeetYouBannerList() ([]entity.Banner, error)

	// @GET("/banner/{bannerType}")
	// g.b.l<List<Banner>> getMeetYouBannerList(@Path("bannerType") @NotNull String str);
	// Renamed from GetMeetYouBannerList
	GetBannerList(bannerType string) ([]entity.Banner, error)

	// @GET("/banner/meeting-you-last-minute-full")
	// g.b.l<List<Banner>> getMeetYouBannerListLastMinute();
	GetMeetYouBannerListLastMinute() ([]entity.Banner, error)

	// @GET("/member/{memberId}/member-live/{contentId}/recommend")
	// g.b.l<List<LiveVideoData>> getMemberCatchupRecommend(@Path("memberId") long j2, @Path("contentId") long j3, @Query("skip") int i2, @Query("take") int i3);
	GetMemberCatchupRecommend(memberId int64, contentId int64, skip int, take int) ([]entity.LiveVideoData, error)

	// @GET("/member/{memberId}/fans/toprank")
	// g.b.l<List<MembersStatsInfo>> getMemberFansTopRank(@Path("memberId") long j2);
	GetMemberFansTopRank(memberId int64) ([]entity.MembersStatsInfo, error)

	// @GET("/member/{memberId}/fans/toprank/summary")
	// g.b.l<List<MembersStatsInfo>> getMemberFansTopRankSummary(@Path("memberId") long j2);
	GetMemberFansTopRankSummary(memberId int64) ([]entity.MembersStatsInfo, error)

	// @GET("/schedules/member-live")
	// g.b.l<List<ScheduleModelInfo>> getMemberLiveSchedule();
	GetMemberLiveSchedule() ([]entity.ScheduleModelInfo, error)

	// @GET("/member/{memberId}/member-lives")
	// g.b.l<List<LiveVideoData>> getMemberLiveVideo(@Path("memberId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetMemberLiveVideo(memberId int64, skip int, take int) ([]entity.LiveVideoData, error)

	// @GET("/member/{memberId}/mini-videos")
	// g.b.l<List<LiveVideoData>> getMemberMiniVideo(@Path("memberId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetMemberMiniVideo(memberId int64, skip int, take int) ([]entity.LiveVideoData, error)

	// @GET("/member/{memberId}/profile")
	// g.b.l<MemberProfile> getMemberProfile(@Path("memberId") long j2);
	GetMemberProfile(memberId int64) (entity.MemberProfile, error)

	// @GET("/member/{memberId}/stats")
	// g.b.l<MemberProfileStatsInfo> getMemberStats(@Path("memberId") long j2);
	GetMemberStats(memberId int64) (entity.MemberProfileStatsInfo, error)

	// @GET("/member/{memberId}/fans/toprank/today")
	// g.b.l<List<MembersStatsInfo>> getMembersFansTopRankToday(@Path("memberId") long j2);
	GetMembersFansTopRankToday(memberId int64) ([]entity.MembersStatsInfo, error)

	// @GET("/member/{memberId}/fans/toprank/today/summary")
	// g.b.l<List<MembersStatsInfo>> getMembersFansTopRankTodaySummary(@Path("memberId") long j2);
	GetMembersFansTopRankTodaySummary(memberId int64) ([]entity.MembersStatsInfo, error)

	// @GET("/content/mini-video/video/{videoId}")
	// g.b.l<LiveVideoData> getMiniVideoById(@Path("videoId") long j2);
	GetMiniVideoById(videoId int64) (entity.LiveVideoData, error)

	// @GET("/content/mini-video/video/{contentId}/comments")
	// g.b.l<List<CommentInfo>> getMiniVideoComments(@Path("contentId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetMiniVideoComments(contentId int64, skip int, take int) ([]entity.CommentInfo, error)

	// @GET("/content/mini-video/video/{videoContentId}/stats")
	// g.b.l<ContentVideoStatsInfo> getMiniVideoStats(@Path("videoContentId") long j2);
	GetMiniVideoStats(videoContentId int64) (entity.ContentVideoStatsInfo, error)

	// @GET("/products/gift/section/timeline-feed")
	// g.b.l<Skus> getNineCookiesGift();
	GetNineCookiesGift() (entity.Skus, error)

	// @GET("/timeline/only/member/{memberId}")
	// g.b.l<TimelineItemModel> getOnlyTimeline(@Path("memberId") long j2, @Nullable @Query("last_id") Long l2, @Query("amount") long j3);
	GetOnlyTimeline(memberId int64, lastId *int64, amount int64) (entity.TimelineItemModel, error)

	// @GET("/products/coin")
	// g.b.l<List<PurchaseCoins>> getProductCoin(@NotNull @Query("app") String str);
	GetProductCoin(app string) ([]entity.PurchaseCoins, error)

	// @GET("/products/gift")
	// g.b.l<retrofit2.Response<ProductGift>> getProductGift();
	GetProductGift() (entity.ProductGift, error)

	// @GET("/products/music")
	// g.b.l<retrofit2.Response<List<MusicAlbumInfo>>> getProductMusic();
	GetProductMusic() ([]entity.MusicAlbumInfo, error)

	// @GET("/promotion")
	// g.b.l<retrofit2.Response<List<DiscoverPromotionListItem>>> getPromotion();
	GetPromotion() ([]entity.DiscoverPromotionListItem, error)

	// @GET("/addressinfo/province-district")
	// g.b.l<List<ProvinceDistrictResponseInfo>> getProvinceDistrict();
	GetProvinceDistrict() ([]entity.ProvinceDistrictResponseInfo, error)

	// @GET("/api/purchase/gifts")
	// g.b.l<List<PurchaseGifts>> getPurchaseGifts();
	GetPurchaseGifts() ([]entity.PurchaseGifts, error)

	// @GET("/media")
	// g.b.l<List<TimelineRecommendMediaModel>> getRecommendMedia(@Query("skip") int i2, @Query("take") int i3);
	GetRecommendMedia(skip int, take int) ([]entity.TimelineRecommendMediaModel, error)

	// @GET("/shop/payment/list")
	// g.b.l<List<ShopPaymentListResponseInfo>> getShopPaymentList();
	GetShopPaymentList() ([]entity.ShopPaymentListResponseInfo, error)

	// @GET("/shop/product/{productId}")
	// g.b.l<ShopProductInfo> getShopProductById(@Path("productId") long j2);
	GetShopProductById(productId int64) (entity.ShopProductInfo, error)

	// @GET("/shop/product/group")
	// g.b.l<List<ShopProductGroupInfo>> getShopProductGroup(@Query("skip") int i2, @Query("take") int i3);
	GetShopProductGroup(skip int, take int) ([]entity.ShopProductGroupInfo, error)

	// @GET("/_debug/product/group")
	// g.b.l<List<ShopProductGroupInfo>> getShopProductGroupDebug(@NotNull @Query("dk") String str, @Query("skip") int i2, @Query("take") int i3);
	GetShopProductGroupDebug(dk string, skip int, take int) ([]entity.ShopProductGroupInfo, error)

	// @GET("/shop/product/list")
	// g.b.l<List<ShopProductInfo>> getShopProductList(@Query("skip") int i2, @Query("take") int i3);
	GetShopProductList(skip int, take int) ([]entity.ShopProductInfo, error)

	// @GET("/shop/product/list/v2")
	// g.b.l<ShopProductInfoV2> getShopProductListV2(@Query("groupId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetShopProductListV2(groupId int64, skip int, take int) (entity.ShopProductInfoV2, error)

	// @GET("/_debug/product/list/v2")
	// g.b.l<ShopProductInfoV2> getShopProductListV2Debug(@NotNull @Query("dk") String str, @Query("groupId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetShopProductListV2Debug(dk string, groupId int64, skip int, take int) (entity.ShopProductInfoV2, error)

	// @GET("/batch/latest")
	// g.b.l<RankingBatchInfo> getSpecialFanDayWinner();
	GetSpecialFanDayWinner() (entity.RankingBatchInfo, error)

	// @GET("/batch/ranking/{date}")
	// g.b.l<List<SpecialFansDayRankingInfo>> getSpecialFansDayByDate(@Path("date") @NotNull String str);
	GetSpecialFansDayByDate(date string) ([]entity.SpecialFansDayRankingInfo, error)

	// @GET("/batch")
	// g.b.l<retrofit2.Response<List<SpecialFansDayHistoryInfo>>> getSpecialFansDayHistory();
	GetSpecialFansDayHistory() ([]entity.SpecialFansDayHistoryInfo, error)

	// @GET("/batch/current")
	// g.b.l<SpecialFansDayRankingList> getSpecialFansDayRanking();
	GetSpecialFansDayRanking() (entity.SpecialFansDayRankingList, error)

	// @GET("/addressinfo/province-district/{districtId}/sub-district")
	// g.b.l<List<SubDistrictInfo>> getSubDistrict(@Path("districtId") long j2);
	GetSubDistrict(districtId int64) ([]entity.SubDistrictInfo, error)

	// @GET("/content/batch-thankyou/content/{batchId}")
	// g.b.l<ThankYouContentInfo> getThankYouContent(@Path("batchId") long j2);
	GetThankYouContent(batchId int64) (entity.ThankYouContentInfo, error)

	// @GET("/content/theater-show/video/{contentId}/comments")
	// g.b.l<List<CommentInfo>> getTheaterComments(@Path("contentId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetTheaterComments(contentId int64, skip int, take int) ([]entity.CommentInfo, error)

	// @GET("/theater/playback/{contentId}")
	// g.b.l<TheaterPlaybackContentInfo> getTheaterPlaybackByContentId(@Path("contentId") long j2);
	GetTheaterPlaybackByContentId(contentId int64) (entity.TheaterPlaybackContentInfo, error)

	// @GET("/theater/playback")
	// g.b.l<List<TheaterPlaybackInfo>> getTheaterPlaybackList(@Query("skip") int i2, @Query("take") int i3);
	GetTheaterPlaybackList(skip int, take int) ([]entity.TheaterPlaybackInfo, error)

	// @GET("/content/theater-show/video/{videoContentId}/stats")
	// g.b.l<ContentVideoStatsInfo> getTheaterVideoStats(@Query("skip") int i2, @Query("take") int i3);
	GetTheaterVideoStats(skip int, take int) (entity.ContentVideoStatsInfo, error)

	// @GET("/content/theater/videos")
	// g.b.l<List<LiveVideoData>> getTheaterVideos(@Query("skip") int i2, @Query("take") int i3);
	GetTheaterVideos(skip int, take int) ([]entity.LiveVideoData, error)

	// @GET("/content/theater/video/{videoId}")
	// g.b.l<LiveVideoData> getTheaterVideosById(@Path("videoId") long j2);
	GetTheaterVideosById(videoId int64) (entity.LiveVideoData, error)

	// @GET("/theme")
	// g.b.l<ThemeInfo> getTheme();
	GetTheme() (entity.ThemeInfo, error)

	// @GET("/timeline/multi/member/{memberId}")
	// g.b.l<TimelineItemModel> getTimelineFeedById(@Path("memberId") @NotNull String str, @Query("amount") long j2, @Nullable @Query("last_id") Long l2);
	GetTimelineFeedById(memberId string, amount int64, lastId *int64) (entity.TimelineItemModel, error)

	// @GET("/timeline/{contentId}/rank/gifts")
	// g.b.l<List<TimelineItemGiftStats>> getTimelineGiftStats(@Path("contentId") long j2);
	GetTimelineGiftStats(contentId int64) ([]entity.TimelineItemGiftStats, error)

	// @GET("/timeline/{itemType}/{itemId}")
	// g.b.l<TimelineFeeds> getTimelineItem(@Path("itemType") @NotNull String str, @Path("itemId") long j2);
	GetTimelineItem(itemType string, itemId int64) (entity.TimelineFeeds, error)

	// @GET("/user/{userId}/liked/videos")
	// g.b.l<List<LiveVideoData>> getUserLikedVideos(@Path("userId") long j2, @Query("skip") int i2, @Query("take") int i3);
	GetUserLikedVideos(userId int64, skip int, take int) ([]entity.LiveVideoData, error)

	// @GET("/user/{userId}/follows/members")
	// g.b.l<List<FollowMemberInfo>> getUserMemeberFollow(@Path("userId") long j2);
	GetUserMemeberFollow(userId int64) ([]entity.FollowMemberInfo, error)

	// @GET("/user/{userId}/profile")
	// g.b.l<PublicUserProfileInfo> getUserProfile(@Path("userId") long j2);
	// Renamed to GetPublicUserProfile
	GetPublicUserProfile(userId int64) (entity.PublicUserProfileInfo, error)

	// @GET("/schedules/member-live/{scheduleId}")
	// g.b.l<ScheduleModelInfo> getWaitingForLiveSchedule(@Path("scheduleId") @Nullable String str);
	GetWaitingForLiveSchedule(scheduleId *string) (entity.ScheduleModelInfo, error)

	// @GET("/version")
	// g.b.l<VersionInfo> version();
	Version() (entity.VersionInfo, error)
}

package theater

import "github.com/dvwzj/iam48-go/api/entity"

type TheaterRepository interface {
	// @GET("/theater-live/{liveId}/stream/{streamId}")
	// g.b.l<LiveStreamInfo> getStream(@Path("liveId") long j2, @Path("streamId") long j3);
	GetStream(liveId int64, streamId int64) (entity.LiveStreamInfo, error)

	// @GET("/live/{liveId}/watch")
	// g.b.l<TheaterWatchInfo> getTheaterWatch(@Path("liveId") int i2);
	GetTheaterWatch(liveId int) (entity.TheaterWatchInfo, error)

	// @POST("/user/{userId}/theater-live/{liveId}/members/all")
	// g.b.l<TheaterGiftInfo> sendTheaterGiftAllMembers(@Path("userId") long j2, @Path("liveId") long j3);
	SendTheaterGiftAllMembers(userId int64, liveId int64) (entity.TheaterGiftInfo, error)

	// @POST("/user/{userId}/theater-live/{liveId}/members")
	// g.b.l<TheaterGiftChooseMembersInfo> sendTheaterGiftChooseMembers(@Path("userId") long j2, @Path("liveId") long j3);
	SendTheaterGiftChooseMembers(userId int64, liveId int64) (entity.TheaterGiftChooseMembersInfo, error)
}

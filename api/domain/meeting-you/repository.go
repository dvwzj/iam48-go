package meetingyou

import (
	"github.com/dvwzj/iam48-go/api/entity"
	"github.com/go-resty/resty/v2"
)

type MeetingYouRepository interface {
	// @POST("/channel/user/join/{memberId}")
	// g.b.l<JoinLobbyInfo> joinLobby(@Path("memberId") @Nullable Long l2);
	JoinLobby(memberId int64) (entity.JoinLobbyInfo, error)

	// @POST("/review/submit")
	// g.b.l<f0> submitReview(@Body @NotNull MeetingReviewRequest meetingReviewRequest);
	SubmitReview(meetingReviewRequest entity.MeetingReviewRequest) (resty.Response, error)
}

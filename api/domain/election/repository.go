package election

import (
	"github.com/dvwzj/iam48-go/api/entity"
	"github.com/go-resty/resty/v2"
)

type ElectionRepository interface {
	// @POST("/claim/v2/shopee")
	// g.b.l<ElectionClaimResponseInfo> claimShopee(@Body @NotNull ElectionClaimBodyInfo electionClaimBodyInfo);
	ClaimShopee(electionClaimBodyInfo entity.ElectionClaimBodyInfo) (entity.ElectionClaimResponseInfo, error)

	// @GET("/code/{codeSetId}/all")
	// g.b.l<ElectionAllCodeSetResponseInfo> getAllElectionCodeById(@Path("codeSetId") long j2);
	GetAllElectionCodeById(codeSetId int64) (entity.ElectionAllCodeSetResponseInfo, error)

	// @GET("/code")
	// g.b.l<List<ElectionCodeResponseInfo>> getElectionCode(@Query("skip") int i2, @Query("take") int i3);
	GetElectionCode(skip int, take int) ([]entity.ElectionCodeResponseInfo, error)

	// @GET("/code/{codeSetId}")
	// g.b.l<List<ElectionCodeSetByIdInfo>> getElectionCodeById(@Path("codeSetId") long j2, @Query("latestCodeId") long j3);
	GetElectionCodeById(codeSetId int64, latestCodeId int64) ([]entity.ElectionCodeSetByIdInfo, error)

	// @GET("/history")
	// g.b.l<List<ElectionHistoryResponseInfo>> getElectionHistory(@Query("skip") int i2, @Query("take") int i3);
	GetElectionHistory(skip int, take int) ([]entity.ElectionHistoryResponseInfo, error)

	// @POST("/code/sent")
	// g.b.l<f0> sentCodeToEmail(@Body @NotNull ElectionSentCodeToEmailBody electionSentCodeToEmailBody);
	SentCodeToEmail(electionSentCodeToEmailBody entity.ElectionSentCodeToEmailBody) (resty.Response, error)
}

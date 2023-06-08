package pdpa

import "github.com/dvwzj/iam48-go/api/entity"

type PdpaRepository interface {
	// @GET("/{appCode}/consent/{ookbeeNumericId}/check")
	// g.b.l<PdpaConsentStatusResponseInfo> checkConsentStatus(@Path("appCode") @NotNull String str, @Path("ookbeeNumericId") long j2);
	CheckConsentStatus(appCode string, ookbeeNumericId int64) (entity.PdpaConsentStatusResponseInfo, error)

	// @GET("/{appCode}/consent/{ookbeeNumericId}/currentversion")
	// g.b.l<PdpaCurrentVersionResponseInfo> checkCurrentVersionPDPA(@Path("appCode") @NotNull String str, @Path("ookbeeNumericId") long j2);
	CheckCurrentVersionPDPA(appCode string, ookbeeNumericId int64) (entity.PdpaCurrentVersionResponseInfo, error)
}

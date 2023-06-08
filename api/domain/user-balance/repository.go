package userbalance

import "github.com/dvwzj/iam48-go/api/entity"

type UserBalanceRepository interface {
	// @GET("/app/{appCode}/user/{userId}/balance")
	// g.b.l<CoreBalance> getBalance(@Path("appCode") @NotNull String str, @Path("userId") int i2);
	GetBalance(appCode string, userId int) (entity.CoreBalance, error)
}

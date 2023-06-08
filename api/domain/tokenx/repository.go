package tokenx

import "github.com/go-resty/resty/v2"

type TokenXRepository interface {
	// @GET("/access-denied")
	// g.b.l<f0> accessDenied();
	AccessDenied() (resty.Response, error)
}

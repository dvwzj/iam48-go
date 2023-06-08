# iam48-go
iAM48 in golang

## WIP

Example
```go
import (
    "log"

    "github.com/dvwzj/iam48-go/utils/json"
    "github.com/dvwzj/iam48-go/session"
)

publicSession, err := session.NewSession()
if err != nil {
    log.Fatal(err)
}
authenticatedSession, err := publicSession.LoginWithEmail(EMAIL, PASSWORD)
if err != nil {
    log.Fatal(err)
}
userProfile, err := authenticatedSession.Service.GetUserProfile(*authenticatedSession.User.Id)
if err != nil {
    log.Fatal(err)
}
b, err := json.Marshal(userProfile)
if err != nil {
    log.Fatal(err)
}
log.Println(string(b))
```
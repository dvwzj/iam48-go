package session

import (
	"os"
	"testing"

	"github.com/dvwzj/iam48-go/utils/json"
	"github.com/joho/godotenv"
)

var (
	EMAIL    = os.Getenv("IAM48_EMAIL")
	PASSWORD = os.Getenv("IAM48_PASSWORD")
)

func init() {
	godotenv.Load("../.env")
	EMAIL = os.Getenv("IAM48_EMAIL")
	PASSWORD = os.Getenv("IAM48_PASSWORD")
}

func TestSession(t *testing.T) {
	publicSession, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	authenticatedSession, err := publicSession.LoginWithEmail(EMAIL, PASSWORD)
	if err != nil {
		t.Fatal(err)
	}
	userProfile, err := authenticatedSession.Service.GetUserProfile(*authenticatedSession.User.Id)
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(userProfile)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}

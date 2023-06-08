package session

import (
	"errors"

	"github.com/dvwzj/iam48-go/api/domain/election"
	meetingyou "github.com/dvwzj/iam48-go/api/domain/meeting-you"
	"github.com/dvwzj/iam48-go/api/domain/pdpa"
	"github.com/dvwzj/iam48-go/api/domain/public"
	"github.com/dvwzj/iam48-go/api/domain/streaming"
	"github.com/dvwzj/iam48-go/api/domain/theater"
	"github.com/dvwzj/iam48-go/api/domain/ticket"
	"github.com/dvwzj/iam48-go/api/domain/tokenx"
	"github.com/dvwzj/iam48-go/api/domain/user"
	userbalance "github.com/dvwzj/iam48-go/api/domain/user-balance"
	"github.com/dvwzj/iam48-go/api/entity"
	"github.com/dvwzj/iam48-go/api/services"
	"github.com/dvwzj/iam48-go/client"
	"github.com/go-resty/resty/v2"
)

type Authorization interface {
	PublicService | AuthenticatedService
	get() any
}

type PublicSession interface {
	public.PublicRepository
}

type PublicService struct {
	*services.PublicService
}

func (p PublicService) get() any {
	return p
}

type AuthenticatedSession interface {
	public.PublicRepository
	election.ElectionRepository
	meetingyou.MeetingYouRepository
	pdpa.PdpaRepository
	streaming.StreamingRepository
	theater.TheaterRepository
	ticket.TicketRepository
	tokenx.TokenXRepository
	userbalance.UserBalanceRepository
	user.UserRepository
}

type AuthenticatedService struct {
	*PublicService
	*services.ElectionService
	*services.MeetingYouService
	*services.PdpaService
	*services.StreamingService
	*services.TheaterService
	*services.TicketService
	*services.TokenXService
	*services.UserBalanceService
	*services.UserService
}

func (a AuthenticatedService) get() any {
	return a
}

type Session[T Authorization] struct {
	AuthEmail    *entity.LoginInfo
	AuthFacebook *entity.LoginFacebookBody
	User         *entity.AuthenResponseInfo
	Service      T
}

func (s *Session[T]) LoginWithEmail(email, password string) (Session[AuthenticatedService], error) {
	loginInfo := entity.LoginInfo{
		Email:    &email,
		Password: &password,
	}
	if public, ok := s.Service.get().(PublicService); !ok {
		return Session[AuthenticatedService]{}, errors.New("session is not Public")
	} else {
		authenticated, err := newAuthenticated(&public)
		if err != nil {
			return Session[AuthenticatedService]{}, err
		}
		user, err := authenticated.(AuthenticatedService).UserService.LoginWithEmail(loginInfo)
		if err != nil {
			return Session[AuthenticatedService]{}, err
		}
		authenticated.(AuthenticatedService).PublicService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).ElectionService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).MeetingYouService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).PdpaService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).StreamingService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).TheaterService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).TicketService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).TokenXService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).UserBalanceService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).UserService.Client.SetAuthToken(*user.Token)
		return Session[AuthenticatedService]{
			AuthEmail: &loginInfo,
			User:      &user,
			Service:   authenticated.(AuthenticatedService),
		}, nil
	}
}

func (s *Session[T]) LoginWithFacebook(email, name, facebookId, accessToken string) (Session[AuthenticatedService], error) {
	loginFacebookBody := entity.LoginFacebookBody{
		Email:       &email,
		Name:        &name,
		FacebookId:  &facebookId,
		AccessToken: &accessToken,
	}
	if public, ok := s.Service.get().(PublicService); !ok {
		return Session[AuthenticatedService]{}, errors.New("session is not Public")
	} else {
		authenticated, err := newAuthenticated(&public)
		if err != nil {
			return Session[AuthenticatedService]{}, err
		}
		user, err := authenticated.(AuthenticatedService).UserService.LoginFacebook(loginFacebookBody)
		if err != nil {
			return Session[AuthenticatedService]{}, err
		}
		authenticated.(AuthenticatedService).PublicService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).ElectionService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).MeetingYouService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).PdpaService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).StreamingService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).TheaterService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).TicketService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).TokenXService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).UserBalanceService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).UserService.Client.SetAuthToken(*user.Token)
		return Session[AuthenticatedService]{
			AuthFacebook: &loginFacebookBody,
			User:         &user,
			Service:      authenticated.(AuthenticatedService),
		}, nil
	}
}

func (s *Session[T]) Logout() (resty.Response, error) {
	if s.User == nil {
		return resty.Response{}, errors.New("user is nil")
	}
	if authenticated, ok := s.Service.get().(AuthenticatedService); !ok {
		return resty.Response{}, errors.New("session is not Authenticated")
	} else {
		logoutDeviceBody := entity.LogoutDeviceBody{
			DeviceId: &client.DeviceId,
		}
		res, err := authenticated.UserService.LogoutDevice(*s.User.Id, logoutDeviceBody)
		if err != nil {
			return resty.Response{}, err
		}
		return res, nil
	}
}

func newPublic() (PublicSession, error) {
	publicService, err := services.NewPublicService()
	if err != nil {
		return PublicService{}, err
	}

	return PublicService{
		PublicService: publicService,
	}, nil
}

func newAuthenticated(publicService *PublicService) (AuthenticatedSession, error) {
	electionService, err := services.NewElectionService()
	if err != nil {
		return AuthenticatedService{}, err
	}
	meetingYouService, err := services.NewMeetingYouService()
	if err != nil {
		return AuthenticatedService{}, err
	}
	pdpaService, err := services.NewPdpaService()
	if err != nil {
		return AuthenticatedService{}, err
	}
	streamingService, err := services.NewStreamingService()
	if err != nil {
		return AuthenticatedService{}, err
	}
	theaterService, err := services.NewTheaterService()
	if err != nil {
		return AuthenticatedService{}, err
	}
	ticketService, err := services.NewTicketService()
	if err != nil {
		return AuthenticatedService{}, err
	}
	tokenXService, err := services.NewTokenXService()
	if err != nil {
		return AuthenticatedService{}, err
	}
	userBalanceService, err := services.NewUserBalanceService()
	if err != nil {
		return AuthenticatedService{}, err
	}
	userService, err := services.NewUserService()
	if err != nil {
		return AuthenticatedService{}, err
	}
	return AuthenticatedService{
		PublicService:      publicService,
		ElectionService:    electionService,
		MeetingYouService:  meetingYouService,
		PdpaService:        pdpaService,
		StreamingService:   streamingService,
		TheaterService:     theaterService,
		TicketService:      ticketService,
		TokenXService:      tokenXService,
		UserBalanceService: userBalanceService,
		UserService:        userService,
	}, nil
}

func NewSession() (Session[PublicService], error) {
	public, err := newPublic()
	if err != nil {
		return Session[PublicService]{}, err
	}
	return Session[PublicService]{
		Service: public.(PublicService),
	}, nil
}

func NewSessionWithEmail(email, password string) (Session[AuthenticatedService], error) {
	publicSession, err := NewSession()
	if err != nil {
		return Session[AuthenticatedService]{}, err
	}
	authenticatedSession, err := publicSession.LoginWithEmail(email, password)
	if err != nil {
		return Session[AuthenticatedService]{}, err
	}
	return authenticatedSession, nil
}

func NewSessionWithFacebook(email, name, facebookId, accessToken string) (Session[AuthenticatedService], error) {
	publicSession, err := NewSession()
	if err != nil {
		return Session[AuthenticatedService]{}, err
	}
	authenticatedSession, err := publicSession.LoginWithFacebook(email, name, facebookId, accessToken)
	if err != nil {
		return Session[AuthenticatedService]{}, err
	}
	return authenticatedSession, nil
}

func NewSessionWithUser(user entity.AuthenResponseInfo) (Session[AuthenticatedService], error) {
	if user.Id == nil {
		return Session[AuthenticatedService]{}, errors.New("user id is nil")
	}
	if user.Token == nil {
		return Session[AuthenticatedService]{}, errors.New("user token is nil")
	}
	if user.RefreshToken == nil {
		return Session[AuthenticatedService]{}, errors.New("user refresh token is nil")
	}
	public, err := newPublic()
	if err != nil {
		return Session[AuthenticatedService]{}, err
	}
	if publicService, ok := public.(PublicService); ok {
		authenticated, err := newAuthenticated(&publicService)
		if err != nil {
			return Session[AuthenticatedService]{}, err
		}
		authenticated.(AuthenticatedService).PublicService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).ElectionService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).MeetingYouService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).PdpaService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).StreamingService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).TheaterService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).TicketService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).TokenXService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).UserBalanceService.Client.SetAuthToken(*user.Token)
		authenticated.(AuthenticatedService).UserService.Client.SetAuthToken(*user.Token)
		return Session[AuthenticatedService]{
			User:    &user,
			Service: authenticated.(AuthenticatedService),
		}, nil
	}
	return Session[AuthenticatedService]{}, errors.New("something went wrong")
}

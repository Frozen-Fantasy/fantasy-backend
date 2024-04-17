package service

import (
	"context"
	"github.com/Frozen-Fantasy/fantasy-backend.git/config"
	"github.com/Frozen-Fantasy/fantasy-backend.git/pkg/models/players"
	"github.com/Frozen-Fantasy/fantasy-backend.git/pkg/models/store"
	"github.com/Frozen-Fantasy/fantasy-backend.git/pkg/models/tournaments"
	"github.com/Frozen-Fantasy/fantasy-backend.git/pkg/models/user"
	"github.com/Frozen-Fantasy/fantasy-backend.git/pkg/storage"
	"github.com/google/uuid"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type User interface {
	SignUp(input user.SignUpInput) error
	SignIn(input user.SignInInput) (user.Tokens, error)
	RefreshTokens(refreshTokenID string) (user.Tokens, error)
	CreateSession(userID uuid.UUID) (user.Tokens, error)
	Logout(refreshTokenID string) error
	SendVerificationCode(email string) error
	CheckEmailVerification(email string, inputCode int) error
	CheckEmailExists(email string) (bool, error)
	CheckNicknameExists(nickname string) (bool, error)
	ChangePassword(inp user.ChangePasswordModel) error
	ForgotPassword(email string) error
	ResetPassword(inp user.ResetPasswordInput) error
	GetUserInfo(userID uuid.UUID) (user.UserInfoModel, error)
	CheckUserDataExists(inp user.UserExistsDataInput) error
	DeleteProfile(userID uuid.UUID) error
	GetCoinTransactions(profileID uuid.UUID) ([]user.CoinTransactionsModel, error)
}

type TokenManager interface {
	CreateJWT(userID string) (int64, string, error)
	ParseJWT(accessToken string) (string, error)
	CreateRefreshToken() (string, error)
}

type Teams interface {
	CreateTeamsNHL(context.Context, []tournaments.Standing) error
	CreateTeamsKHL(ctx context.Context, teams []tournaments.TeamKHL) error
	GetMatchesDay(ctx context.Context, league tournaments.League) ([]tournaments.Matches, error)
	GetTournaments(ctx context.Context, league tournaments.League) ([]tournaments.Tournament, error)
	GetRosterByTournamentID(tournamentID int) (players.TournamentRosterResponse, error)
}

type Store interface {
	GetAllProducts() ([]store.Product, error)
	BuyProduct(buy store.BuyProductModel) error
}

type Players interface {
	CreatePlayers(playersData []players.Player) error
	GetPlayers(playersFilter players.PlayersFilter) ([]players.PlayerResponse, error)
	GetPlayerCards(filter players.PlayerCardsFilter) ([]players.PlayerCardResponse, error)
	CardUnpacking(id int, userID uuid.UUID) error
}

type Services struct {
	User
	TokenManager
	Teams
	Store
	Players
}

type Deps struct {
	Cfg      config.ServiceConfiguration
	Storage  *storage.PostgresStorage
	RStorage *storage.RedisStorage
	Jwt      *Manager
}

func NewServices(deps Deps) *Services {
	userService := NewUserService(deps.Storage, deps.RStorage, deps.Jwt, deps.Cfg)
	storeService := NewStoreService(deps.Storage)
	playersService := NewPlayersService(deps.Storage)
	teamsService := NewTeamsService(deps.Storage, playersService)
	return &Services{
		User:         userService,
		TokenManager: deps.Jwt,
		Teams:        teamsService,
		Store:        storeService,
		Players:      playersService,
	}
}

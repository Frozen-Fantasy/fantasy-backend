// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	tournaments "github.com/Frozen-Fantasy/fantasy-backend.git/pkg/models/tournaments"
	user "github.com/Frozen-Fantasy/fantasy-backend.git/pkg/models/user"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// ChangePassword mocks base method.
func (m *MockUser) ChangePassword(inp user.ChangePasswordModel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserMockRecorder) ChangePassword(inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUser)(nil).ChangePassword), inp)
}

// CheckEmailExists mocks base method.
func (m *MockUser) CheckEmailExists(email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmailExists", email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckEmailExists indicates an expected call of CheckEmailExists.
func (mr *MockUserMockRecorder) CheckEmailExists(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmailExists", reflect.TypeOf((*MockUser)(nil).CheckEmailExists), email)
}

// CheckEmailVerification mocks base method.
func (m *MockUser) CheckEmailVerification(email string, inputCode int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmailVerification", email, inputCode)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckEmailVerification indicates an expected call of CheckEmailVerification.
func (mr *MockUserMockRecorder) CheckEmailVerification(email, inputCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmailVerification", reflect.TypeOf((*MockUser)(nil).CheckEmailVerification), email, inputCode)
}

// CheckNicknameExists mocks base method.
func (m *MockUser) CheckNicknameExists(nickname string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckNicknameExists", nickname)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckNicknameExists indicates an expected call of CheckNicknameExists.
func (mr *MockUserMockRecorder) CheckNicknameExists(nickname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckNicknameExists", reflect.TypeOf((*MockUser)(nil).CheckNicknameExists), nickname)
}

// CheckUserDataExists mocks base method.
func (m *MockUser) CheckUserDataExists(inp user.UserExistsDataInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserDataExists", inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckUserDataExists indicates an expected call of CheckUserDataExists.
func (mr *MockUserMockRecorder) CheckUserDataExists(inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserDataExists", reflect.TypeOf((*MockUser)(nil).CheckUserDataExists), inp)
}

// CreateSession mocks base method.
func (m *MockUser) CreateSession(userID uuid.UUID) (user.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", userID)
	ret0, _ := ret[0].(user.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockUserMockRecorder) CreateSession(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockUser)(nil).CreateSession), userID)
}

// DeleteProfile mocks base method.
func (m *MockUser) DeleteProfile(userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProfile", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProfile indicates an expected call of DeleteProfile.
func (mr *MockUserMockRecorder) DeleteProfile(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProfile", reflect.TypeOf((*MockUser)(nil).DeleteProfile), userID)
}

// ForgotPassword mocks base method.
func (m *MockUser) ForgotPassword(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForgotPassword", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForgotPassword indicates an expected call of ForgotPassword.
func (mr *MockUserMockRecorder) ForgotPassword(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForgotPassword", reflect.TypeOf((*MockUser)(nil).ForgotPassword), email)
}

// GetCoinTransactions mocks base method.
func (m *MockUser) GetCoinTransactions(profileID uuid.UUID) ([]user.CoinTransactionsModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCoinTransactions", profileID)
	ret0, _ := ret[0].([]user.CoinTransactionsModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCoinTransactions indicates an expected call of GetCoinTransactions.
func (mr *MockUserMockRecorder) GetCoinTransactions(profileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCoinTransactions", reflect.TypeOf((*MockUser)(nil).GetCoinTransactions), profileID)
}

// GetUserInfo mocks base method.
func (m *MockUser) GetUserInfo(userID uuid.UUID) (user.UserInfoModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", userID)
	ret0, _ := ret[0].(user.UserInfoModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfo indicates an expected call of GetUserInfo.
func (mr *MockUserMockRecorder) GetUserInfo(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockUser)(nil).GetUserInfo), userID)
}

// Logout mocks base method.
func (m *MockUser) Logout(refreshTokenID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", refreshTokenID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockUserMockRecorder) Logout(refreshTokenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockUser)(nil).Logout), refreshTokenID)
}

// RefreshTokens mocks base method.
func (m *MockUser) RefreshTokens(refreshTokenID string) (user.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTokens", refreshTokenID)
	ret0, _ := ret[0].(user.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshTokens indicates an expected call of RefreshTokens.
func (mr *MockUserMockRecorder) RefreshTokens(refreshTokenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTokens", reflect.TypeOf((*MockUser)(nil).RefreshTokens), refreshTokenID)
}

// ResetPassword mocks base method.
func (m *MockUser) ResetPassword(inp user.ResetPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPassword", inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetPassword indicates an expected call of ResetPassword.
func (mr *MockUserMockRecorder) ResetPassword(inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*MockUser)(nil).ResetPassword), inp)
}

// SendVerificationCode mocks base method.
func (m *MockUser) SendVerificationCode(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendVerificationCode", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendVerificationCode indicates an expected call of SendVerificationCode.
func (mr *MockUserMockRecorder) SendVerificationCode(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendVerificationCode", reflect.TypeOf((*MockUser)(nil).SendVerificationCode), email)
}

// SignIn mocks base method.
func (m *MockUser) SignIn(input user.SignInInput) (user.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", input)
	ret0, _ := ret[0].(user.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUserMockRecorder) SignIn(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUser)(nil).SignIn), input)
}

// SignUp mocks base method.
func (m *MockUser) SignUp(input user.SignUpInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUserMockRecorder) SignUp(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUser)(nil).SignUp), input)
}

// MockTokenManager is a mock of TokenManager interface.
type MockTokenManager struct {
	ctrl     *gomock.Controller
	recorder *MockTokenManagerMockRecorder
}

// MockTokenManagerMockRecorder is the mock recorder for MockTokenManager.
type MockTokenManagerMockRecorder struct {
	mock *MockTokenManager
}

// NewMockTokenManager creates a new mock instance.
func NewMockTokenManager(ctrl *gomock.Controller) *MockTokenManager {
	mock := &MockTokenManager{ctrl: ctrl}
	mock.recorder = &MockTokenManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenManager) EXPECT() *MockTokenManagerMockRecorder {
	return m.recorder
}

// CreateJWT mocks base method.
func (m *MockTokenManager) CreateJWT(userID string) (int64, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJWT", userID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateJWT indicates an expected call of CreateJWT.
func (mr *MockTokenManagerMockRecorder) CreateJWT(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJWT", reflect.TypeOf((*MockTokenManager)(nil).CreateJWT), userID)
}

// CreateRefreshToken mocks base method.
func (m *MockTokenManager) CreateRefreshToken() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRefreshToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRefreshToken indicates an expected call of CreateRefreshToken.
func (mr *MockTokenManagerMockRecorder) CreateRefreshToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRefreshToken", reflect.TypeOf((*MockTokenManager)(nil).CreateRefreshToken))
}

// ParseJWT mocks base method.
func (m *MockTokenManager) ParseJWT(accessToken string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseJWT", accessToken)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseJWT indicates an expected call of ParseJWT.
func (mr *MockTokenManagerMockRecorder) ParseJWT(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseJWT", reflect.TypeOf((*MockTokenManager)(nil).ParseJWT), accessToken)
}

// MockTeams is a mock of Teams interface.
type MockTeams struct {
	ctrl     *gomock.Controller
	recorder *MockTeamsMockRecorder
}

// MockTeamsMockRecorder is the mock recorder for MockTeams.
type MockTeamsMockRecorder struct {
	mock *MockTeams
}

// NewMockTeams creates a new mock instance.
func NewMockTeams(ctrl *gomock.Controller) *MockTeams {
	mock := &MockTeams{ctrl: ctrl}
	mock.recorder = &MockTeamsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeams) EXPECT() *MockTeamsMockRecorder {
	return m.recorder
}

// AddEventsKHL mocks base method.
func (m *MockTeams) AddEventsKHL(ctx context.Context, events []tournaments.EventDataKHL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEventsKHL", ctx, events)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventsKHL indicates an expected call of AddEventsKHL.
func (mr *MockTeamsMockRecorder) AddEventsKHL(ctx, events interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventsKHL", reflect.TypeOf((*MockTeams)(nil).AddEventsKHL), ctx, events)
}

// AddEventsNHL mocks base method.
func (m *MockTeams) AddEventsNHL(ctx context.Context, events []tournaments.Game) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEventsNHL", ctx, events)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventsNHL indicates an expected call of AddEventsNHL.
func (mr *MockTeamsMockRecorder) AddEventsNHL(ctx, events interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventsNHL", reflect.TypeOf((*MockTeams)(nil).AddEventsNHL), ctx, events)
}

// CreateTeamsKHL mocks base method.
func (m *MockTeams) CreateTeamsKHL(ctx context.Context, teams []tournaments.TeamKHL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeamsKHL", ctx, teams)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTeamsKHL indicates an expected call of CreateTeamsKHL.
func (mr *MockTeamsMockRecorder) CreateTeamsKHL(ctx, teams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeamsKHL", reflect.TypeOf((*MockTeams)(nil).CreateTeamsKHL), ctx, teams)
}

// CreateTeamsNHL mocks base method.
func (m *MockTeams) CreateTeamsNHL(arg0 context.Context, arg1 []tournaments.Standing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeamsNHL", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTeamsNHL indicates an expected call of CreateTeamsNHL.
func (mr *MockTeamsMockRecorder) CreateTeamsNHL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeamsNHL", reflect.TypeOf((*MockTeams)(nil).CreateTeamsNHL), arg0, arg1)
}

// CreateTournaments mocks base method.
func (m *MockTeams) CreateTournaments(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTournaments", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTournaments indicates an expected call of CreateTournaments.
func (mr *MockTeamsMockRecorder) CreateTournaments(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTournaments", reflect.TypeOf((*MockTeams)(nil).CreateTournaments), ctx)
}

// GetMatchesDay mocks base method.
func (m *MockTeams) GetMatchesDay(ctx context.Context, league tournaments.League) ([]tournaments.Matches, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchesDay", ctx, league)
	ret0, _ := ret[0].([]tournaments.Matches)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchesDay indicates an expected call of GetMatchesDay.
func (mr *MockTeamsMockRecorder) GetMatchesDay(ctx, league interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchesDay", reflect.TypeOf((*MockTeams)(nil).GetMatchesDay), ctx, league)
}

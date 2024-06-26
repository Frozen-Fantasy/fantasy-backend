package api

import (
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/Frozen-Fantasy/fantasy-backend.git/pkg/models/players"
	"github.com/Frozen-Fantasy/fantasy-backend.git/pkg/models/tournaments"
	"github.com/Frozen-Fantasy/fantasy-backend.git/pkg/service"
	"github.com/Frozen-Fantasy/fantasy-backend.git/pkg/storage"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// CreateTeamsNHL godoc
// @Summary Создание команд NHL
// @Security ApiKeyAuth
// @Schemes
// @Description Добавлят информацию о команде NHL
// @Tags tournament
// @Produce json
// @Success 200
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Router /tournament/create_team_nhl [get]
func (api *Api) CreateTeamsNHL(ctx *gin.Context) {
	_, err := parseUserIDFromContext(ctx)
	if err != nil {
		log.Println("CreateTeamsNHL:", err)
		return
	}

	url := "https://api-web.nhle.com/v1/standings/now"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("CreateTeamsNHL:", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("CreateTeamsNHL:", err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	var standings tournaments.StandingsResponse

	err = decoder.Decode(&standings)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		ctx.JSON(http.StatusBadRequest, getInternalServerError())
		return
	}
	for idT, _ := range standings.Standings {
		standings.Standings[idT].League = tournaments.NHL
	}

	err = api.services.Teams.CreateTeamsNHL(ctx, standings.Standings)
	if err != nil {
		log.Printf("CreateTeamsNHL: %v", err)
		ctx.JSON(http.StatusBadRequest, getInternalServerError())
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

// CreateTeamsKHL godoc
// @Summary Создание команд KHL
// @Security ApiKeyAuth
// @Schemes
// @Description Добавлят информацию о команде KHL
// @Tags tournament
// @Produce json
// @Success 200
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Router /tournament/create_team_khl [get]
func (api *Api) CreateTeamsKHL(ctx *gin.Context) {
	_, err := parseUserIDFromContext(ctx)
	if err != nil {
		log.Println("CreateTeamsKHL:", err)
		return
	}

	url := "https://khl.api.webcaster.pro/api/khl_mobile/teams_v2"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("CreateTeamsKHL:", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("CreateTeamsKHL:", err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	var teamKHL []tournaments.TeamKHL

	err = decoder.Decode(&teamKHL)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		ctx.JSON(http.StatusBadRequest, getInternalServerError())
		return
	}
	for idT, _ := range teamKHL {
		teamKHL[idT].Team.League = tournaments.KHL
		teamKHL[idT].Team.TeamAbbrev = tournaments.KHLAbrev[teamKHL[idT].Team.TeamName]
	}

	err = api.services.Teams.CreateTeamsKHL(ctx, teamKHL)
	if err != nil {
		log.Printf("CreateTeamKHL: %v", err)
		ctx.JSON(http.StatusBadRequest, getInternalServerError())
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

// GetMatches godoc
// @Summary Получение матчей на следующий день
// @Schemes
// @Description Дата берётся автоматически
// @Tags tournament
// @Produce json
// @Success 200 {object} []tournaments.Matches
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Param league path string true "league" Enums(NHL, KHL)
// @Router /tournament/get_matches/{league} [get]
func (api *Api) GetMatches(ctx *gin.Context) {
	//var leagueName tournaments.League
	//var leagueName string
	leagueName := ctx.Param("league")
	if leagueName == "" {
		ctx.JSON(http.StatusBadRequest, getBadRequestError(errors.New("empty league name")))
		return
	}

	league := new(tournaments.League)
	*league = league.GetLeagueId(leagueName)
	matches, err := api.services.Teams.GetMatchesDay(ctx, *league)
	if err != nil {
		log.Println("GetMatches:", err)
		switch err {
		case service.NotFoundMatches:
			ctx.JSON(http.StatusNotFound, getNotFoundError())
			return
		default:
			ctx.JSON(http.StatusInternalServerError, getInternalServerError())
			return
		}
	}

	ctx.JSON(http.StatusOK, matches)
}

// GetTournaments godoc
// @Summary Получение турниров на ближайшие 2 дня
// @Security ApiKeyAuth
// @Schemes
// @Description Дата берётся автоматически
// @Tags tournament
// @Accept json
// @Produce json
// @Success 200 {object} []tournaments.Tournament
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 404 {object} Error
// @Param league path string true "league" Enums(NHL, KHL, Both)
// @Router /tournament/get_tournaments/{league} [get]
func (api *Api) GetTournaments(ctx *gin.Context) {

	_, err := parseUserIDFromContext(ctx)
	if err != nil {
		log.Println("GetTournamentTeam:", err)
		return
	}

	leagueName := ctx.Param("league")
	if leagueName == "" {
		ctx.JSON(http.StatusBadRequest, getBadRequestError(errors.New("empty league name")))
		return
	}

	league := new(tournaments.League)
	*league = league.GetLeagueId(leagueName)
	tournamentsInfo, err := api.services.Tournaments.GetTournaments(ctx, *league)
	if errors.Is(err, service.NotFoundTournaments) {
		log.Printf("GetTournaments: %v", err)
		ctx.JSON(http.StatusNotFound, getNotFoundError())
		return
	}
	if err != nil {
		log.Printf("GetTournaments: %v", err)
		ctx.JSON(http.StatusBadRequest, getInternalServerError())
		return
	}

	ctx.JSON(http.StatusOK, tournamentsInfo)
}

// getTournamentRoster godoc
// @Summary Получение составов на турнир
// @Security ApiKeyAuth
// @Schemes
// @Description Получение составов на турнир
// @Tags tournament
// @Accept json
// @Produce json
// @Param tournamentID query int true "tournamentID"
// @Success 200 {object} players.TournamentRosterResponse
// @Failure 400,401 {object} Error
// @Failure 500 {object} Error
// @Router /tournament/roster [get]
func (api Api) getTournamentRoster(ctx *gin.Context) {
	userID, err := parseUserIDFromContext(ctx)
	if err != nil {
		log.Println("GetTournamentRoster:", err)
		return
	}

	var tournamentID int

	query := ctx.Request.URL.Query()
	if query.Has("tournamentID") {
		id := query.Get("tournamentID")
		tournamentID, err = strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputParametersError))
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputParametersError))
		return
	}

	res, err := api.services.Tournaments.GetRosterByTournamentID(userID, tournamentID)
	if err != nil {
		log.Println("GetTournamentRoster:", err)
		switch err {
		case storage.IncorrectTournamentID:
			ctx.JSON(http.StatusBadRequest, getBadRequestError(err))
			return
		default:
			ctx.JSON(http.StatusInternalServerError, getInternalServerError())
			return
		}
	}

	ctx.JSON(http.StatusOK, res)
}

// createTournamentTeam godoc
// @Summary Создание команды пользователя в турнире
// @Security ApiKeyAuth
// @Schemes
// @Description Создание команды пользователя в турнире
// @Tags tournament
// @Accept json
// @Produce json
// @Param tournamentID query int true "tournamentID"
// @Param data body tournaments.UserTeamInput true "Входные параметры"
// @Success 200 {object} StatusResponse
// @Failure 400,401 {object} Error
// @Failure 500 {object} Error
// @Router /tournament/team/create [POST]
func (api Api) createTournamentTeam(ctx *gin.Context) {
	var inp tournaments.TournamentTeamModel

	userID, err := parseUserIDFromContext(ctx)
	if err != nil {
		log.Println("CreateTournamentTeam:", err)
		return
	}
	inp.ProfileID = userID

	query := ctx.Request.URL.Query()
	if query.Has("tournamentID") {
		id := query.Get("tournamentID")
		inp.TournamentID, err = strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputParametersError))
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputParametersError))
		return
	}

	var bodyInp tournaments.UserTeamInput
	if err = ctx.BindJSON(&bodyInp); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputBodyError))
		return
	}
	inp.UserTeam = bodyInp.Team

	err = api.services.Tournaments.CreateTournamentTeam(inp)
	if err != nil {
		log.Println("CreateTournamentTeam:", err)
		switch err {
		case storage.IncorrectTournamentID,
			service.TeamExpensiveError,
			service.InvalidTournamentTeam,
			service.InvalidTeamPositions,
			service.JoinTimeExpiredError,
			storage.NotEnoughCoinsError,
			service.InvalidPlayersNumber,
			service.TeamAlreadyCreatedError:
			ctx.JSON(http.StatusBadRequest, getBadRequestError(err))
			return
		default:
			ctx.JSON(http.StatusInternalServerError, getInternalServerError())
			return
		}
	}

	ctx.JSON(http.StatusOK, StatusResponse{"ок"})
}

// getTournamentTeam godoc
// @Summary Получение команды пользователя в турнире
// @Security ApiKeyAuth
// @Schemes
// @Description Получение команды пользователя в турнире
// @Tags tournament
// @Accept json
// @Produce json
// @Param tournamentID query int true "tournamentID"
// @Success 200 {object} players.UserTeamResponse
// @Failure 400,401 {object} Error
// @Failure 500 {object} Error
// @Router /tournament/team [GET]
func (api Api) getTournamentTeam(ctx *gin.Context) {
	userID, err := parseUserIDFromContext(ctx)
	if err != nil {
		log.Println("GetTournamentTeam:", err)
		return
	}
	var tournamentID int

	query := ctx.Request.URL.Query()
	if query.Has("tournamentID") {
		id := query.Get("tournamentID")
		tournamentID, err = strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputParametersError))
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputParametersError))
		return
	}

	res, err := api.services.Tournaments.GetTournamentTeam(userID, tournamentID)
	if err != nil {
		switch err {
		case storage.IncorrectTournamentID:
			ctx.JSON(http.StatusBadRequest, getBadRequestError(err))
			return
		default:
			ctx.JSON(http.StatusInternalServerError, getInternalServerError())
			return
		}

	}

	ctx.JSON(http.StatusOK, res)
}

// editTournamentTeam godoc
// @Summary Редактирование команды пользователя в турнире
// @Security ApiKeyAuth
// @Schemes
// @Description Редактирование команды пользователя в турнире
// @Tags tournament
// @Accept json
// @Produce json
// @Param tournamentID query int true "tournamentID"
// @Param data body tournaments.UserTeamInput true "Входные параметры"
// @Success 200 {object} StatusResponse
// @Failure 400,401 {object} Error
// @Failure 500 {object} Error
// @Router /tournament/team/edit [PUT]
func (api Api) editTournamentTeam(ctx *gin.Context) {
	var inp tournaments.TournamentTeamModel

	userID, err := parseUserIDFromContext(ctx)
	if err != nil {
		log.Println("EditTournamentTeam:", err)
		return
	}
	inp.ProfileID = userID

	query := ctx.Request.URL.Query()
	if query.Has("tournamentID") {
		id := query.Get("tournamentID")
		inp.TournamentID, err = strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputParametersError))
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputParametersError))
		return
	}

	var bodyInp tournaments.UserTeamInput
	if err = ctx.BindJSON(&bodyInp); err != nil {
		ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputBodyError))
		return
	}
	inp.UserTeam = bodyInp.Team

	err = api.services.Tournaments.EditTournamentTeam(inp)
	if err != nil {
		log.Println("EditTournamentTeam:", err)
		switch err {
		case storage.IncorrectTournamentID,
			service.TeamExpensiveError,
			service.InvalidTournamentTeam,
			service.InvalidTeamPositions,
			service.JoinTimeExpiredError,
			service.InvalidPlayersNumber,
			service.TeamNotCreatedError:
			ctx.JSON(http.StatusBadRequest, getBadRequestError(err))
			return
		default:
			ctx.JSON(http.StatusInternalServerError, getInternalServerError())
			return
		}
	}

	ctx.JSON(http.StatusOK, StatusResponse{"ок"})
}

type TournamentID struct {
	ID tournaments.ID `uri:"tournament_id" binding:"required"`
}

// GetMatchesByTournId godoc
// @Summary Получение матчей по id турнира
// @Security ApiKeyAuth
// @Schemes
// @Description Возвращается вся необходимая информация о матчах
// @Tags tournament
// @Accept json
// @Produce json
// @Success 200 {object} []tournaments.GetMatchesByTourId
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 404 {object} Error
// @Param tournament_id path int64 true  "id турнира"
// @Router /tournament/matches_by_tournament_id/{tournament_id} [get]
func (api *Api) GetMatchesByTournId(ctx *gin.Context) {
	_, err := parseUserIDFromContext(ctx)
	if err != nil {
		log.Println("GetTournamentTeam:", err)
		return
	}

	var tourId TournamentID
	if err := ctx.ShouldBindUri(&tourId); err != nil {
		ctx.JSON(http.StatusBadRequest, getBadRequestError(err))
		return
	}

	tournInfo, err := api.services.Tournaments.GetMatchesByTournamentsId(ctx, tourId.ID)
	if err != nil {
		log.Printf("GetMatchesByTournId: %v", err)
		if errors.Is(err, service.NotFoundTournamentsById) {
			ctx.JSON(http.StatusNotFound, getNotFoundError())
			return
		}
		ctx.JSON(http.StatusInternalServerError, getInternalServerError())
		return
	}
	ctx.JSON(http.StatusOK, tournInfo)

}

// getTournamentsInfo godoc
// @Summary Получение турниров
// @Security ApiKeyAuth
// @Schemes
// @Description Получение турниров
// @Tags tournament
// @Accept json
// @Produce json
// @Param tournamentID query int false "tournamentID"
// @Param league query string false "league" Enums(NHL, KHL)
// @Param status query string false "status" Enums(not_yet_started, started, finished, active)
// @Param type query string true "type" Enums(all, personal)
// @Success 200 {array} tournaments.Tournament
// @Failure 400,401 {object} Error
// @Failure 500 {object} Error
// @Router /tournaments [GET]
func (api Api) getTournamentsInfo(ctx *gin.Context) {
	userID, err := parseUserIDFromContext(ctx)
	if err != nil {
		log.Println("GetTournamentsInfo:", err)
		return
	}

	var filterTournament tournaments.TournamentFilter
	query := ctx.Request.URL.Query()

	filterTournament.ProfileID = userID
	if query.Has("tournamentID") {
		id := query.Get("tournamentID")
		tournamentID, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputParametersError))
			return
		}
		filterTournament.TournamentID = tournamentID
	}

	if query.Has("status") {
		switch query.Get("status") {
		case "not_yet_started":
			filterTournament.Status = "not_yet_started"
		case "started":
			filterTournament.Status = "started"
		case "finished":
			filterTournament.Status = "finished"
		case "active":
			filterTournament.Status = "active"
		default:
			ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputParametersError))
			return
		}
	}

	if query.Has("league") {
		switch query.Get("league") {
		case "NHL":
			filterTournament.League = tournaments.NHL
		case "KHL":
			filterTournament.League = tournaments.KHL
		default:
			ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputParametersError))
			return
		}
	}

	if query.Has("type") {
		switch query.Get("type") {
		case "all":
			filterTournament.Type = "all"
		case "personal":
			filterTournament.Type = "personal"
		default:
			ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputParametersError))
			return
		}
	}

	res, err := api.services.Tournaments.GetTournamentsInfo(filterTournament)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, getInternalServerError())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// getTournamentResults godoc
// @Summary Получение результатов турнира
// @Security ApiKeyAuth
// @Schemes
// @Description Получение результатов турнира
// @Tags tournament
// @Accept json
// @Produce json
// @Param tournamentID query int true "tournamentID"
// @Success 200 {array} players.TournamentResults
// @Failure 400,401 {object} Error
// @Failure 500 {object} Error
// @Router /tournament/results [get]
func (api Api) getTournamentResults(ctx *gin.Context) {
	_, err := parseUserIDFromContext(ctx)
	if err != nil {
		log.Println("GetTournamentRoster:", err)
		return
	}

	var tournamentID int

	query := ctx.Request.URL.Query()
	if query.Has("tournamentID") {
		id := query.Get("tournamentID")
		tournamentID, err = strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputParametersError))
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, getBadRequestError(InvalidInputParametersError))
		return
	}

	res, err := api.services.Tournaments.GetCachedTournamentResults(tournamentID)
	if err != nil {
		log.Println("GetTournamentResults:", err)
		switch err {
		case storage.IncorrectTournamentID,
			service.TournamentNotFinishedError:
			ctx.JSON(http.StatusBadRequest, getBadRequestError(err))
			return
		default:
			ctx.JSON(http.StatusInternalServerError, getInternalServerError())
			return
		}
	}

	ctx.JSON(http.StatusOK, res)
}

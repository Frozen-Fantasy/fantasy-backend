package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Frozen-Fantasy/fantasy-backend.git/pkg/models/tournaments"
	"github.com/Frozen-Fantasy/fantasy-backend.git/pkg/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// CreateTeamsNHL godoc
// @Summary Создание команд NHL
// @Schemes
// @Description Добавлят информацию о команде NHL
// @Tags tournament
// @Produce json
// @Success 200
// @Failure 400 {object} Error
// @Router /tournament/create_team_nhl [get]
func (api *Api) CreateTeamsNHL(ctx *gin.Context) {

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
// @Schemes
// @Description Добавлят информацию о команде KHL
// @Tags tournament
// @Produce json
// @Success 200
// @Failure 400 {object} Error
// @Router /tournament/create_team_khl [get]
func (api *Api) CreateTeamsKHL(ctx *gin.Context) {

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

// EventsKHL godoc
// @Summary Получение событий на следующий день KHL
// @Schemes
// @Description Добавляет в бд матчи за день
// @Tags tournament
// @Produce json
// @Success 200
// @Failure 400 {object} Error
// @Router /tournament/events_day_khl [get]
func (api *Api) EventsKHL(ctx *gin.Context) {

	curTime := time.Now()
	curTime = curTime.Add(24 * time.Hour)
	startDay := time.Date(curTime.Year(), curTime.Month(), curTime.Day(), 0, 0, 0, 0, time.UTC)
	endDay := time.Date(curTime.Year(), curTime.Month(), curTime.Day(), 23, 59, 59, 0, time.UTC)

	url := fmt.Sprint("https://khl.api.webcaster.pro/api/khl_mobile/events_v2?q[start_at_lt_time_from_unixtime]=", endDay.Unix(), "&order_direction=desc&q[start_at_gt_time_from_unixtime]=", startDay.Unix())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("EventsKHL:", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("EventsKHL:", err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	var eventKHL []tournaments.EventDataKHL

	err = decoder.Decode(&eventKHL)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}

	err = api.services.Teams.AddEventsKHL(ctx, eventKHL)
	if err != nil {
		log.Printf("EventsKHL: %v", err)
		ctx.JSON(http.StatusBadRequest, getInternalServerError())
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

// EventsNHL godoc
// @Summary Получение событий на следующий день NHL
// @Schemes
// @Description Добавляет в бд матчи за день
// @Tags tournament
// @Produce json
// @Success 200
// @Failure 400 {object} Error
// @Router /tournament/events_day_nhl [get]
func (api *Api) EventsNHL(ctx *gin.Context) {
	curTime := time.Now()

	url := fmt.Sprint("https://api-web.nhle.com/v1/schedule/", curTime.Format("2006-01-02"))
	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("EventsNHL:", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("EventsNHL:", err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	var eventNHL tournaments.ScheduleNHL

	err = decoder.Decode(&eventNHL)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}

	err = api.services.Teams.AddEventsNHL(ctx, eventNHL.GameWeeks[0].Games)
	if err != nil {
		log.Printf("EventsNHL: %v", err)
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

	//if err := ctx.BindUri(&leagueName); err != nil {
	//	ctx.JSON(http.StatusBadRequest, getBadRequestError(err))
	//	return
	//}
	//if err := ctx.ShouldBindUri(&leagueName); err != nil {
	//	ctx.JSON(http.StatusBadRequest, getBadRequestError(err))
	//	return
	//}

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

// CreateTournaments godoc
// @Summary Создание турниров на следующий день
// @Schemes
// @Description Дата берётся автоматически
// @Description Создаётся 4 турника 2 НХЛ и 2 КХЛ платный и бесплатный
// @Tags tournament
// @Produce json
// @Success 200
// @Failure 400 {object} Error
// @Router /tournament/create_tournaments [get]
func (api *Api) CreateTournaments(ctx *gin.Context) {
	err := api.services.Teams.CreateTournaments(ctx)
	if err != nil {
		log.Printf("CreateTournaments: %v", err)
		ctx.JSON(http.StatusBadRequest, getInternalServerError())
		return
	}
}

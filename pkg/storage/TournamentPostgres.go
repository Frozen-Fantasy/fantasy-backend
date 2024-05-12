package storage

import (
	"database/sql"
	"errors"
	"github.com/Frozen-Fantasy/fantasy-backend.git/pkg/models/players"
	"github.com/Frozen-Fantasy/fantasy-backend.git/pkg/models/tournaments"
	"github.com/Frozen-Fantasy/fantasy-backend.git/pkg/models/user"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"strconv"
	"strings"
)

var (
	IncorrectTournamentID = errors.New("некорректный id турнира")
)

func (p *PostgresStorage) GetMatchesByTournamentID(tournamentID int) ([]int, error) {
	var matchesIDs []int
	var matchesIDsStr string
	err := p.db.QueryRow("SELECT matches_ids FROM tournaments WHERE id = $1", tournamentID).Scan(&matchesIDsStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return matchesIDs, IncorrectTournamentID
		}
		return matchesIDs, err
	}

	matchesIDsStr = strings.Trim(matchesIDsStr, "{}")
	matchesIDsStr = strings.ReplaceAll(matchesIDsStr, " ", "")
	matchesIDsStrArr := strings.Split(matchesIDsStr, ",")
	for _, idStr := range matchesIDsStrArr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return matchesIDs, err
		}
		matchesIDs = append(matchesIDs, id)
	}

	return matchesIDs, nil
}

func (p *PostgresStorage) GetTeamsByMatches(matchesIDs []int) ([]int, error) {
	var teams []int
	for _, matchID := range matchesIDs {
		var homeTeamApiID, awayTeamApiID, homeTeamID, awayTeamID, league int
		err := p.db.QueryRow("SELECT m.home_team_id, m.away_team_id, t1.team_id, t2.team_id, m.league FROM matches m "+
			"JOIN teams t1 ON m.home_team_id = t1.api_id AND m.league = t1.league "+
			"JOIN teams t2 ON m.away_team_id = t2.api_id AND m.league = t2.league "+
			"WHERE m.id = $1", matchID).Scan(&homeTeamApiID, &awayTeamApiID, &homeTeamID, &awayTeamID, &league)
		if err != nil {
			return teams, err
		}

		teams = append(teams, homeTeamID, awayTeamID)
	}

	return teams, nil
}

func (p *PostgresStorage) GetTeamDataByID(teamID int) (players.TeamData, error) {
	var teamInfo players.TeamData

	err := p.db.QueryRow("SELECT team_id, team_name, team_abbrev FROM teams WHERE team_id = $1", teamID).Scan(&teamInfo.TeamID, &teamInfo.TeamName, &teamInfo.TeamAbbrev)
	if err != nil {
		return teamInfo, err
	}

	return teamInfo, nil
}

func (p *PostgresStorage) GetTournamentDataByID(tournamentID int) (tournaments.Tournament, error) {
	var tournamentInfo tournaments.Tournament

	err := p.db.QueryRow("SELECT id, league, title, matches_ids, started_at, end_at, players_amount, deposit, "+
		"prize_fond, status_tournament FROM tournaments WHERE id = $1", tournamentID).Scan(
		&tournamentInfo.TournamentId,
		&tournamentInfo.League,
		&tournamentInfo.Title,
		&tournamentInfo.MatchesIds,
		&tournamentInfo.TimeStart,
		&tournamentInfo.TimeEnd,
		&tournamentInfo.PlayersAmount,
		&tournamentInfo.Deposit,
		&tournamentInfo.PrizeFond,
		&tournamentInfo.StatusTournament,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return tournamentInfo, IncorrectTournamentID
		}
		return tournamentInfo, err
	}

	return tournamentInfo, nil
}

func (p *PostgresStorage) CreateTournamentTeam(teamInput tournaments.TournamentTeamModel) error {
	tx, err := p.db.Beginx()
	if err != nil {
		return err
	}

	if teamInput.Deposit > 0 {
		coinTr := user.CoinTransactionsModel{
			ProfileID:          teamInput.ProfileID,
			TransactionDetails: "Участие в турнире №" + strconv.Itoa(teamInput.TournamentID),
			Amount:             -teamInput.Deposit,
			Status:             user.SuccessTransaction,
		}
		err = p.UpdateBalance(tx, coinTr.ProfileID, coinTr.Amount)
		if err != nil {
			return err
		}
		err = p.CreateCoinTransaction(tx, coinTr)
		if err != nil {
			return err
		}
		prizeFondQuery := `UPDATE tournaments SET prize_fond = prize_fond + $1 WHERE id = $2`
		_, err = tx.Exec(prizeFondQuery, int(float64(teamInput.Deposit)*1.5), teamInput.TournamentID)
		if err != nil {
			tx.Rollback()
		}
	}

	teamArray := pq.Array(teamInput.UserTeam)
	rosterQuery := `INSERT INTO user_roster (tournament_id, user_id, roster, current_balance) 
              VALUES ($1, $2, $3, $4)`

	_, err = tx.Exec(rosterQuery, teamInput.TournamentID, teamInput.ProfileID, teamArray, 100-teamInput.TeamCost)
	if err != nil {
		tx.Rollback()
		return err
	}

	playersAmountQuery := `UPDATE tournaments SET players_amount = players_amount + 1 WHERE id = $1`
	_, err = tx.Exec(playersAmountQuery, teamInput.TournamentID)
	if err != nil {
		tx.Rollback()
	}

	return tx.Commit()
}

func (p *PostgresStorage) GetTournamentTeam(userID uuid.UUID, tournamentID int) (players.UserTeam, error) {
	var res players.UserTeam
	query := "SELECT roster, current_balance FROM user_roster WHERE tournament_id = $1 AND user_id = $2"

	var rosterStr string
	var currentBalance float64
	err := p.db.QueryRow(query, tournamentID, userID).Scan(&rosterStr, &currentBalance)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, nil
		}
		return res, err
	}

	res.Balance = currentBalance
	rosterStr = strings.Trim(rosterStr, "{}")
	rosterStr = strings.ReplaceAll(rosterStr, " ", "")
	matchesIDsStrArr := strings.Split(rosterStr, ",")
	for _, idStr := range matchesIDsStrArr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return res, err
		}
		res.PlayerIDs = append(res.PlayerIDs, id)
	}

	return res, nil
}

func (p *PostgresStorage) EditTournamentTeam(teamInput tournaments.TournamentTeamModel) error {
	teamArray := pq.Array(teamInput.UserTeam)
	query := `UPDATE user_roster SET roster = $1, current_balance = $2 WHERE tournament_id = $3 AND user_id = $4`

	_, err := p.db.Exec(query, teamArray, 100-teamInput.TeamCost, teamInput.TournamentID, teamInput.ProfileID)
	if err != nil {
		return err
	}

	return nil
}

package flyweight

import (
	"time"
)

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

const (
	TeamA = iota
	TeamB
)

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type teamFlyweightFactory struct {
	createdTeams map[uint64]*Team
}

func (t *teamFlyweightFactory) GetTeam(teamID uint64) *Team {
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}
	team := teamFactory(teamID)
	t.createdTeams[teamID] = &team
	return t.createdTeams[teamID]
}

func teamFactory(teamID uint64) (team Team) {
	switch teamID {
	case TeamA:
		team = Team{
			ID:   1,
			Name: "TeamA",
		}
	case TeamB:
		team = Team{
			ID:   2,
			Name: "TeamB",
		}
	}
	return
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}

func TeamFlyweightFactory() teamFlyweightFactory {
	return teamFlyweightFactory{
		createdTeams: make(map[uint64]*Team, 0),
	}
}

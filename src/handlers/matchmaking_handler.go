package handlers

import "github.com/zsarvas/RL-Discord-Matchmaking/queue"

type MatchObject struct {
	firstTeam  []string
	secondTeam []string
	winner     string
}

type Match interface {
	SetWinner(winner string)
}

type MatchMaker struct {
	activeMatches map[string]Match
}

func NewMatchMaker() *MatchMaker {
	matches := map[string]Match{}

	var matchmaker = &MatchMaker{
		activeMatches: matches,
	}

	return matchmaker
}

func (m *MatchObject) SetWinner(winner string) {
	m.winner = winner
}

func (mm *MatchMaker) addMatch(newMatch MatchObject, matchId string) {
	mm.activeMatches[matchId] = &newMatch
}

func (mm *MatchMaker) DumpQueueIntoMatch(queue *queue.Queue) {
	// player1 := queue.Dequeue()
	// player2 := queue.Dequeue()
	// player3 := queue.Dequeue()
	// player4 := queue.Dequeue()

	// teamOne := []string{queue.Dequeue(), queue.Dequeue()}
	// teamTwo := []string{queue.Dequeue(), queue.Dequeue()}

	// createdMatchObject := MatchObject{
	// 	firstTeam: teamOne,
	// 	secondTeam: teamTwo,
	// 	winner: "",
	// }

	// mm.addMatch(createdMatchObject, )
}

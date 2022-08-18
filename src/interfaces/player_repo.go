package interfaces

import (
	"github.com/zsarvas/RL-Discord-Matchmaking/domain"
)

type PlayerDataHandler interface {
	Add(newPlayer domain.Player)
	GetById(id string) domain.Player
}

type PlayerDataRepo struct {
	dbHandler PlayerDataHandler
}

type PlayerRepo PlayerDataRepo

func NewPlayerRepo(repoHandler PlayerDataHandler) *PlayerRepo {
	dbPlayerRepo := new(PlayerRepo)
	dbPlayerRepo.dbHandler = repoHandler

	return dbPlayerRepo
}

func (repo *PlayerRepo) Store(player domain.Player) {
	repo.dbHandler.Add(player)
}

func (repo *PlayerRepo) Get(playerId string) domain.Player {
	foundPlayer := repo.dbHandler.GetById(playerId)

	return foundPlayer
}

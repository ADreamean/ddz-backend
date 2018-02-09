package game

import "sync"

type gamePool struct {
	lock  sync.Mutex
	games []*Game
}

func (gp *gamePool) Find(roomId int) *Game {
	gp.lock.Lock()
	defer gp.lock.Unlock()
	for _, game := range gp.games {
		if game.roomId == roomId {
			return game
		}
	}

	game := NewGame(roomId)

	gp.games = append(gp.games, game)

	return game
}

func newGamePool(size int) *gamePool {
	return &gamePool{games: make([]*Game, 0, size)}
}

package game

import (
	"golang.org/x/net/websocket"
	"github.com/ADreamean/ddz-backend/login"
	"github.com/go-siris/siris/core/errors"
	"sync"
	"io"
)

var (
	START = "start" //游戏开始
)

type Game struct {
	ctxs        []*Context
	lock        sync.Mutex
	roomId      int
	cards       Cards
	playerCards []Cards
	history     []Cards
	base        int
}

type Context struct {
	con   *websocket.Conn
	user  *login.User
	ready bool
	index int
	role  int
}

func (g *Game) broadcast(data string, except int) {
	for i, ctx := range g.ctxs {
		if i == except {
			continue
		}
		io.WriteString(ctx.con, data)
	}
}

func (g *Game) Join(ctx *Context) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	if len(g.ctxs) >= 3 {
		return errors.New("人数已满！")
	}
	ctx.index = len(g.ctxs)
	g.ctxs = append(g.ctxs, ctx)
	return nil
}

func (g *Game) Turn(ctx *Context, cards Cards) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.playerCards[ctx.index] = g.playerCards[ctx.index].Diff(cards)
	if len(g.playerCards[ctx.index]) > 0 {
		return
	}

	//游戏结束

}

func NewGame(roomId int) *Game {
	return &Game{roomId: roomId, ctxs: make([]*Context, 0, 10)}
}

func (ctx *Context) Ready() {
	ctx.ready = true
}

func (ctx *Context) UnReady() {
	ctx.ready = false
}

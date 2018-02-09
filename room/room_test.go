package room

import (
	"sync"
	"net/http"
	"net/http/httptest"
	"testing"
)

var once sync.Once

func serverStart() {
	mux := http.NewServeMux()
	mux.HandleFunc("/room/list", roomList)
	mux.HandleFunc("/room/create", roomCreate)
	mux.HandleFunc("/room/join", roomJoin)
	mux.HandleFunc("/room/quit", roomQuit)
	ts := httptest.NewServer(mux)
	defer ts.Close()
}

func TestRoomList(t *testing.T) {
	once.Do(serverStart)
}

func TestRoomCreate(t *testing.T) {
	once.Do(serverStart)
}

func TestRoomJoin(t *testing.T) {
	once.Do(serverStart)
}

func TestRoomQuit(t *testing.T) {
	once.Do(serverStart)
}

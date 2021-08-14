package engineio

import (
	"fmt"
	"github.com/chuwt/socketio-go/common/b64id"
	"github.com/chuwt/socketio-go/common/json"
	"net/http"
	"strings"
)

type EngineIO struct {
	b64Id b64id.B64Id

	pingTimeout  int64
	pingInterval int64

	open   func()
	close  func()
	events map[string]func()
}

func NewEngineIO() *EngineIO {
	return &EngineIO{
		pingTimeout:  100000,
		pingInterval: 100000,
		open:         nil,
		close:        nil,
		events:       nil,
	}
}

func (eio *EngineIO) OnOpen(f func()) {
	eio.open = f
}

func (eio *EngineIO) onOpen() []byte {
	if eio.open != nil {
		eio.open()
	}

	handshake := NewHandshake(eio.b64Id.GenerateId(), eio.pingTimeout, eio.pingInterval)
	return append([]byte("0"), handshake.Json()...)
}

func (eio *EngineIO) OnClose(f func()) {
	eio.close = f
}

func (eio *EngineIO) cors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin")
	(*w).Header().Add("Access-Control-Allow-Methods", "*")
}

func (eio *EngineIO) closeConn(w *http.ResponseWriter) {
	conn, _, _ := (*w).(http.Hijacker).Hijack()
	_ = conn.Close()
}

// ServeHTTP implement http.Handler
func (eio *EngineIO) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	eio.cors(&w)

	// urlPath must be /engine.io
	if r.URL.Path != "/engine.io/" {
		eio.closeConn(&w)
		return
	}

	// open
	if !strings.Contains(r.URL.RawQuery, "sid") {
		_, err = w.Write(eio.onOpen())
		if err != nil {
			eio.closeConn(&w)
		}
		return
	} else {

	}
	w.Write([]byte("ok"))
}

func (eio *EngineIO) Run(addr string) error {
	// todo sets some useful config, such as timeout
	server := &http.Server{
		Addr:    addr,
		Handler: eio,
	}
	fmt.Println(fmt.Sprintf("server run at %s", addr))
	return server.ListenAndServe()
}

type Handshake struct {
	// session id
	Sid string `json:"sid"`
	// possible transport upgrades
	Upgrades []string `json:"upgrades"`
	// server configured ping timeout, used for the client to detect that the server is unresponsive
	// milliseconds
	PingTimeout int64 `json:"pingInterval"`
	// server configured ping interval, used for the client to detect that the server is unresponsive
	// milliseconds
	PingInterval int64 `json:"pingTimeout"`
}

func NewHandshake(sid string, pingTimeout, pingInterval int64, upgrades ...string) Handshake {
	handshake := Handshake{
		Sid:          sid,
		Upgrades:     []string{"websocket"},
		PingTimeout:  pingTimeout,
		PingInterval: pingInterval,
	}
	if upgrades != nil {
		handshake.Upgrades = upgrades
	}
	return handshake
}

func (h *Handshake) Json() []byte {
	bytes, _ := json.JSON.Marshal(h)
	return bytes
}

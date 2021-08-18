package engineio

import "github.com/chuwt/socketio-go/internal/json"

const (
	defaultUpgrade = "websocket"
)

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
		Upgrades:     []string{defaultUpgrade},
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

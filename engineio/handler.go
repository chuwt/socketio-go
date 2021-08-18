package engineio

import (
	"net/http"
)

func (eio *EngineIO) OnOpen(f func()) {
	eio.open = f
}

func (eio *EngineIO) onOpen() []byte {
	if eio.open != nil {
		eio.open()
	}
	sid, _ := eio.bid.GenerateId()
	handshake := NewHandshake(sid, eio.pingTimeout, eio.pingInterval)
	return append(Open, handshake.Json()...)
}

func (eio *EngineIO) OnClose(f func()) {
	eio.close = f
}

func (eio *EngineIO) closeConn(w *http.ResponseWriter) {
	conn, _, _ := (*w).(http.Hijacker).Hijack()
	_ = conn.Close()
}

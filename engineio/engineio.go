package engineio

import (
	"fmt"
	"github.com/chuwt/socketio-go/internal/b64id"
	"net/http"
	"strings"
)

type EngineIO struct {
	bid b64id.B64Id

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

func (eio *EngineIO) cors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin")
	(*w).Header().Add("Access-Control-Allow-Methods", "*")
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

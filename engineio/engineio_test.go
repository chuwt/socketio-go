package engineio

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestEngineIo(t *testing.T) {
	engineio := NewEngineIO()
	t.Log(engineio.Run(":9999"))
}

func TestClient(t *testing.T) {
	r, err := http.Get("http://localhost:9999/engine.io/?EIO=4&transport=polling&t=N8hyd6w")
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Body.Close()
	t.Log("response:", string(bytes))
}

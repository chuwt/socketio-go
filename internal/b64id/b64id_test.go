package b64id

import "testing"

func TestGenerateId(t *testing.T) {
	b := NewB64Id()
	for i := 0; i < 1000; i++ {
		t.Log(b.GenerateId())
	}
}

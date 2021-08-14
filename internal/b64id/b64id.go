package b64id

import (
	"encoding/base64"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

// from ASCII 0 to 127
const (
	min = 0
	max = 127
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type B64Id struct {
	sequenceNumber int64
}

func NewB64Id() B64Id {
	// seq start from [0, 10)
	return B64Id{
		sequenceNumber: int64(rand.Intn(10)),
	}
}

func (bid *B64Id) GenerateId() string {
	buf := make([]byte, 15)
	for i := 0; i < 12; i++ {
		buf[i] = byte(min + rand.Intn(max-min))
	}
	big.NewInt(bid.sequenceNumber).FillBytes(buf[12:])
	bid.sequenceNumber = (bid.sequenceNumber + 1) & 0xffffff
	return strings.ReplaceAll(strings.ReplaceAll(base64.StdEncoding.EncodeToString(buf), "/", "_"), "+", "-")
}

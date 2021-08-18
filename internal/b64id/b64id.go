package b64id

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
	"strings"
)

type B64Id struct {
	sequenceNumber int64
}

func NewB64Id() B64Id {
	return B64Id{}
}

func (bid *B64Id) GenerateId() (string, error) {
	buf := make([]byte, 15)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	big.NewInt(bid.sequenceNumber).FillBytes(buf[12:])
	bid.sequenceNumber = (bid.sequenceNumber + 1) & 0xffffff
	return strings.ReplaceAll(strings.ReplaceAll(base64.StdEncoding.EncodeToString(buf), "/", "_"), "+", "-"), nil
}

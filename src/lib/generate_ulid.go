package lib

import (
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
)

func GenerateUlid() string {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())

	val, _ := ulid.New(ms, entropy)

	return val.String()
}

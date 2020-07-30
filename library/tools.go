package library

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

type tools struct {
}

func NewTools() *tools {
	return &tools{}
}

func (t *tools) CreateToken(prefix string) string {
	nanosecond := time.Now().Nanosecond()
	rand.Seed(int64(nanosecond))
	has := md5.Sum([]byte(fmt.Sprintf("%d%d", nanosecond, rand.Int())))
	md5str := fmt.Sprintf("%x", has)
	return prefix + "-" + md5str

}

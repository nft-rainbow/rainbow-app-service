package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

func GenerateIdByTime(prefix string) string {
	suffix, _ := rand.Int(rand.Reader, big.NewInt(999))
	return fmt.Sprintf("%s%13d%03d", prefix, time.Now().UnixMicro(), suffix)
}

func GenerateIDByTimeHash(prefix string, len int) string {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(time.Now().UnixNano()))
	hash := md5.Sum(b)
	hex := hex.EncodeToString(hash[:])
	id := fmt.Sprintf("%s%s", prefix, hex[:len])
	return id
}

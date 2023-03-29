package utils

import (
	"fmt"
	"testing"
)

func TestGenerateIdByTime(t *testing.T) {
	id := GenerateIdByTime("POAP")
	fmt.Println(id)
}

func TestGenerateIdByTimeHex(t *testing.T) {
	id := GenerateIDByTimeHash("POAP", 9)
	fmt.Println(id)
}

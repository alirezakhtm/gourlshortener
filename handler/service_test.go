package handler

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init()  {
	rand.Seed(time.Now().UnixNano())	
}

func TestGenerateRandomAddress(t *testing.T) {
	ans := GenerateRandomAddress()
	fmt.Println("[*] ", ans)
}


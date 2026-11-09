package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	numBits int = 52
)

func gameResult(salt, gameHash []byte) float64 {
	mac := hmac.New(sha256.New, salt)
	mac.Write(gameHash)
	hash := hex.EncodeToString(mac.Sum(nil))

	seed := hash[:numBits/4]
	r, err := strconv.ParseUint(seed, 16, 64)
	if err != nil {
		panic(err)
	}

	x := float64(r) / math.Pow(2, float64(numBits))

	x = 99 / (1 - x)

	result := math.Floor(x)
	return math.Max(1, result/100)
}

func main() {
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}

	gameHash := make([]byte, 32)
	if _, err := rand.Read(gameHash); err != nil {
		panic(err)
	}

	out := os.Stdout

	fmt.Fprint(out, "round,hash,multiplier\n")

	for i := range 1_000 {
		fmt.Fprintf(out, "%d,%s,%.2f\n", i+1, hex.EncodeToString(gameHash), gameResult(salt, gameHash))

		next := sha256.Sum256(gameHash)
		gameHash = next[:]
	}
}

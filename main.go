package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"encoding/json"
	"log"
	"math/rand"

	"time"
)

//go:generate go run scripts/includeJson.go

type Feather struct {
	Time      time.Time `json:"Time"`
	Author    string    `json:"Author"`
	Title     string    `json:"Title"`
	Body      string    `json:"Body"`
	Recipient string    `json:"Recipient Names"`
}

type pillow struct {
	Feathers []Feather `json:"feathers"`
}

func main() {
	pillow := pillow{}
	if err := pluckBirdStuffPillow(feathers, &pillow); err != nil {
		log.Fatalf("Failed to parse json string with error: %s", err)
	}
	metaFeather := pillow.Feathers[rand.Intn(len(pillow.Feathers))]
	log.Print(metaFeather)
}

func pluckBirdStuffPillow(jsonString string, out *pillow) error {
	if err := json.Unmarshal([]byte(jsonString), &out); err != nil {
		return err
	}
	return nil
}

// Seed Random Generator
func init() {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

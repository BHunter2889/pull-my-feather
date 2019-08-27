package main

import (
	"context"
	crand "crypto/rand"
	"encoding/binary"
	"encoding/json"
	"github.com/BHunter2889/da-fish-alexa/alexa"
	"github.com/aws/aws-lambda-go/lambda"
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

func IntentDispatcher(ctx context.Context, request alexa.Request) alexa.Response {
	log.Print("Intent Dispatcher")
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "PullFeatherIntent":
		log.Print("INTENT_DISPATCH: TodaysFishRatingIntent")
		response = HandlePullFeatherIntent(ctx, request)
	case alexa.HelpIntent:
		log.Print("INTENT_DISPATCH: HelpIntent")
		response = HandleHelpIntent(ctx, request)
	case "AboutIntent":
		log.Print("INTENT_DISPATCH: AboutIntent")
		response = HandleAboutIntent(ctx, request)
	default:
		log.Print("INTENT_DISPATCH: Default Response")
		response = HandlePullFeatherIntent(ctx, request)
	}
	log.Print("RESPONSE: ", response)
	return response
}

// TODO
func HandleAboutIntent(ctx context.Context, request alexa.Request) alexa.Response {
	return alexa.Response{}
}

// TODO
func HandleHelpIntent(ctx context.Context, request alexa.Request) alexa.Response {
	return alexa.Response{}
}

func HandlePullFeatherIntent(ctx context.Context, request alexa.Request) alexa.Response {
	pillow := pillow{}
	if err := pluckBirdStuffPillow(feathers, &pillow); err != nil {
		log.Fatalf("Failed to parse json string with error: %s", err)
	}
	metaFeather := pillow.Feathers[rand.Intn(len(pillow.Feathers))]
	log.Print(metaFeather)
	// TODO: Populate response
	return alexa.Response{}
}

func Handler(ctx context.Context, request alexa.Request) (response alexa.Response, error error) {
	log.Print("Begin Handler")
	return IntentDispatcher(ctx, request), nil
}
func main() {
	log.Print("Begin Main")
	lambda.Start(Handler)
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

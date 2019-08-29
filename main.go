package main

import (
	"context"
	crand "crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/BHunter2889/go-alexa-devkit/alexa"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"math/rand"
	"time"
)

//go:generate go run scripts/includeJson.go

const pullFeatherIntentTitle = "Tickled Pink Award Winner"

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
		log.Print("INTENT_DISPATCH: PullFeatherIntent")
		response = HandlePullFeatherIntent(ctx, request)
	case alexa.HelpIntent:
		log.Print("INTENT_DISPATCH: HelpIntent")
		response = HandleHelpIntent(ctx, request)
	case "AboutIntent":
		log.Print("INTENT_DISPATCH: AboutIntent")
		response = HandleAboutIntent(ctx, request)
	default:
		log.Print("INTENT_DISPATCH: Default Response (PullFeatherIntent)")
		response = HandlePullFeatherIntent(ctx, request)
	}
	log.Print("RESPONSE: ", response)
	return response
}

// TODO
func HandleAboutIntent(ctx context.Context, request alexa.Request) alexa.Response {
	return HandlePullFeatherIntent(ctx, request)
}

// TODO
func HandleHelpIntent(ctx context.Context, request alexa.Request) alexa.Response {
	return HandlePullFeatherIntent(ctx, request)
}

func HandlePullFeatherIntent(ctx context.Context, request alexa.Request) alexa.Response {
	pillow := pillow{}
	if err := pluckBirdStuffPillow(feathers, &pillow); err != nil {
		log.Fatalf("Failed to parse json string with error: %s", err)
	}
	metaFeather := pillow.Feathers[rand.Intn(len(pillow.Feathers))]
	log.Print(metaFeather)

	date := metaFeather.Time.Format("Monday, Jan 2, 2006")
	log.Print(date)

	ssmlBuilder := alexa.SSMLBuilder{}
	ssmlBuilder.Say("<prosody pitch='low' volume='soft'> And the </prosody> <prosody pitch='medium' volume='medium'> winner </prosody> <prosody pitch='high' volume='loud' rate='x-slow'> is... </prosody> ") // TODO: Excited
	ssmlBuilder.Pause("300")
	ssmlBuilder.Say(metaFeather.Recipient) // TODO: Exclaim!
	ssmlBuilder.Pause("500")
	ssmlBuilder.Say(fmt.Sprintf("Awarded by %s", metaFeather.Author))
	ssmlBuilder.Pause("750")
	ssmlBuilder.Say(metaFeather.Body)
	ssmlBuilder.Pause("750")
	ssmlBuilder.Say("Congratulations!") // TODO: Exclaim
	ssml := ssmlBuilder.Build()

	speechTransformer := alexa.NewSSMLToSpeechTransformer()
	tl := make([]alexa.Transformer, 0)
	tl = append(tl, speechTransformer)

	rd := alexa.Directive{}
	if err := alexa.ExtractNewRenderDocDirectiveFromString("tickled-pink", aplJson, &rd); err != nil {
		log.Print("ERROR READING APL TEMPLATE", err)
	}

	rd.SetBodyContentTitleText(metaFeather.Recipient)
	rd.SetBodyContentSubtitle(fmt.Sprintf("Awarded By: %s", metaFeather.Author))
	rd.SetBodyContentPrimaryText(metaFeather.Title)
	rd.SetBodyContentSubheader(date)
	rd.AddBodyContentBullets(metaFeather.Body)
	rd.DataSources.TemplateData.Properties.SSML = ssml
	rd.DataSources.TemplateData.Transformers = tl

	resp := alexa.NewAPLResponse(ssml, alexa.NewDirectivesList("Tickled Pink Award Winner", rd))
	return resp
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

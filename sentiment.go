package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"os"

	"github.com/cdipaolo/sentiment"
)

var (
	model sentiment.Models
)

func init() {
	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)
	log.SetFormatter(&log.TextFormatter{})

	var err error
	model, err = sentiment.Restore()
	if err != nil {
		panic(fmt.Sprintf("ERROR: error restoring sentiment model!\n\t%v\n", err))
	}

	http.Handle("/analyze", Post(HandleSentiment))
	http.Handle("/review", Post(HandleGetProductReview))
	http.Handle("/task", Post(HandleHookedRequest))
	http.Handle("/", Get(HandleStatus))
}

func main() {
	flag.Parse()
	err := ParseConfig()
	if err != nil {
		panic(fmt.Sprintf("ERROR: error parsing configuration!\n\t%v\n", err.Error()))
	}

	// server starting message
	log.WithFields(log.Fields{
		"port": Config.portString,
	}).Info("Sentiment analysis is starting...")

	log.Printf("Listening at http://127.0.0.1%v ...\n", Config.portString)
	log.Fatal(http.ListenAndServe(Config.portString, nil))
}

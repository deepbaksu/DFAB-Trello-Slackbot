package main

import (
	"flag"
	"os"
	"time"

	"github.com/dl4ab/DFAB-Trello-Slackbot/internal"
	"github.com/dl4ab/DFAB-Trello-Slackbot/timeutil"

	log "github.com/sirupsen/logrus"
)

var startTimeInDurationString = flag.String("start", "1d", "Start time to search (e.g., 1d means search events from 1 day ago)")

func main() {
	flag.Parse()

	appKey := os.Getenv("TRELLO_APP_KEY")
	token := os.Getenv("TRELLO_TOKEN")
	boardId := os.Getenv("TRELLO_BOARD_ID")
	username := os.Getenv("TRELLO_USERNAME")

	startTimeInDuration, err := timeutil.ParseDuration(*startTimeInDurationString)
	if err != nil {
		log.WithError(err).Fatal("Failed to parse duration string")
	}
	startDate := timeutil.GetBeginningOfDay(timeutil.GetPreviousTime(startTimeInDuration))
	log.Infof("Checking activities since %v", startDate.Local().Format(time.UnixDate))

  app := internal.New(appKey, token, boardId, username, startDate)
  app.Run()
}

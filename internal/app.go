package internal

import (
	"time"

	"github.com/adlio/trello"
)

type App struct {
	appKey string
	token string
	boardId string
	username string

  startDate time.Time

  trelloClient *trello.Client
}

func New(appKey, token, boardId, username string, startDate time.Time) *App {

	client := trello.NewClient(appKey, token)

  return &App {
    appKey: appKey,
    token: token,
    boardId: boardId,
    username: username,
    startDate: startDate,
    trelloClient:  client,
  }
}

func (app *App) Run() {
  interestedActions := app.GetInterestedActions()
  printActions(interestedActions)
}

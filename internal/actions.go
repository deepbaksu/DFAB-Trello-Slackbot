package internal

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adlio/trello"
	"github.com/dl4ab/DFAB-Trello-Slackbot/trelloutils"
	log "github.com/sirupsen/logrus"
)

type InterestedActions map[string]map[string]*trello.Action

func (app *App) GetInterestedActions() InterestedActions {
	trelloBoard, err := app.trelloClient.GetBoard(app.boardId, trello.Defaults())
	if err != nil {
		log.WithField("err", err).Fatalf("client.GetBoard(%v) has failed", app.boardId)
	}

	actions, _ := trelloBoard.GetActions(trello.Defaults())

	// ListName => CardID => Action
	// So that we can get the latest card information.
	interestedActions := make(map[string]map[string]*trello.Action)

	for _, action := range actions {
		if action.Date.After(app.startDate) && trelloutils.GetMemberFromAction(action) == app.username {
			listName, err := trelloutils.GetListNameFromAction(action)
			if err != nil {
				log.WithError(err).Warn("Skipping unknown list name")
				continue
			}
			if action.Data.Card == nil {
				log.WithFields(
					log.Fields{
						"action":   fmt.Sprintf("%+v", action),
						"data":     fmt.Sprintf("%+v", action.Data),
						"data.Old": fmt.Sprintf("%+v", action.Data.Old),
					},
				).Info("Skipping an event that does not contain a card")
				continue
			}

			if _, ok := interestedActions[listName]; !ok {
				interestedActions[listName] = make(map[string]*trello.Action)
			}

			actionData, ok := interestedActions[listName][action.Data.Card.ID]
			if !ok {
				interestedActions[listName][action.Data.Card.ID] = action
			} else if actionData.Date.Before(action.Date) {
				interestedActions[listName][action.Data.Card.ID] = action
			}
		}
	}

	return interestedActions
}

func printActions(interestedActions InterestedActions) {
	// Try to get sorted list names so that it can print in a correct order.
	var listNames []string
	for listName, _ := range interestedActions {
		listNames = append(listNames, listName)
	}

	sort.Strings(listNames)

	for _, listName := range listNames {
		fmt.Println(listName)
		fmt.Println(strings.Repeat("=", len(listName)))

		var cards []string
		for _, action := range interestedActions[listName] {
			cards = append(cards, action.Data.Card.Name)
		}

		trelloutils.PrintCards(cards)
		fmt.Println()
	}
}

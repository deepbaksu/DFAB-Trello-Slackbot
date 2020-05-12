package internal

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/adlio/trello"
	"github.com/dl4ab/DFAB-Trello-Slackbot/trelloutils"
	log "github.com/sirupsen/logrus"
)

type InterestedActions map[string]map[string]*trello.Action

type CardInfo struct {
	id        string
	updatedAt time.Time
	content   string
	listName  string
}

func CardInfoFromAction(action *trello.Action) (*CardInfo, error) {
	listName, err := trelloutils.GetListNameFromAction(action)
	if err != nil {
		return nil, err
	}
	return &CardInfo{
		id:        action.Data.Card.ID,
		updatedAt: action.Date,
		content:   action.Data.Card.Name,
		listName:  listName,
	}, nil
}

// 1. collect all the interesting cards
//    card_id => (action_time, card_content, list_name)
// 2. get all list names and build a map (list_name => card[])
// 3. for each list, print cards
func (app *App) GetInterestedActions() map[string]*CardInfo {
	trelloBoard, err := app.trelloClient.GetBoard(app.boardId, trello.Defaults())
	if err != nil {
		log.WithField("err", err).Fatalf("client.GetBoard(%v) has failed", app.boardId)
	}

	actions, _ := trelloBoard.GetActions(trello.Defaults())

	interestedCards := make(map[string]*CardInfo)

	for _, action := range actions {
		if action.Date.After(app.startDate) && trelloutils.GetMemberFromAction(action) == app.username {
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

			newCardInfo, err := CardInfoFromAction(action)
			if err != nil {
				log.WithError(err).Warn("CardInfoFromAction has failed")
				continue
			}

			existingCardInfo, ok := interestedCards[newCardInfo.id]

			if !ok {
				// not found
				interestedCards[newCardInfo.id] = newCardInfo
				continue
			}

			if existingCardInfo.updatedAt.Before(newCardInfo.updatedAt) {
				interestedCards[newCardInfo.id] = newCardInfo
			}
		}
	}

	return interestedCards
}

func printActions(interestedCards map[string]*CardInfo) {
	// Try to get sorted list names so that it can print in a correct order.
	var listNames []string

	listNameToCardInfo := make(map[string][]string)

	for _, cardInfo := range interestedCards {
		listNames = append(listNames, cardInfo.listName)

		lists, ok := listNameToCardInfo[cardInfo.listName]
		if !ok {
			lists = make([]string, 0)
		}
		listNameToCardInfo[cardInfo.listName] = append(lists, cardInfo.content)
	}

	sort.Strings(listNames)

  prevListName := ""

	for _, listName := range listNames {
    if prevListName == listName {
      // this list is already printed.
      continue
    }
		fmt.Println(listName)
		fmt.Println(strings.Repeat("=", len(listName)))

		var cards []string

		for _, cardContent := range listNameToCardInfo[listName] {
			cards = append(cards, cardContent)
		}

		trelloutils.PrintCards(cards)
    // Add an extra line between lists.
		fmt.Println()

    // track the list name so we don't print the same list multiple times.
    prevListName = listName
	}
}

package trelloutils

import (
	"fmt"

	"github.com/adlio/trello"
)

func PrintCards(cards []string) {
	for _, card := range cards {
		fmt.Printf("- %v\n", card)
	}
}

func GetMemberFromAction(action *trello.Action) string {
	if action.MemberCreator != nil {
		return action.MemberCreator.Username
	}

	if action.Member != nil {
		return action.Member.Username
	}

	return ""
}

func GetListNameFromAction(action *trello.Action) (string, error) {
	if action.Data.List != nil {
		return action.Data.List.Name, nil
	}
	if action.Data.ListAfter != nil {
		return action.Data.ListAfter.Name, nil
	}

	return "", fmt.Errorf("Not able to extract a list name from the action: %+v", action)
}

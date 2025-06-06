package task

import (
	"fmt"
	"strings"
)

type Action int

const (
	ActionAdd Action = iota
	ActionUpdate
	ActionDelete
	ActionList
	ActionMark
)

var actionName = map[Action]string{
	ActionAdd:    "add",
	ActionUpdate: "update",
	ActionDelete: "delete",
	ActionList:   "list",
	ActionMark:   "mark",
}

var nameToAction = map[string]Action{
	"add":    ActionAdd,
	"update": ActionUpdate,
	"delete": ActionDelete,
	"list":   ActionList,
	"mark":   ActionMark,
}

func (a Action) String() string {
	return actionName[a]
}

func ParseAction(input string) (Action, error) {
	a, ok := nameToAction[strings.ToLower(input)]
	if !ok {
		return Action(0), fmt.Errorf("invalid action: %s", input)
	}
	return a, nil
}

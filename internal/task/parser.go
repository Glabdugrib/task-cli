package task

import (
	"fmt"
	"strconv"
)

type ParsedArgs struct {
	Action      Action
	ID          uint
	Description string
	Status      *Status
}

func ValidateArgs(args []string) (ParsedArgs, error) {
	if len(args) < 2 {
		return ParsedArgs{}, fmt.Errorf("you must provide at least one command")
	}

	action, err := ParseAction(args[1])
	if err != nil {
		return ParsedArgs{}, fmt.Errorf(`you must provide a valid action ("add", "update", "delete", "list", "mark"), "%s" provided`, args[1])
	}

	parsed := ParsedArgs{Action: action}

	switch action {
	case ActionAdd:
		if len(args) < 3 {
			return parsed, fmt.Errorf(`"add" action requires a description`)
		}
		parsed.Description = args[2]

	case ActionUpdate:
		if len(args) < 4 {
			return parsed, fmt.Errorf(`"update" action requires an id and a new description`)
		}
		id, err := strconv.ParseUint(args[2], 10, 64)
		if err != nil {
			return parsed, fmt.Errorf("invalid id for update: %v", err)
		}
		parsed.ID = uint(id)
		parsed.Description = args[3]

	case ActionDelete:
		if len(args) < 3 {
			return parsed, fmt.Errorf(`"delete" action requires an id`)
		}
		id, err := strconv.ParseUint(args[2], 10, 64)
		if err != nil {
			return parsed, fmt.Errorf("invalid id for %s: %v", action, err)
		}
		parsed.ID = uint(id)

	case ActionMark:
		if len(args) < 4 {
			return parsed, fmt.Errorf(`"mark" action requires an id and a status`)
		}
		id, err := strconv.ParseUint(args[2], 10, 64)
		if err != nil {
			return parsed, fmt.Errorf("invalid id for %s: %v", action, err)
		}
		parsed.ID = uint(id)

		status, err := ParseStatus(args[3])
		if err != nil {
			return ParsedArgs{}, fmt.Errorf(`you must provide a valid status ("pending", "in_progress", "done"), "%s" provided`, args[3])
		}
		parsed.Status = &status

	case ActionList:
		if len(args) > 2 {
			status, err := ParseStatus(args[2])
			if err != nil {
				return ParsedArgs{}, fmt.Errorf(`you must provide a valid status ("pending", "in_progress", "done"), "%s" provided`, args[2])
			}
			parsed.Status = &status
		}

	default:
		return parsed, fmt.Errorf("unsupported action: %s", action)
	}

	return parsed, nil
}

func PrintArgs(args ParsedArgs) {
	fmt.Println("\nCHOSEN ACTION:")
	fmt.Printf("Action: %s\n", args.Action)
	if args.ID != 0 {
		fmt.Printf("ID: %d\n", args.ID)
	}
	if args.Description != "" {
		fmt.Printf("Description: %s\n", args.Description)
	}
	if args.Status != nil {
		fmt.Printf("Status: %s\n", args.Status.String())
	}
}

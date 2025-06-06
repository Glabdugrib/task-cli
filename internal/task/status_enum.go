package task

import (
	"fmt"
	"strings"
)

type Status int

const (
	StatusPending Status = iota
	StatusInProgress
	StatusDone
)

var statusName = map[Status]string{
	StatusPending:    "pending",
	StatusInProgress: "in_progress",
	StatusDone:       "done",
}

var nameToStatus = map[string]Status{
	"pending":     StatusPending,
	"in_progress": StatusInProgress,
	"done":        StatusDone,
}

// Implement the fmt.Stringer interface for Status
func (s Status) String() string {
	return statusName[s]
}

func ParseStatus(input string) (Status, error) {
	s, ok := nameToStatus[strings.ToLower(input)]
	if !ok {
		return Status(0), fmt.Errorf("invalid status: %s", input)
	}
	return s, nil
}

// Implements the json.Marshaler interface for Status
func (s Status) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, `"%s"`, s.String()), nil
}

// Implements the json.Unmarshaler interface for Status
func (s *Status) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	parsed, err := ParseStatus(str)
	if err != nil {
		return err
	}
	*s = parsed
	return nil
}

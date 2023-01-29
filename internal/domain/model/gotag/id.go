package gotag

import "fmt"

type ID string

var errEmpty = fmt.Errorf("empty ID")

func NewID(id string) (ID, error) {
	if id == "" {
		return ID(""), errEmpty
	}

	return ID(id), nil
}

func (id ID) String() string {
	return string(id)
}

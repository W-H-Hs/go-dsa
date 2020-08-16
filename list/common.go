package list

import (
	"errors"
)

const (
	indexOutOfBoundError  = "index out of bound"
	elementNotInListError = "element not in list"

	invalidIdx = -1
)

type listCommon struct {
	length int
}

func (c *listCommon) Len() int {
	return c.length
}

func (c *listCommon) IsEmpty() bool {
	return c.length == 0
}

func (c *listCommon) indexCheck(idx, size int) error {
	if idx < 0 || idx >= size {
		return errors.New(indexOutOfBoundError)
	}
	return nil
}

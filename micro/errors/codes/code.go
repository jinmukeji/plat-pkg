package codes

import (
	"fmt"
	"strconv"
)

type Code uint32

func (c *Code) Message() string {
	msg := codeToMsg[*c]
	return msg
}

func (c *Code) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	if c == nil {
		return fmt.Errorf("nil receiver passed to UnmarshalJSON")
	}

	if ci, err := strconv.ParseUint(string(b), 10, 32); err != nil {
		if _, ok := codeToMsg[Code(ci)]; !ok {
			return fmt.Errorf("invalid code: %q", ci)
		}

		*c = Code(ci)
		return nil
	}

	return fmt.Errorf("invalid code: %q", string(b))
}

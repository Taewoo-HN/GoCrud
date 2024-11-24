package errors

import "fmt"

const (
	NotFoundUser = iota
	DBError
)

var errMessage = map[int64]string{
	NotFoundUser: "User not found",
	DBError:      "DB error",
}

func Errorf(code int64,
	args ...interface{}) error {
	if message, ok := errMessage[code]; ok {
		return fmt.Errorf(message, args...)

	} else {
		return fmt.Errorf("Not found error Code")
	}
}

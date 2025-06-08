package custom_error

import (
	"errors"
)

var StatusNotActive = errors.New("user status not active")

package slct

import (
	"errors"
	"fmt"
)

var (
	errPkg    = errors.New("select")
	errPrompt = fmt.Errorf("%w: prompt", errPkg)
	errWrite  = fmt.Errorf("%w: unable to write", errPkg)
)

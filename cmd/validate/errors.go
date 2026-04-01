package validate

import (
	"errors"
	"fmt"
)

var (
	errPkg      = errors.New("validate")
	errValidate = fmt.Errorf("%w: validate", errPkg)
	errRead     = fmt.Errorf("%w: read", errPkg)
	errMarshal  = fmt.Errorf("%w: marshal", errPkg)
	errWrite    = fmt.Errorf("%w: unable to write", errPkg)
	errFormat   = fmt.Errorf("%w: incorrect format", errPkg)
)

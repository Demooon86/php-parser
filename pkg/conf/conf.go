package conf

import (
	"github.com/Demooon86/php-parser/pkg/errors"
	"github.com/Demooon86/php-parser/pkg/version"
)

type Config struct {
	Version          *version.Version
	ErrorHandlerFunc func(e *errors.Error)
}

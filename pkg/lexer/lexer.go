package lexer

import (
	"errors"

	"github.com/Demooon86/php-parser/internal/php7"
	"github.com/Demooon86/php-parser/internal/php8"
	"github.com/Demooon86/php-parser/pkg/conf"
	"github.com/Demooon86/php-parser/pkg/token"
)

var ErrVersionOutOfRange = errors.New("the version is out of supported range")

type Lexer interface {
	Lex() *token.Token
}

func New(src []byte, config conf.Config) (Lexer, error) {
	if config.Version.InPhp7Range() {
		return php7.NewLexer(src, config), nil
	}

	if config.Version.InPhp8Range() {
		return php8.NewLexer(src, config), nil
	}

	return nil, ErrVersionOutOfRange
}

package php8_test

import (
	"testing"

	"github.com/VKCOM/php-parser/internal/php8"
	"github.com/VKCOM/php-parser/internal/tester"
	"github.com/VKCOM/php-parser/pkg/conf"
	"github.com/VKCOM/php-parser/pkg/token"
	"gotest.tools/assert"
)

func TestNullsafeMethodCallTokens(t *testing.T) {
	suite := tester.NewLexerTokenStringTestSuite(t)
	suite.UsePHP8()
	suite.Code = "<?php $a?->foo();"
	suite.Expected = []string{
		"$a",
		"?->",
		"foo",
		"(",
		")",
		";",
	}
	suite.Run()
}

func TestNullsafePropertyFetchTokens(t *testing.T) {
	suite := tester.NewLexerTokenStringTestSuite(t)
	suite.UsePHP8()
	suite.Code = "<?php $a?->prop;"
	suite.Expected = []string{
		"$a",
		"?->",
		"prop",
		";",
	}
	suite.Run()
}

func TestNullsafePropertyFetchInStringTokens(t *testing.T) {
	suite := tester.NewLexerTokenStringTestSuite(t)
	suite.UsePHP8()
	suite.Code = "<?php \"$a?->prop\";"
	suite.Expected = []string{
		"\"",
		"$a",
		"?->",
		"prop",
		"\"",
		";",
	}
	suite.Run()
}

func TestNullsafeMethodCallTokensFreeFloating(t *testing.T) {
	suite := tester.NewLexerTokenFreeFloatingTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
	$a ?-> bar ( '' ) ;`

	suite.Expected = [][]*token.Token{
		{
			{
				ID:    token.T_OPEN_TAG,
				Value: []byte("<?php"),
			},
			{
				ID:    token.T_WHITESPACE,
				Value: []byte("\n\t"),
			},
		},
		{
			{
				ID:    token.T_WHITESPACE,
				Value: []byte(" "),
			},
		},
		{
			{
				ID:    token.T_WHITESPACE,
				Value: []byte(" "),
			},
		},
		{
			{
				ID:    token.T_WHITESPACE,
				Value: []byte(" "),
			},
		},
		{
			{
				ID:    token.T_WHITESPACE,
				Value: []byte(" "),
			},
		},
		{
			{
				ID:    token.T_WHITESPACE,
				Value: []byte(" "),
			},
		},
		{
			{
				ID:    token.T_WHITESPACE,
				Value: []byte(" "),
			},
		},
	}

	suite.Run()
}

func TestMatchStringTokens(t *testing.T) {
	suite := tester.NewLexerTokenStringTestSuite(t)
	suite.UsePHP8()
	suite.Code = "<?php match($a) {}"
	suite.Expected = []string{
		"match",
		"(",
		"$a",
		")",
		"{",
		"}",
	}
	suite.Run()
}

func TestMatchWithConditionStringTokens(t *testing.T) {
	suite := tester.NewLexerTokenStringTestSuite(t)
	suite.UsePHP8()
	suite.Code = "<?php match($a) { 10 => 100 }"
	suite.Expected = []string{
		"match",
		"(",
		"$a",
		")",
		"{",
		"10",
		"=>",
		"100",
		"}",
	}
	suite.Run()
}

func TestMatchWithDefaultStringTokens(t *testing.T) {
	suite := tester.NewLexerTokenStringTestSuite(t)
	suite.UsePHP8()
	suite.Code = "<?php match($a) { default => 10 }"
	suite.Expected = []string{
		"match",
		"(",
		"$a",
		")",
		"{",
		"default",
		"=>",
		"10",
		"}",
	}
	suite.Run()
}

func TestAttributeTokens(t *testing.T) {
	suite := tester.NewLexerTokenStringTestSuite(t)
	suite.UsePHP8()
	suite.Code = "<?php #[ FooAttribute]"
	suite.Expected = []string{
		"#[",
		"FooAttribute",
		"]",
	}
	suite.Run()
}

func TestCommentEnd2(t *testing.T) {
	src := `<?php //`

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},

		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
		{
			ID:    token.T_COMMENT,
			Value: []byte("//"),
		},
	}

	lexer := php8.NewLexer([]byte(src), conf.Config{})

	tkn := lexer.Lex()

	actual := tkn.FreeFloating
	for _, v := range actual {
		v.Position = nil
	}

	assert.DeepEqual(t, expected, actual)
}

func TestCommentEnd3(t *testing.T) {
	src := `<?php #`

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},

		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
		{
			ID:    token.T_COMMENT,
			Value: []byte("#"),
		},
	}

	lexer := php8.NewLexer([]byte(src), conf.Config{})

	tkn := lexer.Lex()

	actual := tkn.FreeFloating
	for _, v := range actual {
		v.Position = nil
	}

	assert.DeepEqual(t, expected, actual)
}

func TestAttribute(t *testing.T) {
	src := `<?php #[ FooAttribute]`

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},
		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
	}

	lexer := php8.NewLexer([]byte(src), conf.Config{})

	tkn := lexer.Lex()

	actual := tkn.FreeFloating
	for _, v := range actual {
		v.Position = nil
	}

	assert.DeepEqual(t, expected, actual)

	assert.DeepEqual(t, token.T_ATTRIBUTE, tkn.ID)
	assert.DeepEqual(t, "#[", string(tkn.Value))
}

func TestNamespaceFullyQualifiedTokens(t *testing.T) {
	suite := tester.NewLexerTokenStructTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php use \Foo;`
	suite.Expected = []*token.Token{
		{
			ID:    php8.T_USE,
			Value: []byte("use"),
		},
		{
			ID:    php8.T_NAME_FULLY_QUALIFIED,
			Value: []byte(`\Foo`),
		},
		{
			ID:    ';',
			Value: []byte(";"),
		},
	}
	suite.Run()
}

func TestNamespaceFullyQualifiedWithKeywordsTokens(t *testing.T) {
	suite := tester.NewLexerTokenStructTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php use \Foo\match\fn;`
	suite.Expected = []*token.Token{
		{
			ID:    php8.T_USE,
			Value: []byte("use"),
		},
		{
			ID:    php8.T_NAME_FULLY_QUALIFIED,
			Value: []byte(`\Foo\match\fn`),
		},
		{
			ID:    ';',
			Value: []byte(";"),
		},
	}
	suite.Run()
}

func TestNamespaceQualifiedTokens(t *testing.T) {
	suite := tester.NewLexerTokenStructTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php namespace Boo\Foo;`
	suite.Expected = []*token.Token{
		{
			ID:    php8.T_NAMESPACE,
			Value: []byte("namespace"),
		},
		{
			ID:    php8.T_NAME_QUALIFIED,
			Value: []byte(`Boo\Foo`),
		},
		{
			ID:    ';',
			Value: []byte(";"),
		},
	}
	suite.Run()
}

func TestNamespaceRelativeTokens(t *testing.T) {
	suite := tester.NewLexerTokenStructTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php namespace\match;`
	suite.Expected = []*token.Token{
		{
			ID:    php8.T_NAME_RELATIVE,
			Value: []byte(`namespace\match`),
		},
		{
			ID:    ';',
			Value: []byte(";"),
		},
	}
	suite.Run()
}

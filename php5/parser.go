package php5

import (
	"io"

	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/scanner"
)

func (lval *yySymType) Token(t *scanner.Token) {
	lval.token = t
}

// Parser structure
type Parser struct {
	*scanner.Lexer
	path            string
	currentToken    *scanner.Token
	positionBuilder *parser.PositionBuilder
	errors          []*errors.Error
	rootNode        node.Node
	comments        parser.Comments
}

// NewParser creates and returns new Parser
func NewParser(src io.Reader, path string) *Parser {
	lexer := scanner.NewLexer(src, path)

	return &Parser{
		lexer,
		path,
		nil,
		nil,
		nil,
		nil,
		nil,
	}
}

// Lex proxy to lexer Lex
func (l *Parser) Lex(lval *yySymType) int {
	t := l.Lexer.Lex(lval)
	l.currentToken = lval.token
	return t
}

func (l *Parser) Error(msg string) {
	l.errors = append(l.errors, errors.NewError(msg, l.currentToken))
}

// Parse the php7 Parser entrypoint
func (l *Parser) Parse() int {
	// init
	l.errors = nil
	l.rootNode = nil
	l.comments = parser.Comments{}
	l.positionBuilder = &parser.PositionBuilder{}

	// parse

	return yyParse(l)
}

func (l *Parser) listGetFirstNodeComments(list []node.Node) []*comment.Comment {
	if len(list) == 0 {
		return nil
	}

	node := list[0]

	return l.comments[node]
}

// GetPath return path to file
func (l *Parser) GetPath() string {
	return l.path
}

// GetRootNode returns root node
func (l *Parser) GetRootNode() node.Node {
	return l.rootNode
}

// GetErrors returns errors list
func (l *Parser) GetErrors() []*errors.Error {
	return l.errors
}

// GetComments returns comments list
func (l *Parser) GetComments() parser.Comments {
	return l.comments
}

// helpers

func lastNode(nn []node.Node) node.Node {
	return nn[len(nn)-1]
}

func firstNode(nn []node.Node) node.Node {
	return nn[0]
}

func isDollar(r rune) bool {
	return r == '$'
}

func (p *Parser) returnTokenToPool(yyDollar []yySymType, yyVAL *yySymType) {
	for i := 1; i < len(yyDollar); i++ {
		if yyDollar[i].token != nil {
			p.TokenPool.Put(yyDollar[i].token)
		}
		yyDollar[i].token = nil
	}
	yyVAL.token = nil
}

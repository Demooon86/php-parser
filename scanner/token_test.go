package scanner_test

import (
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/comment"

	"github.com/z7zmey/php-parser/scanner"
)

func TestToken(t *testing.T) {
	tkn := &scanner.Token{
		Value:     `foo`,
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}

	c := []*comment.Comment{
		comment.NewComment("test comment", nil),
	}

	tkn.Comments = c

	if !reflect.DeepEqual(tkn.Comments, c) {
		t.Errorf("comments are not equal\n")
	}

	if tkn.String() != `foo` {
		t.Errorf("token value is not equal\n")
	}
}

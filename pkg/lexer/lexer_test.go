package lexer

import (
	"fmt"
	"os"
	"testing"

	"github.com/fcabjolsky/monkey-interpreter-go/pkg/token"
	"github.com/stretchr/testify/assert"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, test := range tests {
		tok := l.NextToken()
    ok := assert.Equal(t, test.expectedType, tok.Type, fmt.Sprintf("Wrong type %v tok: %v expec: %v", i, tok.Type, test.expectedType))
    if !ok {
      os.Exit(1)
    }
		ok = assert.Equal(t, test.expectedLiteral, tok.LiteralValue, fmt.Sprintf("Wrong literal %v", i))
    if !ok {
      os.Exit(1)
    }

	}

}


func TestNextTokenComplete(t *testing.T) {
	input := `let five_t = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};
let result = add(five, ten);
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five_t"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, test := range tests {
		tok := l.NextToken()

    ok := assert.Equal(t, test.expectedType, tok.Type, fmt.Sprintf("Wrong type %v tok: %v expec: %v", i, tok.Type, test.expectedType))
    if !ok {
      os.Exit(1)
    }
		ok = assert.Equal(t, test.expectedLiteral, tok.LiteralValue, fmt.Sprintf("Wrong literal %v", i))
    if !ok {
      os.Exit(1)
    }

	}


}

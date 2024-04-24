package token

type TokenType byte

const (
	ILLEGAL = 0
	EOF     = 1

	// Identifiers + lilterals
	IDENT = 2
	INT   = 3

	// Operators
	ASSIGN = 4
	PLUS   = 5

	//DELIMITERS
	COMMA     = 6
	SEMICOLON = 7

	LPAREN = 8
	RPAREN = 9
	LBRACE = 10
	RBRACE = 11

	//keyboards
	FUNCTION = 12
	LET      = 13
)

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case IDENT:
		return "IDENT"
	case INT:
		return "INT"
	case ASSIGN:
		return "="
	case PLUS:
		return "+"
	case COMMA:
		return ","
	case SEMICOLON:
		return ";"
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	case LBRACE:
		return "{"
	case RBRACE:
		return "}"
	case FUNCTION:
		return "FUNCTION"
	case LET:
		return "LET"
	}
	return ""

}

type Token struct {
	Type         TokenType
	LiteralValue string
}

func FromByte(literal byte) Token {
	switch literal {
	case '=':
		return Token{Type: ASSIGN, LiteralValue: string(literal)}
	case ';':
		return Token{Type: SEMICOLON, LiteralValue: string(literal)}
	case '(':
		return Token{Type: LPAREN, LiteralValue: string(literal)}
	case ')':
		return Token{Type: RPAREN, LiteralValue: string(literal)}
	case ',':
		return Token{Type: COMMA, LiteralValue: string(literal)}
	case '+':
		return Token{Type: PLUS, LiteralValue: string(literal)}
	case '{':
		return Token{Type: LBRACE, LiteralValue: string(literal)}
	case '}':
		return Token{Type: RBRACE, LiteralValue: string(literal)}
	case 0:
		return Token{Type: EOF, LiteralValue: ""}

	}
	return Token{}
}

func FromString(ident string) Token {
	switch ident {
	case "let":
		return Token{Type: LET, LiteralValue: ident}
	case "fn":
		return Token{Type: FUNCTION, LiteralValue: ident}
	}

	return Token{Type: IDENT, LiteralValue: ident}
}

func FromStringNumber(number string) Token {
	return Token{
		Type:         INT,
		LiteralValue: number,
	}
}


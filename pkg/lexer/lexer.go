package lexer

import (
	"unicode"

	"github.com/fcabjolsky/monkey-interpreter-go/pkg/token"
)

type Lexer struct {
	input        string
	position     int  //current position in input - current char
	readPosition int  // current reading position - after current char
	ch           byte //current char being read
}

func New(input string) *Lexer {
	l := &Lexer{
		input:        input,
		position:     0,
		readPosition: 0,
		ch:           0,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) skipWhiteSpaces() {
	for unicode.IsSpace(rune(l.ch)) {
		l.readChar()
	}
}


func (l *Lexer) readIdentifier() string {
	init := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[init:l.position]
}

func (l *Lexer) readNumber() string {
	init := l.position
	for unicode.IsDigit(rune(l.ch)) {
		l.readChar()
	}
	return l.input[init:l.position]
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhiteSpaces()
	if isLetter(l.ch) {
		identifier := l.readIdentifier()
		return token.FromString(identifier)
	}

  if unicode.IsDigit(rune(l.ch)) {
    number := l.readNumber()
		return token.FromStringNumber(number)
  }


	tok := token.FromByte(l.ch)
	l.readChar()
	return tok
}

func isLetter(ch byte) bool {
  return unicode.IsLetter(rune(ch)) || ch == '_'
}



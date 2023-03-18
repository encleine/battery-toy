package lexer

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/encleine/battery-toy/src/token"
)

type Lexer struct {
	input     *bufio.Reader
	runeCount int
	ch        rune
}

func New(input string) *Lexer {
	I := bufio.NewReader(strings.NewReader(input))
	L := &Lexer{input: I}
	L.readRune()
	return L
}
func (L *Lexer) readRune() {
	r := L.input
	c, _, err := r.ReadRune()

	if err != nil && err != io.EOF {
		os.Exit(1)
		log.Fatal(err)
	}
	L.ch = c
	L.runeCount++
}
func (L *Lexer) peekRune() rune {
	r := L.input
	c, _, err := r.ReadRune()

	if err != nil && err != io.EOF {
		os.Exit(1)
		log.Fatal(err)
	}
	L.unreadRune()
	return c
}
func (L *Lexer) skipSpace() {
	for unicode.IsSpace(L.ch) && L.ch != '\n' {
		L.readRune()
	}
}
func (L *Lexer) skipComment() {
	if L.ch+L.peekRune() == 94 {
		for L.ch != '\n' {
			L.readRune()
		}
	}
}
func (L *Lexer) unreadRune() {
	r := L.input
	r.UnreadRune()
	L.runeCount--
}
func (L *Lexer) PCO(qToken, nqToken token.TokenType) token.Token {
	c := L.ch
	if L.peekRune() == '=' {
		L.readRune()
		return token.NewToken(qToken, string(c)+string(L.ch))
	} else {
		return token.NewToken(nqToken, L.ch)
	}
}
func (L *Lexer) NextToken() token.Token {

	L.skipSpace()
	L.skipComment()
	var tok token.Token

	if val, ok := token.LitToken[L.ch]; ok {
		L.readRune()
		return val
	}

	switch L.ch {
	case '=':
		tok = L.PCO(token.Equal, token.Assign)
	case '!':
		tok = L.PCO(token.NotEqual, token.Bang)
	case '<':
		tok = L.PCO(token.LessEqual, token.Less)
	case '>':
		tok = L.PCO(token.GreaterEqual, token.Greater)
	case '+':
		tok = L.PCO(token.PlusAssign, token.Plus)
	case '-':
		tok = L.PCO(token.MinusAssign, token.Minus)
	case '*':
		tok = L.PCO(token.AsterAssign, token.Aster)
	case '/':
		tok = L.PCO(token.SlahAssign, token.Slash)
	case '^':
		tok = L.PCO(token.PowAssign, token.Pow)
	case '.':
		if unicode.IsDigit(L.peekRune()) {
			return token.OrNumber(L.readNum())
		}
		tok = token.NewToken(token.Dot, '.')
	default:
		switch {
		case unicode.IsLetter(L.ch):
			return token.OrKeyword(L.readIdent())
		case unicode.IsDigit(L.ch):
			return token.OrNumber(L.readNum())
		}
		tok = token.NewToken(token.Undef, L.ch)
	}

	L.readRune()
	return tok
}
func (L *Lexer) readIdent() string {
	var runes []rune
	for unicode.IsLetter(L.ch) || unicode.IsDigit(L.ch) || L.ch == '_' {
		runes = append(runes, L.ch)
		L.readRune()
	}
	return string(runes)
}
func (L *Lexer) readNum() string {
	var runes []rune
	for unicode.IsDigit(L.ch) {
		runes = append(runes, L.ch)
		L.readRune()
	}
	if L.ch == '.' {
		runes = append(runes, L.ch)
		L.readRune()
		for unicode.IsDigit(L.ch) {
			runes = append(runes, L.ch)
			L.readRune()
		}
	}
	return string(runes)
}

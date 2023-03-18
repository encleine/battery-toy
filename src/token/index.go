package token

import "strings"

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF = iota
	Newline
	// singles
	Assign
	Bang
	Less
	Greater
	Plus
	Minus
	Aster
	Slash
	Pow
	Doublequotes
	Singlequotes
	// doubles
	Equal
	NotEqual
	LessEqual
	GreaterEqual
	PlusAssign
	MinusAssign
	AsterAssign
	SlahAssign
	PowAssign
	//
	Semicolon
	Colon
	Lparen
	Rparen
	Lbrace
	Rbrace
	Lsquare
	Rsquare
	Comma
	Dot
	//
	Identi
	Func
	Let
	// Const
	For
	If
	Else
	// Match
	// Case
	Return
	True
	False

	Untyped_int
	Untyped_float
	Untyped_string

	Undef
)

var (
	LitToken = map[rune]Token{
		'\n': NewToken(Newline, ""),
		';':  NewToken(Semicolon, ";"),
		':':  NewToken(Colon, ":"),
		'(':  NewToken(Lparen, "("),
		')':  NewToken(Rparen, ")"),
		'{':  NewToken(Lbrace, "{"),
		'}':  NewToken(Rbrace, "}"),
		'[':  NewToken(Lsquare, "["),
		']':  NewToken(Rsquare, "]"),
		',':  NewToken(Comma, ","),
		0:    NewToken(EOF, ""),
		'"':  NewToken(Doublequotes, "\""),
		'\'': NewToken(Singlequotes, "'"),
	}
	keywords = map[string]TokenType{
		"func": Func,
		"let":  Let,
		// "const": Const,
		"for":  For,
		"if":   If,
		"else": Else,
		// "match":  Match,
		// "case":   Case,
		"return": Return,
		"true":   True,
		"false":  False,
	}
)

func OrKeyword(literal string) Token {
	if val, ok := keywords[literal]; ok {
		return NewToken(val, literal)
	}
	return NewToken(Identi, literal)
}
func OrNumber(literal string) Token {
	if strings.Contains(literal, ".") {
		return NewToken(Untyped_float, literal)
	}
	return NewToken(Untyped_int, literal)
}
func NewToken[T rune | string](Type TokenType, Literal T) Token {
	return Token{Type: Type, Literal: string(Literal)}
}

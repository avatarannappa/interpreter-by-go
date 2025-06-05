package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   TokenType = "ILLEGAL"
	EOF       TokenType = "EOF"
	ASSIGN    TokenType = "="
	PLUS      TokenType = "+"
	MINUS     TokenType = "-"
	BANG      TokenType = "!"
	ASTERISK  TokenType = "*"
	SLASH     TokenType = "/"
	LT        TokenType = "<"
	GT        TokenType = ">"
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"
	FUNCTION  TokenType = "FUNCTION"
	LET       TokenType = "LET"
	IDENT     TokenType = "IDENT"
	RETURN    TokenType = "RETURN"
	TRUE      TokenType = "TRUE"
	FALSE     TokenType = "FALSE"
	IF        TokenType = "IF"
	ELSE      TokenType = "ELSE"
	EQ        TokenType = "=="
	NOT_EQ    TokenType = "!="
	INT       TokenType = "INT"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

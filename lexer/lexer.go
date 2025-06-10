package lexer

import "awesomeProject/token"

type Lexer struct {
	input        string
	position     int  // 当前字符串
	realPosition int  // 当前字符串之后的下一个字符串
	ch           byte // 当前字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // 初始化p，rp，ch
	return l
}

// TODO 支持unicode. ch 改为 rune, 读取字节长度需要判断.
func (l *Lexer) readChar() {
	if l.realPosition >= len(l.input) {
		l.ch = 0 // end
	} else {
		l.ch = l.input[l.realPosition]
	}
	l.position = l.realPosition
	l.realPosition++
}

func (l *Lexer) peekChar() byte {
	if l.realPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.realPosition]
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.EQ, Literal: l.makeTwoCharToken()}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.NOT_EQ, Literal: l.makeTwoCharToken()}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '[':
		tok = newToken(token.LBRANCKET, l.ch)
	case ']':
		tok = newToken(token.RBRANCKET, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
		l.readChar()
	}
}

func (l *Lexer) makeTwoCharToken() string {
	l.readChar()
	return l.input[l.position-1 : l.realPosition]
}

func (l *Lexer) readString() string {
	start := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[start:l.position]
}

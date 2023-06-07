package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	char         byte // current char under examination
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	var currentToken token.Token

	lexer.skipWhitespace()

	switch lexer.char {
	case '=':
		if lexer.peekChar() == '=' {
			char := lexer.char
			lexer.readChar()
			literal := string(char) + string(lexer.char)
			currentToken = token.Token{Type: token.EQ, Literal: literal}
		} else {
			currentToken = newToken(token.ASSIGN, lexer.char)
		}
	case ';':
		currentToken = newToken(token.SEMICOLON, lexer.char)
	case '(':
		currentToken = newToken(token.LPAREN, lexer.char)
	case ')':
		currentToken = newToken(token.RPAREN, lexer.char)
	case ',':
		currentToken = newToken(token.COMMA, lexer.char)
	case '+':
		currentToken = newToken(token.PLUS, lexer.char)
	case '{':
		currentToken = newToken(token.LBRACE, lexer.char)
	case '}':
		currentToken = newToken(token.RBRACE, lexer.char)
	case '!':
		if lexer.peekChar() == '=' {
			char := lexer.char
			lexer.readChar()
			literal := string(char) + string(lexer.char)
			currentToken = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			currentToken = newToken(token.BANG, lexer.char)
		}
	case '-':
		currentToken = newToken(token.MINUS, lexer.char)
	case '/':
		currentToken = newToken(token.SLASH, lexer.char)
	case '*':
		currentToken = newToken(token.ASTERISK, lexer.char)
	case '<':
		currentToken = newToken(token.LT, lexer.char)
	case '>':
		currentToken = newToken(token.GT, lexer.char)
	case 0:
		currentToken.Literal = ""
		currentToken.Type = token.EOF
	default:
		if isLetter(lexer.char) {
			currentToken.Literal = lexer.readIdentifier()
			currentToken.Type = token.LookupIdentifier(currentToken.Literal)
			return currentToken
		} else if isDigit(lexer.char) {
			currentToken.Type = token.INT
			currentToken.Literal = lexer.readNumber()
			return currentToken
		} else {
			currentToken = newToken(token.ILLEGAL, lexer.char)
		}
	}

	lexer.readChar()
	return currentToken
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.char) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.char) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.char == ' ' || lexer.char == '\t' || lexer.char == '\n' || lexer.char == '\r' {
		lexer.readChar()
	}
}

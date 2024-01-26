package lexer

import (
	"monkey/token"
)


type Lexer struct{

    input string
    position int // current position on input
    readposition int //next char in buffer
    ch byte //current char 

}

func New (input string) *Lexer{

    l := &Lexer{input:input}
    l.readChar()
    return l
}

func (l *Lexer)readChar() {
    if l.readposition >= len(l.input){
        l.ch = 0 // equals nul
    } else{
        l.ch = l.input[l.readposition]
    }
    l.position = l.readposition
    l.readposition += 1


}


func (l *Lexer) NextToken() token.Token {

    var tok token.Token

    switch l.ch {
    case '=':tok = newToken(token.ASSIGN, l.ch)
    case ';':
        tok = newToken(token.SEMICOLON, l.ch)
    case '(':
        tok = newToken(token.LPAREN, l.ch)
    case ')':
        tok = newToken(token.RPAREN, l.ch)
    case ',':
        tok = newToken(token.COMMA, l.ch)
    case '+':
        tok = newToken(token.PLUS, l.ch)
    case '{':
        tok = newToken(token.LBRACE, l.ch)
    case '}':
        tok = newToken(token.RBRACE, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    }
    l.readChar()
    return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {

    return token.Token{Type: tokenType, Literal: string(ch)}

}

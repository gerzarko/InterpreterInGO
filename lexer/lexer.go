package lexer


type Lexer struct{

    input string
    position int // current position on input
    readposition int //next char in buffer
    ch byte //current char 

}

func New (input string) *Lexer{

    l := &Lexer{input:input}
    l.ReadChar()
    return l
}

func (l *Lexer)ReadChar() {
    if l.readposition >= len(l.input){
        l.ch = 0 // equals nul
    } else{
        l.ch = l.input[l.readposition]
    }

    l.position = l.readposition
    l.readposition += 1

}

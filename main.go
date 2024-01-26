package main

import (
	"fmt"
	"monkey/lexer"
)

func main() {

	newLexer := lexer.New("hola")
    newLexer.NextToken()
    newLexer.NextToken().Type

	fmt.Printf("%v",typer)

}

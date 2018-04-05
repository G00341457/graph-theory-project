package main

import(
	"fmt"
)

type state struct{
	symbol rune
	edge1 *state
	edge2 *state
}

type nfa struct{
	initial *states
	accept *state
}

func poregtonfa(postfix string){


}

func main(){
	nfa := poregontonfa("lol")
	fmt.Println(nfa)
}
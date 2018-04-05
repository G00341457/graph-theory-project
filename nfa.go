package main

import (
	"fmt"
)

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

type nfa struct {
	initial *state
	accept  *state
}

func poregtonfa(postfix string) {

	nfaStack := []*nfa{}

	for _, r := range postfix {
		switch r {
		case '.':
			frag2 = nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[len(nfaStack)-1]
			frag1 = nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[len(nfaStack)-1]

			frag1.accept.edge1 = frag2.initial

			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':
			frag2 = nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[len(nfaStack)-1]
			frag1 = nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[len(nfaStack)-1]

			initial := state{edge1: frag1.initial, edged2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		case '*':
			frag = nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[len(nfaStack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept
			
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}		
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		}
	}

	return nfaStack[0]
}

func main() {
	nfa := poregontonfa("ab.c*|")
	fmt.Println(nfa)
}

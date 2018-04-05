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

func poregtonfa(postfix string) *nfa {

	nfaStack := []*nfa{}

	for _, r := range postfix {
		switch r {
		case '.':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1.accept.edge1 = frag2.initial

			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		case '*':
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

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

	if len(nfaStack) != 1 {
		fmt.Println("uh oh: ", len(nfaStack), nfaStack)
	}
	return nfaStack[0]
}

func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)

	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	return l
}

func poMatch(po string, s string) bool {
	ismatch := false
	ponfa := poregtonfa(po)

	current := []*state{}
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	for _, r := range s {
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		current, next = next, []*state{}

	}

	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}

	return ismatch
}
func intoPost(infix string) string {

	specials := map[rune]int{'*': 10, '|': 9, '.': 8}

	postfix, s := []rune{}, []rune{}

	for _, r := range infix {
		switch {
		case r == '(':
			s = append(s, r)
		case r == ')':
			for s[len(s)-1] != '(' {
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
			}
			s = s[:len(s)-1]
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)
		default:
			postfix = append(postfix, r)
		}
	}
	for len(s) > 0 {
		postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
	}
	return string(postfix)
}

func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
	fmt.Println("===============================")
	//from regex.go
	fmt.Println(poMatch("ab.c*|", "cccc"))
	//from shunt.go
	//ab.c*
	fmt.Println("infix:   ", "a.b.c")
	fmt.Println("Postfix:   ", intoPost("a.b.c"))
	//abd|.*
	fmt.Println("infix:   ", "a.(b|d)")
	fmt.Println("Postfix:   ", intoPost("a.(b|d)"))
	//abd|.c*
	fmt.Println("infix:   ", "a.(b|d).c")
	fmt.Println("Postfix:   ", intoPost("a.(b|d).c"))
	//abb.+.c
	fmt.Println("infix:   ", "a.(b.b)+.c")
	fmt.Println("Postfix:   ", intoPost("a.(b.b)+.c"))

	fmt.Println("===========================")
	fmt.Println(poMatch("ab.c*|", "cccc"))

}

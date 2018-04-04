package main

import (
	"fmt"
)


func intoPost(infix string)string{

specials := map[rune]int{'*' : 10, '|' : 9, '.' : 8}	

 postfix, s := []rune{}, {}rune{}

 return string(postfix)
}
func main(){

    //ab.c*
	fmt.Println("infix:   ","a.b.c")
	fmt.Println("Postfix:   ", intoPost("a.b.c"))
    //abd|.*
	fmt.Println("infix:   ","a.(b|d)")
	fmt.Println("Postfix:   ", intoPost("a.(b|d)"))
    //abd|.c*
	fmt.Println("infix:   ","a.(b|d).c")
	fmt.Println("Postfix:   ", intoPost("a.(b|d).c"))
    //abb.+.c
	fmt.Println("infix:   ","a.(b.b)+.c")
	fmt.Println("Postfix:   ", intoPost("a.(b.b)+.c"))
}
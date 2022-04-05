package main

import (
	"fmt"
	"strings"
)

func main() {
	var input []string = strings.Split("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.", "")
	tokens := lexer(input)
	run(tokens)
}

type Token struct {
	Type  string
	Value string
}

func lexer(input []string) []Token {
	var res []Token
	for _, v := range input {
		switch v {
		case ">":
			res = append(res, Token{Type: "RIGHT"})
		case "<":
			res = append(res, Token{Type: "LEFT"})
		case "+":
			res = append(res, Token{Type: "INCREMENT"})
		case "-":
			res = append(res, Token{Type: "DECREMENT"})
		case "[":
			res = append(res, Token{Type: "LOOP_START"})
		case "]":
			res = append(res, Token{Type: "LOOP_END"})
		case ".":
			res = append(res, Token{Type: "OUTPUT"})
		case ",":
			res = append(res, Token{Type: "INPUT"})
		}
	}
	return res
}

func run(tokens []Token) {
	var tape map[int]uint8 = map[int]uint8{0: 0}
	var pointer int = 0
	var paren int = 0

	for i := 0; i < len(tokens); {
		curr := tokens[i]
		switch curr.Type {
		case "RIGHT":
			pointer++
			i++
		case "LEFT":
			pointer--
			i++
		case "INCREMENT":
			tape[pointer]++
			i++
		case "DECREMENT":
			tape[pointer]--
			i++
		case "LOOP_START":
			if tape[pointer] == 0 {
				paren++
				for {
					i++
					if tokens[i].Type == "LOOP_START" {
						paren++
					} else if tokens[i].Type == "LOOP_END" {
						paren--
					}
					if paren == 0 {
						break
					}
				}
			} else {
				i++
			}
		case "LOOP_END":
			if tape[pointer] != 0 {
				paren--
				for {
					i--
					if tokens[i].Type == "LOOP_START" {
						paren++
					} else if tokens[i].Type == "LOOP_END" {
						paren--
					}
					if paren == 0 {
						break
					}
				}
			} else {
				i++
			}
		case "OUTPUT":
			fmt.Printf("%c", tape[pointer])
			i++
		case "INPUT":
			var input uint8
			fmt.Scanf("%c", &input)
			tape[pointer] = input
			i++
		}
	}
}
package main

import "fmt"

const (
	DELIM = iota
	NEWLINE
	NESTER

	USERNAME
	PASSWORD
	IDENTIFIER
)

var tokens = [...]string{
	DELIM:   "DELIM",
	NEWLINE: "NEW_LINE",
	NESTER:  "4-SPACES",
	USERNAME: "USERNAME",
	PASSWORD: "PASSWORD",
	IDENTIFIER:  "IDENTIFIER",
}

type tokenId int

type position struct {
	line int
	col  int
}

type token struct {
	kind  tokenId
	value string
	pos   position
}

func (t token) String() string {
	return fmt.Sprintf("{kind: %s, value: %s}", tokens[t.kind], t.value)
}

type tokenizer struct {
	input string
	runes []rune
	// state
	loc          int
	currPosition position
	currToken    string
}

func initTokenizer(input string) *tokenizer {
	runes := []rune(input)
	return &tokenizer{
		input:        input,
		runes:        runes,
		loc:          0,
		currPosition: position{0, 0},
		currToken:    string(runes[0]),
	}
}

func (t *tokenizer) consume() {
	t.loc++
	t.currPosition.col++
	if t.loc < len(t.runes) {
		t.currToken = string(t.runes[t.loc])
	}
}

func (t *tokenizer) peek() string {
	if t.loc < len(t.runes)-1 {
		return string(t.runes[t.loc+1])
	}
	return ""
}

func (t *tokenizer) Tokenize() []token {
	tokens := []token{}
	for t.loc < len(t.runes) {
		switch {
		case t.currToken == " ":
			numSpaces := 1
			t.consume()
			for t.currToken == " " {
				t.consume()
				numSpaces++
			}
			if numSpaces != 4 {
				fmt.Println("total spaces: ", numSpaces)
				panic("4 spaces required for nesting")
			}
			tokens = append(tokens, token{kind: NESTER, value: "    ", pos: t.currPosition})
		case t.currToken == "\n": // new line means we need to advance the currPosition's line num and reset the col num
			t.currPosition.line++
			t.currPosition.col = 0
			t.consume()
		case t.currToken == ":":
			tokens = append(tokens, token{kind: DELIM, value: ":", pos: t.currPosition})
			t.consume() // consume delimeter
			for t.currToken == " " { // consume white space after delimeter
				t.consume()
			}
		default:
			str := ""
			startPos := t.currPosition
			for t.peek() != ":" && t.peek() != "\n" && t.peek() != "" {
				str += t.currToken
				t.consume()
			}
			str += t.currToken
			t.consume()
			if str == "username" {
				tokens = append(tokens, token{kind: USERNAME, value: str, pos: startPos})
			} else if str == "password" {
				tokens = append(tokens, token{kind: PASSWORD, value: str, pos: startPos})
			} else {
				tokens = append(tokens, token{kind: IDENTIFIER, value: str, pos: startPos})
			}
		}
	}
	return tokens
}

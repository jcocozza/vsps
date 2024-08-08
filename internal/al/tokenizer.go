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
	DELIM:      "DELIM",
	NEWLINE:    "NEW_LINE",
	NESTER:     "4-SPACES",
	USERNAME:   "USERNAME",
	PASSWORD:   "PASSWORD",
	IDENTIFIER: "IDENTIFIER",
}

type tokenId int

type position struct {
	line int
	col  int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d", p.line, p.col)
}

type token struct {
	kind  tokenId
	value string
	pos   position
}

func (t token) String() string {
	return fmt.Sprintf("{kind: %s, value: %s, pos: {line: %d, col %d}}", tokens[t.kind], t.value, t.pos.line, t.pos.col)
}

type tokenizer struct {
	input string
	runes []rune
	// state
	loc          int
	currPosition position
	currChar     string
	currTokenStr string
	currTokenPos position
	tokens       []token
}

func initTokenizer(input string) *tokenizer {
	runes := []rune(input)
	return &tokenizer{
		input:        input,
		runes:        runes,
		loc:          0,
		currPosition: position{0, 0},
		currChar:     string(runes[0]),
		currTokenStr: "",
		currTokenPos: position{0, 0},
		tokens:       []token{},
	}
}

func (t *tokenizer) consume() {
	t.loc++
	t.currPosition.col++
	if t.loc < len(t.runes) {
		t.currChar = string(t.runes[t.loc])
	}
}

func (t *tokenizer) pushToken(startPos position) {
	switch t.currTokenStr {
	case "username":
		t.tokens = append(t.tokens, token{kind: USERNAME, value: t.currTokenStr, pos: startPos})
	case "password":
		t.tokens = append(t.tokens, token{kind: PASSWORD, value: t.currTokenStr, pos: startPos})
	default:
		t.tokens = append(t.tokens, token{kind: IDENTIFIER, value: t.currTokenStr, pos: startPos})
	}
	t.currTokenStr = "" // reset token string
	t.currTokenPos = t.currPosition
}

func (t *tokenizer) handleNesting() {
	startPos := t.currPosition
	numSpaces := 1
	t.consume()
	for t.currChar == " " {
		t.consume()
		numSpaces++
	}
	if numSpaces != 4 && numSpaces != 2 {
		fmt.Println("total spaces: ", numSpaces)
		panic(fmt.Sprintf("4 spaces required for nesting %v", t.currPosition))
	}
	t.tokens = append(t.tokens, token{kind: NESTER, value: "    ", pos: startPos})
	t.currTokenPos = t.currPosition
}

func (t *tokenizer) peek() string {
	if t.loc < len(t.runes)-1 {
		return string(t.runes[t.loc+1])
	}
	return ""
}

/*
func (t *tokenizer) Tokenize() []token {
	tokens := []token{}
	for t.loc < len(t.runes) {
		switch {
		case t.currChar == " ":
			startPos := t.currPosition
			numSpaces := 1
			t.consume()
			for t.currChar == " " {
				t.consume()
				numSpaces++
			}
			if numSpaces != 4 && numSpaces != 2 {
				fmt.Println("total spaces: ", numSpaces)
				panic("4 spaces required for nesting")
			}
			tokens = append(tokens, token{kind: NESTER, value: "    ", pos: startPos})
		case t.currChar == "\n": // new line means we need to advance the currPosition's line num and reset the col num
			t.consume()
			t.currPosition.col = 0
			t.currPosition.line++
		case t.currChar == ":":
			if t.peek() == " " || t.peek() == "\n" {
				tokens = append(tokens, token{kind: DELIM, value: ":", pos: t.currPosition})
				t.consume() // consume delimeter
				t.consume() // consume space after ':'
				continue
			}
		default:
			str := ""
			startPos := t.currPosition
			for t.peek() != "\n" && t.peek() != "" {
				str += t.currChar
				t.consume()
			}
			str += t.currChar
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
*/

func (t *tokenizer) Tokenize() []token {
	for t.loc < len(t.runes) {
		switch {
		case t.currChar == "\n": // new line means we need to advance the currPosition's line num and reset the col num
			t.pushToken(t.currTokenPos)
			t.consume()
			t.currPosition.col = 0
			t.currPosition.line++
			t.currTokenPos = t.currPosition
			// if we have nesting, then need to handle it
			if t.currChar == " " {
				t.handleNesting()
			}
		case t.currChar == ":":
			switch t.peek() {
			case " ":
				t.pushToken(t.currTokenPos)
				t.tokens = append(t.tokens, token{kind: DELIM, value: ":", pos: t.currPosition})
				t.consume() // consume ":"
				t.consume() // consume " " after the delimeter
				t.currTokenPos = t.currPosition
			case "\n": // after a delim + newline, we expect nesting
				t.pushToken(t.currTokenPos)
				t.tokens = append(t.tokens, token{kind: DELIM, value: ":", pos: t.currPosition})
				t.consume() // consume delimeter
				t.consume() // consume newline
				t.currPosition.col = 0
				t.currPosition.line++
				t.currTokenPos = t.currPosition
				// handle the nesting
				t.handleNesting()
			case "":
				panic("FREAK OUT")
			default: // it is not a delimeter, so just treat it like a regular char
				t.currTokenStr += t.currChar
				t.consume()
			}
		default:
			t.currTokenStr += t.currChar
			t.consume()
		}
	}
	return t.tokens
}

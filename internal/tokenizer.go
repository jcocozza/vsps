package internal

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
	// the position of the token is from its starting character
	pos position
}

func (t token) String() string {
	return fmt.Sprintf("{kind: %s, value: %s, pos: {line: %d, col %d}}", tokens[t.kind], t.value, t.pos.line, t.pos.col)
}

type tokenizer struct {
	input string
	runes []rune
	// state variables
	loc          int      // the location in the set of characters
	currPosition position // the location in the set of characters as a line num & col num
	currChar     string   // the current character
	currTokenStr string   // the current token's string value
	currTokenPos position // the current token's starting position
	tokens       []token  // the (growing) list of parsed tokens
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

// advance 1 character
//
// increases the currPosition col by 1
func (t *tokenizer) consume() {
	t.loc++
	t.currPosition.col++
	if t.loc < len(t.runes) {
		t.currChar = string(t.runes[t.loc])
	}
}

// add the current token to the list of output tokens
//
// reset the current token string to 0
// update the current token position to be the current position
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

// parse nested input
//
// essentially this deals with account information exluding the account name
func (t *tokenizer) handleNesting() error {
	startPos := t.currPosition
	numSpaces := 1
	t.consume()
	for t.currChar == " " {
		t.consume()
		numSpaces++
	}
	if numSpaces != 4 && numSpaces != 2 {
		return fmt.Errorf("2 or 4 spaces required for nesting, found %d. %v", numSpaces, t.currPosition.col)
	}
	t.tokens = append(t.tokens, token{kind: NESTER, value: "    ", pos: startPos})
	t.currTokenPos = t.currPosition
	return nil
}

// if not at the end, return the next character without advancing
func (t *tokenizer) peek() string {
	if t.loc < len(t.runes)-1 {
		return string(t.runes[t.loc+1])
	}
	return ""
}

// convert the string input to a list of tokens
func (t *tokenizer) Tokenize() ([]token, error) {
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
				err := t.handleNesting()
				if err != nil {
					return nil, err
				}
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
				err := t.handleNesting()
				if err != nil {
					return nil, err
				}
			case "":
				return nil, fmt.Errorf("unexpected file ending: %v", t.currPosition)
			default: // it is not a delimeter, so just treat it like a regular char
				t.currTokenStr += t.currChar
				t.consume()
			}
		default:
			t.currTokenStr += t.currChar
			t.consume()
		}
	}
	return t.tokens, nil
}

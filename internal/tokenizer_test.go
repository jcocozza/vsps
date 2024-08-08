package internal

import (
	"testing"
)

func Test_tokenizer_Tokenize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []token
	}{
		{"test1", "acct_name:\n    username: user\n    password: pass", []token{
			{IDENTIFIER, "acct_name", position{0, 0}},
			{DELIM, ":", position{0, 9}},
			{NESTER, "    ", position{1, 0}},
			{USERNAME, "username", position{1, 4}},
			{DELIM, ":", position{1, 12}},
			{IDENTIFIER, "user", position{1, 14}},
			{NESTER, "    ", position{2, 0}},
			{PASSWORD, "password", position{2, 4}},
			{DELIM, ":", position{2, 12}},
			{IDENTIFIER, "pass", position{2, 14}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tknizer := initTokenizer(tt.input)
			got, _ := tknizer.Tokenize()
			for i, tkn := range got {
				if tkn != tt.want[i] {
					t.Errorf("tokenizer.Tokenize() = %v, want %v", tkn, tt.want[i])
				}
			}
			/*
				fmt.Println(" got: ", got)
				fmt.Println("want: ", tt.want)
			*/
		})
	}
}

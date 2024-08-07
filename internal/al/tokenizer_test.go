package main

import (
	"reflect"
	"testing"
)

func Test_tokenizer_Tokenize(t *testing.T) {
	tests := []struct {
		name string
		input string
		want   []token
	}{
		// TODO: Add test cases.
		{"test1", "acct_name:\n    username: user\n    password: pass", []token{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tknizer := initTokenizer(tt.input)
			if got := tknizer.Tokenize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokenizer.Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}

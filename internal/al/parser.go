package main

import (
	"github.com/jcocozza/vsps/internal"
)

type parser struct {
	tokens []token
	// state
	currToken token
	loc       int
}

func initParser(tokens []token) *parser {
	return &parser{
		loc:       0,
		tokens:    tokens,
		currToken: tokens[0],
	}
}

func (p *parser) consume() {
	p.loc++
	if p.loc < len(p.tokens) {
		p.currToken = p.tokens[p.loc]
	}
}

func (p *parser) peek() token {
	return p.tokens[p.loc+1]
}

func (p *parser) consumeAccount() internal.Account {
	name := p.currToken
	p.consume() // consume name

	if p.currToken.kind != DELIM {
		panic("expected delimeter after account name")
	}
	p.consume() // consume delimeter

	if p.currToken.kind != NESTER {
		panic("expected nested account information")
	}
	p.consume() // consume nesting whitespace

	acct := internal.Account{
		Name:     name.value,
		Username: "",
		Password: "",
		Other:    make(map[string]string),
	}
	for {
		acctParamName := p.currToken
		p.consume()
		if p.currToken.kind != DELIM {
			panic("expected delimeter after param name")
		}
		p.consume() // consume delimeter
		acctParamValue := p.currToken
		p.consume() // consume value

		if acctParamName.kind == USERNAME {
			acct.Username = acctParamValue.value
		} else if acctParamName.kind == PASSWORD {
			acct.Password = acctParamValue.value
		} else {
			err := acct.AddOtherField(acctParamName.value, acctParamValue.value)
			if err != nil {
				panic(err)
			}
		}
		if p.currToken.kind != NESTER {
			return acct
		}
		p.consume() // consume the nester
	}
}

func (p *parser) Parse() internal.Accounts {
	if p.currToken.kind != IDENTIFIER {
		panic("unable to parse when identifier is not first element in file")
	}

	accounts := make(internal.Accounts)
	for p.loc < len(p.tokens) {
		acct := p.consumeAccount()
		err := accounts.Add(acct)
		if err != nil {
			panic(err)
		}
	}
	return accounts
}

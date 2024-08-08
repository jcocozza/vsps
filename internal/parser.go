package internal

import (
	"fmt"
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

func (p *parser) consumeAccount() (Account, error) {
	name := p.currToken
	p.consume() // consume name
	if p.currToken.kind != DELIM {
		return Account{}, fmt.Errorf("%v - expected delimeter after account name", p.currToken.pos)
	}
	p.consume() // consume delimeter
	if p.currToken.kind != NESTER {
		return Account{}, fmt.Errorf("%v - expected nested account information", p.currToken.pos)
	}
	p.consume() // consume nesting whitespace
	acct := Account{
		Name:     name.value,
		Username: "",
		Password: "",
		Other:    make(map[string]string),
	}
	for {
		acctParamName := p.currToken
		p.consume()
		if p.currToken.kind != DELIM {
			return Account{}, fmt.Errorf("%v - expected delimeter after parameter name", p.currToken.pos)
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
			return acct, nil
		}
		p.consume() // consume the nester
	}
}

func (p *parser) parse() (Accounts, error) {
	if p.currToken.kind != IDENTIFIER {
		return nil, fmt.Errorf("%v - unable to parse when identifier is not first element in file", p.currToken.pos)
	}
	accounts := make(Accounts)
	for p.loc < len(p.tokens) {
		acct, err := p.consumeAccount()
		if err != nil {
			return nil, err
		}
		err = accounts.Add(acct)
		if err != nil {
			panic(err)
		}
	}
	return accounts, nil
}

// parse a string input into a set of accounts
func Parser(input string) (Accounts, error) {
	tokens := initTokenizer(input).Tokenize()
	return initParser(tokens).parse()
}

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
		return Account{}, fmt.Errorf("expected delimeter after account name %v", p.currToken.pos)
	}
	p.consume() // consume delimeter
	if p.currToken.kind != NESTER {
		return Account{}, fmt.Errorf("expected nested account information %v", p.currToken.pos)
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
			return Account{}, fmt.Errorf("expected delimeter after parameter name %v", p.currToken.pos)
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
				return Account{}, fmt.Errorf("an error occurred while loading accounts. check if your account file is corrupted: %v %v", err, p.currToken.pos)
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
		return nil, fmt.Errorf("unable to parse when identifier is not first element in file %v", p.currToken.pos)
	}
	accounts := make(Accounts)
	for p.loc < len(p.tokens) {
		acct, err := p.consumeAccount()
		if err != nil {
			return nil, err
		}
		err = accounts.Add(acct)
		if err != nil {
			return nil, fmt.Errorf("an error occurred while loading accounts. check if your account file is corrupted: %v %v", err, p.currToken.pos)
		}
	}
	return accounts, nil
}

// parse a string input into a set of accounts
func Parser(input string) (Accounts, error) {
	tokens, err := initTokenizer(input).Tokenize()
	if err != nil {
		return nil, err
	}
	return initParser(tokens).parse()
}

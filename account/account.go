package account

import (
	"github.com/fatih/color"
	"time"
)

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimestamp struct {
	Account
	timestamp time.Time
}

func NewAccount(login, password, url string) *AccountWithTimestamp {
	return &AccountWithTimestamp{
		Account: Account{
			login:    login,
			password: password,
			url:      url,
		},
		timestamp: time.Now(),
	}
}

func (acc *AccountWithTimestamp) OutPutText() {
	color.Cyan(acc.login)
}

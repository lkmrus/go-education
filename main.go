package main

import (
	"demo/app/account"
)

func main() {
	acc := account.NewAccount("test", "test", "test")
	acc.OutPutText()
}

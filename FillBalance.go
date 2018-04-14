package main

import (
	"fmt"
	"strings"
	"github.com/evilpeach/stellarProject/lib"
)

func main() {

	s := lib.ReadAccounts()
		//fmt.Println("CurrentAddress: ", s)
	all := strings.Split(s, "\n") //seperate all account into slices

	var count = len(all)
	//fmt.Println(count)
	address := make([]string, count)
	for i, e := range all {
		//fmt.Println(e, "  OKK")
		acc := strings.Split(e, " ") //seperate each account's data into slices
		_ = acc[0]
		address[i] = acc[1]
		_ = acc[2]
		fmt.Println("Fill XLM to:", address[i])
	}

	//Fill XLM to balance
	for _, e := range address{
		lib.FillBalance(e, true)
	}

}
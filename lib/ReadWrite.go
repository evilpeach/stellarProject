package lib

import(
	"io/ioutil"
	"fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func WriteAccounts(addr string) {
	//write file
	b := []byte(addr)
	err := ioutil.WriteFile("Accounts.txt", b, 0644)
	check(err)
	fmt.Println("-----Write Completed-----")
	fmt.Println("NewAddress: ", addr)
}

func ReadAccounts() string{
	b, err := ioutil.ReadFile("Accounts.txt")
	check(err)
	fmt.Println("-----Read Completed-----")
	addr := string(b[:])
	return addr
}


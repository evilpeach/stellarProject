package lib

import (
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	h "github.com/stellar/go/clients/horizon"
)

func GetBalance(address string) {
	account, err := h.DefaultTestNetClient.LoadAccount(address)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(account.Balances[0].Balance) 

	//val := reflect.Indirect(reflect.ValueOf(account.Balances[0]))
	/*for i := 0; i < val.NumField(); i++ {
		fmt.Println(val.Field(i))
	}*/
    
    //fmt.Println(val.NumField())
    //fmt.Println(val.Type().Field(0))

    //Balances -> {Balance, Limit, Asset}
    //Asset -> {Type, Code, Issuer}

	for _, e := range account.Balances{
		c := "Native"
		if e.Asset.Code != "" {
			c = e.Asset.Code
		}
		fmt.Println("Amount:", e.Balance, "Code:", c)
	}

    //fmt.Println(account.Balances[0].Asset.Code) //Asset.Type, Asset.Code, Asset.Issuer

}

func FillBalance(addr string, result bool){
	resp, err := http.Get("https://friendbot.stellar.org/?addr=" + addr)

	if(err != nil){
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if(err != nil){
		log.Fatal(err)
	}
	if result{
		fmt.Println(string(body))
	}
	
}
package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/evilpeach/stellarProject/lib"
)

type account struct{
	name 	string
	address string
	seed  	string
} 

var newAccounts bool
var accounts [2]account

func initFlag(){
	flag.BoolVar(&newAccounts, "newAccount", false, "true for create new account or false to not create.")
	flag.StringVar(&accounts[0].name, "source", "Mumu", "Sender name")
	flag.StringVar(&accounts[1].name, "dest", "Momo", "Receiver name")

	flag.Parse()
}
 	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func LoadAccounts(newAcc bool) {

	if newAccounts {
		
		s := "" //string to writing 

		//loop create address
		for i, e := range accounts {
			e.address, e.seed = lib.GetNewKey()
			s += e.name + " " + e.address + " " + e.seed
			if i<len(accounts) - 1 {
				s += "\n"
			}
			
		}

		lib.WriteAccounts(s)

	}else { //Retrieve Exist Accounts
		s := lib.ReadAccounts()
		//fmt.Println("CurrentAddress: ", s)
		all := strings.Split(s, "\n") //seperate all account into slices
		for i, e := range all {
			//fmt.Println(e, "  OKK")
			acc := strings.Split(e, " ") //seperate each account's data into slices
			accounts[i].name = acc[0]
			accounts[i].address = acc[1]
			accounts[i].seed = acc[2]
		}

	}
}


func CreateAccWithBalance(amount string, sourceSeed string){
	address, _ := lib.GetNewKey()

	tx, err := b.Transaction(b.SourceAccount{sourceSeed},
							b.TestNetwork,
							b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
							b.CreateAccount(b.Destination{AddressOrSeed: address},
											b.NativeAmount{Amount: amount}),
							)

	check(err)

	txe, err := tx.Sign(sourceSeed)
	check(err)

	txeB64, err := txe.Base64()
	check(err)

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	check(err)

	log.Println("Successfully create:", address, "and deposit amount is:", amount, " Hash:", resp.Hash)

}

func PathPayment() {

}

func ChangeTrust(trusteeSeed string, issuer string){

	tx, err := b.Transaction(b.SourceAccount{trusteeSeed}, 
								b.TestNetwork, 
								b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
								b.Trust("Gled", "GBG2FDRBZDJ6IROP5XF6P6EKUT23RHFR5F77CNE4BLF5DGAZOVXDJUOX"),
								)
	check(err)

	txe, err := tx.Sign(trusteeSeed)
	check(err)

	txeB64, err := txe.Base64()
	check(err)

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	check(err)

	log.Println("Successfully Trust!! Hash:", resp.Hash)
}

func SendAsset(amount string, sourceSeed string, destAddress string, code string, issuerAddr string){
	tx, err := b.Transaction(b.SourceAccount{sourceSeed}, 
								b.TestNetwork, 
								b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
								b.Payment(b.Destination{AddressOrSeed: destAddress},
											  b.CreditAmount{code, issuerAddr, amount}),
								)

	check(err)

	txe, err := tx.Sign(sourceSeed)
	check(err)

	txeB64, err := txe.Base64()
	check(err)

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	check(err)

	log.Println("Successfully sent ", amount, "lumen to", destAddress, ". Hash:", resp.Hash)
}

func SendLumens(amount string, sourceSeed string, destAddress string){
	tx, err := b.Transaction(b.SourceAccount{sourceSeed}, 
								b.TestNetwork, 
								b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
								b.Payment(b.Destination{AddressOrSeed: destAddress},
											  b.NativeAmount{Amount: amount}),
								)

	check(err)

	txe, err := tx.Sign(sourceSeed)
	check(err)

	txeB64, err := txe.Base64()
	check(err)

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	check(err)

	log.Println("Successfully sent ", amount, "lumen to", destAddress, ". Hash:", resp.Hash)
}


func main() {
	
	issuer := "GBG2FDRBZDJ6IROP5XF6P6EKUT23RHFR5F77CNE4BLF5DGAZOVXDJUOX"

	initFlag()

	//create or retrieve accounts data
	LoadAccounts(newAccounts)

	//create account and deposit
	//CreateAccWithBalance("24", accounts[0].seed)

	//Change Trust
	//ChangeTrust(accounts[0].seed, "sdsdsd")

	//send lumen
	//SendLumens("1", accounts[1].seed, accounts[0].address) /// Amount, Sender's seed, Receiver's address 

	//send from issuer to distributor
	//SendAsset("100000", "SDNYQFFSL7XF7T3JQZJYSRMZTZVI3AFXHGFPICW5HDXDB3WAMR4O5LHG", accounts[0].address, "Gled", issuer)

	//send asset
	SendAsset("10", accounts[0].seed, accounts[1].address, "Gled", issuer)

	//print all account information
	for _,e := range accounts{
		fmt.Println(e)
		lib.GetBalance(e.address)
	}
}

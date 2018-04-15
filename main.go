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

func signAndSubmit(tx *b.TransactionBuilder, seed []string, privateMemo string){

	txe, err := tx.Sign(seed...)
	check(err)

	txeB64, err := txe.Base64()
	check(err)

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	check(err)

	log.Println("Transaction Successfully submitted: ", privateMemo)
	log.Println("Hash:", resp.Hash)
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

	memo := "Create Accounts:" + address + " with " + amount + " lumens"
	multiSigSeed := []string{sourceSeed}
	signAndSubmit(tx, multiSigSeed, memo)
}

func PathPayment() {

}

func Trust(asset b.Asset, limit string, trusteeSeed string){

	tx, err := b.Transaction(b.SourceAccount{trusteeSeed}, 
								b.TestNetwork, 
								b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
								b.Trust("Gled", "GBG2FDRBZDJ6IROP5XF6P6EKUT23RHFR5F77CNE4BLF5DGAZOVXDJUOX"),
								)
	check(err)

	//fmt.Println(tx, "  ", *tx, "  ", &tx)
	memo := "Change Trust"
	multiSigSeed := []string{trusteeSeed}
	signAndSubmit(tx, multiSigSeed, memo)
}

func CreateBuyingOffer(rate b.Rate, amount b.Amount, sourceSeed string){
	tx, err := b.Transaction(b.SourceAccount{sourceSeed}, 
								b.TestNetwork, 
								b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
								b.CreateOffer(rate, amount),
								)

	check(err)
	memo := "Create Buying Offer Successfully: " + rate.Buying.Code 
	multiSigSeed := []string{sourceSeed}
	signAndSubmit(tx, multiSigSeed, memo)
}

func CreateSellingOffer(rate b.Rate, amount b.Amount, sourceSeed string){
	tx, err := b.Transaction(b.SourceAccount{sourceSeed}, 
								b.TestNetwork, 
								b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
								b.CreateOffer(rate, amount),
								)

	check(err)
	memo := "Create Selling Offer Successfully: " + rate.Selling.Code 
	multiSigSeed := []string{sourceSeed}
	signAndSubmit(tx, multiSigSeed, memo)
}

func DeleteOffer(rate b.Rate, offerID b.OfferID, sourceSeed string){
	tx, err := b.Transaction(b.SourceAccount{sourceSeed}, 
								b.TestNetwork, 
								b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
								b.DeleteOffer(rate, offerID),
								)

	check(err)
	memo := "Delete Offer Successfully"
	multiSigSeed := []string{sourceSeed}
	signAndSubmit(tx, multiSigSeed, memo)
}

func SendAsset(amount string, sourceSeed string, destAddress string, code string, issuerAddr string){
	tx, err := b.Transaction(b.SourceAccount{sourceSeed}, 
								b.TestNetwork, 
								b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
								b.Payment(b.Destination{AddressOrSeed: destAddress},
											  b.CreditAmount{code, issuerAddr, amount}),
								)

	check(err)

	memo := "SendAsset:" + code + " to " + destAddress
	multiSigSeed := []string{sourceSeed}
	signAndSubmit(tx, multiSigSeed, memo)
}

func SendLumens(amount string, sourceSeed string, destAddress string){
	tx, err := b.Transaction(b.SourceAccount{sourceSeed}, 
								b.TestNetwork, 
								b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
								b.Payment(b.Destination{AddressOrSeed: destAddress},
											  b.NativeAmount{Amount: amount}),
								)

	check(err)

	memo := "SendLumens " + amount + " to " + destAddress
	multiSigSeed := []string{sourceSeed, "SDQK24QMRQMYKN4BN6U5XHJLNTSDSF4PAJLMCEHYXSEHQCI5JRCMHS5O"}
	signAndSubmit(tx, multiSigSeed, memo)
}

func AddNewSigner(sourceSeed string, signerAddress string){
	tx, err := b.Transaction(b.SourceAccount{sourceSeed}, 
								b.TestNetwork, 
								b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
								b.SetOptions(b.MasterWeight(1),
											b.SetThresholds(2,2,2),
											b.AddSigner(signerAddress, 1)),
								)

	check(err)

	memo := "Successfully Add New Signer!"
	multiSigSeed := []string{sourceSeed}
	signAndSubmit(tx, multiSigSeed, memo)
}

func main() {
	
	//issuer := "GBG2FDRBZDJ6IROP5XF6P6EKUT23RHFR5F77CNE4BLF5DGAZOVXDJUOX"

	initFlag()

	//create or retrieve accounts data
	LoadAccounts(newAccounts)

	//create account and deposit
	//CreateAccWithBalance("2", accounts[0].seed)

	//Trust
	//GledCreditAsset := b.CreditAsset("Gled", issuer)
	//Trust(GledCreditAsset, "500", accounts[1].seed)

	//Create Offer
	/////Buy Gled
	// buyingRate := b.Rate{Selling: b.NativeAsset(), Buying: GledCreditAsset, Price: "10"}
	// CreateBuyingOffer(buyingRate, "10", accounts[1].seed) 	

	// ////Sell Gled
	// sellingRate := b.Rate{Selling: GledCreditAsset, Buying: b.NativeAsset(), Price: "0.5"}
	// CreateSellingOffer(sellingRate, "30", accounts[0].seed)

	//Delete Offer
	//deleteRate := b.Rate{Selling: GledCreditAsset, Buying: b.NativeAsset(), Price: "1"}
	// deleteRate := b.Rate{Selling: b.NativeAsset(), Buying: GledCreditAsset, Price: "10"}
	// DeleteOffer(deleteRate, 341610, accounts[1].seed)

	//Add new signer
	//AddNewSigner(accounts[1].seed, accounts[0].address)

	//send lumen
	SendLumens("20", accounts[1].seed, accounts[0].address) /// Amount, Sender's seed, Receiver's address 

	//send from issuer to distributor
	//SendAsset("1000", "SDNYQFFSL7XF7T3JQZJYSRMZTZVI3AFXHGFPICW5HDXDB3WAMR4O5LHG", accounts[0].address, "Gled", issuer)

	//send asset
	//SendAsset("1", accounts[0].seed, accounts[1].address, "Gled", issuer)

	//print all account information
	for _,e := range accounts{
		fmt.Println(e)
		lib.GetBalance(e.address)
	}
}

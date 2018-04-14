package lib

import (
	"log"
	"github.com/stellar/go/keypair"
	//"reflect"
)

func GetNewKey () (string, string){
	pair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}

	return pair.Address(), pair.Seed()
	
}

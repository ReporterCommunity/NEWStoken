package main

import (
	"fmt"
	"log"

	"github.com/DaveAppleton/ether_go/ethKeys"
)

func roc(keyS string) (key *ethKeys.AccountKey, err error) {
	key = ethKeys.NewKey(keyS)
	if err := key.RestoreOrCreate(); err != nil {
		log.Println(keyS, err)
	}
	return
}

func createSubscribers() bool {
	for i := 0; i < 1000; i++ {
		roc(fmt.Sprintf("userKeys/S%05d", i))
		if i < 50 {
			roc(fmt.Sprintf("userKeys/P%05d", i))
		}
	}
	return true
}

func areContractsSetup() bool {
	fmt.Println("setup")

	return true
}

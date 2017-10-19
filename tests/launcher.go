package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
)

//------- The Canary

var canary *Canary

//------- stuff

var contracts = make(map[string]contract)
var ownerTx *bind.TransactOpts
var reporterToken *ReporterToken
var reporterTokenSale *ReporterTokenSale

var saleAddress common.Address
var coinAddress common.Address

type contract struct {
	Name    string
	TxHash  string
	Address string
}

func (c *contract) SetTxHash(a common.Hash) {
	c.TxHash = fmt.Sprintf("0x%x", a)
}
func (c *contract) SetAddress(a common.Address) {
	c.Address = fmt.Sprintf("0x%x", a)
}
func (c *contract) GetTxHash() common.Hash {
	return common.StringToHash(c.TxHash)
}
func (c *contract) GetAddress() common.Address {
	return common.HexToAddress(c.Address)
}

var ownerAddress common.Address

var strA = []string{"ReporterSale"}

func areContractsLaunched() bool {
	var tx *types.Transaction
	client, err := getClient()
	if err != nil {
		fmt.Println(err)
		return false
	}
	suffix := viper.GetString("suffix")
	ownerTx = role("banker")
	if err != nil {
		log.Printf("Failed to create authorized transactor: %v", err)
		return false
	}
	_, _, canary, err = DeployCanary(ownerTx, client)
	if err != nil {
		log.Fatalf("Failed to deploy: %v", err)
	}

	count := 0
	for _, base := range strA {
		jFile := "contractinfo/" + base + suffix + ".json"

		count++
		var NMContract contract
		NMContract.Name = base
		var address common.Address

		switch base {
		case "ReporterSale":
			address, tx, reporterTokenSale, err = DeployReporterTokenSale(ownerTx, client)

			if err != nil {
				log.Fatalf("Failed to deploy: %v", err)
			}
			saleAddress = address

		}
		waitForMined(tx)
		NMContract.SetTxHash(tx.Hash())
		NMContract.SetAddress(address)
		log.Println("Deployed ", base, " at ", address.Hex())
		log.Println("tx = ", tx.Hash().Hex())
		contracts[base] = NMContract
		serial, _ := json.Marshal(NMContract)
		err = ioutil.WriteFile(jFile, serial, 0644)
		if err != nil {
			log.Fatalf("Error saving contract: %s %v", base, err)
		}
	}
	coinAddress, err = reporterTokenSale.Token(nil)
	if err != nil {
		log.Fatal("Cannot get coin address", err)
	}
	reporterToken, err = NewReporterToken(coinAddress, client)
	if err != nil {
		log.Fatal("Cannot produce coin contract", err)
	}

	fmt.Printf("CrowdSale deployed at : 0x%x\n", saleAddress)
	log.Printf("CrowdSale deployed at : 0x%x\n", saleAddress)
	fmt.Printf("Token deployed at : 0x%x\n", coinAddress)
	log.Printf("Token deployed at : 0x%x\n", coinAddress)

	ltxe(reporterTokenSale.SetWallet(ownerTx, roleAddress("multisig")))
	return true
}

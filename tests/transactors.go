package main

import (
	"log"
	"math/big"

	"github.com/DaveAppleton/ether_go/ethKeys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
)

// func getClient() (client *parityclient.Client, err error) {
// 	endPoint := viper.GetString("IPC_PATH")
// 	if len(endPoint) == 0 {
// 		endPoint = "/Users/daveappleton/Library/Ethereum/geth.ipc"
// 	}
// 	//deadline := time.Now().Add(20 * time.Second)
// 	//ctx, cancel := context.WithDeadline(context.Background(), deadline)
// 	client, err = parityclient.Dial(endPoint)
// 	return
// }

var baseClient *backends.SimulatedBackend

func getClient() (client *backends.SimulatedBackend, err error) {
	if baseClient != nil {
		return baseClient, nil
	}
	funds, _ := new(big.Int).SetString("1000000000000000000000000000", 10)
	bankerKey := roleKey("banker")
	baseClient = backends.NewSimulatedBackend(core.GenesisAlloc{
		bankerKey.PublicKey(): {Balance: funds},
	})
	return baseClient, nil
}

func keyTx(key *ethKeys.AccountKey) *bind.TransactOpts {
	return bind.NewKeyedTransactor(key.GetKey())
}

func role(job string) *bind.TransactOpts {
	return keyTx(roleKey(job))
}

func roleKey(job string) *ethKeys.AccountKey {
	transactorAcc := ethKeys.NewKey("adminKeys/" + job)
	if err := transactorAcc.RestoreOrCreate(); err != nil {
		log.Println(err)
	}
	return transactorAcc
}

func roleAddress(job string) common.Address {
	return roleKey(job).PublicKey()
}

func customerServiceTx() *bind.TransactOpts {
	return role("cs")
}

func customerServiceAddress() common.Address {
	return roleAddress("cs")
}

// ---- user
func userTx(job string) *bind.TransactOpts {
	return keyTx(userKey(job))
}

func userKey(job string) *ethKeys.AccountKey {
	transactorAcc := ethKeys.NewKey("userKeys/" + job)
	if err := transactorAcc.RestoreOrCreate(); err != nil {
		log.Println(err)
	}
	return transactorAcc
}

func userAddress(job string) common.Address {
	return userKey(job).PublicKey()
}

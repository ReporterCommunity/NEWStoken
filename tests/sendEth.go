package main

import (
	"fmt"
	"math/big"

	"github.com/DaveAppleton/ether_go/ethKeys"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/net/context"
)

func sendEthereum(sender *ethKeys.AccountKey, recipient common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	//var ret interface{}
	//var zero interface{}

	ec, _ := getClient()

	nonce, err := ec.PendingNonceAt(context.TODO(), sender.PublicKey())
	gasPrice, err := ec.SuggestGasPrice(context.TODO())
	if err != nil {
		return nil, err
	}
	fmt.Println("Nonce : ", nonce)
	fmt.Println("GasPrice : ", gasPrice)
	//chainID := new(big.Int).SetInt64(viper.GetInt64("ChainId"))
	//s := types.NewEIP155Signer(chainID)
	s := types.HomesteadSigner{}
	if gasLimit.Cmp(new(big.Int)) == 0 {
		gasLimit.SetInt64(121000) // because it is a send - quite standard
	}
	data := common.FromHex("0x")
	t := types.NewTransaction(nonce, recipient, amount, gasLimit, gasPrice, data)
	nt, err := types.SignTx(t, s, sender.GetKey())
	if err != nil {
		return nil, err
	}
	err = ec.SendTransaction(context.TODO(), nt)
	return nt, err

	// rlpEncodedTx, err := rlp.EncodeToBytes(nt)
	// if err != nil {
	// 	return zero, err
	// }
	// err = myEipc.Call(&ret, "eth_sendRawTransaction", common.ToHex(rlpEncodedTx))
	// return ret, err
}

package main

import (
	"context"
	"fmt"
	todo "go-solidity/gen"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	b, err := ioutil.ReadFile("wallet/UTC--2022-04-09T17-02-29.599554300Z--be42d3b5f88536bff88e73dc03960a926d06499e")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial("https://kovan.infura.io/v3/project_id")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	cAdd := common.HexToAddress("0x7A0c50aa3549159fe78Adf18B3B84d009CfabAF2")
	t, err := todo.NewTodo(cAdd, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	tx.GasLimit = 3000000
	tx.GasPrice = gasPrice

	/* Add New Task */
	// tra, err := t.Add(tx, "First Task")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(tra.Hash())

	/* List All Tasks */
	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	tasks, err := t.List(&bind.CallOpts{
		From: add,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tasks)

	/* Update Task */
	// tra, err := t.Update(tx, big.NewInt(0), "update task content")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Update tx", tra.Hash())

	/* Toggle Task */
	// tra, err := t.Toggle(tx, big.NewInt(0))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Toggle tx", tra.Hash())

	/* Remove Task */
	// tra, err := t.Remove(tx, big.NewInt(0))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Remove tx", tra.Hash())
}

// package is main for now but will be core after developement
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jolatechno/libp2p-mpi-core/v3/contract"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough args")
	}

	fmt.Println("\nsender\n") //--------------------------------------

	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice /*= big.NewInt(0) //*/, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721974) // in units
	auth.GasPrice = gasPrice

	transactOpts := bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}

	var ipfs_hash string = "test"
	var kernel_shape []*big.Int = []*big.Int{big.NewInt(1)}
	var keys, values [][]byte = [][]byte{[]byte("test")}, [][]byte{[]byte("test first")}
	var proportion, number *big.Int = big.NewInt(0), big.NewInt(1)

	address, tx, instance, err := contract.DeployTask(auth, client, ipfs_hash, kernel_shape, keys, values, proportion, number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance

	fmt.Println("\nreceptor\n") //--------------------------------------

	task, err := contract.NewTask(address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	value, check, err := task.Read(&bind.CallOpts{}, []byte("test"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("value: %q for key %q, read: %t\n", value, "test", check)

	for i := 0; i < 5; i++ {
		bytes := []byte(fmt.Sprintf("test %d", i))

		_, err = task.Push(&transactOpts, bytes, bytes)
		if err != nil {
			log.Fatal(err)
		}

		value, check, err = task.Read(&bind.CallOpts{}, bytes)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("value: %q for key %q, read: %t\n", value, bytes, check)
	}
}

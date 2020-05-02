// package is main for now but will be core after developement
package main

import (
	"os"
	"context"
	"crypto/ecdsa"
	"log"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jolatechno/libp2p-mpi-core/v3/contract"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough args")
	}

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

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(6721974) // in units
	auth.GasPrice = gasPrice

	fmt.Println(auth)

	var ipfs_hash string = "test"
	var kernel_shape []*big.Int = []*big.Int{big.NewInt(1)}
	var keys, values [][]byte = [][]byte{[]byte{}}, [][]byte{[]byte{}}
	var proportion, number *big.Int = big.NewInt(1), big.NewInt(1)

	address, tx, instance, err := contract.DeployTask(auth, client, ipfs_hash, kernel_shape, keys, values, proportion, number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}
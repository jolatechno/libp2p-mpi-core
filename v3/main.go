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

	fmt.Println("\ninterpreter\n") //--------------------------------------

	var ipfs_hash string = "test"

	interp_address, tx0, interp, err := contract.DeployInterpreter(auth, client, ipfs_hash)

	fmt.Println(interp_address.Hex())
	fmt.Println(tx0.Hash().Hex())

	fmt.Println("\nsender\n") //--------------------------------------

	var kernel_shape []*big.Int = []*big.Int{big.NewInt(1)}
	var keys, values [][]byte = [][]byte{[]byte("test")}, [][]byte{[]byte("test first")}
	var proportion, number *big.Int = big.NewInt(2), big.NewInt(1)

	nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
	auth.Nonce = big.NewInt(int64(nonce))
	transactOpts = bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}

	address, tx, task, err := contract.DeployTask(auth, client, interp_address, kernel_shape, keys, values, proportion, number)
	if err != nil {
		log.Fatal(err)
	}

	_, err = task.Advertise(&transactOpts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	fmt.Println("\ntest\n") //--------------------------------------

	address, ok, err = interp.GetTask(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())

	fmt.Println("\nreceptor\n") //--------------------------------------

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

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
	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough args")
	}

	Shell := shell.NewShell("/ip4/127.0.0.1/tcp/5001")
	inter := &MessageShell{ Shell:Shell }

	client, err := ethclient.Dial("ws://localhost:7545")
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

	fmt.Println("\ninterpreter\n") //--------------------------------------

	_, interp_address, err := NewInterp(auth, client, Shell, "./contract")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(interp_address.Hex())

	interp, err := LoadInterp(client, interp_address)
	if err != nil {
		log.Fatal(err)
	}

	var proportion, number *big.Int = big.NewInt(2), big.NewInt(1)

	stack, err := NewStack(auth, client, proportion, number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\ntask\n") //--------------------------------------

	var kernel_shape []*big.Int = []*big.Int{big.NewInt(1)}
	var keys, values [][]byte = [][]byte{}, [][]byte{}

	for j := 0; j < 3; j++ {
		keys = append(keys, []byte(fmt.Sprintf("key 0 %d", j)))
		values = append(values, []byte(fmt.Sprintf("value 0 %d", j)))
	}

	err = NewTask(auth, client, interp_address, stack,
		kernel_shape, keys, values, inter,
	)
	if err != nil {
		log.Fatal(err)
	}

	address, ok, err := interp.GetTask(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())

	task, err := LoadTask(auth, client, address)
	if err != nil {
		log.Fatal(1, err)
	}

	task.SetIntermediary(inter)

	fmt.Println("\nreceptor\n") //--------------------------------------

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			k := []byte(fmt.Sprintf("key %d %d", i + 1, j))
			v := []byte(fmt.Sprintf("value %d %d", i + 1, j))

			err = task.Push(k, v)
			if err != nil {
				log.Fatal(err)
			}

			read_k := []byte(fmt.Sprintf("key %d %d", i, j))

			value, err := task.Read(read_k)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("value: %q for key %q\n", value, read_k)
		}

		err = task.Done()
		if err != nil {
			log.Fatal(err)
		}

		err = NewTask(auth, client, interp_address, stack,
			kernel_shape, [][]byte{}, [][]byte{}, inter,
		)
		if err != nil {
			log.Fatal(err)
		}

		address, ok, err := interp.GetTask(&bind.CallOpts{})
		if !ok {
			log.Fatal("not interp")
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\n", address.Hex(), "\n")

		task, err := LoadTask(auth, client, address)
		if err != nil {
			log.Fatal(1, err)
		}

		task.SetIntermediary(inter)
	}
}

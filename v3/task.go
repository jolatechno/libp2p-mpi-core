package main

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/jolatechno/libp2p-mpi-core/v3/contract"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func update(auth *bind.TransactOpts, client *ethclient.Client) error {
	/*nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return err
	}*/

	auth.Nonce = big.NewInt(auth.Nonce.Int64() + int64(1)); // = big.NewInt(int64(nonce))
	return nil
}

type pushReq struct {
	Sender common.Address
	Key []byte
	Value []byte
}

type Intermediary interface {
	Write(data []byte) (processed []byte, err error)
	Read(processed []byte) (data []byte, err error)
}

func NewStack(auth *bind.TransactOpts, client *ethclient.Client,
	proportion *big.Int, number *big.Int) (_ *contract.Stack, stackAddr common.Address, _ error) {
	defer recover()

	err := update(auth, client)
	if err != nil {
		return nil, stackAddr, err
	}

	address, _, stack, err := contract.DeployStack(auth, client, proportion, number)
	return stack, address, err
}

func LoadTask(auth *bind.TransactOpts, client *ethclient.Client, taskAddress common.Address) (*Task, error) {
	task, err := contract.NewTask(taskAddress, client)
	if err != nil {
		return nil, err
	}

	return &Task{
		Task: task,
		Auth: auth,
		Client: client,
	}, nil
}

func NewTask(auth *bind.TransactOpts, client *ethclient.Client, ipfsObjectAddress common.Address,
	stack *contract.Stack, stackAddress common.Address,
	kernel_shape []*big.Int, keys [][]byte, values [][]byte,
	inter Intermediary) (_ *Task, err error) {
	defer recover()
	
	for i, val := range values {
		if inter != nil {
			val, err = inter.Write(val)
			if err != nil {
				return nil, err
			}
		}

		err := update(auth, client)
		if err != nil {
			return nil, err
		}

		_, err = stack.Push(auth, auth.From, keys[i], val)
		if err != nil {
			return nil, err
		}
	}

	err = update(auth, client)
	if err != nil {
		return nil, err
	}

	address, _, task, err := contract.DeployTask(auth, client, stackAddress, ipfsObjectAddress, kernel_shape)
	if err != nil {
		return nil, err
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(contract.TaskABI)))
	if err != nil {
		return nil, err
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				panic(err)
			case vLog := <-logs:
				var pushEvent pushReq

				err = contractAbi.Unpack(&pushEvent, "Push", vLog.Data)
				if err != nil {
					panic(err)
				}

				err = update(auth, client)
				if err != nil {
					panic(err)
				}

				_, err := stack.Push(auth, pushEvent.Sender, pushEvent.Key, pushEvent.Value)

				if err != nil {
					panic(err)
				}
			}
		}
	}()
	
	return &Task{
		Task: task,
		Auth: auth,
		Client: client,
		Inter: inter,
	}, nil
}

type Task struct {
	Task *contract.Task
	Inter Intermediary
	Auth *bind.TransactOpts
	Client *ethclient.Client
}

func (t *Task)SetIntermediary(inter Intermediary) {
	t.Inter = inter
}

func (t *Task)Read(address []byte) ([]byte, error) {
	value, check, err := t.Task.Read(&bind.CallOpts{}, address)
	if err != nil {
		return value, err
	}

	if !check {
		return value, errors.New("nil pointer")
	}

	if t.Inter != nil {
		return t.Inter.Read(value)
	}

	return value, nil
}

func (t *Task)Push(address []byte, value []byte) (err error) {
	if t.Inter != nil {
		value, err = t.Inter.Write(value)
		if err != nil {
			return err
		}
	}

	err = update(t.Auth, t.Client)
	if err != nil {
		return err
	}

	_, err = t.Task.Push(t.Auth, address, value)
	return err
}

func (t *Task)Done() (err error) {
	err = update(t.Auth, t.Client)
	if err != nil {
		return err
	}

	_, err = t.Task.Done(t.Auth)
	return err
}
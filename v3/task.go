package main

import (
	"errors"
	"math/big"

	"github.com/jolatechno/libp2p-mpi-core/v3/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	proportion *big.Int, number *big.Int) (_ *contract.Stack, _ error) {
	defer recover()

	err := update(auth, client)
	if err != nil {
		return nil, err
	}

	_, _, stack, err := contract.DeployStack(auth, client, proportion, number)
	return stack, err
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
	stack *contract.Stack,
	kernel_shape []*big.Int, keys [][]byte, values [][]byte,
	inter Intermediary) (err error) {
	defer recover()
	
	for i, val := range values {
		if inter != nil {
			val, err = inter.Write(val)
			if err != nil {
				return err
			}
		}

		err := update(auth, client)
		if err != nil {
			return err
		}

		_, err = stack.Push(auth, auth.From, keys[i], val)
		if err != nil {
			return err
		}
	}

	err = update(auth, client)
	if err != nil {
		return err
	}

	_, err = stack.NewTask(auth, ipfsObjectAddress, kernel_shape)
	return err
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
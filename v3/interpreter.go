package main

import (
	"github.com/jolatechno/libp2p-mpi-core/v3/contract"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func NewInterp(auth *bind.TransactOpts, client *ethclient.Client, Shell *shell.Shell, path string) (_ *contract.Interpreter, interp_address common.Address, err error) {
	ipfs_hash, err := Shell.AddDir(path)
	if err != nil {
		return nil, interp_address, err
	}

	interp_address, _, interp, err := contract.DeployInterpreter(auth, client, ipfs_hash)
	return interp, interp_address, err
}

func LoadInterp(client *ethclient.Client, interp_address common.Address) (*contract.Interpreter, error) {
	return contract.NewInterpreter(interp_address, client)
}
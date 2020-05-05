// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IdentityABI is the input ABI used to generate the binding from.
const IdentityABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_owner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IdentityFuncSigs maps the 4-byte function signature to its string representation.
var IdentityFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"1a695230": "transfer(address)",
}

// IdentityBin is the compiled bytecode used for deploying new contracts.
var IdentityBin = "0x6080604052600080546001600160a01b0319163317905534801561002257600080fd5b50610144806100326000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631a6952301461003b5780638da5cb5b14610063575b600080fd5b6100616004803603602081101561005157600080fd5b50356001600160a01b0316610087565b005b61006b610100565b604080516001600160a01b039092168252519081900360200190f35b6000546001600160a01b031633146100de576040805162461bcd60e51b81526020600482015260156024820152741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03169056fea265627a7a723158202800c969a06feae6706f8c72729d96c451c81ed5b8a3537f74bf115fbb2af7b864736f6c63430005100032"

// DeployIdentity deploys a new Ethereum contract, binding an instance of Identity to it.
func DeployIdentity(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Identity, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// Identity is an auto generated Go binding around an Ethereum contract.
type Identity struct {
	IdentityCaller     // Read-only binding to the contract
	IdentityTransactor // Write-only binding to the contract
	IdentityFilterer   // Log filterer for contract events
}

// IdentityCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentitySession struct {
	Contract     *Identity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityCallerSession struct {
	Contract *IdentityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IdentityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityTransactorSession struct {
	Contract     *IdentityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IdentityRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRaw struct {
	Contract *Identity // Generic contract binding to access the raw methods on
}

// IdentityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityCallerRaw struct {
	Contract *IdentityCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityTransactorRaw struct {
	Contract *IdentityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentity creates a new instance of Identity, bound to a specific deployed contract.
func NewIdentity(address common.Address, backend bind.ContractBackend) (*Identity, error) {
	contract, err := bindIdentity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// NewIdentityCaller creates a new read-only instance of Identity, bound to a specific deployed contract.
func NewIdentityCaller(address common.Address, caller bind.ContractCaller) (*IdentityCaller, error) {
	contract, err := bindIdentity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityCaller{contract: contract}, nil
}

// NewIdentityTransactor creates a new write-only instance of Identity, bound to a specific deployed contract.
func NewIdentityTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityTransactor, error) {
	contract, err := bindIdentity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityTransactor{contract: contract}, nil
}

// NewIdentityFilterer creates a new log filterer instance of Identity, bound to a specific deployed contract.
func NewIdentityFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityFilterer, error) {
	contract, err := bindIdentity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityFilterer{contract: contract}, nil
}

// bindIdentity binds a generic wrapper to an already deployed contract.
func bindIdentity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.IdentityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Identity *IdentityCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Identity *IdentitySession) Owner() (common.Address, error) {
	return _Identity.Contract.Owner(&_Identity.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Identity *IdentityCallerSession) Owner() (common.Address, error) {
	return _Identity.Contract.Owner(&_Identity.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Identity *IdentityTransactor) Transfer(opts *bind.TransactOpts, new_owner common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "transfer", new_owner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Identity *IdentitySession) Transfer(new_owner common.Address) (*types.Transaction, error) {
	return _Identity.Contract.Transfer(&_Identity.TransactOpts, new_owner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Identity *IdentityTransactorSession) Transfer(new_owner common.Address) (*types.Transaction, error) {
	return _Identity.Contract.Transfer(&_Identity.TransactOpts, new_owner)
}

// InterpreterABI is the input ABI used to generate the binding from.
const InterpreterABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"IpfsHash\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contracttask\",\"name\":\"Task\",\"type\":\"address\"}],\"name\":\"advertise\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contracttask\",\"name\":\"Task\",\"type\":\"address\"}],\"name\":\"done\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTask\",\"outputs\":[{\"internalType\":\"contracttask\",\"name\":\"Task\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_owner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// InterpreterFuncSigs maps the 4-byte function signature to its string representation.
var InterpreterFuncSigs = map[string]string{
	"4063d563": "advertise(address)",
	"7be8f86b": "done(address)",
	"a31662d2": "getTask()",
	"8da5cb5b": "owner()",
	"1a695230": "transfer(address)",
}

// InterpreterBin is the compiled bytecode used for deploying new contracts.
var InterpreterBin = "0x6080604052600080546001600160a01b0319163317905534801561002257600080fd5b506040516105ea3803806105ea8339818101604052602081101561004557600080fd5b810190808051604051939291908464010000000082111561006557600080fd5b90830190602082018581111561007a57600080fd5b825164010000000081118282018810171561009457600080fd5b82525081516020918201929091019080838360005b838110156100c15781810151838201526020016100a9565b50505050905090810190601f1680156100ee5780820380516001836020036101000a031916815260200191505b50604052505081516101089150600190602084019061010f565b50506101aa565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061015057805160ff191683800117855561017d565b8280016001018555821561017d579182015b8281111561017d578251825591602001919060010190610162565b5061018992915061018d565b5090565b6101a791905b808211156101895760008155600101610193565b90565b610431806101b96000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631a6952301461005c5780634063d563146100845780637be8f86b146100aa5780638da5cb5b146100d0578063a31662d2146100f4575b600080fd5b6100826004803603602081101561007257600080fd5b50356001600160a01b03166100fc565b005b6100826004803603602081101561009a57600080fd5b50356001600160a01b0316610175565b610082600480360360208110156100c057600080fd5b50356001600160a01b03166101f8565b6100d86102e8565b604080516001600160a01b039092168252519081900360200190f35b6100d86102f7565b6000546001600160a01b03163314610153576040805162461bcd60e51b81526020600482015260156024820152741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6003805460018181019092557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b039093166001600160a01b03199384161790556002805491820181556000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01805490911633179055565b80336001600160a01b0382161461024e576040805162461bcd60e51b81526020600482015260156024820152741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b60005b6003548110156102e357826001600160a01b03166003828154811061027257fe5b6000918252602090912001546001600160a01b031614156102db576003818154811061029a57fe5b600091825260209091200180546001600160a01b031916905560028054829081106102c157fe5b600091825260209091200180546001600160a01b03191690555b600101610251565b505050565b6000546001600160a01b031690565b60035460009061034e576040805162461bcd60e51b815260206004820152601c60248201527f6e6f207461736b2063757272656e74726c7920617661696c61626c6500000000604482015290519081900360640190fd5b60035460009061035d90610387565b90506003818154811061036c57fe5b6000918252602090912001546001600160a01b031691505090565b604080513360601b602080830191909152825180830360140181526034909201909252805191012060009081904390816103bd57fe5b044301604051602001808281526020019150506040516020818303038152906040528051906020012060001c90508281816103f457fe5b06939250505056fea265627a7a7231582040f7639cbb6ae34bc020fa1feafb9ecd88e6ecd238d7294294d8bd2812135c3d64736f6c63430005100032"

// DeployInterpreter deploys a new Ethereum contract, binding an instance of Interpreter to it.
func DeployInterpreter(auth *bind.TransactOpts, backend bind.ContractBackend, IpfsHash string) (common.Address, *types.Transaction, *Interpreter, error) {
	parsed, err := abi.JSON(strings.NewReader(InterpreterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InterpreterBin), backend, IpfsHash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Interpreter{InterpreterCaller: InterpreterCaller{contract: contract}, InterpreterTransactor: InterpreterTransactor{contract: contract}, InterpreterFilterer: InterpreterFilterer{contract: contract}}, nil
}

// Interpreter is an auto generated Go binding around an Ethereum contract.
type Interpreter struct {
	InterpreterCaller     // Read-only binding to the contract
	InterpreterTransactor // Write-only binding to the contract
	InterpreterFilterer   // Log filterer for contract events
}

// InterpreterCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterpreterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterpreterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterpreterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterpreterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterpreterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterpreterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterpreterSession struct {
	Contract     *Interpreter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InterpreterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterpreterCallerSession struct {
	Contract *InterpreterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// InterpreterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterpreterTransactorSession struct {
	Contract     *InterpreterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// InterpreterRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterpreterRaw struct {
	Contract *Interpreter // Generic contract binding to access the raw methods on
}

// InterpreterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterpreterCallerRaw struct {
	Contract *InterpreterCaller // Generic read-only contract binding to access the raw methods on
}

// InterpreterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterpreterTransactorRaw struct {
	Contract *InterpreterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterpreter creates a new instance of Interpreter, bound to a specific deployed contract.
func NewInterpreter(address common.Address, backend bind.ContractBackend) (*Interpreter, error) {
	contract, err := bindInterpreter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Interpreter{InterpreterCaller: InterpreterCaller{contract: contract}, InterpreterTransactor: InterpreterTransactor{contract: contract}, InterpreterFilterer: InterpreterFilterer{contract: contract}}, nil
}

// NewInterpreterCaller creates a new read-only instance of Interpreter, bound to a specific deployed contract.
func NewInterpreterCaller(address common.Address, caller bind.ContractCaller) (*InterpreterCaller, error) {
	contract, err := bindInterpreter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterpreterCaller{contract: contract}, nil
}

// NewInterpreterTransactor creates a new write-only instance of Interpreter, bound to a specific deployed contract.
func NewInterpreterTransactor(address common.Address, transactor bind.ContractTransactor) (*InterpreterTransactor, error) {
	contract, err := bindInterpreter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterpreterTransactor{contract: contract}, nil
}

// NewInterpreterFilterer creates a new log filterer instance of Interpreter, bound to a specific deployed contract.
func NewInterpreterFilterer(address common.Address, filterer bind.ContractFilterer) (*InterpreterFilterer, error) {
	contract, err := bindInterpreter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterpreterFilterer{contract: contract}, nil
}

// bindInterpreter binds a generic wrapper to an already deployed contract.
func bindInterpreter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterpreterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Interpreter *InterpreterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Interpreter.Contract.InterpreterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Interpreter *InterpreterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interpreter.Contract.InterpreterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Interpreter *InterpreterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Interpreter.Contract.InterpreterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Interpreter *InterpreterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Interpreter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Interpreter *InterpreterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interpreter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Interpreter *InterpreterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Interpreter.Contract.contract.Transact(opts, method, params...)
}

// GetTask is a free data retrieval call binding the contract method 0xa31662d2.
//
// Solidity: function getTask() view returns(address Task)
func (_Interpreter *InterpreterCaller) GetTask(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Interpreter.contract.Call(opts, out, "getTask")
	return *ret0, err
}

// GetTask is a free data retrieval call binding the contract method 0xa31662d2.
//
// Solidity: function getTask() view returns(address Task)
func (_Interpreter *InterpreterSession) GetTask() (common.Address, error) {
	return _Interpreter.Contract.GetTask(&_Interpreter.CallOpts)
}

// GetTask is a free data retrieval call binding the contract method 0xa31662d2.
//
// Solidity: function getTask() view returns(address Task)
func (_Interpreter *InterpreterCallerSession) GetTask() (common.Address, error) {
	return _Interpreter.Contract.GetTask(&_Interpreter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Interpreter *InterpreterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Interpreter.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Interpreter *InterpreterSession) Owner() (common.Address, error) {
	return _Interpreter.Contract.Owner(&_Interpreter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Interpreter *InterpreterCallerSession) Owner() (common.Address, error) {
	return _Interpreter.Contract.Owner(&_Interpreter.CallOpts)
}

// Advertise is a paid mutator transaction binding the contract method 0x4063d563.
//
// Solidity: function advertise(address Task) returns()
func (_Interpreter *InterpreterTransactor) Advertise(opts *bind.TransactOpts, Task common.Address) (*types.Transaction, error) {
	return _Interpreter.contract.Transact(opts, "advertise", Task)
}

// Advertise is a paid mutator transaction binding the contract method 0x4063d563.
//
// Solidity: function advertise(address Task) returns()
func (_Interpreter *InterpreterSession) Advertise(Task common.Address) (*types.Transaction, error) {
	return _Interpreter.Contract.Advertise(&_Interpreter.TransactOpts, Task)
}

// Advertise is a paid mutator transaction binding the contract method 0x4063d563.
//
// Solidity: function advertise(address Task) returns()
func (_Interpreter *InterpreterTransactorSession) Advertise(Task common.Address) (*types.Transaction, error) {
	return _Interpreter.Contract.Advertise(&_Interpreter.TransactOpts, Task)
}

// Done is a paid mutator transaction binding the contract method 0x7be8f86b.
//
// Solidity: function done(address Task) returns()
func (_Interpreter *InterpreterTransactor) Done(opts *bind.TransactOpts, Task common.Address) (*types.Transaction, error) {
	return _Interpreter.contract.Transact(opts, "done", Task)
}

// Done is a paid mutator transaction binding the contract method 0x7be8f86b.
//
// Solidity: function done(address Task) returns()
func (_Interpreter *InterpreterSession) Done(Task common.Address) (*types.Transaction, error) {
	return _Interpreter.Contract.Done(&_Interpreter.TransactOpts, Task)
}

// Done is a paid mutator transaction binding the contract method 0x7be8f86b.
//
// Solidity: function done(address Task) returns()
func (_Interpreter *InterpreterTransactorSession) Done(Task common.Address) (*types.Transaction, error) {
	return _Interpreter.Contract.Done(&_Interpreter.TransactOpts, Task)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Interpreter *InterpreterTransactor) Transfer(opts *bind.TransactOpts, new_owner common.Address) (*types.Transaction, error) {
	return _Interpreter.contract.Transact(opts, "transfer", new_owner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Interpreter *InterpreterSession) Transfer(new_owner common.Address) (*types.Transaction, error) {
	return _Interpreter.Contract.Transfer(&_Interpreter.TransactOpts, new_owner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Interpreter *InterpreterTransactorSession) Transfer(new_owner common.Address) (*types.Transaction, error) {
	return _Interpreter.Contract.Transfer(&_Interpreter.TransactOpts, new_owner)
}

// RandomABI is the input ABI used to generate the binding from.
const RandomABI = "[]"

// RandomBin is the compiled bytecode used for deploying new contracts.
var RandomBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a7231582004507314633eb771c8d43f55aeb0c661c189e3377609ac19a77d6c761e7c3e1264736f6c63430005100032"

// DeployRandom deploys a new Ethereum contract, binding an instance of Random to it.
func DeployRandom(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Random, error) {
	parsed, err := abi.JSON(strings.NewReader(RandomABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RandomBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Random{RandomCaller: RandomCaller{contract: contract}, RandomTransactor: RandomTransactor{contract: contract}, RandomFilterer: RandomFilterer{contract: contract}}, nil
}

// Random is an auto generated Go binding around an Ethereum contract.
type Random struct {
	RandomCaller     // Read-only binding to the contract
	RandomTransactor // Write-only binding to the contract
	RandomFilterer   // Log filterer for contract events
}

// RandomCaller is an auto generated read-only Go binding around an Ethereum contract.
type RandomCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomSession struct {
	Contract     *Random           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RandomCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomCallerSession struct {
	Contract *RandomCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RandomTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomTransactorSession struct {
	Contract     *RandomTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RandomRaw is an auto generated low-level Go binding around an Ethereum contract.
type RandomRaw struct {
	Contract *Random // Generic contract binding to access the raw methods on
}

// RandomCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomCallerRaw struct {
	Contract *RandomCaller // Generic read-only contract binding to access the raw methods on
}

// RandomTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomTransactorRaw struct {
	Contract *RandomTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRandom creates a new instance of Random, bound to a specific deployed contract.
func NewRandom(address common.Address, backend bind.ContractBackend) (*Random, error) {
	contract, err := bindRandom(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Random{RandomCaller: RandomCaller{contract: contract}, RandomTransactor: RandomTransactor{contract: contract}, RandomFilterer: RandomFilterer{contract: contract}}, nil
}

// NewRandomCaller creates a new read-only instance of Random, bound to a specific deployed contract.
func NewRandomCaller(address common.Address, caller bind.ContractCaller) (*RandomCaller, error) {
	contract, err := bindRandom(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomCaller{contract: contract}, nil
}

// NewRandomTransactor creates a new write-only instance of Random, bound to a specific deployed contract.
func NewRandomTransactor(address common.Address, transactor bind.ContractTransactor) (*RandomTransactor, error) {
	contract, err := bindRandom(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomTransactor{contract: contract}, nil
}

// NewRandomFilterer creates a new log filterer instance of Random, bound to a specific deployed contract.
func NewRandomFilterer(address common.Address, filterer bind.ContractFilterer) (*RandomFilterer, error) {
	contract, err := bindRandom(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomFilterer{contract: contract}, nil
}

// bindRandom binds a generic wrapper to an already deployed contract.
func bindRandom(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RandomABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Random *RandomRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Random.Contract.RandomCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Random *RandomRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Random.Contract.RandomTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Random *RandomRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Random.Contract.RandomTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Random *RandomCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Random.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Random *RandomTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Random.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Random *RandomTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Random.Contract.contract.Transact(opts, method, params...)
}

// RepositoryABI is the input ABI used to generate the binding from.
const RepositoryABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"IpfsHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"AddInterpreter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"max_size\",\"type\":\"uint256\"}],\"name\":\"getInterpreter\",\"outputs\":[{\"internalType\":\"contractinterpreter\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_owner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RepositoryFuncSigs maps the 4-byte function signature to its string representation.
var RepositoryFuncSigs = map[string]string{
	"500cf9ee": "AddInterpreter(string,uint256)",
	"e197194e": "getInterpreter(uint256)",
	"8da5cb5b": "owner()",
	"1a695230": "transfer(address)",
}

// RepositoryBin is the compiled bytecode used for deploying new contracts.
var RepositoryBin = "0x6080604052600080546001600160a01b0319163317905534801561002257600080fd5b50610b51806100326000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80631a69523014610051578063500cf9ee146100795780638da5cb5b14610121578063e197194e14610145575b600080fd5b6100776004803603602081101561006757600080fd5b50356001600160a01b0316610162565b005b6100776004803603604081101561008f57600080fd5b8101906020810181356401000000008111156100aa57600080fd5b8201836020820111156100bc57600080fd5b803590602001918460018302840111640100000000831117156100de57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506101db915050565b610129610346565b604080516001600160a01b039092168252519081900360200190f35b6101296004803603602081101561015b57600080fd5b5035610355565b6000546001600160a01b031633146101b9576040805162461bcd60e51b81526020600482015260156024820152741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03163314610232576040805162461bcd60e51b81526020600482015260156024820152741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b60008260405161024190610500565b60208082528251818301528251829160408301919085019080838360005b8381101561027757818101518382015260200161025f565b50505050905090810190601f1680156102a45780820380516001836020036101000a031916815260200191505b5092505050604051809103906000f0801580156102c5573d6000803e3d6000fd5b506001805480820182557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319166001600160a01b0393909316929092179091556002805491820181556000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01919091555050565b6000546001600160a01b031690565b600080805b60025481101561039157836002828154811061037257fe5b906000526020600020015411610389576001909101905b60010161035a565b50600081116103d15760405162461bcd60e51b8152600401808060200182810382526025815260200180610af86025913960400191505060405180910390fd5b60006103dc8261048b565b600101905060005b60025481101561044e5784600282815481106103fc57fe5b90600052602060002001541161041457600019909101905b81610446576001818154811061042657fe5b6000918252602090912001546001600160a01b0316935061048692505050565b6001016103e4565b5060405162461bcd60e51b8152600401808060200182810382526025815260200180610af86025913960400191505060405180910390fd5b919050565b604080513360601b602080830191909152825180830360140181526034909201909252805191012060009081904390816104c157fe5b044301604051602001808281526020019150506040516020818303038152906040528051906020012060001c90508281816104f857fe5b069392505050565b6105ea8061050e8339019056fe6080604052600080546001600160a01b0319163317905534801561002257600080fd5b506040516105ea3803806105ea8339818101604052602081101561004557600080fd5b810190808051604051939291908464010000000082111561006557600080fd5b90830190602082018581111561007a57600080fd5b825164010000000081118282018810171561009457600080fd5b82525081516020918201929091019080838360005b838110156100c15781810151838201526020016100a9565b50505050905090810190601f1680156100ee5780820380516001836020036101000a031916815260200191505b50604052505081516101089150600190602084019061010f565b50506101aa565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061015057805160ff191683800117855561017d565b8280016001018555821561017d579182015b8281111561017d578251825591602001919060010190610162565b5061018992915061018d565b5090565b6101a791905b808211156101895760008155600101610193565b90565b610431806101b96000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631a6952301461005c5780634063d563146100845780637be8f86b146100aa5780638da5cb5b146100d0578063a31662d2146100f4575b600080fd5b6100826004803603602081101561007257600080fd5b50356001600160a01b03166100fc565b005b6100826004803603602081101561009a57600080fd5b50356001600160a01b0316610175565b610082600480360360208110156100c057600080fd5b50356001600160a01b03166101f8565b6100d86102e8565b604080516001600160a01b039092168252519081900360200190f35b6100d86102f7565b6000546001600160a01b03163314610153576040805162461bcd60e51b81526020600482015260156024820152741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6003805460018181019092557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b039093166001600160a01b03199384161790556002805491820181556000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01805490911633179055565b80336001600160a01b0382161461024e576040805162461bcd60e51b81526020600482015260156024820152741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b60005b6003548110156102e357826001600160a01b03166003828154811061027257fe5b6000918252602090912001546001600160a01b031614156102db576003818154811061029a57fe5b600091825260209091200180546001600160a01b031916905560028054829081106102c157fe5b600091825260209091200180546001600160a01b03191690555b600101610251565b505050565b6000546001600160a01b031690565b60035460009061034e576040805162461bcd60e51b815260206004820152601c60248201527f6e6f207461736b2063757272656e74726c7920617661696c61626c6500000000604482015290519081900360640190fd5b60035460009061035d90610387565b90506003818154811061036c57fe5b6000918252602090912001546001600160a01b031691505090565b604080513360601b602080830191909152825180830360140181526034909201909252805191012060009081904390816103bd57fe5b044301604051602001808281526020019150506040516020818303038152906040528051906020012060001c90508281816103f457fe5b06939250505056fea265627a7a7231582040f7639cbb6ae34bc020fa1feafb9ecd88e6ecd238d7294294d8bd2812135c3d64736f6c634300051000326e6f20696e746572707265746572207769746820736d616c6c20656e6f7567682073697a65a265627a7a72315820d2d216e27118aa65512d9f77ddede1d46b6b3a8f8bcd8df95b2512983eb7ede664736f6c63430005100032"

// DeployRepository deploys a new Ethereum contract, binding an instance of Repository to it.
func DeployRepository(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Repository, error) {
	parsed, err := abi.JSON(strings.NewReader(RepositoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RepositoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Repository{RepositoryCaller: RepositoryCaller{contract: contract}, RepositoryTransactor: RepositoryTransactor{contract: contract}, RepositoryFilterer: RepositoryFilterer{contract: contract}}, nil
}

// Repository is an auto generated Go binding around an Ethereum contract.
type Repository struct {
	RepositoryCaller     // Read-only binding to the contract
	RepositoryTransactor // Write-only binding to the contract
	RepositoryFilterer   // Log filterer for contract events
}

// RepositoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RepositoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepositoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RepositoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepositoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RepositoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepositorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RepositorySession struct {
	Contract     *Repository       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RepositoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RepositoryCallerSession struct {
	Contract *RepositoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RepositoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RepositoryTransactorSession struct {
	Contract     *RepositoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RepositoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RepositoryRaw struct {
	Contract *Repository // Generic contract binding to access the raw methods on
}

// RepositoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RepositoryCallerRaw struct {
	Contract *RepositoryCaller // Generic read-only contract binding to access the raw methods on
}

// RepositoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RepositoryTransactorRaw struct {
	Contract *RepositoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRepository creates a new instance of Repository, bound to a specific deployed contract.
func NewRepository(address common.Address, backend bind.ContractBackend) (*Repository, error) {
	contract, err := bindRepository(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Repository{RepositoryCaller: RepositoryCaller{contract: contract}, RepositoryTransactor: RepositoryTransactor{contract: contract}, RepositoryFilterer: RepositoryFilterer{contract: contract}}, nil
}

// NewRepositoryCaller creates a new read-only instance of Repository, bound to a specific deployed contract.
func NewRepositoryCaller(address common.Address, caller bind.ContractCaller) (*RepositoryCaller, error) {
	contract, err := bindRepository(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RepositoryCaller{contract: contract}, nil
}

// NewRepositoryTransactor creates a new write-only instance of Repository, bound to a specific deployed contract.
func NewRepositoryTransactor(address common.Address, transactor bind.ContractTransactor) (*RepositoryTransactor, error) {
	contract, err := bindRepository(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RepositoryTransactor{contract: contract}, nil
}

// NewRepositoryFilterer creates a new log filterer instance of Repository, bound to a specific deployed contract.
func NewRepositoryFilterer(address common.Address, filterer bind.ContractFilterer) (*RepositoryFilterer, error) {
	contract, err := bindRepository(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RepositoryFilterer{contract: contract}, nil
}

// bindRepository binds a generic wrapper to an already deployed contract.
func bindRepository(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RepositoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Repository *RepositoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Repository.Contract.RepositoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Repository *RepositoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Repository.Contract.RepositoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Repository *RepositoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Repository.Contract.RepositoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Repository *RepositoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Repository.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Repository *RepositoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Repository.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Repository *RepositoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Repository.Contract.contract.Transact(opts, method, params...)
}

// GetInterpreter is a free data retrieval call binding the contract method 0xe197194e.
//
// Solidity: function getInterpreter(uint256 max_size) view returns(address)
func (_Repository *RepositoryCaller) GetInterpreter(opts *bind.CallOpts, max_size *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Repository.contract.Call(opts, out, "getInterpreter", max_size)
	return *ret0, err
}

// GetInterpreter is a free data retrieval call binding the contract method 0xe197194e.
//
// Solidity: function getInterpreter(uint256 max_size) view returns(address)
func (_Repository *RepositorySession) GetInterpreter(max_size *big.Int) (common.Address, error) {
	return _Repository.Contract.GetInterpreter(&_Repository.CallOpts, max_size)
}

// GetInterpreter is a free data retrieval call binding the contract method 0xe197194e.
//
// Solidity: function getInterpreter(uint256 max_size) view returns(address)
func (_Repository *RepositoryCallerSession) GetInterpreter(max_size *big.Int) (common.Address, error) {
	return _Repository.Contract.GetInterpreter(&_Repository.CallOpts, max_size)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Repository *RepositoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Repository.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Repository *RepositorySession) Owner() (common.Address, error) {
	return _Repository.Contract.Owner(&_Repository.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Repository *RepositoryCallerSession) Owner() (common.Address, error) {
	return _Repository.Contract.Owner(&_Repository.CallOpts)
}

// AddInterpreter is a paid mutator transaction binding the contract method 0x500cf9ee.
//
// Solidity: function AddInterpreter(string IpfsHash, uint256 size) returns()
func (_Repository *RepositoryTransactor) AddInterpreter(opts *bind.TransactOpts, IpfsHash string, size *big.Int) (*types.Transaction, error) {
	return _Repository.contract.Transact(opts, "AddInterpreter", IpfsHash, size)
}

// AddInterpreter is a paid mutator transaction binding the contract method 0x500cf9ee.
//
// Solidity: function AddInterpreter(string IpfsHash, uint256 size) returns()
func (_Repository *RepositorySession) AddInterpreter(IpfsHash string, size *big.Int) (*types.Transaction, error) {
	return _Repository.Contract.AddInterpreter(&_Repository.TransactOpts, IpfsHash, size)
}

// AddInterpreter is a paid mutator transaction binding the contract method 0x500cf9ee.
//
// Solidity: function AddInterpreter(string IpfsHash, uint256 size) returns()
func (_Repository *RepositoryTransactorSession) AddInterpreter(IpfsHash string, size *big.Int) (*types.Transaction, error) {
	return _Repository.Contract.AddInterpreter(&_Repository.TransactOpts, IpfsHash, size)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Repository *RepositoryTransactor) Transfer(opts *bind.TransactOpts, new_owner common.Address) (*types.Transaction, error) {
	return _Repository.contract.Transact(opts, "transfer", new_owner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Repository *RepositorySession) Transfer(new_owner common.Address) (*types.Transaction, error) {
	return _Repository.Contract.Transfer(&_Repository.TransactOpts, new_owner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Repository *RepositoryTransactorSession) Transfer(new_owner common.Address) (*types.Transaction, error) {
	return _Repository.Contract.Transfer(&_Repository.TransactOpts, new_owner)
}

// StackABI is the input ABI used to generate the binding from.
const StackABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Authorized\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"read\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_owner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StackFuncSigs maps the 4-byte function signature to its string representation.
var StackFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"619b40cc": "push(address,bytes,bytes)",
	"8bf4515c": "read(bytes)",
	"1a695230": "transfer(address)",
}

// StackBin is the compiled bytecode used for deploying new contracts.
var StackBin = "0x6080604052600080546001600160a01b0319163317905534801561002257600080fd5b50604051610ac8380380610ac883398101604081905261004191610077565b600180546001600160a01b0319166001600160a01b03929092169190911790556100c5565b8051610071816100ae565b92915050565b60006020828403121561008957600080fd5b60006100958484610066565b949350505050565b60006001600160a01b038216610071565b6100b78161009d565b81146100c257600080fd5b50565b6109f4806100d46000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80631a69523014610051578063619b40cc146100665780638bf4515c146100795780638da5cb5b146100a2575b600080fd5b61006461005f36600461069e565b6100b7565b005b6100646100743660046106c4565b61010c565b61008c610087366004610741565b61028c565b60405161009991906108a6565b60405180910390f35b6100aa610591565b6040516100999190610898565b6000546001600160a01b031633146100ea5760405162461bcd60e51b81526004016100e1906108d7565b60405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b038481169116146101f45760606003836040516101339190610885565b908152604080519182900360209081018320805480830285018301909352828452919083018282801561018f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610171575b50939450600093505050505b81518110156101f1578181815181106101b057fe5b60200260200101516001600160a01b0316856001600160a01b031614156101e95760405162461bcd60e51b81526004016100e1906108b7565b60010161019b565b50505b6003826040516102049190610885565b908152604051602091819003820181208054600181018255600091825292902090910180546001600160a01b0319166001600160a01b03861617905560029061024e908490610885565b90815260405160209181900382019020805460018101808355600092835291839020845192936102859391909201918501906105a1565b5050505050565b60608060028360405161029f9190610885565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156103785760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156103645780601f1061033957610100808354040283529160200191610364565b820191906000526020600020905b81548152906001019060200180831161034757829003601f168201915b5050505050815260200190600101906102cd565b50505050905060606003846040516103909190610885565b90815260408051918290036020908101832080548083028501830190935282845291908301828280156103ec57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116103ce575b50508551939450505081151590506104165760405162461bcd60e51b81526004016100e1906108c7565b606081604051908082528060200260200182016040528015610442578160200160208202803883390190505b50905060005b828110156105735760015484516001600160a01b039091169085908390811061046d57fe5b60200260200101516001600160a01b031614156104a45784818151811061049057fe5b60200260200101519550505050505061058c565b60005b81811161056a578581815181106104ba57fe5b60200260200101516040516104cf9190610885565b60405180910390208683815181106104e357fe5b60200260200101516040516104f89190610885565b604051809103902014156105625782818151811061051257fe5b60209081029190910101805160010190526002840483828151811061053357fe5b602002602001015111156105625785828151811061054d57fe5b6020026020010151965050505050505061058c565b6001016104a7565b50600101610448565b5060405162461bcd60e51b81526004016100e1906108c7565b919050565b6000546001600160a01b03165b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106105e257805160ff191683800117855561060f565b8280016001018555821561060f579182015b8281111561060f5782518255916020019190600101906105f4565b5061061b92915061061f565b5090565b61059e91905b8082111561061b5760008155600101610625565b80356106448161099a565b92915050565b600082601f83011261065b57600080fd5b813561066e6106698261090e565b6108e7565b9150808252602083016020830185838301111561068a57600080fd5b610695838284610954565b50505092915050565b6000602082840312156106b057600080fd5b60006106bc8484610639565b949350505050565b6000806000606084860312156106d957600080fd5b60006106e58686610639565b935050602084013567ffffffffffffffff81111561070257600080fd5b61070e8682870161064a565b925050604084013567ffffffffffffffff81111561072b57600080fd5b6107378682870161064a565b9150509250925092565b60006020828403121561075357600080fd5b813567ffffffffffffffff81111561076a57600080fd5b6106bc8482850161064a565b61077f81610943565b82525050565b600061079082610936565b61079a818561093a565b93506107aa818560208601610960565b6107b381610990565b9093019392505050565b60006107c882610936565b6107d2818561058c565b93506107e2818560208601610960565b9290920192915050565b60006107f960178361093a565b7f63616e27742076616c696461746520796f757273656c66000000000000000000815260200192915050565b600061083260138361093a565b7276616c7565206e6f74206163636f7264696e6760681b815260200192915050565b600061086160158361093a565b741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b815260200192915050565b600061089182846107bd565b9392505050565b602081016106448284610776565b602080825281016108918184610785565b60208082528101610644816107ec565b6020808252810161064481610825565b6020808252810161064481610854565b60405181810167ffffffffffffffff8111828210171561090657600080fd5b604052919050565b600067ffffffffffffffff82111561092557600080fd5b506020601f91909101601f19160190565b5190565b90815260200190565b60006001600160a01b038216610644565b82818337506000910152565b60005b8381101561097b578181015183820152602001610963565b8381111561098a576000848401525b50505050565b601f01601f191690565b6109a381610943565b81146109ae57600080fd5b5056fea365627a7a72315820ab7d59b9b70cc49a5c5507fdb2aced379a69c22d72afb7fc8cbe628aa8826d656c6578706572696d656e74616cf564736f6c63430005100040"

// DeployStack deploys a new Ethereum contract, binding an instance of Stack to it.
func DeployStack(auth *bind.TransactOpts, backend bind.ContractBackend, Authorized common.Address) (common.Address, *types.Transaction, *Stack, error) {
	parsed, err := abi.JSON(strings.NewReader(StackABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StackBin), backend, Authorized)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stack{StackCaller: StackCaller{contract: contract}, StackTransactor: StackTransactor{contract: contract}, StackFilterer: StackFilterer{contract: contract}}, nil
}

// Stack is an auto generated Go binding around an Ethereum contract.
type Stack struct {
	StackCaller     // Read-only binding to the contract
	StackTransactor // Write-only binding to the contract
	StackFilterer   // Log filterer for contract events
}

// StackCaller is an auto generated read-only Go binding around an Ethereum contract.
type StackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StackSession struct {
	Contract     *Stack            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StackCallerSession struct {
	Contract *StackCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StackTransactorSession struct {
	Contract     *StackTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StackRaw is an auto generated low-level Go binding around an Ethereum contract.
type StackRaw struct {
	Contract *Stack // Generic contract binding to access the raw methods on
}

// StackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StackCallerRaw struct {
	Contract *StackCaller // Generic read-only contract binding to access the raw methods on
}

// StackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StackTransactorRaw struct {
	Contract *StackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStack creates a new instance of Stack, bound to a specific deployed contract.
func NewStack(address common.Address, backend bind.ContractBackend) (*Stack, error) {
	contract, err := bindStack(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stack{StackCaller: StackCaller{contract: contract}, StackTransactor: StackTransactor{contract: contract}, StackFilterer: StackFilterer{contract: contract}}, nil
}

// NewStackCaller creates a new read-only instance of Stack, bound to a specific deployed contract.
func NewStackCaller(address common.Address, caller bind.ContractCaller) (*StackCaller, error) {
	contract, err := bindStack(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StackCaller{contract: contract}, nil
}

// NewStackTransactor creates a new write-only instance of Stack, bound to a specific deployed contract.
func NewStackTransactor(address common.Address, transactor bind.ContractTransactor) (*StackTransactor, error) {
	contract, err := bindStack(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StackTransactor{contract: contract}, nil
}

// NewStackFilterer creates a new log filterer instance of Stack, bound to a specific deployed contract.
func NewStackFilterer(address common.Address, filterer bind.ContractFilterer) (*StackFilterer, error) {
	contract, err := bindStack(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StackFilterer{contract: contract}, nil
}

// bindStack binds a generic wrapper to an already deployed contract.
func bindStack(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stack *StackRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stack.Contract.StackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stack *StackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stack.Contract.StackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stack *StackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stack.Contract.StackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stack *StackCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stack.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stack *StackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stack.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stack *StackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stack.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stack *StackCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stack.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stack *StackSession) Owner() (common.Address, error) {
	return _Stack.Contract.Owner(&_Stack.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stack *StackCallerSession) Owner() (common.Address, error) {
	return _Stack.Contract.Owner(&_Stack.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x8bf4515c.
//
// Solidity: function read(bytes key) view returns(bytes value)
func (_Stack *StackCaller) Read(opts *bind.CallOpts, key []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Stack.contract.Call(opts, out, "read", key)
	return *ret0, err
}

// Read is a free data retrieval call binding the contract method 0x8bf4515c.
//
// Solidity: function read(bytes key) view returns(bytes value)
func (_Stack *StackSession) Read(key []byte) ([]byte, error) {
	return _Stack.Contract.Read(&_Stack.CallOpts, key)
}

// Read is a free data retrieval call binding the contract method 0x8bf4515c.
//
// Solidity: function read(bytes key) view returns(bytes value)
func (_Stack *StackCallerSession) Read(key []byte) ([]byte, error) {
	return _Stack.Contract.Read(&_Stack.CallOpts, key)
}

// Push is a paid mutator transaction binding the contract method 0x619b40cc.
//
// Solidity: function push(address sender, bytes key, bytes value) returns()
func (_Stack *StackTransactor) Push(opts *bind.TransactOpts, sender common.Address, key []byte, value []byte) (*types.Transaction, error) {
	return _Stack.contract.Transact(opts, "push", sender, key, value)
}

// Push is a paid mutator transaction binding the contract method 0x619b40cc.
//
// Solidity: function push(address sender, bytes key, bytes value) returns()
func (_Stack *StackSession) Push(sender common.Address, key []byte, value []byte) (*types.Transaction, error) {
	return _Stack.Contract.Push(&_Stack.TransactOpts, sender, key, value)
}

// Push is a paid mutator transaction binding the contract method 0x619b40cc.
//
// Solidity: function push(address sender, bytes key, bytes value) returns()
func (_Stack *StackTransactorSession) Push(sender common.Address, key []byte, value []byte) (*types.Transaction, error) {
	return _Stack.Contract.Push(&_Stack.TransactOpts, sender, key, value)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Stack *StackTransactor) Transfer(opts *bind.TransactOpts, new_owner common.Address) (*types.Transaction, error) {
	return _Stack.contract.Transact(opts, "transfer", new_owner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Stack *StackSession) Transfer(new_owner common.Address) (*types.Transaction, error) {
	return _Stack.Contract.Transfer(&_Stack.TransactOpts, new_owner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Stack *StackTransactorSession) Transfer(new_owner common.Address) (*types.Transaction, error) {
	return _Stack.Contract.Transfer(&_Stack.TransactOpts, new_owner)
}

// StakeABI is the input ABI used to generate the binding from.
const StakeABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"complete\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"delete_idx\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"delete_user\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_owner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StakeFuncSigs maps the 4-byte function signature to its string representation.
var StakeFuncSigs = map[string]string{
	"971d852f": "complete(uint256)",
	"b4bec8e9": "delete_idx(uint256)",
	"f7cfcabf": "delete_user(address)",
	"8da5cb5b": "owner()",
	"590e1ae3": "refund()",
	"1a695230": "transfer(address)",
}

// StakeBin is the compiled bytecode used for deploying new contracts.
var StakeBin = "0x6080604052600080546001600160a01b0319163317905534801561002257600080fd5b50610477806100326000396000f3fe6080604052600436106100555760003560e01c80631a6952301461005a578063590e1ae31461008f5780638da5cb5b14610097578063971d852f146100c8578063b4bec8e9146100e5578063f7cfcabf1461010f575b600080fd5b34801561006657600080fd5b5061008d6004803603602081101561007d57600080fd5b50356001600160a01b0316610142565b005b61008d6101bb565b3480156100a357600080fd5b506100ac610258565b604080516001600160a01b039092168252519081900360200190f35b61008d600480360360208110156100de57600080fd5b5035610267565b3480156100f157600080fd5b5061008d6004803603602081101561010857600080fd5b5035610306565b34801561011b57600080fd5b5061008d6004803603602081101561013257600080fd5b50356001600160a01b031661039f565b6000546001600160a01b03163314610199576040805162461bcd60e51b81526020600482015260156024820152741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b600154600019015b336001600160a01b0316600182815481106101da57fe5b6000918252602090912001546001600160a01b0316141561024f57336001600160a01b03166108fc6002838154811061020f57fe5b90600052602060002001549081150290604051600060405180830381858888f19350505050158015610245573d6000803e3d6000fd5b5061024f81610306565b600019016101c3565b6000546001600160a01b031690565b600154600019015b336001600160a01b03166001828154811061028657fe5b6000918252602090912001546001600160a01b031614156102fd57336001600160a01b03166108fc83600284815481106102bc57fe5b9060005260206000200154019081150290604051600060405180830381858888f193505050501580156102f3573d6000803e3d6000fd5b506102fd81610306565b6000190161026f565b6000546001600160a01b0316331461035d576040805162461bcd60e51b81526020600482015260156024820152741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b6001818154811061036a57fe5b600091825260209091200180546001600160a01b0319169055600280548290811061039157fe5b600091825260208220015550565b6000546001600160a01b031633146103f6576040805162461bcd60e51b81526020600482015260156024820152741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b600154600019015b816001600160a01b03166001828154811061041557fe5b6000918252602090912001546001600160a01b031614156104395761043981610306565b600019016103fe56fea265627a7a72315820b0ba422a5d2d673409e5651b3b72fe8543ad604b1e3273deae94d12495c2058864736f6c63430005100032"

// DeployStake deploys a new Ethereum contract, binding an instance of Stake to it.
func DeployStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Stake, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StakeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stake{StakeCaller: StakeCaller{contract: contract}, StakeTransactor: StakeTransactor{contract: contract}, StakeFilterer: StakeFilterer{contract: contract}}, nil
}

// Stake is an auto generated Go binding around an Ethereum contract.
type Stake struct {
	StakeCaller     // Read-only binding to the contract
	StakeTransactor // Write-only binding to the contract
	StakeFilterer   // Log filterer for contract events
}

// StakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeSession struct {
	Contract     *Stake            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeCallerSession struct {
	Contract *StakeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeTransactorSession struct {
	Contract     *StakeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeRaw struct {
	Contract *Stake // Generic contract binding to access the raw methods on
}

// StakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeCallerRaw struct {
	Contract *StakeCaller // Generic read-only contract binding to access the raw methods on
}

// StakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeTransactorRaw struct {
	Contract *StakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStake creates a new instance of Stake, bound to a specific deployed contract.
func NewStake(address common.Address, backend bind.ContractBackend) (*Stake, error) {
	contract, err := bindStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stake{StakeCaller: StakeCaller{contract: contract}, StakeTransactor: StakeTransactor{contract: contract}, StakeFilterer: StakeFilterer{contract: contract}}, nil
}

// NewStakeCaller creates a new read-only instance of Stake, bound to a specific deployed contract.
func NewStakeCaller(address common.Address, caller bind.ContractCaller) (*StakeCaller, error) {
	contract, err := bindStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeCaller{contract: contract}, nil
}

// NewStakeTransactor creates a new write-only instance of Stake, bound to a specific deployed contract.
func NewStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeTransactor, error) {
	contract, err := bindStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeTransactor{contract: contract}, nil
}

// NewStakeFilterer creates a new log filterer instance of Stake, bound to a specific deployed contract.
func NewStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeFilterer, error) {
	contract, err := bindStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeFilterer{contract: contract}, nil
}

// bindStake binds a generic wrapper to an already deployed contract.
func bindStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.StakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stake.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeSession) Owner() (common.Address, error) {
	return _Stake.Contract.Owner(&_Stake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeCallerSession) Owner() (common.Address, error) {
	return _Stake.Contract.Owner(&_Stake.CallOpts)
}

// Complete is a paid mutator transaction binding the contract method 0x971d852f.
//
// Solidity: function complete(uint256 reward) payable returns()
func (_Stake *StakeTransactor) Complete(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "complete", reward)
}

// Complete is a paid mutator transaction binding the contract method 0x971d852f.
//
// Solidity: function complete(uint256 reward) payable returns()
func (_Stake *StakeSession) Complete(reward *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Complete(&_Stake.TransactOpts, reward)
}

// Complete is a paid mutator transaction binding the contract method 0x971d852f.
//
// Solidity: function complete(uint256 reward) payable returns()
func (_Stake *StakeTransactorSession) Complete(reward *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Complete(&_Stake.TransactOpts, reward)
}

// DeleteIdx is a paid mutator transaction binding the contract method 0xb4bec8e9.
//
// Solidity: function delete_idx(uint256 i) returns()
func (_Stake *StakeTransactor) DeleteIdx(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "delete_idx", i)
}

// DeleteIdx is a paid mutator transaction binding the contract method 0xb4bec8e9.
//
// Solidity: function delete_idx(uint256 i) returns()
func (_Stake *StakeSession) DeleteIdx(i *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.DeleteIdx(&_Stake.TransactOpts, i)
}

// DeleteIdx is a paid mutator transaction binding the contract method 0xb4bec8e9.
//
// Solidity: function delete_idx(uint256 i) returns()
func (_Stake *StakeTransactorSession) DeleteIdx(i *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.DeleteIdx(&_Stake.TransactOpts, i)
}

// DeleteUser is a paid mutator transaction binding the contract method 0xf7cfcabf.
//
// Solidity: function delete_user(address user) returns()
func (_Stake *StakeTransactor) DeleteUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "delete_user", user)
}

// DeleteUser is a paid mutator transaction binding the contract method 0xf7cfcabf.
//
// Solidity: function delete_user(address user) returns()
func (_Stake *StakeSession) DeleteUser(user common.Address) (*types.Transaction, error) {
	return _Stake.Contract.DeleteUser(&_Stake.TransactOpts, user)
}

// DeleteUser is a paid mutator transaction binding the contract method 0xf7cfcabf.
//
// Solidity: function delete_user(address user) returns()
func (_Stake *StakeTransactorSession) DeleteUser(user common.Address) (*types.Transaction, error) {
	return _Stake.Contract.DeleteUser(&_Stake.TransactOpts, user)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() payable returns()
func (_Stake *StakeTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() payable returns()
func (_Stake *StakeSession) Refund() (*types.Transaction, error) {
	return _Stake.Contract.Refund(&_Stake.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() payable returns()
func (_Stake *StakeTransactorSession) Refund() (*types.Transaction, error) {
	return _Stake.Contract.Refund(&_Stake.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Stake *StakeTransactor) Transfer(opts *bind.TransactOpts, new_owner common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "transfer", new_owner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Stake *StakeSession) Transfer(new_owner common.Address) (*types.Transaction, error) {
	return _Stake.Contract.Transfer(&_Stake.TransactOpts, new_owner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Stake *StakeTransactorSession) Transfer(new_owner common.Address) (*types.Transaction, error) {
	return _Stake.Contract.Transfer(&_Stake.TransactOpts, new_owner)
}

// TaskABI is the input ABI used to generate the binding from.
const TaskABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"Bid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Reward\",\"type\":\"uint256\"},{\"internalType\":\"contractinterpreter\",\"name\":\"Ipfs_object\",\"type\":\"address\"},{\"internalType\":\"uint256[][]\",\"name\":\"Task_ids\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256\",\"name\":\"Verify_prob\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Verify_pool\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"complete\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"delete_idx\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"delete_user\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"done\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"task_id\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"keys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"values\",\"type\":\"bytes[]\"}],\"name\":\"finish\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCommand\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"task_id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_owner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TaskFuncSigs maps the 4-byte function signature to its string representation.
var TaskFuncSigs = map[string]string{
	"971d852f": "complete(uint256)",
	"b4bec8e9": "delete_idx(uint256)",
	"f7cfcabf": "delete_user(address)",
	"ae8421e1": "done()",
	"532b0a59": "finish(uint256,bytes[],bytes[])",
	"0110da29": "getCommand()",
	"8da5cb5b": "owner()",
	"590e1ae3": "refund()",
	"1a695230": "transfer(address)",
}

// TaskBin is the compiled bytecode used for deploying new contracts.
var TaskBin = "0x6080604052600080546001600160a01b031916331790553480156200002357600080fd5b5060405162001f4b38038062001f4b833981016040819052620000469162000410565b80826002021115620000755760405162461bcd60e51b81526004016200006c90620005b8565b60405180910390fd5b60038690556004859055825162000094906009906020860190620001d9565b5060058290556006819055600780546001600160a01b0319166001600160a01b038681169190911791829055604051634063d56360e01b8152911690634063d56390620000e690309060040162000596565b600060405180830381600087803b1580156200010157600080fd5b505af115801562000116573d6000803e3d6000fd5b505050503360405162000129906200023d565b62000135919062000586565b604051809103906000f08015801562000152573d6000803e3d6000fd5b50600880546001600160a01b0319166001600160a01b03929092169190911790556200017e306200018a565b50505050505062000670565b6000546001600160a01b03163314620001b75760405162461bcd60e51b81526004016200006c90620005a6565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b8280548282559060005260206000209081019282156200022b579160200282015b828111156200022b57825180516200021a9184916020909101906200024b565b5091602001919060010190620001fa565b506200023992915062000297565b5090565b610ac8806200148383390190565b82805482825590600052602060002090810192821562000289579160200282015b82811115620002895782518255916020019190600101906200026c565b5062000239929150620002c2565b620002bf91905b8082111562000239576000620002b58282620002df565b506001016200029e565b90565b620002bf91905b80821115620002395760008155600101620002c9565b5080546000825590600052602060002090810190620002ff9190620002c2565b50565b600082601f8301126200031457600080fd5b81516200032b6200032582620005f1565b620005ca565b81815260209384019390925082018360005b838110156200036d578151860162000356888262000377565b84525060209283019291909101906001016200033d565b5050505092915050565b600082601f8301126200038957600080fd5b81516200039a6200032582620005f1565b91508181835260208401935060208101905083856020840282011115620003c057600080fd5b60005b838110156200036d5781620003d9888262000403565b8452506020928301929190910190600101620003c3565b8051620003fd816200064e565b92915050565b8051620003fd8162000665565b60008060008060008060c087890312156200042a57600080fd5b600062000438898962000403565b96505060206200044b89828a0162000403565b95505060406200045e89828a01620003f0565b94505060608701516001600160401b038111156200047b57600080fd5b6200048989828a0162000302565b93505060806200049c89828a0162000403565b92505060a0620004af89828a0162000403565b9150509295509295509295565b620004c78162000641565b82525050565b620004c78162000628565b6000620004e760158362000612565b7f73656e646572206e6f7420617574686f72697a65640000000000000000000000815260200192915050565b60006200052260498362000612565b7f746865207665726966792070726f626162696c6974792069732074776f20686981527f676820696e20636f6d7061726169736f6e20746f207468652076657269667920602082015268706f6f6c2073697a6560b81b604082015260600192915050565b60208101620003fd8284620004bc565b60208101620003fd8284620004cd565b60208082528101620003fd81620004d8565b60208082528101620003fd8162000513565b6040518181016001600160401b0381118282101715620005e957600080fd5b604052919050565b60006001600160401b038211156200060857600080fd5b5060209081020190565b90815260200190565b6000620003fd8262000635565b6000620003fd826200061b565b6001600160a01b031690565b6000620003fd8262000628565b620006598162000628565b8114620002ff57600080fd5b6200065981620002bf565b610e0380620006806000396000f3fe6080604052600436106100865760003560e01c80638da5cb5b116100595780638da5cb5b146100f3578063971d852f14610115578063ae8421e114610128578063b4bec8e91461013d578063f7cfcabf1461015d57610086565b80630110da291461008b5780631a695230146100b6578063532b0a59146100d8578063590e1ae3146100eb575b600080fd5b34801561009757600080fd5b506100a061017d565b6040516100ad9190610c98565b60405180910390f35b3480156100c257600080fd5b506100d66100d1366004610947565b610203565b005b6100d66100e636600461098b565b61024f565b6100d6610508565b3480156100ff57600080fd5b506101086105a8565b6040516100ad9190610bdf565b6100d661012336600461096d565b6105b7565b34801561013457600080fd5b506100d6610656565b34801561014957600080fd5b506100d661015836600461096d565b6106e3565b34801561016957600080fd5b506100d6610178366004610947565b61074f565b6009546000906101a85760405162461bcd60e51b815260040161019f90610c58565b60405180910390fd5b60006101cc60096000815481106101bb57fe5b6000918252602090912001546107c5565b905060096000815481106101dc57fe5b9060005260206000200181815481106101f157fe5b90600052602060002001549150505b90565b6000546001600160a01b0316331461022d5760405162461bcd60e51b815260040161019f90610c48565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6003548034146102715760405162461bcd60e51b815260040161019f90610c68565b6001805480820182557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b03191633179055600280549182018155600052347f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace909101556009546102fd5760405162461bcd60e51b815260040161019f90610c78565b815183511461031e5760405162461bcd60e51b815260040161019f90610c38565b6000805b6009548110156103955785600960008154811061033b57fe5b90600052602060002001828154811061035057fe5b9060005260206000200154141561038d576009818154811061036e57fe5b9060005260206000200160006103849190610839565b60019150610395565b600101610322565b50806103b35760405162461bcd60e51b815260040161019f90610c88565b6103be6005546107c5565b6104065760005b6006548110156104045760096000815481106103dd57fe5b600091825260208083209091018054600181810183559184529190922001879055016103c5565b505b60005b84518110156104af5760085485516001600160a01b039091169063619b40cc90339088908590811061043757fe5b602002602001015187858151811061044b57fe5b60200260200101516040518463ffffffff1660e01b815260040161047193929190610bed565b600060405180830381600087803b15801561048b57600080fd5b505af115801561049f573d6000803e3d6000fd5b5050600190920191506104099050565b5060096000815481106104be57fe5b6000918252602090912001546104f25760096000815481106104dc57fe5b9060005260206000200160006104f29190610839565b60095461050157610501610656565b5050505050565b600154600019015b336001600160a01b03166001828154811061052757fe5b6000918252602090912001546001600160a01b0316141561059c57336001600160a01b03166108fc6002838154811061055c57fe5b90600052602060002001549081150290604051600060405180830381858888f19350505050158015610592573d6000803e3d6000fd5b5061059c816106e3565b60001901610510565b50565b6000546001600160a01b031690565b600154600019015b336001600160a01b0316600182815481106105d657fe5b6000918252602090912001546001600160a01b0316141561064d57336001600160a01b03166108fc836002848154811061060c57fe5b9060005260206000200154019081150290604051600060405180830381858888f19350505050158015610643573d6000803e3d6000fd5b5061064d816106e3565b600019016105bf565b303381146106765760405162461bcd60e51b815260040161019f90610c48565b600754604051637be8f86b60e01b81526001600160a01b0390911690637be8f86b906106a6903090600401610c2a565b600060405180830381600087803b1580156106c057600080fd5b505af11580156106d4573d6000803e3d6000fd5b505050506105a56004546105b7565b6000546001600160a01b0316331461070d5760405162461bcd60e51b815260040161019f90610c48565b6001818154811061071a57fe5b600091825260209091200180546001600160a01b0319169055600280548290811061074157fe5b600091825260208220015550565b6000546001600160a01b031633146107795760405162461bcd60e51b815260040161019f90610c48565b600154600019015b816001600160a01b03166001828154811061079857fe5b6000918252602090912001546001600160a01b031614156107bc576107bc816106e3565b60001901610781565b60008043336040516020016107da9190610bb5565b6040516020818303038152906040528051906020012060001c816107fa57fe5b04430160405160200161080d9190610bca565b6040516020818303038152906040528051906020012060001c905082818161083157fe5b069392505050565b50805460008255906000526020600020908101906105a5919061020091905b8082111561086c5760008155600101610858565b5090565b803561087b81610da3565b92915050565b600082601f83011261089257600080fd5b81356108a56108a082610ccd565b610ca6565b81815260209384019390925082018360005b838110156108e357813586016108cd88826108ed565b84525060209283019291909101906001016108b7565b5050505092915050565b600082601f8301126108fe57600080fd5b813561090c6108a082610cee565b9150808252602083016020830185838301111561092857600080fd5b610933838284610d46565b50505092915050565b803561087b81610db7565b60006020828403121561095957600080fd5b60006109658484610870565b949350505050565b60006020828403121561097f57600080fd5b6000610965848461093c565b6000806000606084860312156109a057600080fd5b60006109ac868661093c565b935050602084013567ffffffffffffffff8111156109c957600080fd5b6109d586828701610881565b925050604084013567ffffffffffffffff8111156109f257600080fd5b6109fe86828701610881565b9150509250925092565b610a1181610d34565b82525050565b610a11610a2382610d23565b610d82565b610a1181610d23565b6000610a3c82610d16565b610a468185610d1a565b9350610a56818560208601610d52565b610a5f81610d93565b9093019392505050565b610a1181610d3b565b6000610a7f602283610d1a565b7f6b65797320616e642076616c756573206c656e67746820646f6e2774206d61748152610c6d60f31b602082015260400192915050565b6000610ac3601583610d1a565b741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b815260200192915050565b6000610af4601183610d1a565b702737903a30b9b5903a37903932ba3ab93760791b815260200192915050565b6000610b21601283610d1a565b711d985b1d59481a5cc81b9bdd081c9a59da1d60721b815260200192915050565b6000610b4f600c83610d1a565b6b616c726561647920646f6e6560a01b815260200192915050565b6000610b77601583610d1a565b746e6f20636f72726573706f6e64696e67207461736b60581b815260200192915050565b610a1181610200565b610a11610bb082610200565b610200565b6000610bc18284610a17565b50601401919050565b6000610bd68284610ba4565b50602001919050565b6020810161087b8284610a28565b60608101610bfb8286610a08565b8181036020830152610c0d8185610a31565b90508181036040830152610c218184610a31565b95945050505050565b6020810161087b8284610a69565b6020808252810161087b81610a72565b6020808252810161087b81610ab6565b6020808252810161087b81610ae7565b6020808252810161087b81610b14565b6020808252810161087b81610b42565b6020808252810161087b81610b6a565b6020810161087b8284610b9b565b60405181810167ffffffffffffffff81118282101715610cc557600080fd5b604052919050565b600067ffffffffffffffff821115610ce457600080fd5b5060209081020190565b600067ffffffffffffffff821115610d0557600080fd5b506020601f91909101601f19160190565b5190565b90815260200190565b60006001600160a01b03821661087b565b600061087b825b600061087b82610d23565b82818337506000910152565b60005b83811015610d6d578181015183820152602001610d55565b83811115610d7c576000848401525b50505050565b600061087b82600061087b82610d9d565b601f01601f191690565b60601b90565b610dac81610d23565b81146105a557600080fd5b610dac8161020056fea365627a7a723158203dfa3c7ac3cb21bd54413f21d4c75ba42ed51ea3e6335da5780cbee7ae07006d6c6578706572696d656e74616cf564736f6c634300051000406080604052600080546001600160a01b0319163317905534801561002257600080fd5b50604051610ac8380380610ac883398101604081905261004191610077565b600180546001600160a01b0319166001600160a01b03929092169190911790556100c5565b8051610071816100ae565b92915050565b60006020828403121561008957600080fd5b60006100958484610066565b949350505050565b60006001600160a01b038216610071565b6100b78161009d565b81146100c257600080fd5b50565b6109f4806100d46000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80631a69523014610051578063619b40cc146100665780638bf4515c146100795780638da5cb5b146100a2575b600080fd5b61006461005f36600461069e565b6100b7565b005b6100646100743660046106c4565b61010c565b61008c610087366004610741565b61028c565b60405161009991906108a6565b60405180910390f35b6100aa610591565b6040516100999190610898565b6000546001600160a01b031633146100ea5760405162461bcd60e51b81526004016100e1906108d7565b60405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b038481169116146101f45760606003836040516101339190610885565b908152604080519182900360209081018320805480830285018301909352828452919083018282801561018f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610171575b50939450600093505050505b81518110156101f1578181815181106101b057fe5b60200260200101516001600160a01b0316856001600160a01b031614156101e95760405162461bcd60e51b81526004016100e1906108b7565b60010161019b565b50505b6003826040516102049190610885565b908152604051602091819003820181208054600181018255600091825292902090910180546001600160a01b0319166001600160a01b03861617905560029061024e908490610885565b90815260405160209181900382019020805460018101808355600092835291839020845192936102859391909201918501906105a1565b5050505050565b60608060028360405161029f9190610885565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156103785760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156103645780601f1061033957610100808354040283529160200191610364565b820191906000526020600020905b81548152906001019060200180831161034757829003601f168201915b5050505050815260200190600101906102cd565b50505050905060606003846040516103909190610885565b90815260408051918290036020908101832080548083028501830190935282845291908301828280156103ec57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116103ce575b50508551939450505081151590506104165760405162461bcd60e51b81526004016100e1906108c7565b606081604051908082528060200260200182016040528015610442578160200160208202803883390190505b50905060005b828110156105735760015484516001600160a01b039091169085908390811061046d57fe5b60200260200101516001600160a01b031614156104a45784818151811061049057fe5b60200260200101519550505050505061058c565b60005b81811161056a578581815181106104ba57fe5b60200260200101516040516104cf9190610885565b60405180910390208683815181106104e357fe5b60200260200101516040516104f89190610885565b604051809103902014156105625782818151811061051257fe5b60209081029190910101805160010190526002840483828151811061053357fe5b602002602001015111156105625785828151811061054d57fe5b6020026020010151965050505050505061058c565b6001016104a7565b50600101610448565b5060405162461bcd60e51b81526004016100e1906108c7565b919050565b6000546001600160a01b03165b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106105e257805160ff191683800117855561060f565b8280016001018555821561060f579182015b8281111561060f5782518255916020019190600101906105f4565b5061061b92915061061f565b5090565b61059e91905b8082111561061b5760008155600101610625565b80356106448161099a565b92915050565b600082601f83011261065b57600080fd5b813561066e6106698261090e565b6108e7565b9150808252602083016020830185838301111561068a57600080fd5b610695838284610954565b50505092915050565b6000602082840312156106b057600080fd5b60006106bc8484610639565b949350505050565b6000806000606084860312156106d957600080fd5b60006106e58686610639565b935050602084013567ffffffffffffffff81111561070257600080fd5b61070e8682870161064a565b925050604084013567ffffffffffffffff81111561072b57600080fd5b6107378682870161064a565b9150509250925092565b60006020828403121561075357600080fd5b813567ffffffffffffffff81111561076a57600080fd5b6106bc8482850161064a565b61077f81610943565b82525050565b600061079082610936565b61079a818561093a565b93506107aa818560208601610960565b6107b381610990565b9093019392505050565b60006107c882610936565b6107d2818561058c565b93506107e2818560208601610960565b9290920192915050565b60006107f960178361093a565b7f63616e27742076616c696461746520796f757273656c66000000000000000000815260200192915050565b600061083260138361093a565b7276616c7565206e6f74206163636f7264696e6760681b815260200192915050565b600061086160158361093a565b741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b815260200192915050565b600061089182846107bd565b9392505050565b602081016106448284610776565b602080825281016108918184610785565b60208082528101610644816107ec565b6020808252810161064481610825565b6020808252810161064481610854565b60405181810167ffffffffffffffff8111828210171561090657600080fd5b604052919050565b600067ffffffffffffffff82111561092557600080fd5b506020601f91909101601f19160190565b5190565b90815260200190565b60006001600160a01b038216610644565b82818337506000910152565b60005b8381101561097b578181015183820152602001610963565b8381111561098a576000848401525b50505050565b601f01601f191690565b6109a381610943565b81146109ae57600080fd5b5056fea365627a7a72315820ab7d59b9b70cc49a5c5507fdb2aced379a69c22d72afb7fc8cbe628aa8826d656c6578706572696d656e74616cf564736f6c63430005100040"

// DeployTask deploys a new Ethereum contract, binding an instance of Task to it.
func DeployTask(auth *bind.TransactOpts, backend bind.ContractBackend, Bid *big.Int, Reward *big.Int, Ipfs_object common.Address, Task_ids [][]*big.Int, Verify_prob *big.Int, Verify_pool *big.Int) (common.Address, *types.Transaction, *Task, error) {
	parsed, err := abi.JSON(strings.NewReader(TaskABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TaskBin), backend, Bid, Reward, Ipfs_object, Task_ids, Verify_prob, Verify_pool)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Task{TaskCaller: TaskCaller{contract: contract}, TaskTransactor: TaskTransactor{contract: contract}, TaskFilterer: TaskFilterer{contract: contract}}, nil
}

// Task is an auto generated Go binding around an Ethereum contract.
type Task struct {
	TaskCaller     // Read-only binding to the contract
	TaskTransactor // Write-only binding to the contract
	TaskFilterer   // Log filterer for contract events
}

// TaskCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskSession struct {
	Contract     *Task             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskCallerSession struct {
	Contract *TaskCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TaskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskTransactorSession struct {
	Contract     *TaskTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskRaw struct {
	Contract *Task // Generic contract binding to access the raw methods on
}

// TaskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskCallerRaw struct {
	Contract *TaskCaller // Generic read-only contract binding to access the raw methods on
}

// TaskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskTransactorRaw struct {
	Contract *TaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTask creates a new instance of Task, bound to a specific deployed contract.
func NewTask(address common.Address, backend bind.ContractBackend) (*Task, error) {
	contract, err := bindTask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Task{TaskCaller: TaskCaller{contract: contract}, TaskTransactor: TaskTransactor{contract: contract}, TaskFilterer: TaskFilterer{contract: contract}}, nil
}

// NewTaskCaller creates a new read-only instance of Task, bound to a specific deployed contract.
func NewTaskCaller(address common.Address, caller bind.ContractCaller) (*TaskCaller, error) {
	contract, err := bindTask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskCaller{contract: contract}, nil
}

// NewTaskTransactor creates a new write-only instance of Task, bound to a specific deployed contract.
func NewTaskTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskTransactor, error) {
	contract, err := bindTask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskTransactor{contract: contract}, nil
}

// NewTaskFilterer creates a new log filterer instance of Task, bound to a specific deployed contract.
func NewTaskFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskFilterer, error) {
	contract, err := bindTask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskFilterer{contract: contract}, nil
}

// bindTask binds a generic wrapper to an already deployed contract.
func bindTask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TaskABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Task.Contract.TaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.Contract.TaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Task.Contract.TaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Task.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Task.Contract.contract.Transact(opts, method, params...)
}

// GetCommand is a free data retrieval call binding the contract method 0x0110da29.
//
// Solidity: function getCommand() view returns(uint256 task_id)
func (_Task *TaskCaller) GetCommand(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "getCommand")
	return *ret0, err
}

// GetCommand is a free data retrieval call binding the contract method 0x0110da29.
//
// Solidity: function getCommand() view returns(uint256 task_id)
func (_Task *TaskSession) GetCommand() (*big.Int, error) {
	return _Task.Contract.GetCommand(&_Task.CallOpts)
}

// GetCommand is a free data retrieval call binding the contract method 0x0110da29.
//
// Solidity: function getCommand() view returns(uint256 task_id)
func (_Task *TaskCallerSession) GetCommand() (*big.Int, error) {
	return _Task.Contract.GetCommand(&_Task.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Task *TaskCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Task *TaskSession) Owner() (common.Address, error) {
	return _Task.Contract.Owner(&_Task.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Task *TaskCallerSession) Owner() (common.Address, error) {
	return _Task.Contract.Owner(&_Task.CallOpts)
}

// Complete is a paid mutator transaction binding the contract method 0x971d852f.
//
// Solidity: function complete(uint256 reward) payable returns()
func (_Task *TaskTransactor) Complete(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "complete", reward)
}

// Complete is a paid mutator transaction binding the contract method 0x971d852f.
//
// Solidity: function complete(uint256 reward) payable returns()
func (_Task *TaskSession) Complete(reward *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Complete(&_Task.TransactOpts, reward)
}

// Complete is a paid mutator transaction binding the contract method 0x971d852f.
//
// Solidity: function complete(uint256 reward) payable returns()
func (_Task *TaskTransactorSession) Complete(reward *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Complete(&_Task.TransactOpts, reward)
}

// DeleteIdx is a paid mutator transaction binding the contract method 0xb4bec8e9.
//
// Solidity: function delete_idx(uint256 i) returns()
func (_Task *TaskTransactor) DeleteIdx(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "delete_idx", i)
}

// DeleteIdx is a paid mutator transaction binding the contract method 0xb4bec8e9.
//
// Solidity: function delete_idx(uint256 i) returns()
func (_Task *TaskSession) DeleteIdx(i *big.Int) (*types.Transaction, error) {
	return _Task.Contract.DeleteIdx(&_Task.TransactOpts, i)
}

// DeleteIdx is a paid mutator transaction binding the contract method 0xb4bec8e9.
//
// Solidity: function delete_idx(uint256 i) returns()
func (_Task *TaskTransactorSession) DeleteIdx(i *big.Int) (*types.Transaction, error) {
	return _Task.Contract.DeleteIdx(&_Task.TransactOpts, i)
}

// DeleteUser is a paid mutator transaction binding the contract method 0xf7cfcabf.
//
// Solidity: function delete_user(address user) returns()
func (_Task *TaskTransactor) DeleteUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "delete_user", user)
}

// DeleteUser is a paid mutator transaction binding the contract method 0xf7cfcabf.
//
// Solidity: function delete_user(address user) returns()
func (_Task *TaskSession) DeleteUser(user common.Address) (*types.Transaction, error) {
	return _Task.Contract.DeleteUser(&_Task.TransactOpts, user)
}

// DeleteUser is a paid mutator transaction binding the contract method 0xf7cfcabf.
//
// Solidity: function delete_user(address user) returns()
func (_Task *TaskTransactorSession) DeleteUser(user common.Address) (*types.Transaction, error) {
	return _Task.Contract.DeleteUser(&_Task.TransactOpts, user)
}

// Done is a paid mutator transaction binding the contract method 0xae8421e1.
//
// Solidity: function done() returns()
func (_Task *TaskTransactor) Done(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "done")
}

// Done is a paid mutator transaction binding the contract method 0xae8421e1.
//
// Solidity: function done() returns()
func (_Task *TaskSession) Done() (*types.Transaction, error) {
	return _Task.Contract.Done(&_Task.TransactOpts)
}

// Done is a paid mutator transaction binding the contract method 0xae8421e1.
//
// Solidity: function done() returns()
func (_Task *TaskTransactorSession) Done() (*types.Transaction, error) {
	return _Task.Contract.Done(&_Task.TransactOpts)
}

// Finish is a paid mutator transaction binding the contract method 0x532b0a59.
//
// Solidity: function finish(uint256 task_id, bytes[] keys, bytes[] values) payable returns()
func (_Task *TaskTransactor) Finish(opts *bind.TransactOpts, task_id *big.Int, keys [][]byte, values [][]byte) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "finish", task_id, keys, values)
}

// Finish is a paid mutator transaction binding the contract method 0x532b0a59.
//
// Solidity: function finish(uint256 task_id, bytes[] keys, bytes[] values) payable returns()
func (_Task *TaskSession) Finish(task_id *big.Int, keys [][]byte, values [][]byte) (*types.Transaction, error) {
	return _Task.Contract.Finish(&_Task.TransactOpts, task_id, keys, values)
}

// Finish is a paid mutator transaction binding the contract method 0x532b0a59.
//
// Solidity: function finish(uint256 task_id, bytes[] keys, bytes[] values) payable returns()
func (_Task *TaskTransactorSession) Finish(task_id *big.Int, keys [][]byte, values [][]byte) (*types.Transaction, error) {
	return _Task.Contract.Finish(&_Task.TransactOpts, task_id, keys, values)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() payable returns()
func (_Task *TaskTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() payable returns()
func (_Task *TaskSession) Refund() (*types.Transaction, error) {
	return _Task.Contract.Refund(&_Task.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() payable returns()
func (_Task *TaskTransactorSession) Refund() (*types.Transaction, error) {
	return _Task.Contract.Refund(&_Task.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Task *TaskTransactor) Transfer(opts *bind.TransactOpts, new_owner common.Address) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "transfer", new_owner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Task *TaskSession) Transfer(new_owner common.Address) (*types.Transaction, error) {
	return _Task.Contract.Transfer(&_Task.TransactOpts, new_owner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address new_owner) returns()
func (_Task *TaskTransactorSession) Transfer(new_owner common.Address) (*types.Transaction, error) {
	return _Task.Contract.Transfer(&_Task.TransactOpts, new_owner)
}

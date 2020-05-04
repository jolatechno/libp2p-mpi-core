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
var IdentityBin = "0x6080604052600080546001600160a01b0319163317905534801561002257600080fd5b50610144806100326000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631a6952301461003b5780638da5cb5b14610063575b600080fd5b6100616004803603602081101561005157600080fd5b50356001600160a01b0316610087565b005b61006b610100565b604080516001600160a01b039092168252519081900360200190f35b6000546001600160a01b031633146100de576040805162461bcd60e51b81526020600482015260156024820152741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03169056fea265627a7a72315820f142ffb3d45c621f3f0a57642fd84d811080795fc2176c4148c96f24276f169b64736f6c63430005100032"

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
const InterpreterABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"IpfsHash\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contracttask\",\"name\":\"Task\",\"type\":\"address\"}],\"name\":\"advertise\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contracttask\",\"name\":\"Task\",\"type\":\"address\"}],\"name\":\"done\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTask\",\"outputs\":[{\"internalType\":\"contracttask\",\"name\":\"Task\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_owner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// InterpreterFuncSigs maps the 4-byte function signature to its string representation.
var InterpreterFuncSigs = map[string]string{
	"4063d563": "advertise(address)",
	"7be8f86b": "done(address)",
	"a31662d2": "getTask()",
	"8da5cb5b": "owner()",
	"1a695230": "transfer(address)",
}

// InterpreterBin is the compiled bytecode used for deploying new contracts.
var InterpreterBin = "0x6080604052600080546001600160a01b0319163317905534801561002257600080fd5b5060405161076a38038061076a8339810160408190526100419161014a565b805161005490600190602084019061005b565b5050610203565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061009c57805160ff19168380011785556100c9565b828001600101855582156100c9579182015b828111156100c95782518255916020019190600101906100ae565b506100d59291506100d9565b5090565b6100f391905b808211156100d557600081556001016100df565b90565b600082601f83011261010757600080fd5b815161011a610115826101ac565b610186565b9150808252602083016020830185838301111561013657600080fd5b6101418382846101d3565b50505092915050565b60006020828403121561015c57600080fd5b81516001600160401b0381111561017257600080fd5b61017e848285016100f6565b949350505050565b6040518181016001600160401b03811182821017156101a457600080fd5b604052919050565b60006001600160401b038211156101c257600080fd5b506020601f91909101601f19160190565b60005b838110156101ee5781810151838201526020016101d6565b838111156101fd576000848401525b50505050565b610558806102126000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631a6952301461005c5780634063d563146100715780637be8f86b146100845780638da5cb5b14610097578063a31662d2146100b5575b600080fd5b61006f61006a36600461038e565b6100cb565b005b61006f61007f3660046103b4565b610120565b61006f6100923660046103b4565b6101a3565b61009f610266565b6040516100ac9190610470565b60405180910390f35b6100bd610275565b6040516100ac92919061047e565b6000546001600160a01b031633146100fe5760405162461bcd60e51b81526004016100f5906104a0565b60405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6003805460018181019092557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b039093166001600160a01b03199384161790556002805491820181556000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01805490911633179055565b80336001600160a01b038216146101cc5760405162461bcd60e51b81526004016100f5906104a0565b60005b60035481101561026157826001600160a01b0316600382815481106101f057fe5b6000918252602090912001546001600160a01b03161415610259576003818154811061021857fe5b600091825260209091200180546001600160a01b0319169055600280548290811061023f57fe5b600091825260209091200180546001600160a01b03191690555b6001016101cf565b505050565b6000546001600160a01b031690565b6003546000908190610289575060006102c4565b600354600090610298906102c8565b9050600381815481106102a757fe5b6000918252602090912001546001600160a01b0316925060019150505b9091565b60008042336040516020016102dd9190610446565b6040516020818303038152906040528051906020012060001c816102fd57fe5b0442416040516020016103109190610446565b6040516020818303038152906040528051906020012060001c8161033057fe5b044342010101604051602001610346919061045b565b6040516020818303038152906040528051906020012060001c905082818161036a57fe5b069392505050565b803561037d816104f5565b92915050565b803561037d8161050c565b6000602082840312156103a057600080fd5b60006103ac8484610372565b949350505050565b6000602082840312156103c657600080fd5b60006103ac8484610383565b6103e36103de826104b9565b6104e3565b82525050565b6103e3816104b9565b6103e3816104c4565b6103e3816104c9565b60006104116015836104b0565b741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b815260200192915050565b6103e3610441826104e0565b6104e0565b600061045282846103d2565b50601401919050565b60006104678284610435565b50602001919050565b6020810161037d82846103e9565b6040810161048c82856103fb565b61049960208301846103f2565b9392505050565b6020808252810161037d81610404565b90815260200190565b600061037d826104d4565b151590565b600061037d826104b9565b6001600160a01b031690565b90565b600061037d82600061037d8260601b90565b6104fe816104b9565b811461050957600080fd5b50565b6104fe816104c956fea365627a7a72315820e09a60241b4281c737f2d1e311b1307f80da101a8c41048a7b9b9662b75606b76c6578706572696d656e74616cf564736f6c63430005100040"

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
// Solidity: function getTask() view returns(address Task, bool)
func (_Interpreter *InterpreterCaller) GetTask(opts *bind.CallOpts) (common.Address, bool, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Interpreter.contract.Call(opts, out, "getTask")
	return *ret0, *ret1, err
}

// GetTask is a free data retrieval call binding the contract method 0xa31662d2.
//
// Solidity: function getTask() view returns(address Task, bool)
func (_Interpreter *InterpreterSession) GetTask() (common.Address, bool, error) {
	return _Interpreter.Contract.GetTask(&_Interpreter.CallOpts)
}

// GetTask is a free data retrieval call binding the contract method 0xa31662d2.
//
// Solidity: function getTask() view returns(address Task, bool)
func (_Interpreter *InterpreterCallerSession) GetTask() (common.Address, bool, error) {
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
var RandomBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a72315820a190532ab582958051c73ba3fa1fd3fa7782ea65f6a61acff95b8d7fd5038f5264736f6c63430005100032"

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

// StackABI is the input ABI used to generate the binding from.
const StackABI = "[{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"SafetyProportionTreshold\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"SafetyLengthTreshold\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractinterpreter\",\"name\":\"IpfsObject\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"Kernel_size\",\"type\":\"uint256[]\"}],\"name\":\"newTask\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"read\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_owner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StackFuncSigs maps the 4-byte function signature to its string representation.
var StackFuncSigs = map[string]string{
	"3ec53d59": "newTask(address,uint256[])",
	"8da5cb5b": "owner()",
	"619b40cc": "push(address,bytes,bytes)",
	"8bf4515c": "read(bytes)",
	"1a695230": "transfer(address)",
}

// StackBin is the compiled bytecode used for deploying new contracts.
var StackBin = "0x6080604052600080546001600160a01b0319163317905534801561002257600080fd5b50604051611e14380380611e14833981016040819052610041916100c7565b600480546001600160801b03928316600160801b029383166001600160801b031990911617909116919091179055600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b03191633179055610124565b80516100c18161010d565b92915050565b600080604083850312156100da57600080fd5b60006100e685856100b6565b92505060206100f7858286016100b6565b9150509250929050565b6001600160801b031690565b61011681610101565b811461012157600080fd5b50565b611ce1806101336000396000f3fe60806040523480156200001157600080fd5b50600436106200005e5760003560e01c80631a69523014620000635780633ec53d59146200007c578063619b40cc14620000935780638bf4515c14620000aa5780638da5cb5b14620000da575b600080fd5b6200007a620000743660046200098e565b620000f3565b005b6200007a6200008d36600462000a76565b6200014b565b6200007a620000a4366004620009b7565b62000271565b620000c1620000bb36600462000a3d565b62000474565b604051620000d192919062000c63565b60405180910390f35b620000e4620007c2565b604051620000d1919062000c43565b6000546001600160a01b03163314620001295760405162461bcd60e51b8152600401620001209062000cc3565b60405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03163314620001785760405162461bcd60e51b8152600401620001209062000cc3565b60003083836040516200018b90620007d2565b620001999392919062000c87565b604051809103906000f080158015620001b6573d6000803e3d6000fd5b506040516301a6952360e41b81529091506001600160a01b03821690631a69523090620001e890339060040162000c53565b600060405180830381600087803b1580156200020357600080fd5b505af115801562000218573d6000803e3d6000fd5b5050600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319166001600160a01b03949094169390931790925550505050565b60036000805b8254811015620002c757336001600160a01b03168382815481106200029857fe5b6000918252602090912001546001600160a01b03161415620002be5760019150620002c7565b60010162000277565b5080620002e85760405162461bcd60e51b8152600401620001209062000cc3565b620002f2620007c2565b6001600160a01b0316856001600160a01b031614620003d65760606002856040516200031f919062000c2e565b90815260408051918290036020908101832080548083028501830190935282845291908301828280156200037d57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116200035e575b50939450600093505050505b8151811015620003d357818181518110620003a057fe5b60200260200101516001600160a01b0316876001600160a01b03161415620003ca5750506200046d565b60010162000389565b50505b600284604051620003e8919062000c2e565b908152604051602091819003820181208054600180820183556000928352939091200180546001600160a01b0319166001600160a01b0389161790556200043190869062000c2e565b90815260405160209181900382019020805460018101808355600092835291839020865192936200046a939190920191870190620007e0565b50505b5050505050565b6060600060606001846040516200048c919062000c2e565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156200056b5760008481526020908190208301805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815292830182828015620005565780601f106200052a5761010080835404028352916020019162000556565b820191906000526020600020905b8154815290600101906020018083116200053857829003601f168201915b505050505081526020019060010190620004ba565b505050509050606060028560405162000585919062000c2e565b9081526040805191829003602090810183208054808302850183019093528284529190830182828015620005e357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311620005c4575b5050855160045494955093600160801b90046001600160801b0316841015925062000619915050575060009250620007bd915050565b60608160405190808252806020026020018201604052801562000646578160200160208202803883390190505b50905060005b82811015620007b3576200065f620007c2565b6001600160a01b03168482815181106200067557fe5b60200260200101516001600160a01b03161415620006b6578481815181106200069a57fe5b60200260200101516001819150965096505050505050620007bd565b60005b818111620007a957858181518110620006ce57fe5b6020026020010151604051620006e5919062000c2e565b6040518091039020868381518110620006fa57fe5b602002602001015160405162000711919062000c2e565b60405180910390201415620007a0578281815181106200072d57fe5b602090810291909101018051600101905260045483516001600160801b039182166000198101909216860291908590849081106200076757fe5b60200260200101510210620007a0578582815181106200078357fe5b6020026020010151600181915097509750505050505050620007bd565b600101620006b9565b506001016200064c565b5060009450505050505b915091565b6000546001600160a01b03165b90565b610e8d8062000e1283390190565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200082357805160ff191683800117855562000853565b8280016001018555821562000853579182015b828111156200085357825182559160200191906001019062000836565b506200086192915062000865565b5090565b620007cf91905b808211156200086157600081556001016200086c565b80356200088f8162000de1565b92915050565b600082601f830112620008a757600080fd5b8135620008be620008b88262000cfd565b62000cd5565b91508181835260208401935060208101905083856020840282011115620008e457600080fd5b60005b83811015620009145781620008fd888262000981565b8452506020928301929190910190600101620008e7565b5050505092915050565b600082601f8301126200093057600080fd5b813562000941620008b88262000d1f565b915080825260208301602083018583830111156200095e57600080fd5b6200096b83828462000d98565b50505092915050565b80356200088f8162000dfb565b80356200088f8162000e06565b600060208284031215620009a157600080fd5b6000620009af848462000882565b949350505050565b600080600060608486031215620009cd57600080fd5b6000620009db868662000882565b935050602084013567ffffffffffffffff811115620009f957600080fd5b62000a07868287016200091e565b925050604084013567ffffffffffffffff81111562000a2557600080fd5b62000a33868287016200091e565b9150509250925092565b60006020828403121562000a5057600080fd5b813567ffffffffffffffff81111562000a6857600080fd5b620009af848285016200091e565b6000806040838503121562000a8a57600080fd5b600062000a98858562000974565b925050602083013567ffffffffffffffff81111562000ab657600080fd5b62000ac48582860162000895565b9150509250929050565b600062000adc838362000c23565b505060200190565b62000aef8162000d8b565b82525050565b62000aef8162000d60565b600062000b0d8262000d4e565b62000b19818562000d52565b935062000b268362000d48565b8060005b8381101562000b5a57815162000b41888262000ace565b975062000b4e8362000d48565b92505060010162000b2a565b509495945050505050565b62000aef8162000d6d565b600062000b7d8262000d4e565b62000b89818562000d52565b935062000b9b81856020860162000da4565b62000ba68162000dd7565b9093019392505050565b600062000bbd8262000d4e565b62000bc9818562000d5b565b935062000bdb81856020860162000da4565b9290920192915050565b62000aef8162000d72565b600062000bff60158362000d52565b741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b815260200192915050565b62000aef81620007cf565b600062000c3c828462000bb0565b9392505050565b602081016200088f828462000af5565b602081016200088f828462000ae4565b6040808252810162000c76818562000b70565b905062000c3c602083018462000b65565b6060810162000c97828662000be5565b62000ca6602083018562000be5565b818103604083015262000cba818462000b00565b95945050505050565b602080825281016200088f8162000bf0565b60405181810167ffffffffffffffff8111828210171562000cf557600080fd5b604052919050565b600067ffffffffffffffff82111562000d1557600080fd5b5060209081020190565b600067ffffffffffffffff82111562000d3757600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b60006200088f8262000d7f565b151590565b60006200088f8262000d60565b6001600160a01b031690565b60006200088f8262000d72565b82818337506000910152565b60005b8381101562000dc157818101518382015260200162000da7565b8381111562000dd1576000848401525b50505050565b601f01601f191690565b62000dec8162000d60565b811462000df857600080fd5b50565b62000dec8162000d72565b62000dec81620007cf56fe6080604052600080546001600160a01b031916331790553480156200002357600080fd5b5060405162000e8d38038062000e8d833981016040819052620000469162000291565b80516200005b90600390602084019062000178565b50600280546001600160a01b038085166001600160a01b031992831617909255600180549286169290911691909117815560005b8251811015620000c057828181518110620000a657fe5b60200260200101518202915080806001019150506200008f565b506004819055604080518281526020808402820101909152818015620000f0578160200160208202803883390190505b508051620001079160059160209091019062000178565b50600254604051634063d56360e01b81526001600160a01b0390911690634063d563906200013a9030906004016200030e565b600060405180830381600087803b1580156200015557600080fd5b505af11580156200016a573d6000803e3d6000fd5b5050505050505050620003b1565b828054828255906000526020600020908101928215620001b6579160200282015b82811115620001b657825182559160200191906001019062000199565b50620001c4929150620001c8565b5090565b620001e591905b80821115620001c45760008155600101620001cf565b90565b600082601f830112620001fa57600080fd5b8151620002116200020b8262000345565b6200031e565b915081818352602084019350602081019050838560208402820111156200023757600080fd5b60005b8381101562000267578162000250888262000284565b84525060209283019291909101906001016200023a565b5050505092915050565b80516200027e816200038c565b92915050565b80516200027e81620003a6565b600080600060608486031215620002a757600080fd5b6000620002b5868662000271565b9350506020620002c88682870162000271565b92505060408401516001600160401b03811115620002e557600080fd5b620002f386828701620001e8565b9150509250925092565b620003088162000373565b82525050565b602081016200027e8284620002fd565b6040518181016001600160401b03811182821017156200033d57600080fd5b604052919050565b60006001600160401b038211156200035c57600080fd5b5060209081020190565b60006200027e8262000380565b60006200027e8262000366565b6001600160a01b031690565b620003978162000373565b8114620003a357600080fd5b50565b6200039781620001e5565b610acc80620003c16000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638bf4515c1161005b5780638bf4515c146100c95780638da5cb5b146100ea578063ae8421e1146100ff578063e771ee0d146101075761007d565b80630110da29146100825780631a695230146100a157806329f5e410146100b6575b600080fd5b61008a61011a565b6040516100989291906108d8565b60405180910390f35b6100b46100af3660046105f8565b6101b0565b005b6100b46100c436600461061e565b610205565b6100dc6100d7366004610653565b610279565b604051610098929190610915565b6100f261030c565b604051610098919061088d565b6100b461031b565b6100b46101153660046106d9565b6103a9565b60608060038054905060405190808252806020026020018201604052801561014c578160200160208202803883390190505b50915060005b60035481101561019b5761017c6003828154811061016c57fe5b9060005260206000200154610413565b83828151811061018857fe5b6020908102919091010152600101610152565b50506040805160208101909152600081529091565b6000546001600160a01b031633146101e35760405162461bcd60e51b81526004016101da90610943565b60405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b60006001815b83518110156102545783818151811061022057fe5b60200260200101518202830192506003818154811061023b57fe5b600091825260209091200154919091029060010161020b565b506005828154811061026257fe5b600091825260209091200180546001019055505050565b6001546040516322fd145760e21b81526060916000916001600160a01b0390911690638bf4515c906102af9086906004016108fd565b60006040518083038186803b1580156102c757600080fd5b505afa1580156102db573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526103039190810190610688565b91509150915091565b6000546001600160a01b031690565b6000546001600160a01b031633146103455760405162461bcd60e51b81526004016101da90610943565b600254604051637be8f86b60e01b81526001600160a01b0390911690637be8f86b90610375903090600401610935565b600060405180830381600087803b15801561038f57600080fd5b505af11580156103a3573d6000803e3d6000fd5b50505050565b600154604051631866d03360e21b81526001600160a01b039091169063619b40cc906103dd9033908690869060040161089b565b600060405180830381600087803b1580156103f757600080fd5b505af115801561040b573d6000803e3d6000fd5b505050505050565b60008042336040516020016104289190610863565b6040516020818303038152906040528051906020012060001c8161044857fe5b04424160405160200161045b9190610863565b6040516020818303038152906040528051906020012060001c8161047b57fe5b0443420101016040516020016104919190610878565b6040516020818303038152906040528051906020012060001c90508281816104b557fe5b069392505050565b80356104c881610a60565b92915050565b600082601f8301126104df57600080fd5b81356104f26104ed8261097a565b610953565b9150818183526020840193506020810190508385602084028201111561051757600080fd5b60005b83811015610543578161052d88826105ed565b845250602092830192919091019060010161051a565b5050505092915050565b80516104c881610a77565b600082601f83011261056957600080fd5b81356105776104ed8261099b565b9150808252602083016020830185838301111561059357600080fd5b61059e838284610a07565b50505092915050565b600082601f8301126105b857600080fd5b81516105c66104ed8261099b565b915080825260208301602083018583830111156105e257600080fd5b61059e838284610a13565b80356104c881610a80565b60006020828403121561060a57600080fd5b600061061684846104bd565b949350505050565b60006020828403121561063057600080fd5b813567ffffffffffffffff81111561064757600080fd5b610616848285016104ce565b60006020828403121561066557600080fd5b813567ffffffffffffffff81111561067c57600080fd5b61061684828501610558565b6000806040838503121561069b57600080fd5b825167ffffffffffffffff8111156106b257600080fd5b6106be858286016105a7565b92505060206106cf8582860161054d565b9150509250929050565b600080604083850312156106ec57600080fd5b823567ffffffffffffffff81111561070357600080fd5b61070f85828601610558565b925050602083013567ffffffffffffffff81111561072c57600080fd5b6106cf85828601610558565b60006107448383610849565b505060200190565b610755816109f5565b82525050565b610755610767826109d6565b610a3f565b610755816109d6565b6000610780826109c9565b61078a81856109cd565b9350610795836109c3565b8060005b838110156107c35781516107ad8882610738565b97506107b8836109c3565b925050600101610799565b509495945050505050565b610755816109e1565b60006107e2826109c9565b6107ec81856109cd565b93506107fc818560208601610a13565b61080581610a50565b9093019392505050565b610755816109fc565b60006108256015836109cd565b741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b815260200192915050565b610755816109f2565b61075561085e826109f2565b6109f2565b600061086f828461075b565b50601401919050565b60006108848284610852565b50602001919050565b602081016104c8828461076c565b606081016108a9828661074c565b81810360208301526108bb81856107d7565b905081810360408301526108cf81846107d7565b95945050505050565b604080825281016108e98185610775565b9050818103602083015261061681846107d7565b6020808252810161090e81846107d7565b9392505050565b6040808252810161092681856107d7565b905061090e60208301846107ce565b602081016104c8828461080f565b602080825281016104c881610818565b60405181810167ffffffffffffffff8111828210171561097257600080fd5b604052919050565b600067ffffffffffffffff82111561099157600080fd5b5060209081020190565b600067ffffffffffffffff8211156109b257600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b60006104c8826109e6565b151590565b6001600160a01b031690565b90565b60006104c8825b60006104c8826109d6565b82818337506000910152565b60005b83811015610a2e578181015183820152602001610a16565b838111156103a35750506000910152565b60006104c88260006104c882610a5a565b601f01601f191690565b60601b90565b610a69816109d6565b8114610a7457600080fd5b50565b610a69816109e1565b610a69816109f256fea365627a7a7231582087a02bccb17dd7be4bd57c345da1de1741ab02b54febafa3667de6d3547d4a576c6578706572696d656e74616cf564736f6c63430005100040a365627a7a723158203ed9b943c84a657b025eb37885d71e52b9d959f50d02e57ab66ab3ae38f913586c6578706572696d656e74616cf564736f6c63430005100040"

// DeployStack deploys a new Ethereum contract, binding an instance of Stack to it.
func DeployStack(auth *bind.TransactOpts, backend bind.ContractBackend, SafetyProportionTreshold *big.Int, SafetyLengthTreshold *big.Int) (common.Address, *types.Transaction, *Stack, error) {
	parsed, err := abi.JSON(strings.NewReader(StackABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StackBin), backend, SafetyProportionTreshold, SafetyLengthTreshold)
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
// Solidity: function read(bytes key) view returns(bytes value, bool)
func (_Stack *StackCaller) Read(opts *bind.CallOpts, key []byte) ([]byte, bool, error) {
	var (
		ret0 = new([]byte)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Stack.contract.Call(opts, out, "read", key)
	return *ret0, *ret1, err
}

// Read is a free data retrieval call binding the contract method 0x8bf4515c.
//
// Solidity: function read(bytes key) view returns(bytes value, bool)
func (_Stack *StackSession) Read(key []byte) ([]byte, bool, error) {
	return _Stack.Contract.Read(&_Stack.CallOpts, key)
}

// Read is a free data retrieval call binding the contract method 0x8bf4515c.
//
// Solidity: function read(bytes key) view returns(bytes value, bool)
func (_Stack *StackCallerSession) Read(key []byte) ([]byte, bool, error) {
	return _Stack.Contract.Read(&_Stack.CallOpts, key)
}

// NewTask is a paid mutator transaction binding the contract method 0x3ec53d59.
//
// Solidity: function newTask(address IpfsObject, uint256[] Kernel_size) returns()
func (_Stack *StackTransactor) NewTask(opts *bind.TransactOpts, IpfsObject common.Address, Kernel_size []*big.Int) (*types.Transaction, error) {
	return _Stack.contract.Transact(opts, "newTask", IpfsObject, Kernel_size)
}

// NewTask is a paid mutator transaction binding the contract method 0x3ec53d59.
//
// Solidity: function newTask(address IpfsObject, uint256[] Kernel_size) returns()
func (_Stack *StackSession) NewTask(IpfsObject common.Address, Kernel_size []*big.Int) (*types.Transaction, error) {
	return _Stack.Contract.NewTask(&_Stack.TransactOpts, IpfsObject, Kernel_size)
}

// NewTask is a paid mutator transaction binding the contract method 0x3ec53d59.
//
// Solidity: function newTask(address IpfsObject, uint256[] Kernel_size) returns()
func (_Stack *StackTransactorSession) NewTask(IpfsObject common.Address, Kernel_size []*big.Int) (*types.Transaction, error) {
	return _Stack.Contract.NewTask(&_Stack.TransactOpts, IpfsObject, Kernel_size)
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

// TaskABI is the input ABI used to generate the binding from.
const TaskABI = "[{\"inputs\":[{\"internalType\":\"contractstack\",\"name\":\"Stack\",\"type\":\"address\"},{\"internalType\":\"contractinterpreter\",\"name\":\"IpfsObject\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"Kernel_size\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"kernel_idxs\",\"type\":\"uint256[]\"}],\"name\":\"acceptCommand\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"done\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCommand\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"kernel_idxs\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"read\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_owner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TaskFuncSigs maps the 4-byte function signature to its string representation.
var TaskFuncSigs = map[string]string{
	"29f5e410": "acceptCommand(uint256[])",
	"ae8421e1": "done()",
	"0110da29": "getCommand()",
	"8da5cb5b": "owner()",
	"e771ee0d": "push(bytes,bytes)",
	"8bf4515c": "read(bytes)",
	"1a695230": "transfer(address)",
}

// TaskBin is the compiled bytecode used for deploying new contracts.
var TaskBin = "0x6080604052600080546001600160a01b031916331790553480156200002357600080fd5b5060405162000e8d38038062000e8d833981016040819052620000469162000291565b80516200005b90600390602084019062000178565b50600280546001600160a01b038085166001600160a01b031992831617909255600180549286169290911691909117815560005b8251811015620000c057828181518110620000a657fe5b60200260200101518202915080806001019150506200008f565b506004819055604080518281526020808402820101909152818015620000f0578160200160208202803883390190505b508051620001079160059160209091019062000178565b50600254604051634063d56360e01b81526001600160a01b0390911690634063d563906200013a9030906004016200030e565b600060405180830381600087803b1580156200015557600080fd5b505af11580156200016a573d6000803e3d6000fd5b5050505050505050620003b1565b828054828255906000526020600020908101928215620001b6579160200282015b82811115620001b657825182559160200191906001019062000199565b50620001c4929150620001c8565b5090565b620001e591905b80821115620001c45760008155600101620001cf565b90565b600082601f830112620001fa57600080fd5b8151620002116200020b8262000345565b6200031e565b915081818352602084019350602081019050838560208402820111156200023757600080fd5b60005b8381101562000267578162000250888262000284565b84525060209283019291909101906001016200023a565b5050505092915050565b80516200027e816200038c565b92915050565b80516200027e81620003a6565b600080600060608486031215620002a757600080fd5b6000620002b5868662000271565b9350506020620002c88682870162000271565b92505060408401516001600160401b03811115620002e557600080fd5b620002f386828701620001e8565b9150509250925092565b620003088162000373565b82525050565b602081016200027e8284620002fd565b6040518181016001600160401b03811182821017156200033d57600080fd5b604052919050565b60006001600160401b038211156200035c57600080fd5b5060209081020190565b60006200027e8262000380565b60006200027e8262000366565b6001600160a01b031690565b620003978162000373565b8114620003a357600080fd5b50565b6200039781620001e5565b610acc80620003c16000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638bf4515c1161005b5780638bf4515c146100c95780638da5cb5b146100ea578063ae8421e1146100ff578063e771ee0d146101075761007d565b80630110da29146100825780631a695230146100a157806329f5e410146100b6575b600080fd5b61008a61011a565b6040516100989291906108d8565b60405180910390f35b6100b46100af3660046105f8565b6101b0565b005b6100b46100c436600461061e565b610205565b6100dc6100d7366004610653565b610279565b604051610098929190610915565b6100f261030c565b604051610098919061088d565b6100b461031b565b6100b46101153660046106d9565b6103a9565b60608060038054905060405190808252806020026020018201604052801561014c578160200160208202803883390190505b50915060005b60035481101561019b5761017c6003828154811061016c57fe5b9060005260206000200154610413565b83828151811061018857fe5b6020908102919091010152600101610152565b50506040805160208101909152600081529091565b6000546001600160a01b031633146101e35760405162461bcd60e51b81526004016101da90610943565b60405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b60006001815b83518110156102545783818151811061022057fe5b60200260200101518202830192506003818154811061023b57fe5b600091825260209091200154919091029060010161020b565b506005828154811061026257fe5b600091825260209091200180546001019055505050565b6001546040516322fd145760e21b81526060916000916001600160a01b0390911690638bf4515c906102af9086906004016108fd565b60006040518083038186803b1580156102c757600080fd5b505afa1580156102db573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526103039190810190610688565b91509150915091565b6000546001600160a01b031690565b6000546001600160a01b031633146103455760405162461bcd60e51b81526004016101da90610943565b600254604051637be8f86b60e01b81526001600160a01b0390911690637be8f86b90610375903090600401610935565b600060405180830381600087803b15801561038f57600080fd5b505af11580156103a3573d6000803e3d6000fd5b50505050565b600154604051631866d03360e21b81526001600160a01b039091169063619b40cc906103dd9033908690869060040161089b565b600060405180830381600087803b1580156103f757600080fd5b505af115801561040b573d6000803e3d6000fd5b505050505050565b60008042336040516020016104289190610863565b6040516020818303038152906040528051906020012060001c8161044857fe5b04424160405160200161045b9190610863565b6040516020818303038152906040528051906020012060001c8161047b57fe5b0443420101016040516020016104919190610878565b6040516020818303038152906040528051906020012060001c90508281816104b557fe5b069392505050565b80356104c881610a60565b92915050565b600082601f8301126104df57600080fd5b81356104f26104ed8261097a565b610953565b9150818183526020840193506020810190508385602084028201111561051757600080fd5b60005b83811015610543578161052d88826105ed565b845250602092830192919091019060010161051a565b5050505092915050565b80516104c881610a77565b600082601f83011261056957600080fd5b81356105776104ed8261099b565b9150808252602083016020830185838301111561059357600080fd5b61059e838284610a07565b50505092915050565b600082601f8301126105b857600080fd5b81516105c66104ed8261099b565b915080825260208301602083018583830111156105e257600080fd5b61059e838284610a13565b80356104c881610a80565b60006020828403121561060a57600080fd5b600061061684846104bd565b949350505050565b60006020828403121561063057600080fd5b813567ffffffffffffffff81111561064757600080fd5b610616848285016104ce565b60006020828403121561066557600080fd5b813567ffffffffffffffff81111561067c57600080fd5b61061684828501610558565b6000806040838503121561069b57600080fd5b825167ffffffffffffffff8111156106b257600080fd5b6106be858286016105a7565b92505060206106cf8582860161054d565b9150509250929050565b600080604083850312156106ec57600080fd5b823567ffffffffffffffff81111561070357600080fd5b61070f85828601610558565b925050602083013567ffffffffffffffff81111561072c57600080fd5b6106cf85828601610558565b60006107448383610849565b505060200190565b610755816109f5565b82525050565b610755610767826109d6565b610a3f565b610755816109d6565b6000610780826109c9565b61078a81856109cd565b9350610795836109c3565b8060005b838110156107c35781516107ad8882610738565b97506107b8836109c3565b925050600101610799565b509495945050505050565b610755816109e1565b60006107e2826109c9565b6107ec81856109cd565b93506107fc818560208601610a13565b61080581610a50565b9093019392505050565b610755816109fc565b60006108256015836109cd565b741cd95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b815260200192915050565b610755816109f2565b61075561085e826109f2565b6109f2565b600061086f828461075b565b50601401919050565b60006108848284610852565b50602001919050565b602081016104c8828461076c565b606081016108a9828661074c565b81810360208301526108bb81856107d7565b905081810360408301526108cf81846107d7565b95945050505050565b604080825281016108e98185610775565b9050818103602083015261061681846107d7565b6020808252810161090e81846107d7565b9392505050565b6040808252810161092681856107d7565b905061090e60208301846107ce565b602081016104c8828461080f565b602080825281016104c881610818565b60405181810167ffffffffffffffff8111828210171561097257600080fd5b604052919050565b600067ffffffffffffffff82111561099157600080fd5b5060209081020190565b600067ffffffffffffffff8211156109b257600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b60006104c8826109e6565b151590565b6001600160a01b031690565b90565b60006104c8825b60006104c8826109d6565b82818337506000910152565b60005b83811015610a2e578181015183820152602001610a16565b838111156103a35750506000910152565b60006104c88260006104c882610a5a565b601f01601f191690565b60601b90565b610a69816109d6565b8114610a7457600080fd5b50565b610a69816109e1565b610a69816109f256fea365627a7a7231582087a02bccb17dd7be4bd57c345da1de1741ab02b54febafa3667de6d3547d4a576c6578706572696d656e74616cf564736f6c63430005100040"

// DeployTask deploys a new Ethereum contract, binding an instance of Task to it.
func DeployTask(auth *bind.TransactOpts, backend bind.ContractBackend, Stack common.Address, IpfsObject common.Address, Kernel_size []*big.Int) (common.Address, *types.Transaction, *Task, error) {
	parsed, err := abi.JSON(strings.NewReader(TaskABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TaskBin), backend, Stack, IpfsObject, Kernel_size)
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
// Solidity: function getCommand() view returns(uint256[] kernel_idxs, string error)
func (_Task *TaskCaller) GetCommand(opts *bind.CallOpts) (struct {
	KernelIdxs []*big.Int
	Error      string
}, error) {
	ret := new(struct {
		KernelIdxs []*big.Int
		Error      string
	})
	out := ret
	err := _Task.contract.Call(opts, out, "getCommand")
	return *ret, err
}

// GetCommand is a free data retrieval call binding the contract method 0x0110da29.
//
// Solidity: function getCommand() view returns(uint256[] kernel_idxs, string error)
func (_Task *TaskSession) GetCommand() (struct {
	KernelIdxs []*big.Int
	Error      string
}, error) {
	return _Task.Contract.GetCommand(&_Task.CallOpts)
}

// GetCommand is a free data retrieval call binding the contract method 0x0110da29.
//
// Solidity: function getCommand() view returns(uint256[] kernel_idxs, string error)
func (_Task *TaskCallerSession) GetCommand() (struct {
	KernelIdxs []*big.Int
	Error      string
}, error) {
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

// Read is a free data retrieval call binding the contract method 0x8bf4515c.
//
// Solidity: function read(bytes key) view returns(bytes value, bool)
func (_Task *TaskCaller) Read(opts *bind.CallOpts, key []byte) ([]byte, bool, error) {
	var (
		ret0 = new([]byte)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Task.contract.Call(opts, out, "read", key)
	return *ret0, *ret1, err
}

// Read is a free data retrieval call binding the contract method 0x8bf4515c.
//
// Solidity: function read(bytes key) view returns(bytes value, bool)
func (_Task *TaskSession) Read(key []byte) ([]byte, bool, error) {
	return _Task.Contract.Read(&_Task.CallOpts, key)
}

// Read is a free data retrieval call binding the contract method 0x8bf4515c.
//
// Solidity: function read(bytes key) view returns(bytes value, bool)
func (_Task *TaskCallerSession) Read(key []byte) ([]byte, bool, error) {
	return _Task.Contract.Read(&_Task.CallOpts, key)
}

// AcceptCommand is a paid mutator transaction binding the contract method 0x29f5e410.
//
// Solidity: function acceptCommand(uint256[] kernel_idxs) returns()
func (_Task *TaskTransactor) AcceptCommand(opts *bind.TransactOpts, kernel_idxs []*big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "acceptCommand", kernel_idxs)
}

// AcceptCommand is a paid mutator transaction binding the contract method 0x29f5e410.
//
// Solidity: function acceptCommand(uint256[] kernel_idxs) returns()
func (_Task *TaskSession) AcceptCommand(kernel_idxs []*big.Int) (*types.Transaction, error) {
	return _Task.Contract.AcceptCommand(&_Task.TransactOpts, kernel_idxs)
}

// AcceptCommand is a paid mutator transaction binding the contract method 0x29f5e410.
//
// Solidity: function acceptCommand(uint256[] kernel_idxs) returns()
func (_Task *TaskTransactorSession) AcceptCommand(kernel_idxs []*big.Int) (*types.Transaction, error) {
	return _Task.Contract.AcceptCommand(&_Task.TransactOpts, kernel_idxs)
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

// Push is a paid mutator transaction binding the contract method 0xe771ee0d.
//
// Solidity: function push(bytes key, bytes value) returns()
func (_Task *TaskTransactor) Push(opts *bind.TransactOpts, key []byte, value []byte) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "push", key, value)
}

// Push is a paid mutator transaction binding the contract method 0xe771ee0d.
//
// Solidity: function push(bytes key, bytes value) returns()
func (_Task *TaskSession) Push(key []byte, value []byte) (*types.Transaction, error) {
	return _Task.Contract.Push(&_Task.TransactOpts, key, value)
}

// Push is a paid mutator transaction binding the contract method 0xe771ee0d.
//
// Solidity: function push(bytes key, bytes value) returns()
func (_Task *TaskTransactorSession) Push(key []byte, value []byte) (*types.Transaction, error) {
	return _Task.Contract.Push(&_Task.TransactOpts, key, value)
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

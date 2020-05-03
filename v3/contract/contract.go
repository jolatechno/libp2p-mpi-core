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

// InterpreterABI is the input ABI used to generate the binding from.
const InterpreterABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"IpfsHash\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contracttask\",\"name\":\"Task\",\"type\":\"address\"}],\"name\":\"advertise\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contracttask\",\"name\":\"Task\",\"type\":\"address\"}],\"name\":\"done\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTask\",\"outputs\":[{\"internalType\":\"contracttask\",\"name\":\"Task\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// InterpreterFuncSigs maps the 4-byte function signature to its string representation.
var InterpreterFuncSigs = map[string]string{
	"4063d563": "advertise(address)",
	"7be8f86b": "done(address)",
	"a31662d2": "getTask()",
}

// InterpreterBin is the compiled bytecode used for deploying new contracts.
var InterpreterBin = "0x608060405234801561001057600080fd5b5060405161074838038061074883398101604081905261002f91610138565b8051610042906000906020840190610049565b50506101f1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061008a57805160ff19168380011785556100b7565b828001600101855582156100b7579182015b828111156100b757825182559160200191906001019061009c565b506100c39291506100c7565b5090565b6100e191905b808211156100c357600081556001016100cd565b90565b600082601f8301126100f557600080fd5b81516101086101038261019a565b610174565b9150808252602083016020830185838301111561012457600080fd5b61012f8382846101c1565b50505092915050565b60006020828403121561014a57600080fd5b81516001600160401b0381111561016057600080fd5b61016c848285016100e4565b949350505050565b6040518181016001600160401b038111828210171561019257600080fd5b604052919050565b60006001600160401b038211156101b057600080fd5b506020601f91909101601f19160190565b60005b838110156101dc5781810151838201526020016101c4565b838111156101eb576000848401525b50505050565b610548806102006000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80634063d563146100465780637be8f86b1461005b578063a31662d21461006e575b600080fd5b610059610054366004610409565b61008d565b005b610059610069366004610409565b610161565b610076610263565b604051610084929190610493565b60405180910390f35b600480546001818101909255602081047f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01805460ff601f9093166101000a92909202199091169055600380548083019091557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b039093166001600160a01b03199384161790556002805480830182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180549092163317909155805481019055565b60005b60035481101561025e57816001600160a01b03166003828154811061018557fe5b6000918252602090912001546001600160a01b0316141561025657336001600160a01b0316600282815481106101b757fe5b6000918252602090912001546001600160a01b03161480156102035750600481815481106101e157fe5b90600052602060002090602091828204019190069054906101000a900460ff16155b156102505760016004828154811061021757fe5b60009182526020918290209181049091018054921515601f9092166101000a91820260ff90920219909216179055600180546000190190555b50610260565b600101610164565b505b50565b600080600154600014156102795750600061034a565b600061028660015461034e565b905060005b60035481101561034357811580156102cd5750600481815481106102ab57fe5b90600052602060002090602091828204019190069054906101000a900460ff16155b1561030257600381815481106102df57fe5b6000918252602090912001546001600160a01b031693506001925061034a915050565b6004818154811061030f57fe5b90600052602060002090602091828204019190069054906101000a900460ff1661033b57600019909101905b60010161028b565b5060009150505b9091565b60008042336040516020016103639190610469565b6040516020818303038152906040528051906020012060001c8161038357fe5b0442416040516020016103969190610469565b6040516020818303038152906040528051906020012060001c816103b657fe5b0443420101016040516020016103cc919061047e565b6040516020818303038152906040528051906020012060001c90508281816103f057fe5b069392505050565b8035610403816104f1565b92915050565b60006020828403121561041b57600080fd5b600061042784846103f8565b949350505050565b61044061043b826104b5565b6104df565b82525050565b610440816104c0565b610440816104c5565b610440610464826104dc565b6104dc565b6000610475828461042f565b50601401919050565b600061048a8284610458565b50602001919050565b604081016104a1828561044f565b6104ae6020830184610446565b9392505050565b6000610403826104d0565b151590565b6000610403826104b5565b6001600160a01b031690565b90565b60006104038260006104038260601b90565b6104fa816104c5565b811461026057600080fdfea365627a7a723158208e9ff1c17cee4d113d22576c154b1f6a63c7d40a8824e955c8e84c66ab116ca06c6578706572696d656e74616cf564736f6c63430005100040"

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

// RandomABI is the input ABI used to generate the binding from.
const RandomABI = "[]"

// RandomBin is the compiled bytecode used for deploying new contracts.
var RandomBin = "0x6080604052348015600f57600080fd5b50604c80601d6000396000f3fe6080604052600080fdfea365627a7a7231582041f9e0432d97d72bd6e47d6f8cc05b9056155f068d06c525053ee63172da6d296c6578706572696d656e74616cf564736f6c63430005100040"

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
const StackABI = "[{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"SafetyProportionTreshold\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"SafetyLengthTreshold\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"read\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StackFuncSigs maps the 4-byte function signature to its string representation.
var StackFuncSigs = map[string]string{
	"619b40cc": "push(address,bytes,bytes)",
	"8bf4515c": "read(bytes)",
}

// StackBin is the compiled bytecode used for deploying new contracts.
var StackBin = "0x608060405234801561001057600080fd5b506040516109e43803806109e483398101604081905261002f91610085565b600080546001600160a01b03191633179055600380546001600160801b03928316600160801b029383166001600160801b0319909116179091169190911790556100e2565b805161007f816100cb565b92915050565b6000806040838503121561009857600080fd5b60006100a48585610074565b92505060206100b585828601610074565b9150509250929050565b6001600160801b031690565b6100d4816100bf565b81146100df57600080fd5b50565b6108f3806100f16000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063619b40cc1461003b5780638bf4515c14610050575b600080fd5b61004e61004936600461063c565b61007a565b005b61006361005e3660046106b9565b610213565b6040516100719291906107a6565b60405180910390f35b6000546001600160a01b031633146100ad5760405162461bcd60e51b81526004016100a4906107c6565b60405180910390fd5b6000546001600160a01b038481169116146101845760606002836040516100d49190610793565b908152604080519182900360209081018320805480830285018301909352828452919083018282801561013057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610112575b50939450600093505050505b81518110156101815781818151811061015157fe5b60200260200101516001600160a01b0316336001600160a01b0316141561017957505061020e565b60010161013c565b50505b6002826040516101949190610793565b90815260405190819003602090810182208054600181810183556000928352929091200180546001600160a01b03191633179055906101d4908490610793565b908152604051602091819003820190208054600181018083556000928352918390208451929361020b93919092019185019061053c565b50505b505050565b6060600060606001846040516102299190610793565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156103025760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156102ee5780601f106102c3576101008083540402835291602001916102ee565b820191906000526020600020905b8154815290600101906020018083116102d157829003601f168201915b505050505081526020019060010190610257565b505050509050606060028560405161031a9190610793565b908152604080519182900360209081018320805480830285018301909352828452919083018282801561037657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610358575b5050855160035494955093600160801b90046001600160801b031684101592506103aa915050575060009250610537915050565b6060816040519080825280602002602001820160405280156103d6578160200160208202803883390190505b50905060005b8281101561052d5760005484516001600160a01b039091169085908390811061040157fe5b60200260200101516001600160a01b0316141561043f5784818151811061042457fe5b60200260200101516001819150965096505050505050610537565b60005b8181116105245785818151811061045557fe5b602002602001015160405161046a9190610793565b604051809103902086838151811061047e57fe5b60200260200101516040516104939190610793565b6040518091039020141561051c578281815181106104ad57fe5b602090810291909101018051600101905260035483516001600160801b039182166000198101909216860291908590849081106104e657fe5b6020026020010151021061051c5785828151811061050057fe5b6020026020010151600181915097509750505050505050610537565b600101610442565b506001016103dc565b5060009450505050505b915091565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061057d57805160ff19168380011785556105aa565b828001600101855582156105aa579182015b828111156105aa57825182559160200191906001019061058f565b506105b69291506105ba565b5090565b6105d491905b808211156105b657600081556001016105c0565b90565b80356105e281610899565b92915050565b600082601f8301126105f957600080fd5b813561060c610607826107fd565b6107d6565b9150808252602083016020830185838301111561062857600080fd5b610633838284610853565b50505092915050565b60008060006060848603121561065157600080fd5b600061065d86866105d7565b935050602084013567ffffffffffffffff81111561067a57600080fd5b610686868287016105e8565b925050604084013567ffffffffffffffff8111156106a357600080fd5b6106af868287016105e8565b9150509250925092565b6000602082840312156106cb57600080fd5b813567ffffffffffffffff8111156106e257600080fd5b6106ee848285016105e8565b949350505050565b6106ff81610842565b82525050565b600061071082610825565b61071a8185610829565b935061072a81856020860161085f565b6107338161088f565b9093019392505050565b600061074882610825565b6107528185610832565b935061076281856020860161085f565b9290920192915050565b6000610779600b83610829565b6a6e6f74206f776e6572202160a81b815260200192915050565b600061079f828461073d565b9392505050565b604080825281016107b78185610705565b905061079f60208301846106f6565b602080825281016105e28161076c565b60405181810167ffffffffffffffff811182821017156107f557600080fd5b604052919050565b600067ffffffffffffffff82111561081457600080fd5b506020601f91909101601f19160190565b5190565b90815260200190565b919050565b60006105e282610847565b151590565b6001600160a01b031690565b82818337506000910152565b60005b8381101561087a578181015183820152602001610862565b83811115610889576000848401525b50505050565b601f01601f191690565b6108a281610837565b81146108ad57600080fd5b5056fea365627a7a7231582064f0ffa66ed62cc655103fbe5339c417ff73d5220f08d39378eb36114472c8186c6578706572696d656e74616cf564736f6c63430005100040"

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

// TaskABI is the input ABI used to generate the binding from.
const TaskABI = "[{\"inputs\":[{\"internalType\":\"contractstack\",\"name\":\"Stack\",\"type\":\"address\"},{\"internalType\":\"contractinterpreter\",\"name\":\"IpfsObject\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"Kernel_size\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"Push\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"kernel_idxs\",\"type\":\"uint256[]\"}],\"name\":\"acceptCommand\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"done\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCommand\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"kernel_idxs\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"read\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TaskFuncSigs maps the 4-byte function signature to its string representation.
var TaskFuncSigs = map[string]string{
	"29f5e410": "acceptCommand(uint256[])",
	"ae8421e1": "done()",
	"0110da29": "getCommand()",
	"e771ee0d": "push(bytes,bytes)",
	"8bf4515c": "read(bytes)",
}

// TaskBin is the compiled bytecode used for deploying new contracts.
var TaskBin = "0x60806040523480156200001157600080fd5b5060405162000d5c38038062000d5c833981016040819052620000349162000287565b8051620000499060039060208401906200016e565b5060008054336001600160a01b03199182161782556002805482166001600160a01b038681169190911790915560018054909216908616178155905b8251811015620000b6578281815181106200009c57fe5b602002602001015182029150808060010191505062000085565b506004819055604080518281526020808402820101909152818015620000e6578160200160208202803883390190505b508051620000fd916005916020909101906200016e565b50600254604051634063d56360e01b81526001600160a01b0390911690634063d563906200013090309060040162000304565b600060405180830381600087803b1580156200014b57600080fd5b505af115801562000160573d6000803e3d6000fd5b5050505050505050620003a7565b828054828255906000526020600020908101928215620001ac579160200282015b82811115620001ac5782518255916020019190600101906200018f565b50620001ba929150620001be565b5090565b620001db91905b80821115620001ba5760008155600101620001c5565b90565b600082601f830112620001f057600080fd5b81516200020762000201826200033b565b62000314565b915081818352602084019350602081019050838560208402820111156200022d57600080fd5b60005b838110156200025d57816200024688826200027a565b845250602092830192919091019060010162000230565b5050505092915050565b8051620002748162000382565b92915050565b805162000274816200039c565b6000806000606084860312156200029d57600080fd5b6000620002ab868662000267565b9350506020620002be8682870162000267565b92505060408401516001600160401b03811115620002db57600080fd5b620002e986828701620001de565b9150509250925092565b620002fe8162000369565b82525050565b60208101620002748284620002f3565b6040518181016001600160401b03811182821017156200033357600080fd5b604052919050565b60006001600160401b038211156200035257600080fd5b5060209081020190565b6000620002748262000376565b600062000274826200035c565b6001600160a01b031690565b6200038d8162000369565b81146200039957600080fd5b50565b6200038d81620001db565b6109a580620003b76000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630110da291461005c57806329f5e4101461007b5780638bf4515c14610090578063ae8421e1146100b1578063e771ee0d146100b9575b600080fd5b6100646100cc565b6040516100729291906107ba565b60405180910390f35b61008e610089366004610519565b610162565b005b6100a361009e366004610556565b6101d6565b6040516100729291906107f7565b61008e610269565b61008e6100c73660046105dc565b610300565b6060806003805490506040519080825280602002602001820160405280156100fe578160200160208202803883390190505b50915060005b60035481101561014d5761012e6003828154811061011e57fe5b906000526020600020015461033f565b83828151811061013a57fe5b6020908102919091010152600101610104565b50506040805160208101909152600081529091565b60006001815b83518110156101b15783818151811061017d57fe5b60200260200101518202830192506003818154811061019857fe5b6000918252602090912001549190910290600101610168565b50600582815481106101bf57fe5b600091825260209091200180546001019055505050565b6001546040516322fd145760e21b81526060916000916001600160a01b0390911690638bf4515c9061020c9086906004016107df565b60006040518083038186803b15801561022457600080fd5b505afa158015610238573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610260919081019061058b565b91509150915091565b6000546001600160a01b0316331461029c5760405162461bcd60e51b815260040161029390610825565b60405180910390fd5b600254604051637be8f86b60e01b81526001600160a01b0390911690637be8f86b906102cc903090600401610817565b600060405180830381600087803b1580156102e657600080fd5b505af11580156102fa573d6000803e3d6000fd5b50505050565b7f0381ce894b22a3a2930bf6352fe67aa0d5af69d769ac34b044c53189c507a3a13383836040516103339392919061077d565b60405180910390a15050565b60008042336040516020016103549190610753565b6040516020818303038152906040528051906020012060001c8161037457fe5b0442416040516020016103879190610753565b6040516020818303038152906040528051906020012060001c816103a757fe5b0443420101016040516020016103bd9190610768565b6040516020818303038152906040528051906020012060001c90508281816103e157fe5b069392505050565b600082601f8301126103fa57600080fd5b813561040d6104088261085c565b610835565b9150818183526020840193506020810190508385602084028201111561043257600080fd5b60005b8381101561045e5781610448888261050e565b8452506020928301929190910190600101610435565b5050505092915050565b805161047381610942565b92915050565b600082601f83011261048a57600080fd5b81356104986104088261087d565b915080825260208301602083018583830111156104b457600080fd5b6104bf8382846108e9565b50505092915050565b600082601f8301126104d957600080fd5b81516104e76104088261087d565b9150808252602083016020830185838301111561050357600080fd5b6104bf8382846108f5565b803561047381610959565b60006020828403121561052b57600080fd5b813567ffffffffffffffff81111561054257600080fd5b61054e848285016103e9565b949350505050565b60006020828403121561056857600080fd5b813567ffffffffffffffff81111561057f57600080fd5b61054e84828501610479565b6000806040838503121561059e57600080fd5b825167ffffffffffffffff8111156105b557600080fd5b6105c1858286016104c8565b92505060206105d285828601610468565b9150509250929050565b600080604083850312156105ef57600080fd5b823567ffffffffffffffff81111561060657600080fd5b61061285828601610479565b925050602083013567ffffffffffffffff81111561062f57600080fd5b6105d285828601610479565b60006106478383610739565b505060200190565b610658816108d7565b82525050565b61065861066a826108b8565b610921565b600061067a826108ab565b61068481856108af565b935061068f836108a5565b8060005b838110156106bd5781516106a7888261063b565b97506106b2836108a5565b925050600101610693565b509495945050505050565b610658816108c3565b60006106dc826108ab565b6106e681856108af565b93506106f68185602086016108f5565b6106ff81610932565b9093019392505050565b610658816108de565b600061071f600b836108af565b6a6e6f74206f776e6572202160a81b815260200192915050565b610658816108d4565b61065861074e826108d4565b6108d4565b600061075f828461065e565b50601401919050565b60006107748284610742565b50602001919050565b6060810161078b828661064f565b818103602083015261079d81856106d1565b905081810360408301526107b181846106d1565b95945050505050565b604080825281016107cb818561066f565b9050818103602083015261054e81846106d1565b602080825281016107f081846106d1565b9392505050565b6040808252810161080881856106d1565b90506107f060208301846106c8565b602081016104738284610709565b6020808252810161047381610712565b60405181810167ffffffffffffffff8111828210171561085457600080fd5b604052919050565b600067ffffffffffffffff82111561087357600080fd5b5060209081020190565b600067ffffffffffffffff82111561089457600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b6000610473826108c8565b151590565b6001600160a01b031690565b90565b6000610473825b6000610473826108b8565b82818337506000910152565b60005b838110156109105781810151838201526020016108f8565b838111156102fa5750506000910152565b60006104738260006104738261093c565b601f01601f191690565b60601b90565b61094b816108c3565b811461095657600080fd5b50565b61094b816108d456fea365627a7a723158202f1132f5fd2f45cda035474a458e91bae25e0f55f608ca1c9e7f7a518ce8840b6c6578706572696d656e74616cf564736f6c63430005100040"

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

// TaskPushIterator is returned from FilterPush and is used to iterate over the raw logs and unpacked data for Push events raised by the Task contract.
type TaskPushIterator struct {
	Event *TaskPush // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaskPushIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPush)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaskPush)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaskPushIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPushIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPush represents a Push event raised by the Task contract.
type TaskPush struct {
	Sender common.Address
	Key    []byte
	Value  []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPush is a free log retrieval operation binding the contract event 0x0381ce894b22a3a2930bf6352fe67aa0d5af69d769ac34b044c53189c507a3a1.
//
// Solidity: event Push(address sender, bytes key, bytes value)
func (_Task *TaskFilterer) FilterPush(opts *bind.FilterOpts) (*TaskPushIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "Push")
	if err != nil {
		return nil, err
	}
	return &TaskPushIterator{contract: _Task.contract, event: "Push", logs: logs, sub: sub}, nil
}

// WatchPush is a free log subscription operation binding the contract event 0x0381ce894b22a3a2930bf6352fe67aa0d5af69d769ac34b044c53189c507a3a1.
//
// Solidity: event Push(address sender, bytes key, bytes value)
func (_Task *TaskFilterer) WatchPush(opts *bind.WatchOpts, sink chan<- *TaskPush) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "Push")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPush)
				if err := _Task.contract.UnpackLog(event, "Push", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePush is a log parse operation binding the contract event 0x0381ce894b22a3a2930bf6352fe67aa0d5af69d769ac34b044c53189c507a3a1.
//
// Solidity: event Push(address sender, bytes key, bytes value)
func (_Task *TaskFilterer) ParsePush(log types.Log) (*TaskPush, error) {
	event := new(TaskPush)
	if err := _Task.contract.UnpackLog(event, "Push", log); err != nil {
		return nil, err
	}
	return event, nil
}

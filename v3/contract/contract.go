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
var InterpreterBin = "0x608060405234801561001057600080fd5b5060405161071538038061071583398101604081905261002f91610138565b8051610042906000906020840190610049565b50506101f1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061008a57805160ff19168380011785556100b7565b828001600101855582156100b7579182015b828111156100b757825182559160200191906001019061009c565b506100c39291506100c7565b5090565b6100e191905b808211156100c357600081556001016100cd565b90565b600082601f8301126100f557600080fd5b81516101086101038261019a565b610174565b9150808252602083016020830185838301111561012457600080fd5b61012f8382846101c1565b50505092915050565b60006020828403121561014a57600080fd5b81516001600160401b0381111561016057600080fd5b61016c848285016100e4565b949350505050565b6040518181016001600160401b038111828210171561019257600080fd5b604052919050565b60006001600160401b038211156101b057600080fd5b506020601f91909101601f19160190565b60005b838110156101dc5781810151838201526020016101c4565b838111156101eb576000848401525b50505050565b610515806102006000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80634063d563146100465780637be8f86b1461005b578063a31662d21461006e575b600080fd5b6100596100543660046103d6565b61008d565b005b6100596100693660046103d6565b610161565b610076610263565b604051610084929190610460565b60405180910390f35b600480546001818101909255602081047f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01805460ff601f9093166101000a92909202199091169055600380548083019091557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b039093166001600160a01b03199384161790556002805480830182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180549092163317909155805481019055565b60005b60035481101561025e57816001600160a01b03166003828154811061018557fe5b6000918252602090912001546001600160a01b0316141561025657336001600160a01b0316600282815481106101b757fe5b6000918252602090912001546001600160a01b03161480156102035750600481815481106101e157fe5b90600052602060002090602091828204019190069054906101000a900460ff16155b156102505760016004828154811061021757fe5b60009182526020918290209181049091018054921515601f9092166101000a91820260ff90920219909216179055600180546000190190555b50610260565b600101610164565b505b50565b6000806001546000141561027957506000610313565b6000610286600154610317565b905060005b60035481101561030c57816102ca57600381815481106102a757fe5b6000918252602090912001546001600160a01b0316935060019250610313915050565b600481815481106102d757fe5b90600052602060002090602091828204019190069054906101000a900460ff161561030457600019909101905b60010161028b565b5060009150505b9091565b60008043423360405160200161032d9190610436565b6040516020818303038152906040528051906020012060001c8161034d57fe5b044542416040516020016103619190610436565b6040516020818303038152906040528051906020012060001c8161038157fe5b0444420101010101604051602001610399919061044b565b6040516020818303038152906040528051906020012060001c90508281816103bd57fe5b069392505050565b80356103d0816104be565b92915050565b6000602082840312156103e857600080fd5b60006103f484846103c5565b949350505050565b61040d61040882610482565b6104ac565b82525050565b61040d8161048d565b61040d81610492565b61040d610431826104a9565b6104a9565b600061044282846103fc565b50601401919050565b60006104578284610425565b50602001919050565b6040810161046e828561041c565b61047b6020830184610413565b9392505050565b60006103d08261049d565b151590565b60006103d082610482565b6001600160a01b031690565b90565b60006103d08260006103d08260601b90565b6104c781610492565b811461026057600080fdfea365627a7a72315820d47b87f8416a9d3942fbb5111d38a68e07c9afdbb060a0fb7335a78e164234f86c6578706572696d656e74616cf564736f6c63430005100040"

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
var RandomBin = "0x6080604052348015600f57600080fd5b50604c80601d6000396000f3fe6080604052600080fdfea365627a7a72315820d34fed5a046b44afca7598e45ac8b8512a1a1c0639906b7b1237d0febe014c456c6578706572696d656e74616cf564736f6c63430005100040"

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

// TaskABI is the input ABI used to generate the binding from.
const TaskABI = "[{\"inputs\":[{\"internalType\":\"contractinterpreter\",\"name\":\"IpfsObject\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"Kernel_size\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"stackAddresses\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"stackValue\",\"type\":\"bytes[]\"},{\"internalType\":\"uint128\",\"name\":\"SafetyProportionTreshold\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"SafetyLengthTreshold\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"kernel_idxs\",\"type\":\"uint256[]\"}],\"name\":\"acceptCommand\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"advertise\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"done\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCommand\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"kernel_idxs\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"read\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TaskFuncSigs maps the 4-byte function signature to its string representation.
var TaskFuncSigs = map[string]string{
	"29f5e410": "acceptCommand(uint256[])",
	"0f3023ef": "advertise()",
	"ae8421e1": "done()",
	"0110da29": "getCommand()",
	"e771ee0d": "push(bytes,bytes)",
	"8bf4515c": "read(bytes)",
}

// TaskBin is the compiled bytecode used for deploying new contracts.
var TaskBin = "0x60806040523480156200001157600080fd5b50604051620013f0380380620013f083398101604081905262000034916200046c565b600480546001600160801b03838116600160801b028186166001600160801b0319909316929092171617905584516200007590600590602088019062000219565b5060008054336001600160a01b0319918216178255600380549091166001600160a01b0389161790556001905b8651811015620000d357868181518110620000b957fe5b6020026020010151820291508080600101915050620000a2565b50600681905560408051828152602080840282010190915281801562000103578160200160208202803883390190505b5080516200011a9160079160209091019062000219565b508351855114156200020c5760005b85518110156200020a5760018682815181106200014257fe5b60200260200101516040516200015991906200057d565b90815260200160405180910390208582815181106200017457fe5b6020908102919091018101518254600181018085556000948552938390208251620001a6949190920192019062000269565b50506002868281518110620001b757fe5b6020026020010151604051620001ce91906200057d565b9081526040519081900360209081019091208054600181810183556000928352929091200180546001600160a01b031916331790550162000129565b505b50505050505050620006a0565b82805482825590600052602060002090810192821562000257579160200282015b82811115620002575782518255916020019190600101906200023a565b5062000265929150620002db565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002ac57805160ff191683800117855562000257565b82800160010185558215620002575791820182811115620002575782518255916020019190600101906200023a565b620002f891905b80821115620002655760008155600101620002e2565b90565b600082601f8301126200030d57600080fd5b8151620003246200031e82620005b9565b62000592565b81815260209384019390925082018360005b838110156200036657815186016200034f8882620003e9565b845250602092830192919091019060010162000336565b5050505092915050565b600082601f8301126200038257600080fd5b8151620003936200031e82620005b9565b91508181835260208401935060208101905083856020840282011115620003b957600080fd5b60005b83811015620003665781620003d288826200045f565b8452506020928301929190910190600101620003bc565b600082601f830112620003fb57600080fd5b81516200040c6200031e82620005da565b915080825260208301602083018583830111156200042957600080fd5b620004368382846200063d565b50505092915050565b80516200044c8162000670565b92915050565b80516200044c816200068a565b80516200044c8162000695565b60008060008060008060c087890312156200048657600080fd5b60006200049489896200043f565b96505060208701516001600160401b03811115620004b157600080fd5b620004bf89828a0162000370565b95505060408701516001600160401b03811115620004dc57600080fd5b620004ea89828a01620002fb565b94505060608701516001600160401b038111156200050757600080fd5b6200051589828a01620002fb565b93505060806200052889828a0162000452565b92505060a06200053b89828a0162000452565b9150509295509295509295565b6000620005558262000602565b62000561818562000606565b9350620005738185602086016200063d565b9290920192915050565b60006200058b828462000548565b9392505050565b6040518181016001600160401b0381118282101715620005b157600080fd5b604052919050565b60006001600160401b03821115620005d057600080fd5b5060209081020190565b60006001600160401b03821115620005f157600080fd5b506020601f91909101601f19160190565b5190565b919050565b60006200044c8262000631565b60006200044c826200060b565b6001600160801b031690565b6001600160a01b031690565b60005b838110156200065a57818101518382015260200162000640565b838111156200066a576000848401525b50505050565b6200067b8162000618565b81146200068757600080fd5b50565b6200067b8162000625565b6200067b81620002f8565b610d4080620006b06000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630110da29146100675780630f3023ef1461008657806329f5e410146100905780638bf4515c146100a3578063ae8421e1146100c4578063e771ee0d146100cc575b600080fd5b61006f6100df565b60405161007d929190610b8b565b60405180910390f35b61008e610175565b005b61008e61009e36600461095c565b6101f1565b6100b66100b1366004610999565b610265565b60405161007d929190610bb0565b61008e61058e565b61008e6100da3660046109ce565b6105d5565b606080600580549050604051908082528060200260200182016040528015610111578160200160208202803883390190505b50915060005b600554811015610160576101416005828154811061013157fe5b9060005260206000200154610734565b83828151811061014d57fe5b6020908102919091010152600101610117565b50506040805160208101909152600081529091565b6000546001600160a01b0316331461018c576101ef565b600354604051634063d56360e01b81526001600160a01b0390911690634063d563906101bc903090600401610bd0565b600060405180830381600087803b1580156101d657600080fd5b505af11580156101ea573d6000803e3d6000fd5b505050505b565b60006001815b83518110156102405783818151811061020c57fe5b60200260200101518202830192506005818154811061022757fe5b60009182526020909120015491909102906001016101f7565b506007828154811061024e57fe5b600091825260209091200180546001019055505050565b60606000606060018460405161027b9190610b63565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156103545760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156103405780601f1061031557610100808354040283529160200191610340565b820191906000526020600020905b81548152906001019060200180831161032357829003601f168201915b5050505050815260200190600101906102a9565b505050509050606060028560405161036c9190610b63565b90815260408051918290036020908101832080548083028501830190935282845291908301828280156103c857602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116103aa575b5050855160045494955093600160801b90046001600160801b031684101592506103fc915050575060009250610589915050565b606081604051908082528060200260200182016040528015610428578160200160208202803883390190505b50905060005b8281101561057f5760005484516001600160a01b039091169085908390811061045357fe5b60200260200101516001600160a01b031614156104915784818151811061047657fe5b60200260200101516001819150965096505050505050610589565b60005b818111610576578581815181106104a757fe5b60200260200101516040516104bc9190610b63565b60405180910390208683815181106104d057fe5b60200260200101516040516104e59190610b63565b6040518091039020141561056e578281815181106104ff57fe5b602090810291909101018051600101905260045483516001600160801b0391821660001981019092168602919085908490811061053857fe5b6020026020010151021061056e5785828151811061055257fe5b6020026020010151600181915097509750505050505050610589565b600101610494565b5060010161042e565b5060009450505050505b915091565b6000546001600160a01b031633146105a5576101ef565b600354604051637be8f86b60e01b81526001600160a01b0390911690637be8f86b906101bc903090600401610bd0565b6000546001600160a01b031633146106a95760606002836040516105f99190610b63565b908152604080519182900360209081018320805480830285018301909352828452919083018282801561065557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610637575b50939450600093505050505b81518110156106a65781818151811061067657fe5b60200260200101516001600160a01b0316336001600160a01b0316141561069e575050610730565b600101610661565b50505b6002826040516106b99190610b63565b90815260405190819003602090810182208054600181810183556000928352929091200180546001600160a01b03191633179055906106f9908490610b63565b90815260405160209181900382019020805460018101808355600092835291839020845192936101ea9391909201918501906107e2565b5050565b60008043423360405160200161074a9190610b4e565b6040516020818303038152906040528051906020012060001c8161076a57fe5b0445424160405160200161077e9190610b4e565b6040516020818303038152906040528051906020012060001c8161079e57fe5b04444201010101016040516020016107b69190610b76565b6040516020818303038152906040528051906020012060001c90508281816107da57fe5b069392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061082357805160ff1916838001178555610850565b82800160010185558215610850579182015b82811115610850578251825591602001919060010190610835565b5061085c929150610860565b5090565b61087a91905b8082111561085c5760008155600101610866565b90565b600082601f83011261088e57600080fd5b81356108a161089c82610c05565b610bde565b915081818352602084019350602081019050838560208402820111156108c657600080fd5b60005b838110156108f257816108dc888261094b565b84525060209283019291909101906001016108c9565b5050505092915050565b600082601f83011261090d57600080fd5b813561091b61089c82610c26565b9150808252602083016020830185838301111561093757600080fd5b610942838284610c8d565b50505092915050565b803561095681610ce6565b92915050565b60006020828403121561096e57600080fd5b813567ffffffffffffffff81111561098557600080fd5b6109918482850161087d565b949350505050565b6000602082840312156109ab57600080fd5b813567ffffffffffffffff8111156109c257600080fd5b610991848285016108fc565b600080604083850312156109e157600080fd5b823567ffffffffffffffff8111156109f857600080fd5b610a04858286016108fc565b925050602083013567ffffffffffffffff811115610a2157600080fd5b610a2d858286016108fc565b9150509250929050565b6000610a438383610b34565b505060200190565b610a5c610a5782610c66565b610cc5565b82525050565b6000610a6d82610c54565b610a778185610c58565b9350610a8283610c4e565b8060005b83811015610ab0578151610a9a8882610a37565b9750610aa583610c4e565b925050600101610a86565b509495945050505050565b610a5c81610c71565b6000610acf82610c54565b610ad98185610c58565b9350610ae9818560208601610c99565b610af281610cd6565b9093019392505050565b6000610b0782610c54565b610b118185610c61565b9350610b21818560208601610c99565b9290920192915050565b610a5c81610c82565b610a5c8161087a565b610a5c610b498261087a565b61087a565b6000610b5a8284610a4b565b50601401919050565b6000610b6f8284610afc565b9392505050565b6000610b828284610b3d565b50602001919050565b60408082528101610b9c8185610a62565b905081810360208301526109918184610ac4565b60408082528101610bc18185610ac4565b9050610b6f6020830184610abb565b602081016109568284610b2b565b60405181810167ffffffffffffffff81118282101715610bfd57600080fd5b604052919050565b600067ffffffffffffffff821115610c1c57600080fd5b5060209081020190565b600067ffffffffffffffff821115610c3d57600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b600061095682610c76565b151590565b6001600160a01b031690565b600061095682610c66565b82818337506000910152565b60005b83811015610cb4578181015183820152602001610c9c565b838111156101ea5750506000910152565b600061095682600061095682610ce0565b601f01601f191690565b60601b90565b610cef8161087a565b8114610cfa57600080fd5b5056fea365627a7a72315820398a8205c47f1f2a55988668e7260b34e58dafeba8c3cf3b66fa70eb0b9815146c6578706572696d656e74616cf564736f6c63430005100040"

// DeployTask deploys a new Ethereum contract, binding an instance of Task to it.
func DeployTask(auth *bind.TransactOpts, backend bind.ContractBackend, IpfsObject common.Address, Kernel_size []*big.Int, stackAddresses [][]byte, stackValue [][]byte, SafetyProportionTreshold *big.Int, SafetyLengthTreshold *big.Int) (common.Address, *types.Transaction, *Task, error) {
	parsed, err := abi.JSON(strings.NewReader(TaskABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TaskBin), backend, IpfsObject, Kernel_size, stackAddresses, stackValue, SafetyProportionTreshold, SafetyLengthTreshold)
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

// Advertise is a paid mutator transaction binding the contract method 0x0f3023ef.
//
// Solidity: function advertise() returns()
func (_Task *TaskTransactor) Advertise(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "advertise")
}

// Advertise is a paid mutator transaction binding the contract method 0x0f3023ef.
//
// Solidity: function advertise() returns()
func (_Task *TaskSession) Advertise() (*types.Transaction, error) {
	return _Task.Contract.Advertise(&_Task.TransactOpts)
}

// Advertise is a paid mutator transaction binding the contract method 0x0f3023ef.
//
// Solidity: function advertise() returns()
func (_Task *TaskTransactorSession) Advertise() (*types.Transaction, error) {
	return _Task.Contract.Advertise(&_Task.TransactOpts)
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

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

// TaskABI is the input ABI used to generate the binding from.
const TaskABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"IpfsHash\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"Kernel_size\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"stackAddresses\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"stackValue\",\"type\":\"bytes[]\"},{\"internalType\":\"uint128\",\"name\":\"SafetyProportionTreshold\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"SafetyLengthTreshold\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"kernel_idxs\",\"type\":\"uint256[]\"}],\"name\":\"acceptCommand\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCommand\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"kernel_idxs\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"read\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TaskFuncSigs maps the 4-byte function signature to its string representation.
var TaskFuncSigs = map[string]string{
	"29f5e410": "acceptCommand(uint256[])",
	"0110da29": "getCommand()",
	"e771ee0d": "push(bytes,bytes)",
	"8bf4515c": "read(bytes)",
}

// TaskBin is the compiled bytecode used for deploying new contracts.
var TaskBin = "0x60806040523480156200001157600080fd5b506040516200122e3803806200122e83398101604081905262000034916200045f565b600480546001600160801b03838116600160801b028186166001600160801b0319909316929092171617905585516200007590600390602089019062000219565b5084516200008b9060059060208801906200029e565b50600080546001600160a01b031916331781556001905b8651811015620000d357868181518110620000b957fe5b6020026020010151820291508080600101915050620000a2565b50600681905560408051828152602080840282010190915281801562000103578160200160208202803883390190505b5080516200011a916007916020909101906200029e565b508351855114156200020c5760005b85518110156200020a5760018682815181106200014257fe5b602002602001015160405162000159919062000587565b90815260200160405180910390208582815181106200017457fe5b6020908102919091018101518254600181018085556000948552938390208251620001a6949190920192019062000219565b50506002868281518110620001b757fe5b6020026020010151604051620001ce919062000587565b9081526040519081900360209081019091208054600181810183556000928352929091200180546001600160a01b031916331790550162000129565b505b5050505050505062000679565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200025c57805160ff19168380011785556200028c565b828001600101855582156200028c579182015b828111156200028c5782518255916020019190600101906200026f565b506200029a929150620002db565b5090565b8280548282559060005260206000209081019282156200028c57916020028201828111156200028c5782518255916020019190600101906200026f565b620002f891905b808211156200029a5760008155600101620002e2565b90565b600082601f8301126200030d57600080fd5b8151620003246200031e82620005c3565b6200059c565b81815260209384019390925082018360005b838110156200036657815186016200034f8882620003e9565b845250602092830192919091019060010162000336565b5050505092915050565b600082601f8301126200038257600080fd5b8151620003936200031e82620005c3565b91508181835260208401935060208101905083856020840282011115620003b957600080fd5b60005b83811015620003665781620003d2888262000452565b8452506020928301929190910190600101620003bc565b600082601f830112620003fb57600080fd5b81516200040c6200031e82620005e4565b915080825260208301602083018583830111156200042957600080fd5b6200043683828462000621565b50505092915050565b80516200044c8162000654565b92915050565b80516200044c816200066e565b60008060008060008060c087890312156200047957600080fd5b86516001600160401b038111156200049057600080fd5b6200049e89828a01620003e9565b96505060208701516001600160401b03811115620004bb57600080fd5b620004c989828a0162000370565b95505060408701516001600160401b03811115620004e657600080fd5b620004f489828a01620002fb565b94505060608701516001600160401b038111156200051157600080fd5b6200051f89828a01620002fb565b93505060806200053289828a016200043f565b92505060a06200054589828a016200043f565b9150509295509295509295565b60006200055f826200060c565b6200056b818562000610565b93506200057d81856020860162000621565b9290920192915050565b600062000595828462000552565b9392505050565b6040518181016001600160401b0381118282101715620005bb57600080fd5b604052919050565b60006001600160401b03821115620005da57600080fd5b5060209081020190565b60006001600160401b03821115620005fb57600080fd5b506020601f91909101601f19160190565b5190565b919050565b6001600160801b031690565b60005b838110156200063e57818101518382015260200162000624565b838111156200064e576000848401525b50505050565b6200065f8162000615565b81146200066b57600080fd5b50565b6200065f81620002f8565b610ba580620006896000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630110da291461005157806329f5e410146100705780638bf4515c14610085578063e771ee0d146100a6575b600080fd5b6100596100b9565b604051610067929190610a33565b60405180910390f35b61008361007e366004610859565b6101e0565b005b610098610093366004610896565b610254565b604051610067929190610a58565b6100836100b43660046108cb565b61057b565b6060806060600580548060200260200160405190810160405280929190818152602001828054801561010a57602002820191906000526020600020905b8154815260200190600101908083116100f6575b505060055460408051828152602080840282010190915294955060009493509150508015610142578160200160208202803883390190505b50935060005b82518110156101ad5782818151811061015d57fe5b6020026020010151828161016d57fe5b0685828151811061017a57fe5b60200260200101818152505082818151811061019257fe5b602002602001015182816101a257fe5b049150600101610148565b50506040805180820190915260138152721b9bdd081e595d081a5b5c1b195b595b9d1959606a1b60208201529150509091565b60006001815b835181101561022f578381815181106101fb57fe5b60200260200101518202830192506005818154811061021657fe5b60009182526020909120015491909102906001016101e6565b506007828154811061023d57fe5b600091825260209091200180546001019055505050565b60606000606060018460405161026a9190610a20565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156103435760008481526020908190208301805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529283018282801561032f5780601f106103045761010080835404028352916020019161032f565b820191906000526020600020905b81548152906001019060200180831161031257829003601f168201915b505050505081526020019060010190610298565b505050509050606060028560405161035b9190610a20565b90815260408051918290036020908101832080548083028501830190935282845291908301828280156103b757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610399575b5050855160045494955093600160801b90046001600160801b031684101592506103eb915050575060009250610576915050565b606081604051908082528060200260200182016040528015610417578160200160208202803883390190505b50905060005b8281101561056c5760005484516001600160a01b039091169085908390811061044257fe5b60200260200101516001600160a01b031614156104805784818151811061046557fe5b60200260200101516001819150965096505050505050610576565b60005b8181116105635785818151811061049657fe5b60200260200101516040516104ab9190610a20565b60405180910390208683815181106104bf57fe5b60200260200101516040516104d49190610a20565b6040518091039020141561055b578281815181106104ee57fe5b602090810291909101018051600101905260045483516001600160801b0390911690859085908490811061051e57fe5b60200260200101518161052d57fe5b041061055b5785828151811061053f57fe5b6020026020010151600181915097509750505050505050610576565b600101610483565b5060010161041d565b5060009450505050505b915091565b606060028360405161058d9190610a20565b90815260408051918290036020908101832080548083028501830190935282845291908301828280156105e957602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116105cb575b505060005493945050506001600160a01b03909116331490506106505760005b815181101561064e5781818151811061061e57fe5b60200260200101516001600160a01b0316336001600160a01b031614156106465750506106db565b600101610609565b505b6002836040516106609190610a20565b90815260405190819003602090810182208054600181810183556000928352929091200180546001600160a01b03191633179055906106a0908590610a20565b90815260405160209181900382019020805460018101808355600092835291839020855192936106d79391909201918601906106df565b5050505b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061072057805160ff191683800117855561074d565b8280016001018555821561074d579182015b8281111561074d578251825591602001919060010190610732565b5061075992915061075d565b5090565b61077791905b808211156107595760008155600101610763565b90565b600082601f83011261078b57600080fd5b813561079e61079982610a9f565b610a78565b915081818352602084019350602081019050838560208402820111156107c357600080fd5b60005b838110156107ef57816107d98882610848565b84525060209283019291909101906001016107c6565b5050505092915050565b600082601f83011261080a57600080fd5b813561081861079982610ac0565b9150808252602083016020830185838301111561083457600080fd5b61083f838284610b05565b50505092915050565b803561085381610b4b565b92915050565b60006020828403121561086b57600080fd5b813567ffffffffffffffff81111561088257600080fd5b61088e8482850161077a565b949350505050565b6000602082840312156108a857600080fd5b813567ffffffffffffffff8111156108bf57600080fd5b61088e848285016107f9565b600080604083850312156108de57600080fd5b823567ffffffffffffffff8111156108f557600080fd5b610901858286016107f9565b925050602083013567ffffffffffffffff81111561091e57600080fd5b61092a858286016107f9565b9150509250929050565b60006109408383610a17565b505060200190565b600061095382610aee565b61095d8185610af2565b935061096883610ae8565b8060005b838110156109965781516109808882610934565b975061098b83610ae8565b92505060010161096c565b509495945050505050565b6109aa81610b00565b82525050565b60006109bb82610aee565b6109c58185610af2565b93506109d5818560208601610b11565b6109de81610b41565b9093019392505050565b60006109f382610aee565b6109fd8185610afb565b9350610a0d818560208601610b11565b9290920192915050565b6109aa81610777565b6000610a2c82846109e8565b9392505050565b60408082528101610a448185610948565b9050818103602083015261088e81846109b0565b60408082528101610a6981856109b0565b9050610a2c60208301846109a1565b60405181810167ffffffffffffffff81118282101715610a9757600080fd5b604052919050565b600067ffffffffffffffff821115610ab657600080fd5b5060209081020190565b600067ffffffffffffffff821115610ad757600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b151590565b82818337506000910152565b60005b83811015610b2c578181015183820152602001610b14565b83811115610b3b576000848401525b50505050565b601f01601f191690565b610b5481610777565b8114610b5f57600080fd5b5056fea365627a7a723158203420e5adf52052f618041078a94fb801c2e5c0ce48646d2b8b60ae81a9bbf5ed6c6578706572696d656e74616cf564736f6c63430005100040"

// DeployTask deploys a new Ethereum contract, binding an instance of Task to it.
func DeployTask(auth *bind.TransactOpts, backend bind.ContractBackend, IpfsHash string, Kernel_size []*big.Int, stackAddresses [][]byte, stackValue [][]byte, SafetyProportionTreshold *big.Int, SafetyLengthTreshold *big.Int) (common.Address, *types.Transaction, *Task, error) {
	parsed, err := abi.JSON(strings.NewReader(TaskABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TaskBin), backend, IpfsHash, Kernel_size, stackAddresses, stackValue, SafetyProportionTreshold, SafetyLengthTreshold)
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

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
const TaskABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"IpfsHash\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"Kernel_size\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"stackAddresses\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"stackValue\",\"type\":\"bytes[]\"},{\"internalType\":\"uint128\",\"name\":\"SafetyProportionTreshold\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"SafetyLengthTreshold\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[],\"name\":\"getCommand\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"kernel_idxs\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"read\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TaskFuncSigs maps the 4-byte function signature to its string representation.
var TaskFuncSigs = map[string]string{
	"0110da29": "getCommand()",
	"e771ee0d": "push(bytes,bytes)",
	"8bf4515c": "read(bytes)",
}

// TaskBin is the compiled bytecode used for deploying new contracts.
var TaskBin = "0x60806040523480156200001157600080fd5b5060405162000ff338038062000ff383398101604081905262000034916200045f565b600480546001600160801b03838116600160801b028186166001600160801b0319909316929092171617905585516200007590600390602089019062000219565b5084516200008b9060059060208801906200029e565b50600080546001600160a01b031916331781556001905b8651811015620000d357868181518110620000b957fe5b6020026020010151820291508080600101915050620000a2565b50600681905560408051828152602080840282010190915281801562000103578160200160208202803883390190505b5080516200011a916007916020909101906200029e565b508351855114156200020c5760005b85518110156200020a5760018682815181106200014257fe5b602002602001015160405162000159919062000587565b90815260200160405180910390208582815181106200017457fe5b6020908102919091018101518254600181018085556000948552938390208251620001a6949190920192019062000219565b50506002868281518110620001b757fe5b6020026020010151604051620001ce919062000587565b9081526040519081900360209081019091208054600181810183556000928352929091200180546001600160a01b031916331790550162000129565b505b5050505050505062000679565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200025c57805160ff19168380011785556200028c565b828001600101855582156200028c579182015b828111156200028c5782518255916020019190600101906200026f565b506200029a929150620002db565b5090565b8280548282559060005260206000209081019282156200028c57916020028201828111156200028c5782518255916020019190600101906200026f565b620002f891905b808211156200029a5760008155600101620002e2565b90565b600082601f8301126200030d57600080fd5b8151620003246200031e82620005c3565b6200059c565b81815260209384019390925082018360005b838110156200036657815186016200034f8882620003e9565b845250602092830192919091019060010162000336565b5050505092915050565b600082601f8301126200038257600080fd5b8151620003936200031e82620005c3565b91508181835260208401935060208101905083856020840282011115620003b957600080fd5b60005b83811015620003665781620003d2888262000452565b8452506020928301929190910190600101620003bc565b600082601f830112620003fb57600080fd5b81516200040c6200031e82620005e4565b915080825260208301602083018583830111156200042957600080fd5b6200043683828462000621565b50505092915050565b80516200044c8162000654565b92915050565b80516200044c816200066e565b60008060008060008060c087890312156200047957600080fd5b86516001600160401b038111156200049057600080fd5b6200049e89828a01620003e9565b96505060208701516001600160401b03811115620004bb57600080fd5b620004c989828a0162000370565b95505060408701516001600160401b03811115620004e657600080fd5b620004f489828a01620002fb565b94505060608701516001600160401b038111156200051157600080fd5b6200051f89828a01620002fb565b93505060806200053289828a016200043f565b92505060a06200054589828a016200043f565b9150509295509295509295565b60006200055f826200060c565b6200056b818562000610565b93506200057d81856020860162000621565b9290920192915050565b600062000595828462000552565b9392505050565b6040518181016001600160401b0381118282101715620005bb57600080fd5b604052919050565b60006001600160401b03821115620005da57600080fd5b5060209081020190565b60006001600160401b03821115620005fb57600080fd5b506020601f91909101601f19160190565b5190565b919050565b6001600160801b031690565b60005b838110156200063e57818101518382015260200162000624565b838111156200064e576000848401525b50505050565b6200065f8162000615565b81146200066b57600080fd5b50565b6200065f81620002f8565b61096a80620006896000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630110da29146100465780638bf4515c14610065578063e771ee0d14610086575b600080fd5b61004e61009b565b60405161005c929190610830565b60405180910390f35b61007861007336600461068b565b6101de565b60405161005c929190610855565b6100996100943660046106c8565b610438565b005b606080606060058054806020026020016040519081016040528092919081815260200182805480156100ec57602002820191906000526020600020905b8154815260200190600101908083116100d8575b505060055460408051828152602080840282010190915294955060009493509150508015610124578160200160208202803883390190505b5093506007818154811061013457fe5b60009182526020822001805460010190555b82518110156101ab5782818151811061015b57fe5b6020026020010151828161016b57fe5b0685828151811061017857fe5b60200260200101818152505082818151811061019057fe5b602002602001015182816101a057fe5b049150600101610146565b50506040805180820190915260138152721b9bdd081e595d081a5b5c1b195b595b9d1959606a1b60208201529150509091565b6060600060606001846040516101f4919061081d565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156102cd5760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156102b95780601f1061028e576101008083540402835291602001916102b9565b820191906000526020600020905b81548152906001019060200180831161029c57829003601f168201915b505050505081526020019060010190610222565b5050825160045493945092600160801b90046001600160801b031683101591506102ff90505750600091506104339050565b60608160405190808252806020026020018201604052801561032b578160200160208202803883390190505b50905060015b8281101561042a5760005b8181116104215784818151811061034f57fe5b6020026020010151604051610364919061081d565b604051809103902085838151811061037857fe5b602002602001015160405161038d919061081d565b60405180910390201415610419578281815181106103a757fe5b602090810291909101018051600101905260045483516001600160801b039091169085908590849081106103d757fe5b6020026020010151816103e657fe5b041115610414578482815181106103f957fe5b60200260200101516001819150965096505050505050610433565b610421565b60010161033c565b50600101610331565b50600093505050505b915091565b606060028360405161044a919061081d565b90815260408051918290036020908101832080548083028501830190935282845291908301828280156104a657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610488575b505060005493945050506001600160a01b039091163314905061050d5760005b815181101561050b578181815181106104db57fe5b60200260200101516001600160a01b0316336001600160a01b03161415610503575050610598565b6001016104c6565b505b60028360405161051d919061081d565b90815260405190819003602090810182208054600181810183556000928352929091200180546001600160a01b031916331790559061055d90859061081d565b908152604051602091819003820190208054600181018083556000928352918390208551929361059493919092019186019061059c565b5050505b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106105dd57805160ff191683800117855561060a565b8280016001018555821561060a579182015b8281111561060a5782518255916020019190600101906105ef565b5061061692915061061a565b5090565b61063491905b808211156106165760008155600101610620565b90565b600082601f83011261064857600080fd5b813561065b6106568261089c565b610875565b9150808252602083016020830185838301111561067757600080fd5b6106828382846108e1565b50505092915050565b60006020828403121561069d57600080fd5b813567ffffffffffffffff8111156106b457600080fd5b6106c084828501610637565b949350505050565b600080604083850312156106db57600080fd5b823567ffffffffffffffff8111156106f257600080fd5b6106fe85828601610637565b925050602083013567ffffffffffffffff81111561071b57600080fd5b61072785828601610637565b9150509250929050565b600061073d8383610814565b505060200190565b6000610750826108ca565b61075a81856108ce565b9350610765836108c4565b8060005b8381101561079357815161077d8882610731565b9750610788836108c4565b925050600101610769565b509495945050505050565b6107a7816108dc565b82525050565b60006107b8826108ca565b6107c281856108ce565b93506107d28185602086016108ed565b6107db8161091d565b9093019392505050565b60006107f0826108ca565b6107fa81856108d7565b935061080a8185602086016108ed565b9290920192915050565b6107a781610634565b600061082982846107e5565b9392505050565b604080825281016108418185610745565b905081810360208301526106c081846107ad565b6040808252810161086681856107ad565b9050610829602083018461079e565b60405181810167ffffffffffffffff8111828210171561089457600080fd5b604052919050565b600067ffffffffffffffff8211156108b357600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b151590565b82818337506000910152565b60005b838110156109085781810151838201526020016108f0565b83811115610917576000848401525b50505050565b601f01601f19169056fea365627a7a723158209caf5a6eae8425ab529272881b43cf97d6e3ffdb58c12b2e6b78a1522f1e159c6c6578706572696d656e74616cf564736f6c63430005100040"

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

// GetCommand is a paid mutator transaction binding the contract method 0x0110da29.
//
// Solidity: function getCommand() returns(uint256[] kernel_idxs, string error)
func (_Task *TaskTransactor) GetCommand(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "getCommand")
}

// GetCommand is a paid mutator transaction binding the contract method 0x0110da29.
//
// Solidity: function getCommand() returns(uint256[] kernel_idxs, string error)
func (_Task *TaskSession) GetCommand() (*types.Transaction, error) {
	return _Task.Contract.GetCommand(&_Task.TransactOpts)
}

// GetCommand is a paid mutator transaction binding the contract method 0x0110da29.
//
// Solidity: function getCommand() returns(uint256[] kernel_idxs, string error)
func (_Task *TaskTransactorSession) GetCommand() (*types.Transaction, error) {
	return _Task.Contract.GetCommand(&_Task.TransactOpts)
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

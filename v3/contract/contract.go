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
var InterpreterBin = "0x608060405234801561001057600080fd5b5060405161059938038061059983398101604081905261002f91610138565b8051610042906000906020840190610049565b50506101f1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061008a57805160ff19168380011785556100b7565b828001600101855582156100b7579182015b828111156100b757825182559160200191906001019061009c565b506100c39291506100c7565b5090565b6100e191905b808211156100c357600081556001016100cd565b90565b600082601f8301126100f557600080fd5b81516101086101038261019a565b610174565b9150808252602083016020830185838301111561012457600080fd5b61012f8382846101c1565b50505092915050565b60006020828403121561014a57600080fd5b81516001600160401b0381111561016057600080fd5b61016c848285016100e4565b949350505050565b6040518181016001600160401b038111828210171561019257600080fd5b604052919050565b60006001600160401b038211156101b057600080fd5b506020601f91909101601f19160190565b60005b838110156101dc5781810151838201526020016101c4565b838111156101eb576000848401525b50505050565b610399806102006000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80634063d563146100465780637be8f86b1461005b578063a31662d21461006e575b600080fd5b6100596100543660046102bb565b61008d565b005b6100596100693660046102bb565b610161565b610076610263565b6040516100849291906102f9565b60405180910390f35b600480546001818101909255602081047f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01805460ff601f9093166101000a92909202199091169055600380548083019091557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b039093166001600160a01b03199384161790556002805480830182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180549092163317909155805481019055565b60005b60035481101561025e57816001600160a01b03166003828154811061018557fe5b6000918252602090912001546001600160a01b0316141561025657336001600160a01b0316600282815481106101b757fe5b6000918252602090912001546001600160a01b03161480156102035750600481815481106101e157fe5b90600052602060002090602091828204019190069054906101000a900460ff16155b156102505760016004828154811061021757fe5b60009182526020918290209181049091018054921515601f9092166101000a91820260ff90920219909216179055600180546000190190555b50610260565b600101610164565b505b50565b6003546000908190610277575060006102a6565b60008090506003818154811061028957fe5b6000918252602090912001546001600160a01b0316925060019150505b9091565b80356102b581610342565b92915050565b6000602082840312156102cd57600080fd5b60006102d984846102aa565b949350505050565b6102ea81610326565b82525050565b6102ea8161032b565b6040810161030782856102f0565b61031460208301846102e1565b9392505050565b60006102b582610336565b151590565b60006102b58261031b565b6001600160a01b031690565b61034b8161032b565b811461026057600080fdfea365627a7a72315820db66d09e94cabb9bbafa7cbf0472a8d9fcb250afe74e86e12ba0ba6eb8d94ab86c6578706572696d656e74616cf564736f6c63430005100040"

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
var TaskBin = "0x60806040523480156200001157600080fd5b506040516200136b3803806200136b83398101604081905262000034916200046c565b600480546001600160801b03838116600160801b028186166001600160801b0319909316929092171617905584516200007590600590602088019062000219565b5060008054336001600160a01b0319918216178255600380549091166001600160a01b0389161790556001905b8651811015620000d357868181518110620000b957fe5b6020026020010151820291508080600101915050620000a2565b50600681905560408051828152602080840282010190915281801562000103578160200160208202803883390190505b5080516200011a9160079160209091019062000219565b508351855114156200020c5760005b85518110156200020a5760018682815181106200014257fe5b60200260200101516040516200015991906200057d565b90815260200160405180910390208582815181106200017457fe5b6020908102919091018101518254600181018085556000948552938390208251620001a6949190920192019062000269565b50506002868281518110620001b757fe5b6020026020010151604051620001ce91906200057d565b9081526040519081900360209081019091208054600181810183556000928352929091200180546001600160a01b031916331790550162000129565b505b50505050505050620006a0565b82805482825590600052602060002090810192821562000257579160200282015b82811115620002575782518255916020019190600101906200023a565b5062000265929150620002db565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002ac57805160ff191683800117855562000257565b82800160010185558215620002575791820182811115620002575782518255916020019190600101906200023a565b620002f891905b80821115620002655760008155600101620002e2565b90565b600082601f8301126200030d57600080fd5b8151620003246200031e82620005b9565b62000592565b81815260209384019390925082018360005b838110156200036657815186016200034f8882620003e9565b845250602092830192919091019060010162000336565b5050505092915050565b600082601f8301126200038257600080fd5b8151620003936200031e82620005b9565b91508181835260208401935060208101905083856020840282011115620003b957600080fd5b60005b83811015620003665781620003d288826200045f565b8452506020928301929190910190600101620003bc565b600082601f830112620003fb57600080fd5b81516200040c6200031e82620005da565b915080825260208301602083018583830111156200042957600080fd5b620004368382846200063d565b50505092915050565b80516200044c8162000670565b92915050565b80516200044c816200068a565b80516200044c8162000695565b60008060008060008060c087890312156200048657600080fd5b60006200049489896200043f565b96505060208701516001600160401b03811115620004b157600080fd5b620004bf89828a0162000370565b95505060408701516001600160401b03811115620004dc57600080fd5b620004ea89828a01620002fb565b94505060608701516001600160401b038111156200050757600080fd5b6200051589828a01620002fb565b93505060806200052889828a0162000452565b92505060a06200053b89828a0162000452565b9150509295509295509295565b6000620005558262000602565b62000561818562000606565b9350620005738185602086016200063d565b9290920192915050565b60006200058b828462000548565b9392505050565b6040518181016001600160401b0381118282101715620005b157600080fd5b604052919050565b60006001600160401b03821115620005d057600080fd5b5060209081020190565b60006001600160401b03821115620005f157600080fd5b506020601f91909101601f19160190565b5190565b919050565b60006200044c8262000631565b60006200044c826200060b565b6001600160801b031690565b6001600160a01b031690565b60005b838110156200065a57818101518382015260200162000640565b838111156200066a576000848401525b50505050565b6200067b8162000618565b81146200068757600080fd5b50565b6200067b8162000625565b6200067b81620002f8565b610cbb80620006b06000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630110da29146100675780630f3023ef1461008657806329f5e410146100905780638bf4515c146100a3578063ae8421e1146100c4578063e771ee0d146100cc575b600080fd5b61006f6100df565b60405161007d929190610b22565b60405180910390f35b61008e610206565b005b61008e61009e36600461093f565b610282565b6100b66100b136600461097c565b6102f6565b60405161007d929190610b47565b61008e61061f565b61008e6100da3660046109b1565b610666565b6060806060600580548060200260200160405190810160405280929190818152602001828054801561013057602002820191906000526020600020905b81548152602001906001019080831161011c575b505060055460408051828152602080840282010190915294955060009493509150508015610168578160200160208202803883390190505b50935060005b82518110156101d35782818151811061018357fe5b6020026020010151828161019357fe5b068582815181106101a057fe5b6020026020010181815250508281815181106101b857fe5b602002602001015182816101c857fe5b04915060010161016e565b50506040805180820190915260138152721b9bdd081e595d081a5b5c1b195b595b9d1959606a1b60208201529150509091565b6000546001600160a01b0316331461021d57610280565b600354604051634063d56360e01b81526001600160a01b0390911690634063d5639061024d903090600401610b67565b600060405180830381600087803b15801561026757600080fd5b505af115801561027b573d6000803e3d6000fd5b505050505b565b60006001815b83518110156102d15783818151811061029d57fe5b6020026020010151820283019250600581815481106102b857fe5b6000918252602090912001549190910290600101610288565b50600782815481106102df57fe5b600091825260209091200180546001019055505050565b60606000606060018460405161030c9190610b0f565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156103e55760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156103d15780601f106103a6576101008083540402835291602001916103d1565b820191906000526020600020905b8154815290600101906020018083116103b457829003601f168201915b50505050508152602001906001019061033a565b50505050905060606002856040516103fd9190610b0f565b908152604080519182900360209081018320805480830285018301909352828452919083018282801561045957602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161043b575b5050855160045494955093600160801b90046001600160801b0316841015925061048d91505057506000925061061a915050565b6060816040519080825280602002602001820160405280156104b9578160200160208202803883390190505b50905060005b828110156106105760005484516001600160a01b03909116908590839081106104e457fe5b60200260200101516001600160a01b031614156105225784818151811061050757fe5b6020026020010151600181915096509650505050505061061a565b60005b8181116106075785818151811061053857fe5b602002602001015160405161054d9190610b0f565b604051809103902086838151811061056157fe5b60200260200101516040516105769190610b0f565b604051809103902014156105ff5782818151811061059057fe5b602090810291909101018051600101905260045483516001600160801b039182166000198101909216860291908590849081106105c957fe5b602002602001015102106105ff578582815181106105e357fe5b602002602001015160018191509750975050505050505061061a565b600101610525565b506001016104bf565b5060009450505050505b915091565b6000546001600160a01b0316331461063657610280565b600354604051637be8f86b60e01b81526001600160a01b0390911690637be8f86b9061024d903090600401610b67565b6000546001600160a01b0316331461073a57606060028360405161068a9190610b0f565b90815260408051918290036020908101832080548083028501830190935282845291908301828280156106e657602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116106c8575b50939450600093505050505b81518110156107375781818151811061070757fe5b60200260200101516001600160a01b0316336001600160a01b0316141561072f5750506107c1565b6001016106f2565b50505b60028260405161074a9190610b0f565b90815260405190819003602090810182208054600181810183556000928352929091200180546001600160a01b031916331790559061078a908490610b0f565b908152604051602091819003820190208054600181018083556000928352918390208451929361027b9391909201918501906107c5565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061080657805160ff1916838001178555610833565b82800160010185558215610833579182015b82811115610833578251825591602001919060010190610818565b5061083f929150610843565b5090565b61085d91905b8082111561083f5760008155600101610849565b90565b600082601f83011261087157600080fd5b813561088461087f82610b9c565b610b75565b915081818352602084019350602081019050838560208402820111156108a957600080fd5b60005b838110156108d557816108bf888261092e565b84525060209283019291909101906001016108ac565b5050505092915050565b600082601f8301126108f057600080fd5b81356108fe61087f82610bbd565b9150808252602083016020830185838301111561091a57600080fd5b610925838284610c1f565b50505092915050565b803561093981610c61565b92915050565b60006020828403121561095157600080fd5b813567ffffffffffffffff81111561096857600080fd5b61097484828501610860565b949350505050565b60006020828403121561098e57600080fd5b813567ffffffffffffffff8111156109a557600080fd5b610974848285016108df565b600080604083850312156109c457600080fd5b823567ffffffffffffffff8111156109db57600080fd5b6109e7858286016108df565b925050602083013567ffffffffffffffff811115610a0457600080fd5b610a10858286016108df565b9150509250929050565b6000610a268383610b06565b505060200190565b6000610a3982610beb565b610a438185610bef565b9350610a4e83610be5565b8060005b83811015610a7c578151610a668882610a1a565b9750610a7183610be5565b925050600101610a52565b509495945050505050565b610a9081610bfd565b82525050565b6000610aa182610beb565b610aab8185610bef565b9350610abb818560208601610c2b565b610ac481610c57565b9093019392505050565b6000610ad982610beb565b610ae38185610bf8565b9350610af3818560208601610c2b565b9290920192915050565b610a9081610c0e565b610a908161085d565b6000610b1b8284610ace565b9392505050565b60408082528101610b338185610a2e565b905081810360208301526109748184610a96565b60408082528101610b588185610a96565b9050610b1b6020830184610a87565b602081016109398284610afd565b60405181810167ffffffffffffffff81118282101715610b9457600080fd5b604052919050565b600067ffffffffffffffff821115610bb357600080fd5b5060209081020190565b600067ffffffffffffffff821115610bd457600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b151590565b6001600160a01b031690565b600061093982600061093982610c02565b82818337506000910152565b60005b83811015610c46578181015183820152602001610c2e565b8381111561027b5750506000910152565b601f01601f191690565b610c6a8161085d565b8114610c7557600080fd5b5056fea365627a7a72315820cbcf58a416268cfff515a3c0cf00ff8a7e433c4f6f280d35ea25a543e2b2cfc76c6578706572696d656e74616cf564736f6c63430005100040"

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

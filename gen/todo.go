// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todo

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TodoTask is an auto generated low-level Go binding around an user-defined struct.
type TodoTask struct {
	Content string
	Status  bool
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"toggle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611042806100606000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100d85780639507d39a146100f6578063b0c8f9dc14610126578063f745630f146101425761007d565b80630f560cd7146100825780634cc82215146100a0578063751ef753146100bc575b600080fd5b61008a61015e565b6040516100979190610b1f565b60405180910390f35b6100ba60048036038101906100b59190610b8b565b6102c2565b005b6100d660048036038101906100d19190610b8b565b61043d565b005b6100e0610508565b6040516100ed9190610bf9565b60405180910390f35b610110600480360381019061010b9190610b8b565b61052c565b60405161011d9190610c51565b60405180910390f35b610140600480360381019061013b9190610da8565b61066d565b005b61015c60048036038101906101579190610df1565b61074a565b005b60603373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146101b857600080fd5b6001805480602002602001604051908101604052809291908181526020016000905b828210156102b9578382906000526020600020906002020160405180604001604052908160008201805461020d90610e7c565b80601f016020809104026020016040519081016040528092919081815260200182805461023990610e7c565b80156102865780601f1061025b57610100808354040283529160200191610286565b820191906000526020600020905b81548152906001019060200180831161026957829003601f168201915b505050505081526020016001820160009054906101000a900460ff161515151581525050815260200190600101906101da565b50505050905090565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461031a57600080fd5b60008190505b600180805490506103319190610edc565b8110156103ea57600180826103469190610f10565b8154811061035757610356610f66565b5b90600052602060002090600202016001828154811061037957610378610f66565b5b9060005260206000209060020201600082018160000190805461039b90610e7c565b6103a69291906107e0565b506001820160009054906101000a900460ff168160010160006101000a81548160ff02191690831515021790555090505080806103e290610f95565b915050610320565b5060018054806103fd576103fc610fdd565b5b600190038181906000526020600020906002020160008082016000610422919061086d565b6001820160006101000a81549060ff02191690555050905550565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461049557600080fd5b600181815481106104a9576104a8610f66565b5b906000526020600020906002020160010160009054906101000a900460ff1615600182815481106104dd576104dc610f66565b5b906000526020600020906002020160010160006101000a81548160ff02191690831515021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6105346108ad565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461058c57600080fd5b600182815481106105a05761059f610f66565b5b90600052602060002090600202016040518060400160405290816000820180546105c990610e7c565b80601f01602080910402602001604051908101604052809291908181526020018280546105f590610e7c565b80156106425780601f1061061757610100808354040283529160200191610642565b820191906000526020600020905b81548152906001019060200180831161062557829003601f168201915b505050505081526020016001820160009054906101000a900460ff1615151515815250509050919050565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146106c557600080fd5b6001604051806040016040528083815260200160001515815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000190805190602001906107249291906108c9565b5060208201518160010160006101000a81548160ff021916908315150217905550505050565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146107a257600080fd5b80600183815481106107b7576107b6610f66565b5b906000526020600020906002020160000190805190602001906107db9291906108c9565b505050565b8280546107ec90610e7c565b90600052602060002090601f01602090048101928261080e576000855561085c565b82601f1061081f578054855561085c565b8280016001018555821561085c57600052602060002091601f016020900482015b8281111561085b578254825591600101919060010190610840565b5b509050610869919061094f565b5090565b50805461087990610e7c565b6000825580601f1061088b57506108aa565b601f0160209004906000526020600020908101906108a9919061094f565b5b50565b6040518060400160405280606081526020016000151581525090565b8280546108d590610e7c565b90600052602060002090601f0160209004810192826108f7576000855561093e565b82601f1061091057805160ff191683800117855561093e565b8280016001018555821561093e579182015b8281111561093d578251825591602001919060010190610922565b5b50905061094b919061094f565b5090565b5b80821115610968576000816000905550600101610950565b5090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156109d25780820151818401526020810190506109b7565b838111156109e1576000848401525b50505050565b6000601f19601f8301169050919050565b6000610a0382610998565b610a0d81856109a3565b9350610a1d8185602086016109b4565b610a26816109e7565b840191505092915050565b60008115159050919050565b610a4681610a31565b82525050565b60006040830160008301518482036000860152610a6982826109f8565b9150506020830151610a7e6020860182610a3d565b508091505092915050565b6000610a958383610a4c565b905092915050565b6000602082019050919050565b6000610ab58261096c565b610abf8185610977565b935083602082028501610ad185610988565b8060005b85811015610b0d5784840389528151610aee8582610a89565b9450610af983610a9d565b925060208a01995050600181019050610ad5565b50829750879550505050505092915050565b60006020820190508181036000830152610b398184610aaa565b905092915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610b6881610b55565b8114610b7357600080fd5b50565b600081359050610b8581610b5f565b92915050565b600060208284031215610ba157610ba0610b4b565b5b6000610baf84828501610b76565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610be382610bb8565b9050919050565b610bf381610bd8565b82525050565b6000602082019050610c0e6000830184610bea565b92915050565b60006040830160008301518482036000860152610c3182826109f8565b9150506020830151610c466020860182610a3d565b508091505092915050565b60006020820190508181036000830152610c6b8184610c14565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610cb5826109e7565b810181811067ffffffffffffffff82111715610cd457610cd3610c7d565b5b80604052505050565b6000610ce7610b41565b9050610cf38282610cac565b919050565b600067ffffffffffffffff821115610d1357610d12610c7d565b5b610d1c826109e7565b9050602081019050919050565b82818337600083830152505050565b6000610d4b610d4684610cf8565b610cdd565b905082815260208101848484011115610d6757610d66610c78565b5b610d72848285610d29565b509392505050565b600082601f830112610d8f57610d8e610c73565b5b8135610d9f848260208601610d38565b91505092915050565b600060208284031215610dbe57610dbd610b4b565b5b600082013567ffffffffffffffff811115610ddc57610ddb610b50565b5b610de884828501610d7a565b91505092915050565b60008060408385031215610e0857610e07610b4b565b5b6000610e1685828601610b76565b925050602083013567ffffffffffffffff811115610e3757610e36610b50565b5b610e4385828601610d7a565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610e9457607f821691505b602082108103610ea757610ea6610e4d565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610ee782610b55565b9150610ef283610b55565b925082821015610f0557610f04610ead565b5b828203905092915050565b6000610f1b82610b55565b9150610f2683610b55565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610f5b57610f5a610ead565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000610fa082610b55565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610fd257610fd1610ead565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea2646970667358221220293efa2919e783fbb797d109f5040e263722855da156fd6d977b767cc7374a9a64736f6c634300080d0033",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TodoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoCaller) Get(opts *bind.CallOpts, _id *big.Int) (TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "get", _id)

	if err != nil {
		return *new(TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new(TodoTask)).(*TodoTask)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoSession) Get(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoCallerSession) Get(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _id)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCaller) List(opts *bind.CallOpts) ([]TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTask)).(*[]TodoTask)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCallerSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCallerSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoTransactor) Add(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "add", _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoSession) Add(_content string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoTransactorSession) Add(_content string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactor) Remove(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "remove", _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactorSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Todo *TodoTransactor) Toggle(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "toggle", _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Todo *TodoSession) Toggle(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Toggle(&_Todo.TransactOpts, _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Todo *TodoTransactorSession) Toggle(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Toggle(&_Todo.TransactOpts, _id)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactor) Update(opts *bind.TransactOpts, _id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "update", _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactorSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}

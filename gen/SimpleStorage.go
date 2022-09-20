// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SimpleStorage

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

// SimpleStorageMetaData contains all meta data concerning the SimpleStorage contract.
var SimpleStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_favoriteNumber\",\"type\":\"uint256\"}],\"name\":\"addPerson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"nameToFavoriteNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"people\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"favoriteNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_favoriteNumber\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061093b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632e64cec11461005c5780636057361d1461007a5780636f760f41146100965780638bab8dd5146100b25780639e7a13ad146100e2575b600080fd5b610064610113565b60405161007191906102b2565b60405180910390f35b610094600480360381019061008f919061030d565b61011c565b005b6100b060048036038101906100ab9190610480565b610126565b005b6100cc60048036038101906100c791906104dc565b6101af565b6040516100d991906102b2565b60405180910390f35b6100fc60048036038101906100f7919061030d565b6101dd565b60405161010a9291906105a4565b60405180910390f35b60008054905090565b8060008190555050565b6001604051806040016040528083815260200184815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000155602082015181600101908161018591906107e0565b5050508060028360405161019991906108ee565b9081526020016040518091039020819055505050565b6002818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b600181815481106101ed57600080fd5b906000526020600020906002020160009150905080600001549080600101805461021690610603565b80601f016020809104026020016040519081016040528092919081815260200182805461024290610603565b801561028f5780601f106102645761010080835404028352916020019161028f565b820191906000526020600020905b81548152906001019060200180831161027257829003601f168201915b5050505050905082565b6000819050919050565b6102ac81610299565b82525050565b60006020820190506102c760008301846102a3565b92915050565b6000604051905090565b600080fd5b600080fd5b6102ea81610299565b81146102f557600080fd5b50565b600081359050610307816102e1565b92915050565b600060208284031215610323576103226102d7565b5b6000610331848285016102f8565b91505092915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61038d82610344565b810181811067ffffffffffffffff821117156103ac576103ab610355565b5b80604052505050565b60006103bf6102cd565b90506103cb8282610384565b919050565b600067ffffffffffffffff8211156103eb576103ea610355565b5b6103f482610344565b9050602081019050919050565b82818337600083830152505050565b600061042361041e846103d0565b6103b5565b90508281526020810184848401111561043f5761043e61033f565b5b61044a848285610401565b509392505050565b600082601f8301126104675761046661033a565b5b8135610477848260208601610410565b91505092915050565b60008060408385031215610497576104966102d7565b5b600083013567ffffffffffffffff8111156104b5576104b46102dc565b5b6104c185828601610452565b92505060206104d2858286016102f8565b9150509250929050565b6000602082840312156104f2576104f16102d7565b5b600082013567ffffffffffffffff8111156105105761050f6102dc565b5b61051c84828501610452565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561055f578082015181840152602081019050610544565b60008484015250505050565b600061057682610525565b6105808185610530565b9350610590818560208601610541565b61059981610344565b840191505092915050565b60006040820190506105b960008301856102a3565b81810360208301526105cb818461056b565b90509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061061b57607f821691505b60208210810361062e5761062d6105d4565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026106967fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610659565b6106a08683610659565b95508019841693508086168417925050509392505050565b6000819050919050565b60006106dd6106d86106d384610299565b6106b8565b610299565b9050919050565b6000819050919050565b6106f7836106c2565b61070b610703826106e4565b848454610666565b825550505050565b600090565b610720610713565b61072b8184846106ee565b505050565b5b8181101561074f57610744600082610718565b600181019050610731565b5050565b601f8211156107945761076581610634565b61076e84610649565b8101602085101561077d578190505b61079161078985610649565b830182610730565b50505b505050565b600082821c905092915050565b60006107b760001984600802610799565b1980831691505092915050565b60006107d083836107a6565b9150826002028217905092915050565b6107e982610525565b67ffffffffffffffff81111561080257610801610355565b5b61080c8254610603565b610817828285610753565b600060209050601f83116001811461084a5760008415610838578287015190505b61084285826107c4565b8655506108aa565b601f19841661085886610634565b60005b828110156108805784890151825560018201915060208501945060208101905061085b565b8683101561089d5784890151610899601f8916826107a6565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b60006108c882610525565b6108d281856108b2565b93506108e2818560208601610541565b80840191505092915050565b60006108fa82846108bd565b91508190509291505056fea26469706673582212204d872ac483d4dfe29e5799ff17f27700a362fd9ba4068705bd504302008de60d64736f6c63430008110033",
}

// SimpleStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleStorageMetaData.ABI instead.
var SimpleStorageABI = SimpleStorageMetaData.ABI

// SimpleStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleStorageMetaData.Bin instead.
var SimpleStorageBin = SimpleStorageMetaData.Bin

// DeploySimpleStorage deploys a new Ethereum contract, binding an instance of SimpleStorage to it.
func DeploySimpleStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleStorage, error) {
	parsed, err := SimpleStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleStorage{SimpleStorageCaller: SimpleStorageCaller{contract: contract}, SimpleStorageTransactor: SimpleStorageTransactor{contract: contract}, SimpleStorageFilterer: SimpleStorageFilterer{contract: contract}}, nil
}

// SimpleStorage is an auto generated Go binding around an Ethereum contract.
type SimpleStorage struct {
	SimpleStorageCaller     // Read-only binding to the contract
	SimpleStorageTransactor // Write-only binding to the contract
	SimpleStorageFilterer   // Log filterer for contract events
}

// SimpleStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleStorageSession struct {
	Contract     *SimpleStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleStorageCallerSession struct {
	Contract *SimpleStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleStorageTransactorSession struct {
	Contract     *SimpleStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleStorageRaw struct {
	Contract *SimpleStorage // Generic contract binding to access the raw methods on
}

// SimpleStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleStorageCallerRaw struct {
	Contract *SimpleStorageCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleStorageTransactorRaw struct {
	Contract *SimpleStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleStorage creates a new instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorage(address common.Address, backend bind.ContractBackend) (*SimpleStorage, error) {
	contract, err := bindSimpleStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleStorage{SimpleStorageCaller: SimpleStorageCaller{contract: contract}, SimpleStorageTransactor: SimpleStorageTransactor{contract: contract}, SimpleStorageFilterer: SimpleStorageFilterer{contract: contract}}, nil
}

// NewSimpleStorageCaller creates a new read-only instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageCaller(address common.Address, caller bind.ContractCaller) (*SimpleStorageCaller, error) {
	contract, err := bindSimpleStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageCaller{contract: contract}, nil
}

// NewSimpleStorageTransactor creates a new write-only instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleStorageTransactor, error) {
	contract, err := bindSimpleStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageTransactor{contract: contract}, nil
}

// NewSimpleStorageFilterer creates a new log filterer instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleStorageFilterer, error) {
	contract, err := bindSimpleStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageFilterer{contract: contract}, nil
}

// bindSimpleStorage binds a generic wrapper to an already deployed contract.
func bindSimpleStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleStorage *SimpleStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleStorage.Contract.SimpleStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleStorage *SimpleStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SimpleStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleStorage *SimpleStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SimpleStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleStorage *SimpleStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleStorage *SimpleStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleStorage *SimpleStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleStorage.Contract.contract.Transact(opts, method, params...)
}

// NameToFavoriteNumber is a free data retrieval call binding the contract method 0x8bab8dd5.
//
// Solidity: function nameToFavoriteNumber(string ) view returns(uint256)
func (_SimpleStorage *SimpleStorageCaller) NameToFavoriteNumber(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _SimpleStorage.contract.Call(opts, &out, "nameToFavoriteNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NameToFavoriteNumber is a free data retrieval call binding the contract method 0x8bab8dd5.
//
// Solidity: function nameToFavoriteNumber(string ) view returns(uint256)
func (_SimpleStorage *SimpleStorageSession) NameToFavoriteNumber(arg0 string) (*big.Int, error) {
	return _SimpleStorage.Contract.NameToFavoriteNumber(&_SimpleStorage.CallOpts, arg0)
}

// NameToFavoriteNumber is a free data retrieval call binding the contract method 0x8bab8dd5.
//
// Solidity: function nameToFavoriteNumber(string ) view returns(uint256)
func (_SimpleStorage *SimpleStorageCallerSession) NameToFavoriteNumber(arg0 string) (*big.Int, error) {
	return _SimpleStorage.Contract.NameToFavoriteNumber(&_SimpleStorage.CallOpts, arg0)
}

// People is a free data retrieval call binding the contract method 0x9e7a13ad.
//
// Solidity: function people(uint256 ) view returns(uint256 favoriteNumber, string name)
func (_SimpleStorage *SimpleStorageCaller) People(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FavoriteNumber *big.Int
	Name           string
}, error) {
	var out []interface{}
	err := _SimpleStorage.contract.Call(opts, &out, "people", arg0)

	outstruct := new(struct {
		FavoriteNumber *big.Int
		Name           string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FavoriteNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// People is a free data retrieval call binding the contract method 0x9e7a13ad.
//
// Solidity: function people(uint256 ) view returns(uint256 favoriteNumber, string name)
func (_SimpleStorage *SimpleStorageSession) People(arg0 *big.Int) (struct {
	FavoriteNumber *big.Int
	Name           string
}, error) {
	return _SimpleStorage.Contract.People(&_SimpleStorage.CallOpts, arg0)
}

// People is a free data retrieval call binding the contract method 0x9e7a13ad.
//
// Solidity: function people(uint256 ) view returns(uint256 favoriteNumber, string name)
func (_SimpleStorage *SimpleStorageCallerSession) People(arg0 *big.Int) (struct {
	FavoriteNumber *big.Int
	Name           string
}, error) {
	return _SimpleStorage.Contract.People(&_SimpleStorage.CallOpts, arg0)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_SimpleStorage *SimpleStorageCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleStorage.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_SimpleStorage *SimpleStorageSession) Retrieve() (*big.Int, error) {
	return _SimpleStorage.Contract.Retrieve(&_SimpleStorage.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_SimpleStorage *SimpleStorageCallerSession) Retrieve() (*big.Int, error) {
	return _SimpleStorage.Contract.Retrieve(&_SimpleStorage.CallOpts)
}

// AddPerson is a paid mutator transaction binding the contract method 0x6f760f41.
//
// Solidity: function addPerson(string _name, uint256 _favoriteNumber) returns()
func (_SimpleStorage *SimpleStorageTransactor) AddPerson(opts *bind.TransactOpts, _name string, _favoriteNumber *big.Int) (*types.Transaction, error) {
	return _SimpleStorage.contract.Transact(opts, "addPerson", _name, _favoriteNumber)
}

// AddPerson is a paid mutator transaction binding the contract method 0x6f760f41.
//
// Solidity: function addPerson(string _name, uint256 _favoriteNumber) returns()
func (_SimpleStorage *SimpleStorageSession) AddPerson(_name string, _favoriteNumber *big.Int) (*types.Transaction, error) {
	return _SimpleStorage.Contract.AddPerson(&_SimpleStorage.TransactOpts, _name, _favoriteNumber)
}

// AddPerson is a paid mutator transaction binding the contract method 0x6f760f41.
//
// Solidity: function addPerson(string _name, uint256 _favoriteNumber) returns()
func (_SimpleStorage *SimpleStorageTransactorSession) AddPerson(_name string, _favoriteNumber *big.Int) (*types.Transaction, error) {
	return _SimpleStorage.Contract.AddPerson(&_SimpleStorage.TransactOpts, _name, _favoriteNumber)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 _favoriteNumber) returns()
func (_SimpleStorage *SimpleStorageTransactor) Store(opts *bind.TransactOpts, _favoriteNumber *big.Int) (*types.Transaction, error) {
	return _SimpleStorage.contract.Transact(opts, "store", _favoriteNumber)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 _favoriteNumber) returns()
func (_SimpleStorage *SimpleStorageSession) Store(_favoriteNumber *big.Int) (*types.Transaction, error) {
	return _SimpleStorage.Contract.Store(&_SimpleStorage.TransactOpts, _favoriteNumber)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 _favoriteNumber) returns()
func (_SimpleStorage *SimpleStorageTransactorSession) Store(_favoriteNumber *big.Int) (*types.Transaction, error) {
	return _SimpleStorage.Contract.Store(&_SimpleStorage.TransactOpts, _favoriteNumber)
}

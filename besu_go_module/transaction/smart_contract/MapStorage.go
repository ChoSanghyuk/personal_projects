// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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
	_ = abi.ConvertType
)

// MapStoreMetaData contains all meta data concerning the MapStore contract.
var MapStoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_key\",\"type\":\"uint256\"}],\"name\":\"getAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_key\",\"type\":\"uint256\"}],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getPriceList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_age\",\"type\":\"uint256\"}],\"name\":\"setAgeList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_key\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setNameList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_itemName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPriceList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "608060405234801561001057600080fd5b506109cc806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630e37008a146100675780631bbefaf1146100975780635eac9098146100b35780636b8ff574146100e3578063a266f6c214610113578063fe093d3b1461012f575b600080fd5b610081600480360381019061007c91906102e6565b61014b565b60405161008e9190610322565b60405180910390f35b6100b160048036038101906100ac9190610483565b610168565b005b6100cd60048036038101906100c891906104df565b61018d565b6040516100da9190610322565b60405180910390f35b6100fd60048036038101906100f891906102e6565b6101b4565b60405161010a91906105a7565b60405180910390f35b61012d600480360381019061012891906105c9565b610259565b005b61014960048036038101906101449190610625565b610280565b005b600060026000838152602001908152602001600020549050919050565b806001600084815260200190815260200160002090816101889190610871565b505050565b6000808260405161019e919061097f565b9081526020016040518091039020549050919050565b60606001600083815260200190815260200160002080546101d490610694565b80601f016020809104026020016040519081016040528092919081815260200182805461020090610694565b801561024d5780601f106102225761010080835404028352916020019161024d565b820191906000526020600020905b81548152906001019060200180831161023057829003601f168201915b50505050509050919050565b8060008360405161026a919061097f565b9081526020016040518091039020819055505050565b8060026000848152602001908152602001600020819055505050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6102c3816102b0565b81146102ce57600080fd5b50565b6000813590506102e0816102ba565b92915050565b6000602082840312156102fc576102fb6102a6565b5b600061030a848285016102d1565b91505092915050565b61031c816102b0565b82525050565b60006020820190506103376000830184610313565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61039082610347565b810181811067ffffffffffffffff821117156103af576103ae610358565b5b80604052505050565b60006103c261029c565b90506103ce8282610387565b919050565b600067ffffffffffffffff8211156103ee576103ed610358565b5b6103f782610347565b9050602081019050919050565b82818337600083830152505050565b6000610426610421846103d3565b6103b8565b90508281526020810184848401111561044257610441610342565b5b61044d848285610404565b509392505050565b600082601f83011261046a5761046961033d565b5b813561047a848260208601610413565b91505092915050565b6000806040838503121561049a576104996102a6565b5b60006104a8858286016102d1565b925050602083013567ffffffffffffffff8111156104c9576104c86102ab565b5b6104d585828601610455565b9150509250929050565b6000602082840312156104f5576104f46102a6565b5b600082013567ffffffffffffffff811115610513576105126102ab565b5b61051f84828501610455565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610562578082015181840152602081019050610547565b60008484015250505050565b600061057982610528565b6105838185610533565b9350610593818560208601610544565b61059c81610347565b840191505092915050565b600060208201905081810360008301526105c1818461056e565b905092915050565b600080604083850312156105e0576105df6102a6565b5b600083013567ffffffffffffffff8111156105fe576105fd6102ab565b5b61060a85828601610455565b925050602061061b858286016102d1565b9150509250929050565b6000806040838503121561063c5761063b6102a6565b5b600061064a858286016102d1565b925050602061065b858286016102d1565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806106ac57607f821691505b6020821081036106bf576106be610665565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026107277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826106ea565b61073186836106ea565b95508019841693508086168417925050509392505050565b6000819050919050565b600061076e610769610764846102b0565b610749565b6102b0565b9050919050565b6000819050919050565b61078883610753565b61079c61079482610775565b8484546106f7565b825550505050565b600090565b6107b16107a4565b6107bc81848461077f565b505050565b5b818110156107e0576107d56000826107a9565b6001810190506107c2565b5050565b601f821115610825576107f6816106c5565b6107ff846106da565b8101602085101561080e578190505b61082261081a856106da565b8301826107c1565b50505b505050565b600082821c905092915050565b60006108486000198460080261082a565b1980831691505092915050565b60006108618383610837565b9150826002028217905092915050565b61087a82610528565b67ffffffffffffffff81111561089357610892610358565b5b61089d8254610694565b6108a88282856107e4565b600060209050601f8311600181146108db57600084156108c9578287015190505b6108d38582610855565b86555061093b565b601f1984166108e9866106c5565b60005b82811015610911578489015182556001820191506020850194506020810190506108ec565b8683101561092e578489015161092a601f891682610837565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b600061095982610528565b6109638185610943565b9350610973818560208601610544565b80840191505092915050565b600061098b828461094e565b91508190509291505056fea264697066735822122012167c78aeb84dbe98ce331af0560ae50eb8e31956354ffd3020f3c8be25ebf764736f6c63430008130033",
}

// MapStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use MapStoreMetaData.ABI instead.
var MapStoreABI = MapStoreMetaData.ABI

// MapStoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MapStoreMetaData.Bin instead.
var MapStoreBin = MapStoreMetaData.Bin

// DeployMapStore deploys a new Ethereum contract, binding an instance of MapStore to it.
func DeployMapStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MapStore, error) {
	parsed, err := MapStoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MapStoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MapStore{MapStoreCaller: MapStoreCaller{contract: contract}, MapStoreTransactor: MapStoreTransactor{contract: contract}, MapStoreFilterer: MapStoreFilterer{contract: contract}}, nil
}

// MapStore is an auto generated Go binding around an Ethereum contract.
type MapStore struct {
	MapStoreCaller     // Read-only binding to the contract
	MapStoreTransactor // Write-only binding to the contract
	MapStoreFilterer   // Log filterer for contract events
}

// MapStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type MapStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MapStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MapStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MapStoreSession struct {
	Contract     *MapStore         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MapStoreCallerSession struct {
	Contract *MapStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MapStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MapStoreTransactorSession struct {
	Contract     *MapStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MapStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type MapStoreRaw struct {
	Contract *MapStore // Generic contract binding to access the raw methods on
}

// MapStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MapStoreCallerRaw struct {
	Contract *MapStoreCaller // Generic read-only contract binding to access the raw methods on
}

// MapStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MapStoreTransactorRaw struct {
	Contract *MapStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMapStore creates a new instance of MapStore, bound to a specific deployed contract.
func NewMapStore(address common.Address, backend bind.ContractBackend) (*MapStore, error) {
	contract, err := bindMapStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MapStore{MapStoreCaller: MapStoreCaller{contract: contract}, MapStoreTransactor: MapStoreTransactor{contract: contract}, MapStoreFilterer: MapStoreFilterer{contract: contract}}, nil
}

// NewMapStoreCaller creates a new read-only instance of MapStore, bound to a specific deployed contract.
func NewMapStoreCaller(address common.Address, caller bind.ContractCaller) (*MapStoreCaller, error) {
	contract, err := bindMapStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MapStoreCaller{contract: contract}, nil
}

// NewMapStoreTransactor creates a new write-only instance of MapStore, bound to a specific deployed contract.
func NewMapStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*MapStoreTransactor, error) {
	contract, err := bindMapStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MapStoreTransactor{contract: contract}, nil
}

// NewMapStoreFilterer creates a new log filterer instance of MapStore, bound to a specific deployed contract.
func NewMapStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*MapStoreFilterer, error) {
	contract, err := bindMapStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MapStoreFilterer{contract: contract}, nil
}

// bindMapStore binds a generic wrapper to an already deployed contract.
func bindMapStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MapStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MapStore *MapStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MapStore.Contract.MapStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MapStore *MapStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapStore.Contract.MapStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MapStore *MapStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MapStore.Contract.MapStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MapStore *MapStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MapStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MapStore *MapStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MapStore *MapStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MapStore.Contract.contract.Transact(opts, method, params...)
}

// GetAge is a free data retrieval call binding the contract method 0x0e37008a.
//
// Solidity: function getAge(uint256 _key) view returns(uint256)
func (_MapStore *MapStoreCaller) GetAge(opts *bind.CallOpts, _key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MapStore.contract.Call(opts, &out, "getAge", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAge is a free data retrieval call binding the contract method 0x0e37008a.
//
// Solidity: function getAge(uint256 _key) view returns(uint256)
func (_MapStore *MapStoreSession) GetAge(_key *big.Int) (*big.Int, error) {
	return _MapStore.Contract.GetAge(&_MapStore.CallOpts, _key)
}

// GetAge is a free data retrieval call binding the contract method 0x0e37008a.
//
// Solidity: function getAge(uint256 _key) view returns(uint256)
func (_MapStore *MapStoreCallerSession) GetAge(_key *big.Int) (*big.Int, error) {
	return _MapStore.Contract.GetAge(&_MapStore.CallOpts, _key)
}

// GetName is a free data retrieval call binding the contract method 0x6b8ff574.
//
// Solidity: function getName(uint256 _key) view returns(string)
func (_MapStore *MapStoreCaller) GetName(opts *bind.CallOpts, _key *big.Int) (string, error) {
	var out []interface{}
	err := _MapStore.contract.Call(opts, &out, "getName", _key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x6b8ff574.
//
// Solidity: function getName(uint256 _key) view returns(string)
func (_MapStore *MapStoreSession) GetName(_key *big.Int) (string, error) {
	return _MapStore.Contract.GetName(&_MapStore.CallOpts, _key)
}

// GetName is a free data retrieval call binding the contract method 0x6b8ff574.
//
// Solidity: function getName(uint256 _key) view returns(string)
func (_MapStore *MapStoreCallerSession) GetName(_key *big.Int) (string, error) {
	return _MapStore.Contract.GetName(&_MapStore.CallOpts, _key)
}

// GetPriceList is a free data retrieval call binding the contract method 0x5eac9098.
//
// Solidity: function getPriceList(string _key) view returns(uint256)
func (_MapStore *MapStoreCaller) GetPriceList(opts *bind.CallOpts, _key string) (*big.Int, error) {
	var out []interface{}
	err := _MapStore.contract.Call(opts, &out, "getPriceList", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceList is a free data retrieval call binding the contract method 0x5eac9098.
//
// Solidity: function getPriceList(string _key) view returns(uint256)
func (_MapStore *MapStoreSession) GetPriceList(_key string) (*big.Int, error) {
	return _MapStore.Contract.GetPriceList(&_MapStore.CallOpts, _key)
}

// GetPriceList is a free data retrieval call binding the contract method 0x5eac9098.
//
// Solidity: function getPriceList(string _key) view returns(uint256)
func (_MapStore *MapStoreCallerSession) GetPriceList(_key string) (*big.Int, error) {
	return _MapStore.Contract.GetPriceList(&_MapStore.CallOpts, _key)
}

// SetAgeList is a paid mutator transaction binding the contract method 0xfe093d3b.
//
// Solidity: function setAgeList(uint256 _key, uint256 _age) returns()
func (_MapStore *MapStoreTransactor) SetAgeList(opts *bind.TransactOpts, _key *big.Int, _age *big.Int) (*types.Transaction, error) {
	return _MapStore.contract.Transact(opts, "setAgeList", _key, _age)
}

// SetAgeList is a paid mutator transaction binding the contract method 0xfe093d3b.
//
// Solidity: function setAgeList(uint256 _key, uint256 _age) returns()
func (_MapStore *MapStoreSession) SetAgeList(_key *big.Int, _age *big.Int) (*types.Transaction, error) {
	return _MapStore.Contract.SetAgeList(&_MapStore.TransactOpts, _key, _age)
}

// SetAgeList is a paid mutator transaction binding the contract method 0xfe093d3b.
//
// Solidity: function setAgeList(uint256 _key, uint256 _age) returns()
func (_MapStore *MapStoreTransactorSession) SetAgeList(_key *big.Int, _age *big.Int) (*types.Transaction, error) {
	return _MapStore.Contract.SetAgeList(&_MapStore.TransactOpts, _key, _age)
}

// SetNameList is a paid mutator transaction binding the contract method 0x1bbefaf1.
//
// Solidity: function setNameList(uint256 _key, string _name) returns()
func (_MapStore *MapStoreTransactor) SetNameList(opts *bind.TransactOpts, _key *big.Int, _name string) (*types.Transaction, error) {
	return _MapStore.contract.Transact(opts, "setNameList", _key, _name)
}

// SetNameList is a paid mutator transaction binding the contract method 0x1bbefaf1.
//
// Solidity: function setNameList(uint256 _key, string _name) returns()
func (_MapStore *MapStoreSession) SetNameList(_key *big.Int, _name string) (*types.Transaction, error) {
	return _MapStore.Contract.SetNameList(&_MapStore.TransactOpts, _key, _name)
}

// SetNameList is a paid mutator transaction binding the contract method 0x1bbefaf1.
//
// Solidity: function setNameList(uint256 _key, string _name) returns()
func (_MapStore *MapStoreTransactorSession) SetNameList(_key *big.Int, _name string) (*types.Transaction, error) {
	return _MapStore.Contract.SetNameList(&_MapStore.TransactOpts, _key, _name)
}

// SetPriceList is a paid mutator transaction binding the contract method 0xa266f6c2.
//
// Solidity: function setPriceList(string _itemName, uint256 _price) returns()
func (_MapStore *MapStoreTransactor) SetPriceList(opts *bind.TransactOpts, _itemName string, _price *big.Int) (*types.Transaction, error) {
	return _MapStore.contract.Transact(opts, "setPriceList", _itemName, _price)
}

// SetPriceList is a paid mutator transaction binding the contract method 0xa266f6c2.
//
// Solidity: function setPriceList(string _itemName, uint256 _price) returns()
func (_MapStore *MapStoreSession) SetPriceList(_itemName string, _price *big.Int) (*types.Transaction, error) {
	return _MapStore.Contract.SetPriceList(&_MapStore.TransactOpts, _itemName, _price)
}

// SetPriceList is a paid mutator transaction binding the contract method 0xa266f6c2.
//
// Solidity: function setPriceList(string _itemName, uint256 _price) returns()
func (_MapStore *MapStoreTransactorSession) SetPriceList(_itemName string, _price *big.Int) (*types.Transaction, error) {
	return _MapStore.Contract.SetPriceList(&_MapStore.TransactOpts, _itemName, _price)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dominionApp

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

// ChannelAllocation is an auto generated low-level Go binding around an user-defined struct.
type ChannelAllocation struct {
	Assets   []common.Address
	Balances [][]*big.Int
	Locked   []ChannelSubAlloc
}

// ChannelParams is an auto generated low-level Go binding around an user-defined struct.
type ChannelParams struct {
	ChallengeDuration *big.Int
	Nonce             *big.Int
	Participants      []common.Address
	App               common.Address
	LedgerChannel     bool
	VirtualChannel    bool
}

// ChannelState is an auto generated low-level Go binding around an user-defined struct.
type ChannelState struct {
	ChannelID [32]byte
	Version   uint64
	Outcome   ChannelAllocation
	AppData   []byte
	IsFinal   bool
}

// ChannelSubAlloc is an auto generated low-level Go binding around an user-defined struct.
type ChannelSubAlloc struct {
	ID       [32]byte
	Balances []*big.Int
	IndexMap []uint16
}

// AppMetaData contains all meta data concerning the App contract.
var AppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ledgerChannel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualChannel\",\"type\":\"bool\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"to\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"actorIdx\",\"type\":\"uint256\"}],\"name\":\"validTransition\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0d1feb4f": "validTransition((uint256,uint256,address[],address,bool,bool),(bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),(bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),uint256)",
	},
}

// AppABI is the input ABI used to generate the binding from.
// Deprecated: Use AppMetaData.ABI instead.
var AppABI = AppMetaData.ABI

// Deprecated: Use AppMetaData.Sigs instead.
// AppFuncSigs maps the 4-byte function signature to its string representation.
var AppFuncSigs = AppMetaData.Sigs

// App is an auto generated Go binding around an Ethereum contract.
type App struct {
	AppCaller     // Read-only binding to the contract
	AppTransactor // Write-only binding to the contract
	AppFilterer   // Log filterer for contract events
}

// AppCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppSession struct {
	Contract     *App              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppCallerSession struct {
	Contract *AppCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppTransactorSession struct {
	Contract     *AppTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppRaw struct {
	Contract *App // Generic contract binding to access the raw methods on
}

// AppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppCallerRaw struct {
	Contract *AppCaller // Generic read-only contract binding to access the raw methods on
}

// AppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppTransactorRaw struct {
	Contract *AppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApp creates a new instance of App, bound to a specific deployed contract.
func NewApp(address common.Address, backend bind.ContractBackend) (*App, error) {
	contract, err := bindApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &App{AppCaller: AppCaller{contract: contract}, AppTransactor: AppTransactor{contract: contract}, AppFilterer: AppFilterer{contract: contract}}, nil
}

// NewAppCaller creates a new read-only instance of App, bound to a specific deployed contract.
func NewAppCaller(address common.Address, caller bind.ContractCaller) (*AppCaller, error) {
	contract, err := bindApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppCaller{contract: contract}, nil
}

// NewAppTransactor creates a new write-only instance of App, bound to a specific deployed contract.
func NewAppTransactor(address common.Address, transactor bind.ContractTransactor) (*AppTransactor, error) {
	contract, err := bindApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppTransactor{contract: contract}, nil
}

// NewAppFilterer creates a new log filterer instance of App, bound to a specific deployed contract.
func NewAppFilterer(address common.Address, filterer bind.ContractFilterer) (*AppFilterer, error) {
	contract, err := bindApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppFilterer{contract: contract}, nil
}

// bindApp binds a generic wrapper to an already deployed contract.
func bindApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_App *AppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _App.Contract.AppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_App *AppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _App.Contract.AppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_App *AppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _App.Contract.AppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_App *AppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _App.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_App *AppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _App.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_App *AppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _App.Contract.contract.Transact(opts, method, params...)
}

// ValidTransition is a free data retrieval call binding the contract method 0x0d1feb4f.
//
// Solidity: function validTransition((uint256,uint256,address[],address,bool,bool) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) from, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) to, uint256 actorIdx) pure returns()
func (_App *AppCaller) ValidTransition(opts *bind.CallOpts, params ChannelParams, from ChannelState, to ChannelState, actorIdx *big.Int) error {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "validTransition", params, from, to, actorIdx)

	if err != nil {
		return err
	}

	return err

}

// ValidTransition is a free data retrieval call binding the contract method 0x0d1feb4f.
//
// Solidity: function validTransition((uint256,uint256,address[],address,bool,bool) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) from, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) to, uint256 actorIdx) pure returns()
func (_App *AppSession) ValidTransition(params ChannelParams, from ChannelState, to ChannelState, actorIdx *big.Int) error {
	return _App.Contract.ValidTransition(&_App.CallOpts, params, from, to, actorIdx)
}

// ValidTransition is a free data retrieval call binding the contract method 0x0d1feb4f.
//
// Solidity: function validTransition((uint256,uint256,address[],address,bool,bool) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) from, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) to, uint256 actorIdx) pure returns()
func (_App *AppCallerSession) ValidTransition(params ChannelParams, from ChannelState, to ChannelState, actorIdx *big.Int) error {
	return _App.Contract.ValidTransition(&_App.CallOpts, params, from, to, actorIdx)
}

// ArrayMetaData contains all meta data concerning the Array contract.
var ArrayMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a582595d42977fa7d3db10e23e726e6f19611d0555df6a35e6d2d923f9fa37cd64736f6c63430007060033",
}

// ArrayABI is the input ABI used to generate the binding from.
// Deprecated: Use ArrayMetaData.ABI instead.
var ArrayABI = ArrayMetaData.ABI

// ArrayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ArrayMetaData.Bin instead.
var ArrayBin = ArrayMetaData.Bin

// DeployArray deploys a new Ethereum contract, binding an instance of Array to it.
func DeployArray(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Array, error) {
	parsed, err := ArrayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArrayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Array{ArrayCaller: ArrayCaller{contract: contract}, ArrayTransactor: ArrayTransactor{contract: contract}, ArrayFilterer: ArrayFilterer{contract: contract}}, nil
}

// Array is an auto generated Go binding around an Ethereum contract.
type Array struct {
	ArrayCaller     // Read-only binding to the contract
	ArrayTransactor // Write-only binding to the contract
	ArrayFilterer   // Log filterer for contract events
}

// ArrayCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArrayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArrayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArrayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArraySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArraySession struct {
	Contract     *Array            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArrayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArrayCallerSession struct {
	Contract *ArrayCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArrayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArrayTransactorSession struct {
	Contract     *ArrayTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArrayRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArrayRaw struct {
	Contract *Array // Generic contract binding to access the raw methods on
}

// ArrayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArrayCallerRaw struct {
	Contract *ArrayCaller // Generic read-only contract binding to access the raw methods on
}

// ArrayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArrayTransactorRaw struct {
	Contract *ArrayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArray creates a new instance of Array, bound to a specific deployed contract.
func NewArray(address common.Address, backend bind.ContractBackend) (*Array, error) {
	contract, err := bindArray(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Array{ArrayCaller: ArrayCaller{contract: contract}, ArrayTransactor: ArrayTransactor{contract: contract}, ArrayFilterer: ArrayFilterer{contract: contract}}, nil
}

// NewArrayCaller creates a new read-only instance of Array, bound to a specific deployed contract.
func NewArrayCaller(address common.Address, caller bind.ContractCaller) (*ArrayCaller, error) {
	contract, err := bindArray(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArrayCaller{contract: contract}, nil
}

// NewArrayTransactor creates a new write-only instance of Array, bound to a specific deployed contract.
func NewArrayTransactor(address common.Address, transactor bind.ContractTransactor) (*ArrayTransactor, error) {
	contract, err := bindArray(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArrayTransactor{contract: contract}, nil
}

// NewArrayFilterer creates a new log filterer instance of Array, bound to a specific deployed contract.
func NewArrayFilterer(address common.Address, filterer bind.ContractFilterer) (*ArrayFilterer, error) {
	contract, err := bindArray(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArrayFilterer{contract: contract}, nil
}

// bindArray binds a generic wrapper to an already deployed contract.
func bindArray(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArrayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Array *ArrayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Array.Contract.ArrayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Array *ArrayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Array.Contract.ArrayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Array *ArrayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Array.Contract.ArrayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Array *ArrayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Array.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Array *ArrayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Array.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Array *ArrayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Array.Contract.contract.Transact(opts, method, params...)
}

// CardLibMetaData contains all meta data concerning the CardLib contract.
var CardLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122092696a4e16bc44ae84689a1892eccb0ebb3ca1afe543b0bf733f1b9672ac0f0d64736f6c63430007060033",
}

// CardLibABI is the input ABI used to generate the binding from.
// Deprecated: Use CardLibMetaData.ABI instead.
var CardLibABI = CardLibMetaData.ABI

// CardLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CardLibMetaData.Bin instead.
var CardLibBin = CardLibMetaData.Bin

// DeployCardLib deploys a new Ethereum contract, binding an instance of CardLib to it.
func DeployCardLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CardLib, error) {
	parsed, err := CardLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CardLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CardLib{CardLibCaller: CardLibCaller{contract: contract}, CardLibTransactor: CardLibTransactor{contract: contract}, CardLibFilterer: CardLibFilterer{contract: contract}}, nil
}

// CardLib is an auto generated Go binding around an Ethereum contract.
type CardLib struct {
	CardLibCaller     // Read-only binding to the contract
	CardLibTransactor // Write-only binding to the contract
	CardLibFilterer   // Log filterer for contract events
}

// CardLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type CardLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CardLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CardLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CardLibSession struct {
	Contract     *CardLib          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CardLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CardLibCallerSession struct {
	Contract *CardLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CardLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CardLibTransactorSession struct {
	Contract     *CardLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CardLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type CardLibRaw struct {
	Contract *CardLib // Generic contract binding to access the raw methods on
}

// CardLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CardLibCallerRaw struct {
	Contract *CardLibCaller // Generic read-only contract binding to access the raw methods on
}

// CardLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CardLibTransactorRaw struct {
	Contract *CardLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCardLib creates a new instance of CardLib, bound to a specific deployed contract.
func NewCardLib(address common.Address, backend bind.ContractBackend) (*CardLib, error) {
	contract, err := bindCardLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CardLib{CardLibCaller: CardLibCaller{contract: contract}, CardLibTransactor: CardLibTransactor{contract: contract}, CardLibFilterer: CardLibFilterer{contract: contract}}, nil
}

// NewCardLibCaller creates a new read-only instance of CardLib, bound to a specific deployed contract.
func NewCardLibCaller(address common.Address, caller bind.ContractCaller) (*CardLibCaller, error) {
	contract, err := bindCardLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CardLibCaller{contract: contract}, nil
}

// NewCardLibTransactor creates a new write-only instance of CardLib, bound to a specific deployed contract.
func NewCardLibTransactor(address common.Address, transactor bind.ContractTransactor) (*CardLibTransactor, error) {
	contract, err := bindCardLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CardLibTransactor{contract: contract}, nil
}

// NewCardLibFilterer creates a new log filterer instance of CardLib, bound to a specific deployed contract.
func NewCardLibFilterer(address common.Address, filterer bind.ContractFilterer) (*CardLibFilterer, error) {
	contract, err := bindCardLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CardLibFilterer{contract: contract}, nil
}

// bindCardLib binds a generic wrapper to an already deployed contract.
func bindCardLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CardLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardLib *CardLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CardLib.Contract.CardLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardLib *CardLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardLib.Contract.CardLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardLib *CardLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardLib.Contract.CardLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardLib *CardLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CardLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardLib *CardLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardLib *CardLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardLib.Contract.contract.Transact(opts, method, params...)
}

// ChannelMetaData contains all meta data concerning the Channel contract.
var ChannelMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f93f0d8575f02336f16b9622fb98f990fafade4dad5f855e281c82cc1e7be50364736f6c63430007060033",
}

// ChannelABI is the input ABI used to generate the binding from.
// Deprecated: Use ChannelMetaData.ABI instead.
var ChannelABI = ChannelMetaData.ABI

// ChannelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChannelMetaData.Bin instead.
var ChannelBin = ChannelMetaData.Bin

// DeployChannel deploys a new Ethereum contract, binding an instance of Channel to it.
func DeployChannel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Channel, error) {
	parsed, err := ChannelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChannelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Channel{ChannelCaller: ChannelCaller{contract: contract}, ChannelTransactor: ChannelTransactor{contract: contract}, ChannelFilterer: ChannelFilterer{contract: contract}}, nil
}

// Channel is an auto generated Go binding around an Ethereum contract.
type Channel struct {
	ChannelCaller     // Read-only binding to the contract
	ChannelTransactor // Write-only binding to the contract
	ChannelFilterer   // Log filterer for contract events
}

// ChannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChannelSession struct {
	Contract     *Channel          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChannelCallerSession struct {
	Contract *ChannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ChannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChannelTransactorSession struct {
	Contract     *ChannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ChannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChannelRaw struct {
	Contract *Channel // Generic contract binding to access the raw methods on
}

// ChannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChannelCallerRaw struct {
	Contract *ChannelCaller // Generic read-only contract binding to access the raw methods on
}

// ChannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChannelTransactorRaw struct {
	Contract *ChannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChannel creates a new instance of Channel, bound to a specific deployed contract.
func NewChannel(address common.Address, backend bind.ContractBackend) (*Channel, error) {
	contract, err := bindChannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Channel{ChannelCaller: ChannelCaller{contract: contract}, ChannelTransactor: ChannelTransactor{contract: contract}, ChannelFilterer: ChannelFilterer{contract: contract}}, nil
}

// NewChannelCaller creates a new read-only instance of Channel, bound to a specific deployed contract.
func NewChannelCaller(address common.Address, caller bind.ContractCaller) (*ChannelCaller, error) {
	contract, err := bindChannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelCaller{contract: contract}, nil
}

// NewChannelTransactor creates a new write-only instance of Channel, bound to a specific deployed contract.
func NewChannelTransactor(address common.Address, transactor bind.ContractTransactor) (*ChannelTransactor, error) {
	contract, err := bindChannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelTransactor{contract: contract}, nil
}

// NewChannelFilterer creates a new log filterer instance of Channel, bound to a specific deployed contract.
func NewChannelFilterer(address common.Address, filterer bind.ContractFilterer) (*ChannelFilterer, error) {
	contract, err := bindChannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChannelFilterer{contract: contract}, nil
}

// bindChannel binds a generic wrapper to an already deployed contract.
func bindChannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Channel *ChannelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Channel.Contract.ChannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Channel *ChannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Channel.Contract.ChannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Channel *ChannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Channel.Contract.ChannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Channel *ChannelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Channel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Channel *ChannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Channel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Channel *ChannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Channel.Contract.contract.Transact(opts, method, params...)
}

// ConstantMetaData contains all meta data concerning the Constant contract.
var ConstantMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207f90049c02018178f1cdd7dfe76a24facc961024a236ed63488e7a2af19b674464736f6c63430007060033",
}

// ConstantABI is the input ABI used to generate the binding from.
// Deprecated: Use ConstantMetaData.ABI instead.
var ConstantABI = ConstantMetaData.ABI

// ConstantBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConstantMetaData.Bin instead.
var ConstantBin = ConstantMetaData.Bin

// DeployConstant deploys a new Ethereum contract, binding an instance of Constant to it.
func DeployConstant(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Constant, error) {
	parsed, err := ConstantMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConstantBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Constant{ConstantCaller: ConstantCaller{contract: contract}, ConstantTransactor: ConstantTransactor{contract: contract}, ConstantFilterer: ConstantFilterer{contract: contract}}, nil
}

// Constant is an auto generated Go binding around an Ethereum contract.
type Constant struct {
	ConstantCaller     // Read-only binding to the contract
	ConstantTransactor // Write-only binding to the contract
	ConstantFilterer   // Log filterer for contract events
}

// ConstantCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConstantCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConstantTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConstantFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConstantSession struct {
	Contract     *Constant         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConstantCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConstantCallerSession struct {
	Contract *ConstantCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ConstantTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConstantTransactorSession struct {
	Contract     *ConstantTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ConstantRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConstantRaw struct {
	Contract *Constant // Generic contract binding to access the raw methods on
}

// ConstantCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConstantCallerRaw struct {
	Contract *ConstantCaller // Generic read-only contract binding to access the raw methods on
}

// ConstantTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConstantTransactorRaw struct {
	Contract *ConstantTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConstant creates a new instance of Constant, bound to a specific deployed contract.
func NewConstant(address common.Address, backend bind.ContractBackend) (*Constant, error) {
	contract, err := bindConstant(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Constant{ConstantCaller: ConstantCaller{contract: contract}, ConstantTransactor: ConstantTransactor{contract: contract}, ConstantFilterer: ConstantFilterer{contract: contract}}, nil
}

// NewConstantCaller creates a new read-only instance of Constant, bound to a specific deployed contract.
func NewConstantCaller(address common.Address, caller bind.ContractCaller) (*ConstantCaller, error) {
	contract, err := bindConstant(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConstantCaller{contract: contract}, nil
}

// NewConstantTransactor creates a new write-only instance of Constant, bound to a specific deployed contract.
func NewConstantTransactor(address common.Address, transactor bind.ContractTransactor) (*ConstantTransactor, error) {
	contract, err := bindConstant(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConstantTransactor{contract: contract}, nil
}

// NewConstantFilterer creates a new log filterer instance of Constant, bound to a specific deployed contract.
func NewConstantFilterer(address common.Address, filterer bind.ContractFilterer) (*ConstantFilterer, error) {
	contract, err := bindConstant(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConstantFilterer{contract: contract}, nil
}

// bindConstant binds a generic wrapper to an already deployed contract.
func bindConstant(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConstantABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constant *ConstantRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Constant.Contract.ConstantCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constant *ConstantRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constant.Contract.ConstantTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constant *ConstantRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constant.Contract.ConstantTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constant *ConstantCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Constant.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constant *ConstantTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constant.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constant *ConstantTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constant.Contract.contract.Transact(opts, method, params...)
}

// ConvertMetaData contains all meta data concerning the Convert contract.
var ConvertMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122049c52c39066175bf59d20edd30b267af213f72a320ce66aa2993f105f090558d64736f6c63430007060033",
}

// ConvertABI is the input ABI used to generate the binding from.
// Deprecated: Use ConvertMetaData.ABI instead.
var ConvertABI = ConvertMetaData.ABI

// ConvertBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConvertMetaData.Bin instead.
var ConvertBin = ConvertMetaData.Bin

// DeployConvert deploys a new Ethereum contract, binding an instance of Convert to it.
func DeployConvert(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Convert, error) {
	parsed, err := ConvertMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConvertBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Convert{ConvertCaller: ConvertCaller{contract: contract}, ConvertTransactor: ConvertTransactor{contract: contract}, ConvertFilterer: ConvertFilterer{contract: contract}}, nil
}

// Convert is an auto generated Go binding around an Ethereum contract.
type Convert struct {
	ConvertCaller     // Read-only binding to the contract
	ConvertTransactor // Write-only binding to the contract
	ConvertFilterer   // Log filterer for contract events
}

// ConvertCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConvertCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvertTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConvertTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvertFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConvertFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvertSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConvertSession struct {
	Contract     *Convert          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConvertCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConvertCallerSession struct {
	Contract *ConvertCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ConvertTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConvertTransactorSession struct {
	Contract     *ConvertTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConvertRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConvertRaw struct {
	Contract *Convert // Generic contract binding to access the raw methods on
}

// ConvertCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConvertCallerRaw struct {
	Contract *ConvertCaller // Generic read-only contract binding to access the raw methods on
}

// ConvertTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConvertTransactorRaw struct {
	Contract *ConvertTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConvert creates a new instance of Convert, bound to a specific deployed contract.
func NewConvert(address common.Address, backend bind.ContractBackend) (*Convert, error) {
	contract, err := bindConvert(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Convert{ConvertCaller: ConvertCaller{contract: contract}, ConvertTransactor: ConvertTransactor{contract: contract}, ConvertFilterer: ConvertFilterer{contract: contract}}, nil
}

// NewConvertCaller creates a new read-only instance of Convert, bound to a specific deployed contract.
func NewConvertCaller(address common.Address, caller bind.ContractCaller) (*ConvertCaller, error) {
	contract, err := bindConvert(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConvertCaller{contract: contract}, nil
}

// NewConvertTransactor creates a new write-only instance of Convert, bound to a specific deployed contract.
func NewConvertTransactor(address common.Address, transactor bind.ContractTransactor) (*ConvertTransactor, error) {
	contract, err := bindConvert(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConvertTransactor{contract: contract}, nil
}

// NewConvertFilterer creates a new log filterer instance of Convert, bound to a specific deployed contract.
func NewConvertFilterer(address common.Address, filterer bind.ContractFilterer) (*ConvertFilterer, error) {
	contract, err := bindConvert(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConvertFilterer{contract: contract}, nil
}

// bindConvert binds a generic wrapper to an already deployed contract.
func bindConvert(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConvertABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Convert *ConvertRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Convert.Contract.ConvertCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Convert *ConvertRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Convert.Contract.ConvertTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Convert *ConvertRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Convert.Contract.ConvertTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Convert *ConvertCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Convert.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Convert *ConvertTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Convert.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Convert *ConvertTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Convert.Contract.contract.Transact(opts, method, params...)
}

// DataLibMetaData contains all meta data concerning the DataLib contract.
var DataLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122004d8cec990397880c2baee0365e6b6731b1c4f17d603d4b37a3b7c82eaccbf6864736f6c63430007060033",
}

// DataLibABI is the input ABI used to generate the binding from.
// Deprecated: Use DataLibMetaData.ABI instead.
var DataLibABI = DataLibMetaData.ABI

// DataLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DataLibMetaData.Bin instead.
var DataLibBin = DataLibMetaData.Bin

// DeployDataLib deploys a new Ethereum contract, binding an instance of DataLib to it.
func DeployDataLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataLib, error) {
	parsed, err := DataLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DataLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataLib{DataLibCaller: DataLibCaller{contract: contract}, DataLibTransactor: DataLibTransactor{contract: contract}, DataLibFilterer: DataLibFilterer{contract: contract}}, nil
}

// DataLib is an auto generated Go binding around an Ethereum contract.
type DataLib struct {
	DataLibCaller     // Read-only binding to the contract
	DataLibTransactor // Write-only binding to the contract
	DataLibFilterer   // Log filterer for contract events
}

// DataLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataLibSession struct {
	Contract     *DataLib          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataLibCallerSession struct {
	Contract *DataLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DataLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataLibTransactorSession struct {
	Contract     *DataLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DataLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataLibRaw struct {
	Contract *DataLib // Generic contract binding to access the raw methods on
}

// DataLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataLibCallerRaw struct {
	Contract *DataLibCaller // Generic read-only contract binding to access the raw methods on
}

// DataLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataLibTransactorRaw struct {
	Contract *DataLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataLib creates a new instance of DataLib, bound to a specific deployed contract.
func NewDataLib(address common.Address, backend bind.ContractBackend) (*DataLib, error) {
	contract, err := bindDataLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataLib{DataLibCaller: DataLibCaller{contract: contract}, DataLibTransactor: DataLibTransactor{contract: contract}, DataLibFilterer: DataLibFilterer{contract: contract}}, nil
}

// NewDataLibCaller creates a new read-only instance of DataLib, bound to a specific deployed contract.
func NewDataLibCaller(address common.Address, caller bind.ContractCaller) (*DataLibCaller, error) {
	contract, err := bindDataLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataLibCaller{contract: contract}, nil
}

// NewDataLibTransactor creates a new write-only instance of DataLib, bound to a specific deployed contract.
func NewDataLibTransactor(address common.Address, transactor bind.ContractTransactor) (*DataLibTransactor, error) {
	contract, err := bindDataLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataLibTransactor{contract: contract}, nil
}

// NewDataLibFilterer creates a new log filterer instance of DataLib, bound to a specific deployed contract.
func NewDataLibFilterer(address common.Address, filterer bind.ContractFilterer) (*DataLibFilterer, error) {
	contract, err := bindDataLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataLibFilterer{contract: contract}, nil
}

// bindDataLib binds a generic wrapper to an already deployed contract.
func bindDataLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataLib *DataLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataLib.Contract.DataLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataLib *DataLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataLib.Contract.DataLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataLib *DataLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataLib.Contract.DataLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataLib *DataLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataLib *DataLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataLib *DataLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataLib.Contract.contract.Transact(opts, method, params...)
}

// DeckLibMetaData contains all meta data concerning the DeckLib contract.
var DeckLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122094a8b5d035a43a2bd13df780ff092e1c702e11cee2ee0f474892063439ed6d2364736f6c63430007060033",
}

// DeckLibABI is the input ABI used to generate the binding from.
// Deprecated: Use DeckLibMetaData.ABI instead.
var DeckLibABI = DeckLibMetaData.ABI

// DeckLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DeckLibMetaData.Bin instead.
var DeckLibBin = DeckLibMetaData.Bin

// DeployDeckLib deploys a new Ethereum contract, binding an instance of DeckLib to it.
func DeployDeckLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DeckLib, error) {
	parsed, err := DeckLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DeckLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DeckLib{DeckLibCaller: DeckLibCaller{contract: contract}, DeckLibTransactor: DeckLibTransactor{contract: contract}, DeckLibFilterer: DeckLibFilterer{contract: contract}}, nil
}

// DeckLib is an auto generated Go binding around an Ethereum contract.
type DeckLib struct {
	DeckLibCaller     // Read-only binding to the contract
	DeckLibTransactor // Write-only binding to the contract
	DeckLibFilterer   // Log filterer for contract events
}

// DeckLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeckLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeckLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeckLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeckLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeckLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeckLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeckLibSession struct {
	Contract     *DeckLib          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeckLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeckLibCallerSession struct {
	Contract *DeckLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DeckLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeckLibTransactorSession struct {
	Contract     *DeckLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DeckLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeckLibRaw struct {
	Contract *DeckLib // Generic contract binding to access the raw methods on
}

// DeckLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeckLibCallerRaw struct {
	Contract *DeckLibCaller // Generic read-only contract binding to access the raw methods on
}

// DeckLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeckLibTransactorRaw struct {
	Contract *DeckLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeckLib creates a new instance of DeckLib, bound to a specific deployed contract.
func NewDeckLib(address common.Address, backend bind.ContractBackend) (*DeckLib, error) {
	contract, err := bindDeckLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeckLib{DeckLibCaller: DeckLibCaller{contract: contract}, DeckLibTransactor: DeckLibTransactor{contract: contract}, DeckLibFilterer: DeckLibFilterer{contract: contract}}, nil
}

// NewDeckLibCaller creates a new read-only instance of DeckLib, bound to a specific deployed contract.
func NewDeckLibCaller(address common.Address, caller bind.ContractCaller) (*DeckLibCaller, error) {
	contract, err := bindDeckLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeckLibCaller{contract: contract}, nil
}

// NewDeckLibTransactor creates a new write-only instance of DeckLib, bound to a specific deployed contract.
func NewDeckLibTransactor(address common.Address, transactor bind.ContractTransactor) (*DeckLibTransactor, error) {
	contract, err := bindDeckLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeckLibTransactor{contract: contract}, nil
}

// NewDeckLibFilterer creates a new log filterer instance of DeckLib, bound to a specific deployed contract.
func NewDeckLibFilterer(address common.Address, filterer bind.ContractFilterer) (*DeckLibFilterer, error) {
	contract, err := bindDeckLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeckLibFilterer{contract: contract}, nil
}

// bindDeckLib binds a generic wrapper to an already deployed contract.
func bindDeckLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeckLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeckLib *DeckLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeckLib.Contract.DeckLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeckLib *DeckLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeckLib.Contract.DeckLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeckLib *DeckLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeckLib.Contract.DeckLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeckLib *DeckLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeckLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeckLib *DeckLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeckLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeckLib *DeckLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeckLib.Contract.contract.Transact(opts, method, params...)
}

// DominionAppMetaData contains all meta data concerning the DominionApp contract.
var DominionAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ledgerChannel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualChannel\",\"type\":\"bool\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"to\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"signerIdx\",\"type\":\"uint256\"}],\"name\":\"validTransition\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0d1feb4f": "validTransition((uint256,uint256,address[],address,bool,bool),(bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),(bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610cdd806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80630d1feb4f14610030575b600080fd5b61004361003e366004610bc8565b610045565b005b6000604051806020016040528061009f8680606001906100659190610c5b565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061013392505050565b9052905060006100ae826101e5565b9050600060405180602001604052806100d08880606001906100659190610c5b565b9052905060006100df826101e5565b9050600060405180602001604052806101018980606001906100659190610c5b565b905290506000610110826101e5565b90506001815160200151600881111561012557fe5b505050505050505050505050565b60606000825167ffffffffffffffff8111801561014f57600080fd5b50604051908082528060200260200182016040528015610179578160200160208202803683370190505b50905060005b83518163ffffffff1610156101de57838163ffffffff16815181106101a057fe5b602001015160f81c60f81b828263ffffffff16815181106101bd57fe5b6001600160f81b03199092166020928302919091019091015260010161017f565b5092915050565b6101ed610aa4565b60006101f883610290565b90506000610205846102bb565b6040805160028082526060820190925291925060009190816020015b610229610add565b81526020019060019003908161022157905050905060005b600281101561025c57610253866102da565b50600101610241565b506000610268866102f9565b6040805160808101825295865260208601949094529284019190915250606082015292915050565b610298610b24565b60006102b46102af846102aa86610318565b61036c565b6103a2565b9392505050565b6102c3610b4d565b60006102b46102d5846102aa86610318565b6104b9565b6102e2610add565b60006102b46102f4846102aa86610318565b610540565b610301610b60565b60006102b4610313846102aa86610318565b61073b565b600080826000015160008151811061032c57fe5b6020026020010151905061034b836000015160018560000151516107ba565b808452805160009061035957fe5b602002602001015160f81c915050919050565b60606000610383846000015160008560ff166107ba565b845180519192506103989160ff8616906107ba565b8452905092915050565b6103aa610b24565b6040805160088082526101208201909252600091602082016101008036833701905050905060005b60038451038160ff16101561042a57610403848260030160ff16815181106103f657fe5b6020026020010151610865565b828260ff168151811061041257fe5b911515602092830291909101909101526001016103d2565b5060405180608001604052808460008151811061044357fe5b602002602001015160f81c60ff1681526020018460018151811061046357fe5b602002602001015160f81c60ff16600881111561047c57fe5b600881111561048757fe5b81526020016104a861049c86600260036107ba565b6000815181106103f657fe5b151581526020019190915292915050565b6104c1610b4d565b6040805160066020820181815261010083018452600093839290830160c0803683375050509052905060005b600660ff821610156101de57838160ff168151811061050857fe5b602002602001015160f81c82600001518260ff168151811061052657fe5b60ff909216602092830291909101909101526001016104ed565b610548610add565b60008260008151811061055757fe5b602002602001015160f81c9050600061058061057b8560018560010160ff166107ba565b61086e565b9050610594848360010160ff1686516107ba565b93506000846000815181106105a557fe5b602002602001015160f81c905060006105c961057b8760018560010160ff166107ba565b90506105dd868360010160ff1688516107ba565b95506000866000815181106105ee57fe5b602002602001015160f81c9050600061061261057b8960018560010160ff166107ba565b9050610626888360010160ff168a516107ba565b975060008860008151811061063757fe5b602002602001015160f81c9050600061065b61057b8b60018560010160ff166107ba565b905061066f8a8560010160ff168c516107ba565b60408051600480825260a08201909252919b50600091906020820160808036833701905050905060008b6000815181106106a557fe5b602002602001015160f81c905060005b8160ff168160ff161015610709578c8160010160ff16815181106106d557fe5b602002602001015160f81c838260ff16815181106106ef57fe5b60ff909216602092830291909101909101526001016106b5565b50506040805160a081018252988952602089019690965294870192909252506060850152506080830152509392505050565b610743610b60565b815160608080601460ff8516106107645761076186600060146107ba565b92505b602860ff85161061077f5761077c86601460286107ba565b91505b603c60ff85161061079a57610797866028603c6107ba565b90505b604080516060810182529384526020840192909252908201529392505050565b6060600083830367ffffffffffffffff811180156107d757600080fd5b50604051908082528060200260200182016040528015610801578160200160208202803683370190505b50905060005b81518160ff16101561085c57858160ff1686018151811061082457fe5b6020026020010151828260ff168151811061083b57fe5b6001600160f81b031990921660209283029190910190910152600101610807565b50949350505050565b60f81c60011490565b610876610b4d565b6000825167ffffffffffffffff8111801561089057600080fd5b506040519080825280602002602001820160405280156108ca57816020015b6108b7610b81565b8152602001906001900390816108af5790505b50905060005b83518160ff16101561091a576108f86108f3858360ff168460010160ff166107ba565b61092f565b828260ff168151811061090757fe5b60209081029190910101526001016108d0565b50604080516020810190915290815292915050565b610937610b81565b6000808360008151811061094757fe5b602002602001015160f81c90506000806000806000600581111561096757fe5b60ff168560ff1614156109895760009550600193506000925060009150610a5b565b600186600581111561099757fe5b60ff1614156109b55760019550600293506000925060019150610a5b565b60028660058111156109c357fe5b60ff1614156109e15760029550600393506000925060029150610a5b565b60038660058111156109ef57fe5b60ff161415610a0b575060039450600091506001905080610a5b565b6004866005811115610a1957fe5b60ff161415610a35575060049450600091506002905080610a5b565b60058681811115610a4257fe5b60ff161415610a5b575060059450600091506006905060035b60006040518060a00160405280886005811115610a7457fe5b815260ff968716602082015292861660408401529385166060830152509216608090920191909152949350505050565b6040518060800160405280610ab7610b24565b8152602001610ac4610b4d565b815260200160608152602001610ad8610b60565b905290565b6040518060a00160405280610af0610b4d565b8152602001610afd610b4d565b8152602001610b0a610b4d565b8152602001610b17610b4d565b8152602001606081525090565b604080516080810190915260008082526020820190815260006020820152606060409091015290565b6040518060200160405280606081525090565b60405180606001604052806060815260200160608152602001606081525090565b6040805160a081019091528060008152600060208201819052604082018190526060820181905260809091015290565b600060a08284031215610bc2578081fd5b50919050565b60008060008060808587031215610bdd578384fd5b843567ffffffffffffffff80821115610bf4578586fd5b9086019060c08289031215610c07578586fd5b90945060208601359080821115610c1c578485fd5b610c2888838901610bb1565b94506040870135915080821115610c3d578384fd5b50610c4a87828801610bb1565b949793965093946060013593505050565b6000808335601e19843603018112610c71578283fd5b83018035915067ffffffffffffffff821115610c8b578283fd5b602001915036819003821315610ca057600080fd5b925092905056fea264697066735822122024231772a7931c96ec54b3d5df5d6bd75200109edc162165d59502851d2904db64736f6c63430007060033",
}

// DominionAppABI is the input ABI used to generate the binding from.
// Deprecated: Use DominionAppMetaData.ABI instead.
var DominionAppABI = DominionAppMetaData.ABI

// Deprecated: Use DominionAppMetaData.Sigs instead.
// DominionAppFuncSigs maps the 4-byte function signature to its string representation.
var DominionAppFuncSigs = DominionAppMetaData.Sigs

// DominionAppBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DominionAppMetaData.Bin instead.
var DominionAppBin = DominionAppMetaData.Bin

// DeployDominionApp deploys a new Ethereum contract, binding an instance of DominionApp to it.
func DeployDominionApp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DominionApp, error) {
	parsed, err := DominionAppMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DominionAppBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DominionApp{DominionAppCaller: DominionAppCaller{contract: contract}, DominionAppTransactor: DominionAppTransactor{contract: contract}, DominionAppFilterer: DominionAppFilterer{contract: contract}}, nil
}

// DominionApp is an auto generated Go binding around an Ethereum contract.
type DominionApp struct {
	DominionAppCaller     // Read-only binding to the contract
	DominionAppTransactor // Write-only binding to the contract
	DominionAppFilterer   // Log filterer for contract events
}

// DominionAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type DominionAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DominionAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DominionAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DominionAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DominionAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DominionAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DominionAppSession struct {
	Contract     *DominionApp      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DominionAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DominionAppCallerSession struct {
	Contract *DominionAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DominionAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DominionAppTransactorSession struct {
	Contract     *DominionAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DominionAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type DominionAppRaw struct {
	Contract *DominionApp // Generic contract binding to access the raw methods on
}

// DominionAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DominionAppCallerRaw struct {
	Contract *DominionAppCaller // Generic read-only contract binding to access the raw methods on
}

// DominionAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DominionAppTransactorRaw struct {
	Contract *DominionAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDominionApp creates a new instance of DominionApp, bound to a specific deployed contract.
func NewDominionApp(address common.Address, backend bind.ContractBackend) (*DominionApp, error) {
	contract, err := bindDominionApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DominionApp{DominionAppCaller: DominionAppCaller{contract: contract}, DominionAppTransactor: DominionAppTransactor{contract: contract}, DominionAppFilterer: DominionAppFilterer{contract: contract}}, nil
}

// NewDominionAppCaller creates a new read-only instance of DominionApp, bound to a specific deployed contract.
func NewDominionAppCaller(address common.Address, caller bind.ContractCaller) (*DominionAppCaller, error) {
	contract, err := bindDominionApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DominionAppCaller{contract: contract}, nil
}

// NewDominionAppTransactor creates a new write-only instance of DominionApp, bound to a specific deployed contract.
func NewDominionAppTransactor(address common.Address, transactor bind.ContractTransactor) (*DominionAppTransactor, error) {
	contract, err := bindDominionApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DominionAppTransactor{contract: contract}, nil
}

// NewDominionAppFilterer creates a new log filterer instance of DominionApp, bound to a specific deployed contract.
func NewDominionAppFilterer(address common.Address, filterer bind.ContractFilterer) (*DominionAppFilterer, error) {
	contract, err := bindDominionApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DominionAppFilterer{contract: contract}, nil
}

// bindDominionApp binds a generic wrapper to an already deployed contract.
func bindDominionApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DominionAppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DominionApp *DominionAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DominionApp.Contract.DominionAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DominionApp *DominionAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DominionApp.Contract.DominionAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DominionApp *DominionAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DominionApp.Contract.DominionAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DominionApp *DominionAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DominionApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DominionApp *DominionAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DominionApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DominionApp *DominionAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DominionApp.Contract.contract.Transact(opts, method, params...)
}

// ValidTransition is a free data retrieval call binding the contract method 0x0d1feb4f.
//
// Solidity: function validTransition((uint256,uint256,address[],address,bool,bool) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) from, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) to, uint256 signerIdx) pure returns()
func (_DominionApp *DominionAppCaller) ValidTransition(opts *bind.CallOpts, params ChannelParams, from ChannelState, to ChannelState, signerIdx *big.Int) error {
	var out []interface{}
	err := _DominionApp.contract.Call(opts, &out, "validTransition", params, from, to, signerIdx)

	if err != nil {
		return err
	}

	return err

}

// ValidTransition is a free data retrieval call binding the contract method 0x0d1feb4f.
//
// Solidity: function validTransition((uint256,uint256,address[],address,bool,bool) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) from, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) to, uint256 signerIdx) pure returns()
func (_DominionApp *DominionAppSession) ValidTransition(params ChannelParams, from ChannelState, to ChannelState, signerIdx *big.Int) error {
	return _DominionApp.Contract.ValidTransition(&_DominionApp.CallOpts, params, from, to, signerIdx)
}

// ValidTransition is a free data retrieval call binding the contract method 0x0d1feb4f.
//
// Solidity: function validTransition((uint256,uint256,address[],address,bool,bool) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) from, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) to, uint256 signerIdx) pure returns()
func (_DominionApp *DominionAppCallerSession) ValidTransition(params ChannelParams, from ChannelState, to ChannelState, signerIdx *big.Int) error {
	return _DominionApp.Contract.ValidTransition(&_DominionApp.CallOpts, params, from, to, signerIdx)
}

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208d53096eb43ff4f8eefef5aeb62c2c0151a130301976dbb20ab9dc0ae2437d8964736f6c63430007060033",
}

// ECDSAABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAMetaData.ABI instead.
var ECDSAABI = ECDSAMetaData.ABI

// ECDSABin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSAMetaData.Bin instead.
var ECDSABin = ECDSAMetaData.Bin

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// PileLibMetaData contains all meta data concerning the PileLib contract.
var PileLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207113688eb09bc065025c4ababbb77efedcf6ab8b3eceaaec9130e942b5c94c1a64736f6c63430007060033",
}

// PileLibABI is the input ABI used to generate the binding from.
// Deprecated: Use PileLibMetaData.ABI instead.
var PileLibABI = PileLibMetaData.ABI

// PileLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PileLibMetaData.Bin instead.
var PileLibBin = PileLibMetaData.Bin

// DeployPileLib deploys a new Ethereum contract, binding an instance of PileLib to it.
func DeployPileLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PileLib, error) {
	parsed, err := PileLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PileLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PileLib{PileLibCaller: PileLibCaller{contract: contract}, PileLibTransactor: PileLibTransactor{contract: contract}, PileLibFilterer: PileLibFilterer{contract: contract}}, nil
}

// PileLib is an auto generated Go binding around an Ethereum contract.
type PileLib struct {
	PileLibCaller     // Read-only binding to the contract
	PileLibTransactor // Write-only binding to the contract
	PileLibFilterer   // Log filterer for contract events
}

// PileLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type PileLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PileLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PileLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PileLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PileLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PileLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PileLibSession struct {
	Contract     *PileLib          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PileLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PileLibCallerSession struct {
	Contract *PileLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PileLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PileLibTransactorSession struct {
	Contract     *PileLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PileLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type PileLibRaw struct {
	Contract *PileLib // Generic contract binding to access the raw methods on
}

// PileLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PileLibCallerRaw struct {
	Contract *PileLibCaller // Generic read-only contract binding to access the raw methods on
}

// PileLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PileLibTransactorRaw struct {
	Contract *PileLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPileLib creates a new instance of PileLib, bound to a specific deployed contract.
func NewPileLib(address common.Address, backend bind.ContractBackend) (*PileLib, error) {
	contract, err := bindPileLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PileLib{PileLibCaller: PileLibCaller{contract: contract}, PileLibTransactor: PileLibTransactor{contract: contract}, PileLibFilterer: PileLibFilterer{contract: contract}}, nil
}

// NewPileLibCaller creates a new read-only instance of PileLib, bound to a specific deployed contract.
func NewPileLibCaller(address common.Address, caller bind.ContractCaller) (*PileLibCaller, error) {
	contract, err := bindPileLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PileLibCaller{contract: contract}, nil
}

// NewPileLibTransactor creates a new write-only instance of PileLib, bound to a specific deployed contract.
func NewPileLibTransactor(address common.Address, transactor bind.ContractTransactor) (*PileLibTransactor, error) {
	contract, err := bindPileLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PileLibTransactor{contract: contract}, nil
}

// NewPileLibFilterer creates a new log filterer instance of PileLib, bound to a specific deployed contract.
func NewPileLibFilterer(address common.Address, filterer bind.ContractFilterer) (*PileLibFilterer, error) {
	contract, err := bindPileLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PileLibFilterer{contract: contract}, nil
}

// bindPileLib binds a generic wrapper to an already deployed contract.
func bindPileLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PileLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PileLib *PileLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PileLib.Contract.PileLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PileLib *PileLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PileLib.Contract.PileLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PileLib *PileLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PileLib.Contract.PileLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PileLib *PileLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PileLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PileLib *PileLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PileLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PileLib *PileLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PileLib.Contract.contract.Transact(opts, method, params...)
}

// RNGLibMetaData contains all meta data concerning the RNGLib contract.
var RNGLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122043c67652d88153a1b3e0360ea05f373a4d86cb0c8d7ac8f1314bfbf9800ff31964736f6c63430007060033",
}

// RNGLibABI is the input ABI used to generate the binding from.
// Deprecated: Use RNGLibMetaData.ABI instead.
var RNGLibABI = RNGLibMetaData.ABI

// RNGLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RNGLibMetaData.Bin instead.
var RNGLibBin = RNGLibMetaData.Bin

// DeployRNGLib deploys a new Ethereum contract, binding an instance of RNGLib to it.
func DeployRNGLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RNGLib, error) {
	parsed, err := RNGLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RNGLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RNGLib{RNGLibCaller: RNGLibCaller{contract: contract}, RNGLibTransactor: RNGLibTransactor{contract: contract}, RNGLibFilterer: RNGLibFilterer{contract: contract}}, nil
}

// RNGLib is an auto generated Go binding around an Ethereum contract.
type RNGLib struct {
	RNGLibCaller     // Read-only binding to the contract
	RNGLibTransactor // Write-only binding to the contract
	RNGLibFilterer   // Log filterer for contract events
}

// RNGLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type RNGLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RNGLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RNGLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RNGLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RNGLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RNGLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RNGLibSession struct {
	Contract     *RNGLib           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RNGLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RNGLibCallerSession struct {
	Contract *RNGLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RNGLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RNGLibTransactorSession struct {
	Contract     *RNGLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RNGLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type RNGLibRaw struct {
	Contract *RNGLib // Generic contract binding to access the raw methods on
}

// RNGLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RNGLibCallerRaw struct {
	Contract *RNGLibCaller // Generic read-only contract binding to access the raw methods on
}

// RNGLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RNGLibTransactorRaw struct {
	Contract *RNGLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRNGLib creates a new instance of RNGLib, bound to a specific deployed contract.
func NewRNGLib(address common.Address, backend bind.ContractBackend) (*RNGLib, error) {
	contract, err := bindRNGLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RNGLib{RNGLibCaller: RNGLibCaller{contract: contract}, RNGLibTransactor: RNGLibTransactor{contract: contract}, RNGLibFilterer: RNGLibFilterer{contract: contract}}, nil
}

// NewRNGLibCaller creates a new read-only instance of RNGLib, bound to a specific deployed contract.
func NewRNGLibCaller(address common.Address, caller bind.ContractCaller) (*RNGLibCaller, error) {
	contract, err := bindRNGLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RNGLibCaller{contract: contract}, nil
}

// NewRNGLibTransactor creates a new write-only instance of RNGLib, bound to a specific deployed contract.
func NewRNGLibTransactor(address common.Address, transactor bind.ContractTransactor) (*RNGLibTransactor, error) {
	contract, err := bindRNGLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RNGLibTransactor{contract: contract}, nil
}

// NewRNGLibFilterer creates a new log filterer instance of RNGLib, bound to a specific deployed contract.
func NewRNGLibFilterer(address common.Address, filterer bind.ContractFilterer) (*RNGLibFilterer, error) {
	contract, err := bindRNGLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RNGLibFilterer{contract: contract}, nil
}

// bindRNGLib binds a generic wrapper to an already deployed contract.
func bindRNGLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RNGLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RNGLib *RNGLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RNGLib.Contract.RNGLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RNGLib *RNGLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RNGLib.Contract.RNGLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RNGLib *RNGLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RNGLib.Contract.RNGLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RNGLib *RNGLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RNGLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RNGLib *RNGLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RNGLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RNGLib *RNGLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RNGLib.Contract.contract.Transact(opts, method, params...)
}

// ReaderLibMetaData contains all meta data concerning the ReaderLib contract.
var ReaderLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204ddc7ea37fea533d59fcd257682dd28cd228ecfce281954c75826819888f655164736f6c63430007060033",
}

// ReaderLibABI is the input ABI used to generate the binding from.
// Deprecated: Use ReaderLibMetaData.ABI instead.
var ReaderLibABI = ReaderLibMetaData.ABI

// ReaderLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReaderLibMetaData.Bin instead.
var ReaderLibBin = ReaderLibMetaData.Bin

// DeployReaderLib deploys a new Ethereum contract, binding an instance of ReaderLib to it.
func DeployReaderLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReaderLib, error) {
	parsed, err := ReaderLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReaderLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReaderLib{ReaderLibCaller: ReaderLibCaller{contract: contract}, ReaderLibTransactor: ReaderLibTransactor{contract: contract}, ReaderLibFilterer: ReaderLibFilterer{contract: contract}}, nil
}

// ReaderLib is an auto generated Go binding around an Ethereum contract.
type ReaderLib struct {
	ReaderLibCaller     // Read-only binding to the contract
	ReaderLibTransactor // Write-only binding to the contract
	ReaderLibFilterer   // Log filterer for contract events
}

// ReaderLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReaderLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReaderLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReaderLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReaderLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReaderLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReaderLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReaderLibSession struct {
	Contract     *ReaderLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReaderLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReaderLibCallerSession struct {
	Contract *ReaderLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ReaderLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReaderLibTransactorSession struct {
	Contract     *ReaderLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ReaderLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReaderLibRaw struct {
	Contract *ReaderLib // Generic contract binding to access the raw methods on
}

// ReaderLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReaderLibCallerRaw struct {
	Contract *ReaderLibCaller // Generic read-only contract binding to access the raw methods on
}

// ReaderLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReaderLibTransactorRaw struct {
	Contract *ReaderLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReaderLib creates a new instance of ReaderLib, bound to a specific deployed contract.
func NewReaderLib(address common.Address, backend bind.ContractBackend) (*ReaderLib, error) {
	contract, err := bindReaderLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReaderLib{ReaderLibCaller: ReaderLibCaller{contract: contract}, ReaderLibTransactor: ReaderLibTransactor{contract: contract}, ReaderLibFilterer: ReaderLibFilterer{contract: contract}}, nil
}

// NewReaderLibCaller creates a new read-only instance of ReaderLib, bound to a specific deployed contract.
func NewReaderLibCaller(address common.Address, caller bind.ContractCaller) (*ReaderLibCaller, error) {
	contract, err := bindReaderLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReaderLibCaller{contract: contract}, nil
}

// NewReaderLibTransactor creates a new write-only instance of ReaderLib, bound to a specific deployed contract.
func NewReaderLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ReaderLibTransactor, error) {
	contract, err := bindReaderLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReaderLibTransactor{contract: contract}, nil
}

// NewReaderLibFilterer creates a new log filterer instance of ReaderLib, bound to a specific deployed contract.
func NewReaderLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ReaderLibFilterer, error) {
	contract, err := bindReaderLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReaderLibFilterer{contract: contract}, nil
}

// bindReaderLib binds a generic wrapper to an already deployed contract.
func bindReaderLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReaderLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReaderLib *ReaderLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReaderLib.Contract.ReaderLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReaderLib *ReaderLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReaderLib.Contract.ReaderLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReaderLib *ReaderLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReaderLib.Contract.ReaderLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReaderLib *ReaderLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReaderLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReaderLib *ReaderLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReaderLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReaderLib *ReaderLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReaderLib.Contract.contract.Transact(opts, method, params...)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d17b09d1099c654bbea20d5517c1fba64df872d94b281a0d9655cb8827a7555064736f6c63430007060033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SigMetaData contains all meta data concerning the Sig contract.
var SigMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122052da2a290f91c7c19584c810f9c6e7f81a783a561a1807e893ebb3b2d226d4cb64736f6c63430007060033",
}

// SigABI is the input ABI used to generate the binding from.
// Deprecated: Use SigMetaData.ABI instead.
var SigABI = SigMetaData.ABI

// SigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SigMetaData.Bin instead.
var SigBin = SigMetaData.Bin

// DeploySig deploys a new Ethereum contract, binding an instance of Sig to it.
func DeploySig(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sig, error) {
	parsed, err := SigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SigBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sig{SigCaller: SigCaller{contract: contract}, SigTransactor: SigTransactor{contract: contract}, SigFilterer: SigFilterer{contract: contract}}, nil
}

// Sig is an auto generated Go binding around an Ethereum contract.
type Sig struct {
	SigCaller     // Read-only binding to the contract
	SigTransactor // Write-only binding to the contract
	SigFilterer   // Log filterer for contract events
}

// SigCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigSession struct {
	Contract     *Sig              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigCallerSession struct {
	Contract *SigCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigTransactorSession struct {
	Contract     *SigTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigRaw struct {
	Contract *Sig // Generic contract binding to access the raw methods on
}

// SigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigCallerRaw struct {
	Contract *SigCaller // Generic read-only contract binding to access the raw methods on
}

// SigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigTransactorRaw struct {
	Contract *SigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSig creates a new instance of Sig, bound to a specific deployed contract.
func NewSig(address common.Address, backend bind.ContractBackend) (*Sig, error) {
	contract, err := bindSig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sig{SigCaller: SigCaller{contract: contract}, SigTransactor: SigTransactor{contract: contract}, SigFilterer: SigFilterer{contract: contract}}, nil
}

// NewSigCaller creates a new read-only instance of Sig, bound to a specific deployed contract.
func NewSigCaller(address common.Address, caller bind.ContractCaller) (*SigCaller, error) {
	contract, err := bindSig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigCaller{contract: contract}, nil
}

// NewSigTransactor creates a new write-only instance of Sig, bound to a specific deployed contract.
func NewSigTransactor(address common.Address, transactor bind.ContractTransactor) (*SigTransactor, error) {
	contract, err := bindSig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigTransactor{contract: contract}, nil
}

// NewSigFilterer creates a new log filterer instance of Sig, bound to a specific deployed contract.
func NewSigFilterer(address common.Address, filterer bind.ContractFilterer) (*SigFilterer, error) {
	contract, err := bindSig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigFilterer{contract: contract}, nil
}

// bindSig binds a generic wrapper to an already deployed contract.
func bindSig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sig *SigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sig.Contract.SigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sig *SigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sig.Contract.SigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sig *SigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sig.Contract.SigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sig *SigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sig *SigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sig *SigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sig.Contract.contract.Transact(opts, method, params...)
}

// StockLibMetaData contains all meta data concerning the StockLib contract.
var StockLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200009aebe27b0523f78859058ee998b7babaece7fd86b181822748aefe0926c1664736f6c63430007060033",
}

// StockLibABI is the input ABI used to generate the binding from.
// Deprecated: Use StockLibMetaData.ABI instead.
var StockLibABI = StockLibMetaData.ABI

// StockLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StockLibMetaData.Bin instead.
var StockLibBin = StockLibMetaData.Bin

// DeployStockLib deploys a new Ethereum contract, binding an instance of StockLib to it.
func DeployStockLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StockLib, error) {
	parsed, err := StockLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StockLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StockLib{StockLibCaller: StockLibCaller{contract: contract}, StockLibTransactor: StockLibTransactor{contract: contract}, StockLibFilterer: StockLibFilterer{contract: contract}}, nil
}

// StockLib is an auto generated Go binding around an Ethereum contract.
type StockLib struct {
	StockLibCaller     // Read-only binding to the contract
	StockLibTransactor // Write-only binding to the contract
	StockLibFilterer   // Log filterer for contract events
}

// StockLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type StockLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StockLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StockLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StockLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StockLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StockLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StockLibSession struct {
	Contract     *StockLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StockLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StockLibCallerSession struct {
	Contract *StockLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StockLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StockLibTransactorSession struct {
	Contract     *StockLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StockLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type StockLibRaw struct {
	Contract *StockLib // Generic contract binding to access the raw methods on
}

// StockLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StockLibCallerRaw struct {
	Contract *StockLibCaller // Generic read-only contract binding to access the raw methods on
}

// StockLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StockLibTransactorRaw struct {
	Contract *StockLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStockLib creates a new instance of StockLib, bound to a specific deployed contract.
func NewStockLib(address common.Address, backend bind.ContractBackend) (*StockLib, error) {
	contract, err := bindStockLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StockLib{StockLibCaller: StockLibCaller{contract: contract}, StockLibTransactor: StockLibTransactor{contract: contract}, StockLibFilterer: StockLibFilterer{contract: contract}}, nil
}

// NewStockLibCaller creates a new read-only instance of StockLib, bound to a specific deployed contract.
func NewStockLibCaller(address common.Address, caller bind.ContractCaller) (*StockLibCaller, error) {
	contract, err := bindStockLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StockLibCaller{contract: contract}, nil
}

// NewStockLibTransactor creates a new write-only instance of StockLib, bound to a specific deployed contract.
func NewStockLibTransactor(address common.Address, transactor bind.ContractTransactor) (*StockLibTransactor, error) {
	contract, err := bindStockLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StockLibTransactor{contract: contract}, nil
}

// NewStockLibFilterer creates a new log filterer instance of StockLib, bound to a specific deployed contract.
func NewStockLibFilterer(address common.Address, filterer bind.ContractFilterer) (*StockLibFilterer, error) {
	contract, err := bindStockLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StockLibFilterer{contract: contract}, nil
}

// bindStockLib binds a generic wrapper to an already deployed contract.
func bindStockLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StockLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StockLib *StockLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StockLib.Contract.StockLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StockLib *StockLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StockLib.Contract.StockLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StockLib *StockLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StockLib.Contract.StockLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StockLib *StockLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StockLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StockLib *StockLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StockLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StockLib *StockLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StockLib.Contract.contract.Transact(opts, method, params...)
}

// TurnLibMetaData contains all meta data concerning the TurnLib contract.
var TurnLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d614f89623f2acab4b3e114cd73d282e0573c14e66270070da4818922fce408864736f6c63430007060033",
}

// TurnLibABI is the input ABI used to generate the binding from.
// Deprecated: Use TurnLibMetaData.ABI instead.
var TurnLibABI = TurnLibMetaData.ABI

// TurnLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TurnLibMetaData.Bin instead.
var TurnLibBin = TurnLibMetaData.Bin

// DeployTurnLib deploys a new Ethereum contract, binding an instance of TurnLib to it.
func DeployTurnLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TurnLib, error) {
	parsed, err := TurnLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TurnLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TurnLib{TurnLibCaller: TurnLibCaller{contract: contract}, TurnLibTransactor: TurnLibTransactor{contract: contract}, TurnLibFilterer: TurnLibFilterer{contract: contract}}, nil
}

// TurnLib is an auto generated Go binding around an Ethereum contract.
type TurnLib struct {
	TurnLibCaller     // Read-only binding to the contract
	TurnLibTransactor // Write-only binding to the contract
	TurnLibFilterer   // Log filterer for contract events
}

// TurnLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type TurnLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TurnLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TurnLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TurnLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TurnLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TurnLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TurnLibSession struct {
	Contract     *TurnLib          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TurnLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TurnLibCallerSession struct {
	Contract *TurnLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TurnLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TurnLibTransactorSession struct {
	Contract     *TurnLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TurnLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type TurnLibRaw struct {
	Contract *TurnLib // Generic contract binding to access the raw methods on
}

// TurnLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TurnLibCallerRaw struct {
	Contract *TurnLibCaller // Generic read-only contract binding to access the raw methods on
}

// TurnLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TurnLibTransactorRaw struct {
	Contract *TurnLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTurnLib creates a new instance of TurnLib, bound to a specific deployed contract.
func NewTurnLib(address common.Address, backend bind.ContractBackend) (*TurnLib, error) {
	contract, err := bindTurnLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TurnLib{TurnLibCaller: TurnLibCaller{contract: contract}, TurnLibTransactor: TurnLibTransactor{contract: contract}, TurnLibFilterer: TurnLibFilterer{contract: contract}}, nil
}

// NewTurnLibCaller creates a new read-only instance of TurnLib, bound to a specific deployed contract.
func NewTurnLibCaller(address common.Address, caller bind.ContractCaller) (*TurnLibCaller, error) {
	contract, err := bindTurnLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TurnLibCaller{contract: contract}, nil
}

// NewTurnLibTransactor creates a new write-only instance of TurnLib, bound to a specific deployed contract.
func NewTurnLibTransactor(address common.Address, transactor bind.ContractTransactor) (*TurnLibTransactor, error) {
	contract, err := bindTurnLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TurnLibTransactor{contract: contract}, nil
}

// NewTurnLibFilterer creates a new log filterer instance of TurnLib, bound to a specific deployed contract.
func NewTurnLibFilterer(address common.Address, filterer bind.ContractFilterer) (*TurnLibFilterer, error) {
	contract, err := bindTurnLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TurnLibFilterer{contract: contract}, nil
}

// bindTurnLib binds a generic wrapper to an already deployed contract.
func bindTurnLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TurnLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TurnLib *TurnLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TurnLib.Contract.TurnLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TurnLib *TurnLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TurnLib.Contract.TurnLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TurnLib *TurnLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TurnLib.Contract.TurnLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TurnLib *TurnLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TurnLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TurnLib *TurnLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TurnLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TurnLib *TurnLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TurnLib.Contract.contract.Transact(opts, method, params...)
}

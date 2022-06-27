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

// DominionAppMetaData contains all meta data concerning the DominionApp contract.
var DominionAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ledgerChannel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualChannel\",\"type\":\"bool\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"to\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"signerIdx\",\"type\":\"uint256\"}],\"name\":\"validTransition\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610cdd806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80630d1feb4f14610030575b600080fd5b61004361003e366004610bc8565b610045565b005b6000604051806020016040528061009f8680606001906100659190610c5b565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061013392505050565b9052905060006100ae826101e5565b9050600060405180602001604052806100d08880606001906100659190610c5b565b9052905060006100df826101e5565b9050600060405180602001604052806101018980606001906100659190610c5b565b905290506000610110826101e5565b90506001815160200151600881111561012557fe5b505050505050505050505050565b60606000825167ffffffffffffffff8111801561014f57600080fd5b50604051908082528060200260200182016040528015610179578160200160208202803683370190505b50905060005b83518163ffffffff1610156101de57838163ffffffff16815181106101a057fe5b602001015160f81c60f81b828263ffffffff16815181106101bd57fe5b6001600160f81b03199092166020928302919091019091015260010161017f565b5092915050565b6101ed610aa4565b60006101f883610290565b90506000610205846102bb565b6040805160028082526060820190925291925060009190816020015b610229610add565b81526020019060019003908161022157905050905060005b600281101561025c57610253866102da565b50600101610241565b506000610268866102f9565b6040805160808101825295865260208601949094529284019190915250606082015292915050565b610298610b24565b60006102b46102af846102aa86610318565b61036c565b6103a2565b9392505050565b6102c3610b4d565b60006102b46102d5846102aa86610318565b6104b9565b6102e2610add565b60006102b46102f4846102aa86610318565b610540565b610301610b60565b60006102b4610313846102aa86610318565b61073b565b600080826000015160008151811061032c57fe5b6020026020010151905061034b836000015160018560000151516107ba565b808452805160009061035957fe5b602002602001015160f81c915050919050565b60606000610383846000015160008560ff166107ba565b845180519192506103989160ff8616906107ba565b8452905092915050565b6103aa610b24565b6040805160088082526101208201909252600091602082016101008036833701905050905060005b60038451038160ff16101561042a57610403848260030160ff16815181106103f657fe5b6020026020010151610865565b828260ff168151811061041257fe5b911515602092830291909101909101526001016103d2565b5060405180608001604052808460008151811061044357fe5b602002602001015160f81c60ff1681526020018460018151811061046357fe5b602002602001015160f81c60ff16600881111561047c57fe5b600881111561048757fe5b81526020016104a861049c86600260036107ba565b6000815181106103f657fe5b151581526020019190915292915050565b6104c1610b4d565b6040805160066020820181815261010083018452600093839290830160c0803683375050509052905060005b600660ff821610156101de57838160ff168151811061050857fe5b602002602001015160f81c82600001518260ff168151811061052657fe5b60ff909216602092830291909101909101526001016104ed565b610548610add565b60008260008151811061055757fe5b602002602001015160f81c9050600061058061057b8560018560010160ff166107ba565b61086e565b9050610594848360010160ff1686516107ba565b93506000846000815181106105a557fe5b602002602001015160f81c905060006105c961057b8760018560010160ff166107ba565b90506105dd868360010160ff1688516107ba565b95506000866000815181106105ee57fe5b602002602001015160f81c9050600061061261057b8960018560010160ff166107ba565b9050610626888360010160ff168a516107ba565b975060008860008151811061063757fe5b602002602001015160f81c9050600061065b61057b8b60018560010160ff166107ba565b905061066f8a8560010160ff168c516107ba565b60408051600480825260a08201909252919b50600091906020820160808036833701905050905060008b6000815181106106a557fe5b602002602001015160f81c905060005b8160ff168160ff161015610709578c8160010160ff16815181106106d557fe5b602002602001015160f81c838260ff16815181106106ef57fe5b60ff909216602092830291909101909101526001016106b5565b50506040805160a081018252988952602089019690965294870192909252506060850152506080830152509392505050565b610743610b60565b815160608080601460ff8516106107645761076186600060146107ba565b92505b602860ff85161061077f5761077c86601460286107ba565b91505b603c60ff85161061079a57610797866028603c6107ba565b90505b604080516060810182529384526020840192909252908201529392505050565b6060600083830367ffffffffffffffff811180156107d757600080fd5b50604051908082528060200260200182016040528015610801578160200160208202803683370190505b50905060005b81518160ff16101561085c57858160ff1686018151811061082457fe5b6020026020010151828260ff168151811061083b57fe5b6001600160f81b031990921660209283029190910190910152600101610807565b50949350505050565b60f81c60011490565b610876610b4d565b6000825167ffffffffffffffff8111801561089057600080fd5b506040519080825280602002602001820160405280156108ca57816020015b6108b7610b81565b8152602001906001900390816108af5790505b50905060005b83518160ff16101561091a576108f86108f3858360ff168460010160ff166107ba565b61092f565b828260ff168151811061090757fe5b60209081029190910101526001016108d0565b50604080516020810190915290815292915050565b610937610b81565b6000808360008151811061094757fe5b602002602001015160f81c90506000806000806000600581111561096757fe5b60ff168560ff1614156109895760009550600193506000925060009150610a5b565b600186600581111561099757fe5b60ff1614156109b55760019550600293506000925060019150610a5b565b60028660058111156109c357fe5b60ff1614156109e15760029550600393506000925060029150610a5b565b60038660058111156109ef57fe5b60ff161415610a0b575060039450600091506001905080610a5b565b6004866005811115610a1957fe5b60ff161415610a35575060049450600091506002905080610a5b565b60058681811115610a4257fe5b60ff161415610a5b575060059450600091506006905060035b60006040518060a00160405280886005811115610a7457fe5b815260ff968716602082015292861660408401529385166060830152509216608090920191909152949350505050565b6040518060800160405280610ab7610b24565b8152602001610ac4610b4d565b815260200160608152602001610ad8610b60565b905290565b6040518060a00160405280610af0610b4d565b8152602001610afd610b4d565b8152602001610b0a610b4d565b8152602001610b17610b4d565b8152602001606081525090565b604080516080810190915260008082526020820190815260006020820152606060409091015290565b6040518060200160405280606081525090565b60405180606001604052806060815260200160608152602001606081525090565b6040805160a081019091528060008152600060208201819052604082018190526060820181905260809091015290565b600060a08284031215610bc2578081fd5b50919050565b60008060008060808587031215610bdd578384fd5b843567ffffffffffffffff80821115610bf4578586fd5b9086019060c08289031215610c07578586fd5b90945060208601359080821115610c1c578485fd5b610c2888838901610bb1565b94506040870135915080821115610c3d578384fd5b50610c4a87828801610bb1565b949793965093946060013593505050565b6000808335601e19843603018112610c71578283fd5b83018035915067ffffffffffffffff821115610c8b578283fd5b602001915036819003821315610ca057600080fd5b925092905056fea26469706673582212200bbf6d9a9fdfa01ffd47d1bddf2f34e0f1ebfa64c6972c5c28268ef40a8f58bf64736f6c63430007060033",
}

// DominionAppABI is the input ABI used to generate the binding from.
// Deprecated: Use DominionAppMetaData.ABI instead.
var DominionAppABI = DominionAppMetaData.ABI

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

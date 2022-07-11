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
	Bin: "0x608060405234801561001057600080fd5b506110fa806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80630d1feb4f14610030575b600080fd5b61004361003e366004610f37565b610045565b005b6000604051806020016040528061009f868060600190610065919061107f565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506101ac92505050565b9052905060006100ae8261025e565b9050600060405180602001604052806100d0888060600190610065919061107f565b9052905060006100df8261025e565b905060006040518060200160405280610101898060600190610065919061107f565b9052905060006101108261025e565b80515190915060ff1687146101405760405162461bcd60e51b815260040161013790610ffa565b60405180910390fd5b600261014f60408c018c611031565b90501461016e5760405162461bcd60e51b815260040161013790610fca565b845160600151815160200151600881111561018557fe5b8151811061018f57fe5b60200260200101516101a057600080fd5b50505050505050505050565b60606000825167ffffffffffffffff811180156101c857600080fd5b506040519080825280602002602001820160405280156101f2578160200160208202803683370190505b50905060005b83518163ffffffff16101561025757838163ffffffff168151811061021957fe5b602001015160f81c60f81b828263ffffffff168151811061023657fe5b6001600160f81b0319909216602092830291909101909101526001016101f8565b5092915050565b610266610dd8565b600061027183610309565b9050600061027e84610334565b6040805160028082526060820190925291925060009190816020015b6102a2610e11565b81526020019060019003908161029a57905050905060005b60028110156102d5576102cc86610353565b506001016102ba565b5060006102e186610372565b6040805160808101825295865260208601949094529284019190915250606082015292915050565b610311610e58565b600061032d6103288461032386610391565b6103cb565b610403565b9392505050565b61033c610e88565b600061032d61034e8461032386610391565b610581565b61035b610e11565b600061032d61036d8461032386610391565b61068b565b61037a610ea2565b600061032d61038c8461032386610391565b610886565b6000806103ad6103a88460000151600060026108e4565b61098f565b835180519192506103c0916002906108e4565b835261032d81610a36565b606060006103e3846000015160008561ffff166108e4565b845180519192506103f99161ffff8616906108e4565b8452905092915050565b61040b610e58565b6040805160088082526101208201909252600091602082016101008036833701905050905060008360038151811061043f57fe5b602002602001015160f81c905060005b8160ff168160ff1610156104a65761047f858260040160ff168151811061047257fe5b6020026020010151610a3d565b838260ff168151811061048e57fe5b9115156020928302919091019091015260010161044f565b506000848260030160ff16815181106104bb57fe5b602002602001015160f81c905060006104e76104e2878560030160010160ff1689516108e4565b610a46565b90506040518060a001604052808760008151811061050157fe5b602002602001015160f81c60ff1681526020018760018151811061052157fe5b602002602001015160f81c60ff16600881111561053a57fe5b600881111561054557fe5b815260200161056661055a89600260036108e4565b60008151811061047257fe5b15158152602081019590955260409094015250909392505050565b610589610e88565b60408051600680825260e082019092526000916020820160c080368337505060408051600680825260e0820190925292935060009291506020820160c08036833701905050905060005b600660ff8216101561062257848160ff16815181106105ee57fe5b602002602001015160f81c838260ff168151811061060857fe5b60ff909216602092830291909101909101526001016105d3565b5060005b6006811015610672578481600660ff16018151811061064157fe5b602002602001015160f81c82828151811061065857fe5b60ff90921660209283029190910190910152600101610626565b5060408051808201909152918252602082015292915050565b610693610e11565b6000826000815181106106a257fe5b602002602001015160f81c905060006106cb6106c68560018560010160ff166108e4565b610b9f565b90506106df848360010160ff1686516108e4565b93506000846000815181106106f057fe5b602002602001015160f81c905060006107146106c68760018560010160ff166108e4565b9050610728868360010160ff1688516108e4565b955060008660008151811061073957fe5b602002602001015160f81c9050600061075d6106c68960018560010160ff166108e4565b9050610771888360010160ff168a516108e4565b975060008860008151811061078257fe5b602002602001015160f81c905060006107a66106c68b60018560010160ff166108e4565b90506107ba8a8560010160ff168c516108e4565b60408051600480825260a08201909252919b50600091906020820160808036833701905050905060008b6000815181106107f057fe5b602002602001015160f81c905060005b8160ff168160ff161015610854578c8160010160ff168151811061082057fe5b602002602001015160f81c838260ff168151811061083a57fe5b60ff90921660209283029190910190910152600101610800565b50506040805160a081018252988952602089019690965294870192909252506060850152506080830152509392505050565b61088e610ea2565b8151606080806108a186600060206108e4565b92506108b086602060a06108e4565b91506108c08660a06101206108e4565b60408051606081018252948552602085019390935291830191909152509392505050565b6060600083830367ffffffffffffffff8111801561090157600080fd5b5060405190808252806020026020018201604052801561092b578160200160208202803683370190505b50905060005b81518160ff16101561098657858160ff1686018151811061094e57fe5b6020026020010151828260ff168151811061096557fe5b6001600160f81b031990921660209283029190910190910152600101610931565b50949350505050565b60606000825167ffffffffffffffff811180156109ab57600080fd5b506040519080825280601f01601f1916602001820160405280156109d6576020820181803683370190505b50905060005b83518163ffffffff16101561025757838163ffffffff16815181106109fd57fe5b6020026020010151828263ffffffff1681518110610a1757fe5b60200101906001600160f81b031916908160001a9053506001016109dc565b6020015190565b60f81c60011490565b610a4e610ec3565b600082600081518110610a5d57fe5b602002602001015160f81c9050600083600181518110610a7957fe5b602002602001015160f81c60ff16600f811115610a9257fe5b9050600084600281518110610aa357fe5b602002602001015160f81c60ff16600f811115610abc57fe5b905060006003840360ff1667ffffffffffffffff81118015610add57600080fd5b50604051908082528060200260200182016040528015610b07578160200160208202803683370190505b50905060005b6003850360ff168160ff161015610b6457868160030160ff1681518110610b3057fe5b602002602001015160f81c828260ff1681518110610b4a57fe5b60ff90921660209283029190910190910152600101610b0d565b50604051806060016040528084600f811115610b7c57fe5b815260200183600f811115610b8d57fe5b81526020019190915295945050505050565b610ba7610edd565b6000825167ffffffffffffffff81118015610bc157600080fd5b50604051908082528060200260200182016040528015610bfb57816020015b610be8610ef0565b815260200190600190039081610be05790505b50905060005b83518160ff161015610c4b57610c29610c24858360ff168460010160ff166108e4565b610c60565b828260ff1681518110610c3857fe5b6020908102919091010152600101610c01565b50604080516020810190915290815292915050565b610c68610ef0565b60008083600081518110610c7857fe5b602002602001015160f81c90506000806000806000600f811115610c9857fe5b60ff168560ff161415610cba5760009550600193506000925060009150610d8f565b600186600f811115610cc857fe5b60ff161415610ce65760019550600293506000925060039150610d8f565b600286600f811115610cf457fe5b60ff161415610d125760029550600393506000925060069150610d8f565b600386600f811115610d2057fe5b60ff161415610d3d57506003945060009150600290506001610d8f565b600486600f811115610d4b57fe5b60ff161415610d6857506004945060009150600590506003610d8f565b600586600f811115610d7657fe5b60ff161415610d8f575060059450600091506008905060065b60006040518060a0016040528088600f811115610da857fe5b815260ff968716602082015292861660408401529385166060830152509216608090920191909152949350505050565b6040518060800160405280610deb610e58565b8152602001610df8610e88565b815260200160608152602001610e0c610ea2565b905290565b6040518060a00160405280610e24610edd565b8152602001610e31610edd565b8152602001610e3e610edd565b8152602001610e4b610edd565b8152602001606081525090565b6040805160a081019091526000808252602082019081526000602082015260606040820181905201610e0c610ec3565b604051806040016040528060608152602001606081525090565b60405180606001604052806060815260200160608152602001606081525090565b604080516060810190915280600081526020016000610e4b565b6040518060200160405280606081525090565b6040805160a081019091528060008152600060208201819052604082018190526060820181905260809091015290565b600060a08284031215610f31578081fd5b50919050565b60008060008060808587031215610f4c578384fd5b843567ffffffffffffffff80821115610f63578586fd5b9086019060c08289031215610f76578586fd5b90945060208601359080821115610f8b578485fd5b610f9788838901610f20565b94506040870135915080821115610fac578384fd5b50610fb987828801610f20565b949793965093946060013593505050565b6020808252601690820152754e756d626572206f66207061727469636970616e747360501b604082015260600190565b60208082526017908201527f5369676e6572206973206e6f74206e6578744163746f72000000000000000000604082015260600190565b6000808335601e19843603018112611047578283fd5b83018035915067ffffffffffffffff821115611061578283fd5b602090810192508102360382131561107857600080fd5b9250929050565b6000808335601e19843603018112611095578283fd5b83018035915067ffffffffffffffff8211156110af578283fd5b60200191503681900382131561107857600080fdfea26469706673582212205454e507c87a85c6841f9b3ff3d4a1e0c6e36597b8c401fadeb1f48116efc29564736f6c63430007060033",
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

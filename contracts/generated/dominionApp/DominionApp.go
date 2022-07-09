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
	Bin: "0x608060405234801561001057600080fd5b50610e96806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80630d1feb4f14610030575b600080fd5b61004361003e366004610cd3565b610045565b005b6000604051806020016040528061009f8680606001906100659190610e1b565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506101ac92505050565b9052905060006100ae8261025e565b9050600060405180602001604052806100d08880606001906100659190610e1b565b9052905060006100df8261025e565b9050600060405180602001604052806101018980606001906100659190610e1b565b9052905060006101108261025e565b80515190915060ff1687146101405760405162461bcd60e51b815260040161013790610d96565b60405180910390fd5b600261014f60408c018c610dcd565b90501461016e5760405162461bcd60e51b815260040161013790610d66565b845160600151815160200151600881111561018557fe5b8151811061018f57fe5b60200260200101516101a057600080fd5b50505050505050505050565b60606000825167ffffffffffffffff811180156101c857600080fd5b506040519080825280602002602001820160405280156101f2578160200160208202803683370190505b50905060005b83518163ffffffff16101561025757838163ffffffff168151811061021957fe5b602001015160f81c60f81b828263ffffffff168151811061023657fe5b6001600160f81b0319909216602092830291909101909101526001016101f8565b5092915050565b610266610baf565b600061027183610309565b9050600061027e84610334565b6040805160028082526060820190925291925060009190816020015b6102a2610be8565b81526020019060019003908161029a57905050905060005b60028110156102d5576102cc86610353565b506001016102ba565b5060006102e186610372565b6040805160808101825295865260208601949094529284019190915250606082015292915050565b610311610c2f565b600061032d6103288461032386610391565b6103cb565b610403565b9392505050565b61033c610c58565b600061032d61034e8461032386610391565b610536565b61035b610be8565b600061032d61036d8461032386610391565b6105bd565b61037a610c6b565b600061032d61038c8461032386610391565b6107b8565b6000806103ad6103a8846000015160006002610816565b6108c1565b835180519192506103c091600290610816565b835261032d81610968565b606060006103e3846000015160008561ffff16610816565b845180519192506103f99161ffff861690610816565b8452905092915050565b61040b610c2f565b6040805160088082526101208201909252600091602082016101008036833701905050905060008360038151811061043f57fe5b602002602001015160f81c905060005b8160ff168160ff1610156104a65761047f858260040160ff168151811061047257fe5b602002602001015161096f565b838260ff168151811061048e57fe5b9115156020928302919091019091015260010161044f565b506040518060800160405280856000815181106104bf57fe5b602002602001015160f81c60ff168152602001856001815181106104df57fe5b602002602001015160f81c60ff1660088111156104f857fe5b600881111561050357fe5b81526020016105246105188760026003610816565b60008151811061047257fe5b15158152602001929092525092915050565b61053e610c58565b6040805160066020820181815261010083018452600093839290830160c0803683375050509052905060005b600660ff8216101561025757838160ff168151811061058557fe5b602002602001015160f81c82600001518260ff16815181106105a357fe5b60ff9092166020928302919091019091015260010161056a565b6105c5610be8565b6000826000815181106105d457fe5b602002602001015160f81c905060006105fd6105f88560018560010160ff16610816565b610978565b9050610611848360010160ff168651610816565b935060008460008151811061062257fe5b602002602001015160f81c905060006106466105f88760018560010160ff16610816565b905061065a868360010160ff168851610816565b955060008660008151811061066b57fe5b602002602001015160f81c9050600061068f6105f88960018560010160ff16610816565b90506106a3888360010160ff168a51610816565b97506000886000815181106106b457fe5b602002602001015160f81c905060006106d86105f88b60018560010160ff16610816565b90506106ec8a8560010160ff168c51610816565b60408051600480825260a08201909252919b50600091906020820160808036833701905050905060008b60008151811061072257fe5b602002602001015160f81c905060005b8160ff168160ff161015610786578c8160010160ff168151811061075257fe5b602002602001015160f81c838260ff168151811061076c57fe5b60ff90921660209283029190910190910152600101610732565b50506040805160a081018252988952602089019690965294870192909252506060850152506080830152509392505050565b6107c0610c6b565b8151606080806107d38660006020610816565b92506107e286602060a0610816565b91506107f28660a0610120610816565b60408051606081018252948552602085019390935291830191909152509392505050565b6060600083830367ffffffffffffffff8111801561083357600080fd5b5060405190808252806020026020018201604052801561085d578160200160208202803683370190505b50905060005b81518160ff1610156108b857858160ff1686018151811061088057fe5b6020026020010151828260ff168151811061089757fe5b6001600160f81b031990921660209283029190910190910152600101610863565b50949350505050565b60606000825167ffffffffffffffff811180156108dd57600080fd5b506040519080825280601f01601f191660200182016040528015610908576020820181803683370190505b50905060005b83518163ffffffff16101561025757838163ffffffff168151811061092f57fe5b6020026020010151828263ffffffff168151811061094957fe5b60200101906001600160f81b031916908160001a90535060010161090e565b6020015190565b60f81c60011490565b610980610c58565b6000825167ffffffffffffffff8111801561099a57600080fd5b506040519080825280602002602001820160405280156109d457816020015b6109c1610c8c565b8152602001906001900390816109b95790505b50905060005b83518160ff161015610a2457610a026109fd858360ff168460010160ff16610816565b610a39565b828260ff1681518110610a1157fe5b60209081029190910101526001016109da565b50604080516020810190915290815292915050565b610a41610c8c565b60008083600081518110610a5157fe5b602002602001015160f81c90506000806000806000600f811115610a7157fe5b60ff168560ff161415610a935760009550600193506000925060009150610b66565b600186600f811115610aa157fe5b60ff161415610abf5760019550600293506000925060019150610b66565b600286600f811115610acd57fe5b60ff161415610aeb5760029550600393506000925060029150610b66565b600386600f811115610af957fe5b60ff161415610b15575060039450600091506001905080610b66565b600486600f811115610b2357fe5b60ff161415610b3f575060049450600091506002905080610b66565b600586600f811115610b4d57fe5b60ff161415610b66575060059450600091506006905060035b60006040518060a0016040528088600f811115610b7f57fe5b815260ff968716602082015292861660408401529385166060830152509216608090920191909152949350505050565b6040518060800160405280610bc2610c2f565b8152602001610bcf610c58565b815260200160608152602001610be3610c6b565b905290565b6040518060a00160405280610bfb610c58565b8152602001610c08610c58565b8152602001610c15610c58565b8152602001610c22610c58565b8152602001606081525090565b604080516080810190915260008082526020820190815260006020820152606060409091015290565b6040518060200160405280606081525090565b60405180606001604052806060815260200160608152602001606081525090565b6040805160a081019091528060008152600060208201819052604082018190526060820181905260809091015290565b600060a08284031215610ccd578081fd5b50919050565b60008060008060808587031215610ce8578384fd5b843567ffffffffffffffff80821115610cff578586fd5b9086019060c08289031215610d12578586fd5b90945060208601359080821115610d27578485fd5b610d3388838901610cbc565b94506040870135915080821115610d48578384fd5b50610d5587828801610cbc565b949793965093946060013593505050565b6020808252601690820152754e756d626572206f66207061727469636970616e747360501b604082015260600190565b60208082526017908201527f5369676e6572206973206e6f74206e6578744163746f72000000000000000000604082015260600190565b6000808335601e19843603018112610de3578283fd5b83018035915067ffffffffffffffff821115610dfd578283fd5b6020908101925081023603821315610e1457600080fd5b9250929050565b6000808335601e19843603018112610e31578283fd5b83018035915067ffffffffffffffff821115610e4b578283fd5b602001915036819003821315610e1457600080fdfea264697066735822122081c2f6b35df69c6c97bef203471204523df0421dccc122ced57f45607500eadf64736f6c63430007060033",
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

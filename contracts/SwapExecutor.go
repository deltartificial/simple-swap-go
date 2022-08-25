// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swapexecutor

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

// SwapExecutorMetaData contains all meta data concerning the SwapExecutor contract.
var SwapExecutorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"swapTokensOnUniswapV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610e9b806100606000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633e0c0629146100515780635f3e849f1461006d5780639b96eece14610089578063c1eaf593146100b9575b600080fd5b61006b600480360381019061006691906107a5565b6100d5565b005b610087600480360381019061008291906107e5565b610217565b005b6100a3600480360381019061009e9190610838565b61036d565b6040516100b09190610874565b60405180910390f35b6100d360048036038101906100ce919061088f565b6103f0565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610163576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015a90610979565b60405180910390fd5b6000808373ffffffffffffffffffffffffffffffffffffffff168360405161018a906109ca565b60006040518083038185875af1925050503d80600081146101c7576040519150601f19603f3d011682016040523d82523d6000602084013e6101cc565b606091505b509150915081610211576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020890610a2b565b60405180910390fd5b50505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029c90610979565b60405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b81526004016102e2929190610a5a565b6020604051808303816000875af1158015610301573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103259190610abb565b905080610367576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035e90610a2b565b60405180910390fd5b50505050565b60008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016103a89190610ae8565b602060405180830381865afa1580156103c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e99190610b18565b9050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461047e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047590610979565b60405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff166323b872dd3330876040518463ffffffff1660e01b81526004016104bb93929190610b45565b6020604051808303816000875af11580156104da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104fe9190610abb565b508573ffffffffffffffffffffffffffffffffffffffff1663095ea7b382866040518363ffffffff1660e01b815260040161053a929190610a5a565b6020604051808303816000875af1158015610559573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057d9190610abb565b506060600267ffffffffffffffff81111561059b5761059a610b7c565b5b6040519080825280602002602001820160405280156105c95781602001602082028036833780820191505090505b50905086816000815181106105e1576105e0610bab565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505085816001815181106106305761062f610bab565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff166338ed173986868487426040518663ffffffff1660e01b81526004016106ab959493929190610c98565b6000604051808303816000875af11580156106ca573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906106f39190610e1c565b5050505050505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061073c82610711565b9050919050565b61074c81610731565b811461075757600080fd5b50565b60008135905061076981610743565b92915050565b6000819050919050565b6107828161076f565b811461078d57600080fd5b50565b60008135905061079f81610779565b92915050565b600080604083850312156107bc576107bb610707565b5b60006107ca8582860161075a565b92505060206107db85828601610790565b9150509250929050565b6000806000606084860312156107fe576107fd610707565b5b600061080c8682870161075a565b935050602061081d8682870161075a565b925050604061082e86828701610790565b9150509250925092565b60006020828403121561084e5761084d610707565b5b600061085c8482850161075a565b91505092915050565b61086e8161076f565b82525050565b60006020820190506108896000830184610865565b92915050565b60008060008060008060c087890312156108ac576108ab610707565b5b60006108ba89828a0161075a565b96505060206108cb89828a0161075a565b95505060406108dc89828a01610790565b94505060606108ed89828a01610790565b93505060806108fe89828a0161075a565b92505060a061090f89828a0161075a565b9150509295509295509295565b600082825260208201905092915050565b7f214558454355544f520000000000000000000000000000000000000000000000600082015250565b600061096360098361091c565b915061096e8261092d565b602082019050919050565b6000602082019050818103600083015261099281610956565b9050919050565b600081905092915050565b50565b60006109b4600083610999565b91506109bf826109a4565b600082019050919050565b60006109d5826109a7565b9150819050919050565b7f215452414e534645520000000000000000000000000000000000000000000000600082015250565b6000610a1560098361091c565b9150610a20826109df565b602082019050919050565b60006020820190508181036000830152610a4481610a08565b9050919050565b610a5481610731565b82525050565b6000604082019050610a6f6000830185610a4b565b610a7c6020830184610865565b9392505050565b60008115159050919050565b610a9881610a83565b8114610aa357600080fd5b50565b600081519050610ab581610a8f565b92915050565b600060208284031215610ad157610ad0610707565b5b6000610adf84828501610aa6565b91505092915050565b6000602082019050610afd6000830184610a4b565b92915050565b600081519050610b1281610779565b92915050565b600060208284031215610b2e57610b2d610707565b5b6000610b3c84828501610b03565b91505092915050565b6000606082019050610b5a6000830186610a4b565b610b676020830185610a4b565b610b746040830184610865565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610c0f81610731565b82525050565b6000610c218383610c06565b60208301905092915050565b6000602082019050919050565b6000610c4582610bda565b610c4f8185610be5565b9350610c5a83610bf6565b8060005b83811015610c8b578151610c728882610c15565b9750610c7d83610c2d565b925050600181019050610c5e565b5085935050505092915050565b600060a082019050610cad6000830188610865565b610cba6020830187610865565b8181036040830152610ccc8186610c3a565b9050610cdb6060830185610a4b565b610ce86080830184610865565b9695505050505050565b600080fd5b6000601f19601f8301169050919050565b610d1182610cf7565b810181811067ffffffffffffffff82111715610d3057610d2f610b7c565b5b80604052505050565b6000610d436106fd565b9050610d4f8282610d08565b919050565b600067ffffffffffffffff821115610d6f57610d6e610b7c565b5b602082029050602081019050919050565b600080fd5b6000610d98610d9384610d54565b610d39565b90508083825260208201905060208402830185811115610dbb57610dba610d80565b5b835b81811015610de45780610dd08882610b03565b845260208401935050602081019050610dbd565b5050509392505050565b600082601f830112610e0357610e02610cf2565b5b8151610e13848260208601610d85565b91505092915050565b600060208284031215610e3257610e31610707565b5b600082015167ffffffffffffffff811115610e5057610e4f61070c565b5b610e5c84828501610dee565b9150509291505056fea2646970667358221220e2e79ef40129b04baac92de978d7981455ee5517b66091cf492994df879f3d6264736f6c634300080e0033",
}

// SwapExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapExecutorMetaData.ABI instead.
var SwapExecutorABI = SwapExecutorMetaData.ABI

// SwapExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwapExecutorMetaData.Bin instead.
var SwapExecutorBin = SwapExecutorMetaData.Bin

// DeploySwapExecutor deploys a new Ethereum contract, binding an instance of SwapExecutor to it.
func DeploySwapExecutor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SwapExecutor, error) {
	parsed, err := SwapExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwapExecutorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwapExecutor{SwapExecutorCaller: SwapExecutorCaller{contract: contract}, SwapExecutorTransactor: SwapExecutorTransactor{contract: contract}, SwapExecutorFilterer: SwapExecutorFilterer{contract: contract}}, nil
}

// SwapExecutor is an auto generated Go binding around an Ethereum contract.
type SwapExecutor struct {
	SwapExecutorCaller     // Read-only binding to the contract
	SwapExecutorTransactor // Write-only binding to the contract
	SwapExecutorFilterer   // Log filterer for contract events
}

// SwapExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapExecutorSession struct {
	Contract     *SwapExecutor     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapExecutorCallerSession struct {
	Contract *SwapExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SwapExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapExecutorTransactorSession struct {
	Contract     *SwapExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SwapExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapExecutorRaw struct {
	Contract *SwapExecutor // Generic contract binding to access the raw methods on
}

// SwapExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapExecutorCallerRaw struct {
	Contract *SwapExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// SwapExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapExecutorTransactorRaw struct {
	Contract *SwapExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapExecutor creates a new instance of SwapExecutor, bound to a specific deployed contract.
func NewSwapExecutor(address common.Address, backend bind.ContractBackend) (*SwapExecutor, error) {
	contract, err := bindSwapExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapExecutor{SwapExecutorCaller: SwapExecutorCaller{contract: contract}, SwapExecutorTransactor: SwapExecutorTransactor{contract: contract}, SwapExecutorFilterer: SwapExecutorFilterer{contract: contract}}, nil
}

// NewSwapExecutorCaller creates a new read-only instance of SwapExecutor, bound to a specific deployed contract.
func NewSwapExecutorCaller(address common.Address, caller bind.ContractCaller) (*SwapExecutorCaller, error) {
	contract, err := bindSwapExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapExecutorCaller{contract: contract}, nil
}

// NewSwapExecutorTransactor creates a new write-only instance of SwapExecutor, bound to a specific deployed contract.
func NewSwapExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapExecutorTransactor, error) {
	contract, err := bindSwapExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapExecutorTransactor{contract: contract}, nil
}

// NewSwapExecutorFilterer creates a new log filterer instance of SwapExecutor, bound to a specific deployed contract.
func NewSwapExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapExecutorFilterer, error) {
	contract, err := bindSwapExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapExecutorFilterer{contract: contract}, nil
}

// bindSwapExecutor binds a generic wrapper to an already deployed contract.
func bindSwapExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapExecutorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapExecutor *SwapExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapExecutor.Contract.SwapExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapExecutor *SwapExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapExecutor.Contract.SwapExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapExecutor *SwapExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapExecutor.Contract.SwapExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapExecutor *SwapExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapExecutor *SwapExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapExecutor *SwapExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapExecutor.Contract.contract.Transact(opts, method, params...)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x9b96eece.
//
// Solidity: function getBalanceOf(address _token) view returns(uint256)
func (_SwapExecutor *SwapExecutorCaller) GetBalanceOf(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwapExecutor.contract.Call(opts, &out, "getBalanceOf", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceOf is a free data retrieval call binding the contract method 0x9b96eece.
//
// Solidity: function getBalanceOf(address _token) view returns(uint256)
func (_SwapExecutor *SwapExecutorSession) GetBalanceOf(_token common.Address) (*big.Int, error) {
	return _SwapExecutor.Contract.GetBalanceOf(&_SwapExecutor.CallOpts, _token)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x9b96eece.
//
// Solidity: function getBalanceOf(address _token) view returns(uint256)
func (_SwapExecutor *SwapExecutorCallerSession) GetBalanceOf(_token common.Address) (*big.Int, error) {
	return _SwapExecutor.Contract.GetBalanceOf(&_SwapExecutor.CallOpts, _token)
}

// RecoverETH is a paid mutator transaction binding the contract method 0x3e0c0629.
//
// Solidity: function recoverETH(address _to, uint256 _amount) returns()
func (_SwapExecutor *SwapExecutorTransactor) RecoverETH(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapExecutor.contract.Transact(opts, "recoverETH", _to, _amount)
}

// RecoverETH is a paid mutator transaction binding the contract method 0x3e0c0629.
//
// Solidity: function recoverETH(address _to, uint256 _amount) returns()
func (_SwapExecutor *SwapExecutorSession) RecoverETH(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapExecutor.Contract.RecoverETH(&_SwapExecutor.TransactOpts, _to, _amount)
}

// RecoverETH is a paid mutator transaction binding the contract method 0x3e0c0629.
//
// Solidity: function recoverETH(address _to, uint256 _amount) returns()
func (_SwapExecutor *SwapExecutorTransactorSession) RecoverETH(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapExecutor.Contract.RecoverETH(&_SwapExecutor.TransactOpts, _to, _amount)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x5f3e849f.
//
// Solidity: function recoverTokens(address _token, address _to, uint256 _amount) returns()
func (_SwapExecutor *SwapExecutorTransactor) RecoverTokens(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapExecutor.contract.Transact(opts, "recoverTokens", _token, _to, _amount)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x5f3e849f.
//
// Solidity: function recoverTokens(address _token, address _to, uint256 _amount) returns()
func (_SwapExecutor *SwapExecutorSession) RecoverTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapExecutor.Contract.RecoverTokens(&_SwapExecutor.TransactOpts, _token, _to, _amount)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x5f3e849f.
//
// Solidity: function recoverTokens(address _token, address _to, uint256 _amount) returns()
func (_SwapExecutor *SwapExecutorTransactorSession) RecoverTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapExecutor.Contract.RecoverTokens(&_SwapExecutor.TransactOpts, _token, _to, _amount)
}

// SwapTokensOnUniswapV2 is a paid mutator transaction binding the contract method 0xc1eaf593.
//
// Solidity: function swapTokensOnUniswapV2(address _tokenIn, address _tokenOut, uint256 _amountIn, uint256 _amountOutMin, address _to, address _router) returns()
func (_SwapExecutor *SwapExecutorTransactor) SwapTokensOnUniswapV2(opts *bind.TransactOpts, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int, _amountOutMin *big.Int, _to common.Address, _router common.Address) (*types.Transaction, error) {
	return _SwapExecutor.contract.Transact(opts, "swapTokensOnUniswapV2", _tokenIn, _tokenOut, _amountIn, _amountOutMin, _to, _router)
}

// SwapTokensOnUniswapV2 is a paid mutator transaction binding the contract method 0xc1eaf593.
//
// Solidity: function swapTokensOnUniswapV2(address _tokenIn, address _tokenOut, uint256 _amountIn, uint256 _amountOutMin, address _to, address _router) returns()
func (_SwapExecutor *SwapExecutorSession) SwapTokensOnUniswapV2(_tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int, _amountOutMin *big.Int, _to common.Address, _router common.Address) (*types.Transaction, error) {
	return _SwapExecutor.Contract.SwapTokensOnUniswapV2(&_SwapExecutor.TransactOpts, _tokenIn, _tokenOut, _amountIn, _amountOutMin, _to, _router)
}

// SwapTokensOnUniswapV2 is a paid mutator transaction binding the contract method 0xc1eaf593.
//
// Solidity: function swapTokensOnUniswapV2(address _tokenIn, address _tokenOut, uint256 _amountIn, uint256 _amountOutMin, address _to, address _router) returns()
func (_SwapExecutor *SwapExecutorTransactorSession) SwapTokensOnUniswapV2(_tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int, _amountOutMin *big.Int, _to common.Address, _router common.Address) (*types.Transaction, error) {
	return _SwapExecutor.Contract.SwapTokensOnUniswapV2(&_SwapExecutor.TransactOpts, _tokenIn, _tokenOut, _amountIn, _amountOutMin, _to, _router)
}

// @author - @deltartificial
// @dev - Simple Buy Tokens on Uniswap V2 Router using a contract in Solidity or Vyper

package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	swapexecutor "github.com/deltartificial/simple-buy-go/contracts"
)

var (
    HTTP_RPC string = "https://bsc-dataseed.binance.org/" // http rpc
    PRIVATE string = "" // private key of the account to use
    CONTRACT_ADDRESS string = "" // address of the contract to use
    GAS_LIMIT = 300000 // 300.000 gas
    GAS_PRICE = 10// 10 gwei
    TOKEN_IN string = "" // address of the token to swap in
    TOKEN_OUT string= "" // address of the token to swap out
    AMOUNT_IN = 100 // in wei
	AMOUNT_OUT_MIN = 0 // in wei
	TO string = CONTRACT_ADDRESS // address receiver
	ROUTER string = "" // address of the router to use
	CHAINID int64 = 56 // chain id of the network
)

func main() {

    client, err := ethclient.Dial(HTTP_RPC)
    if err != nil {
        log.Fatal(err, "could not connect to ethereum client")
    }

    privateKey, err := crypto.HexToECDSA(PRIVATE)
    if err != nil {
        log.Fatal(err, "could not parse private key")
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal(err, "could not get public key")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err, "could not get nonce")
    }

	chainId := big.NewInt(CHAINID)
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err, "could not create auth")
	}

    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0) // in wei
    auth.GasLimit = uint64(GAS_LIMIT) // in units
    auth.GasPrice = big.NewInt(int64(GAS_PRICE * 1000000000)) // in wei

    address := common.HexToAddress(CONTRACT_ADDRESS)

    instance, err := swapexecutor.NewSwapExecutor(address, client)
    if err != nil {
        log.Fatal(err, "could not create instance")
    }

    tokenIn := common.HexToAddress(TOKEN_IN)
    tokenOut := common.HexToAddress(TOKEN_OUT)
	amountIn := big.NewInt(int64(AMOUNT_IN))
	amountOutMin := big.NewInt(int64(AMOUNT_OUT_MIN))
	to := common.HexToAddress(TO)
	router := common.HexToAddress(ROUTER)

    tx, err := instance.SwapTokensOnUniswapV2(auth, tokenIn, tokenOut, amountIn, amountOutMin, to, router)
    if err != nil {
        log.Fatal(err, "could not send swap")
    }

    fmt.Printf("Swap Transaction Sent: %s", tx.Hash().Hex())

}
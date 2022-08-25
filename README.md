# Simple Swap Go •
This project is a simple demonstration of a swap on Uniswap V2 (and all forks) using Golang. It does a simple swap using a smart-contract "SwapExecutor" developed in Vyper and Solidity.

### Environment Variables

- **HTTP_RPC** - HTTP Rpc Endpoint
- **PRIVATE** - Private key for the contract executor.
- **CONTRACT_ADDRESS** - Swap Executor contract address.
- **GAS_LIMIT** - Gas limit of the swap transaction.
- **GAS_PRICE** - Gas price of the swap transaction.
- **TOKEN_IN** - Address of the token in (token to buy with).
- **TOKEN_OUT** - Address of the token out (token to get).
- **AMOUNT_IN** - Amount of token in (in wei).
- **AMOUNT_OUT_MIN** - Amount of token out min (in wei, you can set to 0).
- **TO** - Address who receives the tokens (you can set the contract).
- **ROUTER** - Router Address. (ex PancakeSwap, UniswapV2 router).
- **CHAINDID** - Chain Id of current network. (ex BSC : 56)

### Setup 

```sh
git clone https://github.com/deltartificial/simple-buy-go
cd simple-buy-go
```

### Usage

1. Generate a new bot wallet address and extract the private key into a raw 32-byte format.
2. Deploy the included SwapExecutor.sol or SwapExecutor.vy to the blockchain.
3. Transfer tokens to the newly wallet address (the swap executor contract creator).
4. Customize the environment variables according to your choices.

- You can withdraw the tokens from the contract by calling "recoverTokens()" or "recoverETH()"

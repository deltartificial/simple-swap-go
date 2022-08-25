# @version ^0.3.3

# @author - @deltartificial
# @dev This contract is used to execute swaps between two tokens on Uniswap V2 Router.

from vyper.interfaces import ERC20

executor: public(address)

@external
def __init__():
    self.executor = msg.sender

#///////////////////////////////////////////////////////////////
#                       SWAP UNISWAP V2 
#//////////////////////////////////////////////////////////////

@external
def swapTokensOnUniswapV2(
    _tokenIn: address, 
    _tokenOut: address, 
    _amountIn: uint256, 
    _amountOutMin: uint256, 
    _to: address, 
    _router: address):
    assert self.executor == msg.sender,"!EXECUTOR"
    ERC20(_tokenIn).transferFrom(msg.sender, self, _amountIn)
    ERC20(_tokenIn).approve(_router, _amountIn)

    res: Bytes[128] = raw_call(
        _router,
        concat(
            method_id("swapExactTokensForTokens(uint256,uint256,address[],address,uint256)"),
            convert(_amountIn, bytes32),       # amount in
            convert(_amountOutMin, bytes32),   # amount out min
            convert(160, bytes32),             # path[] offset (5 * 32, 5 = number of func args)
            convert(_to, bytes32),             # to
            convert(block.timestamp, bytes32), # deadline
            convert(2, bytes32),               # path[] length
            convert(_tokenIn, bytes32),        # path[] token in    
            convert(_tokenOut, bytes32)        # path[] token out
    ),
    max_outsize=128
  )

#///////////////////////////////////////////////////////////////
#                      PUBLIC VIEW FUNCTIONS
#///////////////////////////////////////////////////////////////

@external
def getBalanceOf(_token: address) -> uint256:
    return ERC20(_token).balanceOf(self)


#///////////////////////////////////////////////////////////////
#                       EXTERNAL FUNCTIONS
#///////////////////////////////////////////////////////////////

@external
def recoverTokens(_token: address, _amount: uint256, _to: address):
    assert self.executor == msg.sender,"!EXECUTOR"
    sent: bool = ERC20(_token).transfer(_to, _amount)
    assert sent, "!TRANSFER"

@external
def recoverETH(_amount: uint256, _to: address):
    assert self.executor == msg.sender,"!EXECUTOR"
    send(_to, _amount)


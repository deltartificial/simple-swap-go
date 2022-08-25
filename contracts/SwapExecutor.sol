// @author: @deltartificial
// @dev This contract is used to execute swaps between two tokens on Uniswap V2 Router.

pragma solidity ^0.8.14;

import "interfaces/IERC20.sol";
import "interfaces/IUniswapV2.sol";

contract SwapExecutor {

    address private _executor;

    constructor() {
        _executor = msg.sender;
    }

    modifier OnlyExecutor {
        require (msg.sender == _executor, "!EXECUTOR");
        _;
    }

    /*//////////////////////////////////////////////////////////////
                             SWAP UNISWAP V2 
    //////////////////////////////////////////////////////////////*/
    
    function swapTokensOnUniswapV2(
        address _tokenIn,
        address _tokenOut,
        uint _amountIn,
        uint _amountOutMin,
        address _to,
        address _router
    ) external OnlyExecutor {
        IERC20(_tokenIn).transferFrom(msg.sender, address(this), _amountIn);
        IERC20(_tokenIn).approve(_router, _amountIn);

        address[] memory path;
        path = new address[](2);
        path[0] = _tokenIn;
        path[1] = _tokenOut;
        

        IUniswapV2Router(_router).swapExactTokensForTokens(
            _amountIn,
            _amountOutMin,
            path,
            _to,
            block.timestamp
        );
    }

    /*//////////////////////////////////////////////////////////////
                          PUBLIC VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    function getBalanceOf(address _token) external view returns(uint256) {
        return IERC20(_token).balanceOf(address(this));
    }

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    function recoverTokens(
        address _token,
        address _to,
        uint256 _amount
    ) external OnlyExecutor {
        (bool sent) = IERC20(_token).transfer( _to, _amount);
        require(sent, "!TRANSFER");
    }

    function recoverETH(
        address _to,
        uint256 _amount
    ) external OnlyExecutor {
        (bool sent, bytes memory data) = _to.call{value: _amount}("");
        require(sent, "!TRANSFER");
    }
}